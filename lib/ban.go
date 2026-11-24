package lib

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

const BanMemberResetMs int64 = 24 * 3600 * 1000

var reDurasiBan = regexp.MustCompile(`(?i)^(\d+)(m|j|h)$`)

func ParseDurasiBan(s string) (int64, bool) {
	cocok := reDurasiBan.FindStringSubmatch(strings.TrimSpace(s))
	if cocok == nil {
		return 0, false
	}
	n, err := strconv.ParseInt(cocok[1], 10, 64)
	if err != nil {
		return 0, false
	}
	var pengali int64
	switch strings.ToLower(cocok[2]) {
	case "m":
		pengali = 60000
	case "j":
		pengali = 3600000
	case "h":
		pengali = 86400000
	}
	ms := n * pengali
	if ms < 60000 {
		return 0, false
	}
	return ms, true
}

func FormatDurasiBan(ms int64) string {
	if ms >= 86400000 {
		h := ms / 86400000
		sj := (ms % 86400000) / 3600000
		if sj > 0 {
			return strconv.FormatInt(h, 10) + " hari " + strconv.FormatInt(sj, 10) + "j"
		}
		return strconv.FormatInt(h, 10) + " hari"
	}
	j := ms / 3600000
	mnt := (ms % 3600000) / 60000
	if j > 0 && mnt > 0 {
		return strconv.FormatInt(j, 10) + "j " + strconv.FormatInt(mnt, 10) + "m"
	}
	if j > 0 {
		return strconv.FormatInt(j, 10) + "j"
	}
	return strconv.FormatInt(mnt, 10) + "m"
}

func BanJid(e any) string {
	switch t := e.(type) {
	case string:
		return t
	case map[string]any:
		s, _ := t["j"].(string)
		return s
	}
	return ""
}

func banWaktu(e any, now int64) (int64, int64) {
	if _, ok := e.(string); ok {
		return now, BanMemberResetMs
	}
	if t, ok := e.(map[string]any); ok {
		var waktu, dur int64
		if f, ok := t["t"].(float64); ok {
			waktu = int64(f)
		}
		if f, ok := t["d"].(float64); ok {
			dur = int64(f)
		}
		if dur <= 0 {
			dur = BanMemberResetMs
		}
		return waktu, dur
	}
	return now, BanMemberResetMs
}

func BanSisaMs(e any, now int64) int64 {
	t, d := banWaktu(e, now)
	sisa := d - (now - t)
	if sisa < 0 {
		return 0
	}
	return sisa
}

func DaftarBan(g *DataGroup, grup string) []any {
	if g == nil || g.BanMember == nil {
		return nil
	}
	daftar, _ := g.BanMember[grup].([]any)
	return daftar
}

func PruneBanMember(g *DataGroup, grup string) {
	daftar := DaftarBan(g, grup)
	if len(daftar) == 0 {
		return
	}
	now := time.Now().UnixMilli()
	baru := []any{}
	berubah := false
	for _, e := range daftar {
		j := BanJid(e)
		if j == "" {
			berubah = true
			continue
		}
		if s, ok := e.(string); ok {
			berubah = true
			baru = append(baru, map[string]any{"j": s, "t": float64(now)})
			continue
		}
		t, d := banWaktu(e, now)
		if now-t > d {
			berubah = true
			continue
		}
		baru = append(baru, e)
	}
	if berubah {
		if g.BanMember == nil {
			g.BanMember = map[string]any{}
		}
		g.BanMember[grup] = baru
		_ = g.Simpan()
	}
}

func BanAktif(g *DataGroup, grup string, identitas ...string) bool {
	daftar := DaftarBan(g, grup)
	now := time.Now().UnixMilli()
	for _, e := range daftar {
		n := NomorSaja(BanJid(e))
		if n == "" {
			continue
		}
		for _, id := range identitas {
			if id != "" && n == id {
				if _, isStr := e.(string); isStr {
					return true
				}
				t, d := banWaktu(e, now)
				if now-t <= d {
					return true
				}
			}
		}
	}
	return false
}

func LepasBan(g *DataGroup, grup string, identitas ...string) bool {
	daftar := DaftarBan(g, grup)
	if len(daftar) == 0 {
		return false
	}
	baru := []any{}
	for _, e := range daftar {
		n := NomorSaja(BanJid(e))
		kena := false
		for _, id := range identitas {
			if id != "" && n == id {
				kena = true
			}
		}
		if !kena {
			baru = append(baru, e)
		}
	}
	if len(baru) == len(daftar) {
		return false
	}
	if g.BanMember == nil {
		g.BanMember = map[string]any{}
	}
	g.BanMember[grup] = baru
	_ = g.Simpan()
	return true
}
