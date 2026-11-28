package group

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"

	"gow1/lib"
)

const grupKhususPecat = "6285757196481-1625886618@g.us"

func identitasDari(k *lib.Konteks, gc *lib.GrupInfo, target string) []string {
	set := map[string]bool{}
	if n := lib.NomorSaja(target); n != "" {
		set[n] = true
	}
	if gc != nil {
		if p := cariPeserta(gc, target); p != nil {
			for _, x := range []string{p.Nomor, p.LID, p.JID} {
				if n := lib.NomorSaja(x); n != "" {
					set[n] = true
				}
			}
		}
	}
	hasil := []string{}
	for n := range set {
		hasil = append(hasil, n)
	}
	return hasil
}

func targetBanDariPesan(k *lib.Konteks) string {
	if k.M.Quoted != nil && k.M.Quoted.Sender != "" {
		return k.M.Quoted.Sender
	}
	if len(k.M.Mentions) > 0 {
		return k.M.Mentions[0]
	}
	for _, a := range strings.Fields(k.M.Q) {
		if _, ok := lib.ParseDurasiBan(a); ok {
			continue
		}
		digit := nonDigit.ReplaceAllString(a, "")
		if len(digit) >= 5 {
			return digit + "@s.whatsapp.net"
		}
	}
	return ""
}

func TanganiAgrup(k *lib.Konteks) error {
	m := k.M
	if !m.IsGroup {
		return k.Conn.OK(m, "Bukan grup")
	}
	cmd := m.Command

	if strings.Contains(cmd, "setwelcome") {
		if m.Q == "" {
			if k.Group.Welcome != nil {
				delete(k.Group.Welcome, m.From)
			}
			_ = k.Group.Simpan()
			return k.Conn.OK(m, "Welcome dimatikn soalnya kosonh")
		}
		if k.Group.Welcome == nil {
			k.Group.Welcome = map[string]any{}
		}
		k.Group.Welcome[m.From] = m.Q
		_ = k.Group.Simpan()
		return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: m.Q}, m)
	}
	if strings.Contains(cmd, "setgoodbye") {
		if m.Q == "" {
			if k.Group.Goodbye != nil {
				delete(k.Group.Goodbye, m.From)
			}
			_ = k.Group.Simpan()
			return k.Conn.OK(m, "Goodbye dimatikn soalnya kosonh")
		}
		if k.Group.Goodbye == nil {
			k.Group.Goodbye = map[string]any{}
		}
		k.Group.Goodbye[m.From] = m.Q
		_ = k.Group.Simpan()
		return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: m.Q}, m)
	}

	if strings.Contains(cmd, "banmember") && !strings.Contains(cmd, "listbanmember") {
		gc, err := k.Conn.GetGroupInfo(m.From)
		if err != nil {
			return err
		}
		if !botAdmin(k, gc) {
			return k.Conn.OK(m, "Gw bukan admin")
		}
		if !pengirimAdmin(k, gc) && !pengirimOwner(k) {
			return k.Conn.OK(m, "Lu bukan admin")
		}
	}

	if strings.Contains(cmd, "unbanaku") {
		lib.PruneBanMember(k.Group, m.From)
		saya := []string{lib.NomorSaja(m.Sender), lib.NomorSaja(m.SenderAlt)}
		if lib.LepasBan(k.Group, m.From, saya...) {
			return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: "hoki lu nemu fitur rahasia ini jir — ban lu gw lepas. jangan ngespam lagi"}, nil)
		}
		return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: "lu gak lagi ke-ban jir, ngapain unban"}, nil)
	}

	if strings.Contains(cmd, "unbanmember") {
		lib.PruneBanMember(k.Group, m.From)
		if len(lib.DaftarBan(k.Group, m.From)) == 0 {
			return k.Conn.OK(m, "Belum ada yang dibanned")
		}
		target := targetBanDariPesan(k)
		if target == "" {
			return k.Conn.OK(m, "Tag atau reply orangnya")
		}
		gc, _ := k.Conn.GetGroupInfo(m.From)
		kenal := identitasDari(k, gc, target)
		phoneNum := lib.NomorSaja(target)
		if lib.LepasBan(k.Group, m.From, kenal...) {
			teks := fmt.Sprintf("@%s dilepas dari ban", phoneNum)
			return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: teks, Mentions: []string{target}}, m)
		}
		return k.Conn.OK(m, "Dia gak dibanned kok")
	} else if strings.Contains(cmd, "listban") {
		lib.PruneBanMember(k.Group, m.From)
		daftar := lib.DaftarBan(k.Group, m.From)
		if len(daftar) == 0 {
			return k.Conn.OK(m, "Belum ada yang dibanned")
		}
		now := time.Now().UnixMilli()
		baris := []string{}
		mentions := []string{}
		for i, x := range daftar {
			sisa := lib.BanSisaMs(x, now)
			sj := sisa / 3600000
			sm := (sisa % 3600000) / 60000
			baris = append(baris, fmt.Sprintf("%d. @%s (%dj %dm lagi)", i+1, lib.NomorSaja(lib.BanJid(x)), sj, sm))
			mentions = append(mentions, lib.BanJid(x))
		}
		teks := "Banned:\n" + strings.Join(baris, "\n")
		return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: teks, Mentions: mentions}, m)
	} else if strings.Contains(cmd, "banmember") {
		lib.PruneBanMember(k.Group, m.From)
		var durMs int64
		for _, a := range strings.Fields(m.Q) {
			if d, ok := lib.ParseDurasiBan(a); ok {
				durMs = d
				break
			}
		}
		target := targetBanDariPesan(k)
		if target == "" {
			return k.Conn.OK(m, "Tag atau reply orangnya")
		}
		gc, _ := k.Conn.GetGroupInfo(m.From)
		kenal := identitasDari(k, gc, target)
		kenalPapa := identitasDari(k, gc, lib.NomorOwner()+"@s.whatsapp.net")
		for _, n := range kenal {
			for _, p := range kenalPapa {
				if n == p {
					return k.Conn.OK(m, "itu papa gw anjir, aku gamau jadi anak durhaka")
				}
			}
		}
		if lib.BanAktif(k.Group, m.From, kenal...) {
			return k.Conn.OK(m, "Udah dibanned sebelumnya")
		}
		entri := map[string]any{"j": target, "t": float64(time.Now().UnixMilli())}
		if durMs > 0 {
			entri["d"] = float64(durMs)
		}
		daftar := append(lib.DaftarBan(k.Group, m.From), entri)
		if k.Group.BanMember == nil {
			k.Group.BanMember = map[string]any{}
		}
		k.Group.BanMember[m.From] = daftar
		_ = k.Group.Simpan()
		label := "24 jam"
		if durMs > 0 {
			label = lib.FormatDurasiBan(durMs)
		}
		phoneNum := lib.NomorSaja(target)
		teks := fmt.Sprintf("@%s dibanned dari chat (%s)", phoneNum, label)
		return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: teks, Mentions: []string{target}}, m)
	}
	return nil
}

func jidStrings(jids []types.JID) []string {
	hasil := make([]string, len(jids))
	for i, j := range jids {
		hasil[i] = j.String()
	}
	return hasil
}

func teksSambutan(template string, jids []types.JID) (string, []string) {
	sebutan := make([]string, len(jids))
	for i, j := range jids {
		sebutan[i] = "@" + lib.NomorSaja(j.String())
	}
	return strings.ReplaceAll(template, "@", strings.Join(sebutan, ",")), jidStrings(jids)
}

func AksiGrup(k *lib.Konteks, evt *events.GroupInfo) error {
	if evt == nil {
		return nil
	}
	jidGrup := evt.JID.String()

	if len(evt.Leave) > 0 {
		botIkutKeluar := false
		for _, j := range evt.Leave {
			n := lib.NomorSaja(j.String())
			if n == k.Conn.BotNomor() || n == lib.NomorSaja(k.Conn.BotLID()) {
				botIkutKeluar = true
				break
			}
		}
		if botIkutKeluar {
			pengkick := ""
			if evt.SenderPN != nil {
				pengkick = evt.SenderPN.String()
			} else if evt.Sender != nil {
				pengkick = evt.Sender.String()
			}
			if strings.HasSuffix(pengkick, "@s.whatsapp.net") {
				_ = k.Conn.SendMessage(pengkick, lib.IsiPesan{Text: "maksud lu apa ngekick gw"}, nil)
			}
		}
	}

	if jidGrup == grupKhususPecat {
		pelaku := ""
		if evt.Sender != nil {
			pelaku = evt.Sender.String()
		}
		pelakuPN := ""
		if evt.SenderPN != nil {
			pelakuPN = evt.SenderPN.String()
		}
		pelakuOwner := adalahOwner(k, pelakuPN) || adalahOwner(k, pelaku)
		if len(evt.Promote) > 0 && !pelakuOwner {
			target := append([]string{pelaku}, jidStrings(evt.Promote)...)
			_ = k.Conn.UpdateGroupParticipants(jidGrup, target, "demote")
			teks := fmt.Sprintf("@%s dipecat karena melanggar", lib.NomorSaja(pelaku))
			_ = k.Conn.SendMessage(jidGrup, lib.IsiPesan{Text: teks, Mentions: []string{pelaku}}, nil)
		}
		if len(evt.Demote) > 0 && !pelakuOwner {
			_ = k.Conn.UpdateGroupParticipants(jidGrup, []string{pelaku}, "demote")
			_ = k.Conn.UpdateGroupParticipants(jidGrup, jidStrings(evt.Demote), "promote")
			teks := fmt.Sprintf("@%s dipecat karena melanggar", lib.NomorSaja(pelaku))
			_ = k.Conn.SendMessage(jidGrup, lib.IsiPesan{Text: teks, Mentions: []string{pelaku}}, nil)
		}
		return nil
	}

	if len(evt.Leave) > 0 && k.Group.Goodbye != nil {
		if teks, ok := k.Group.Goodbye[jidGrup].(string); ok && teks != "" {
			isi, mentions := teksSambutan(teks, evt.Leave)
			_ = k.Conn.SendMessage(jidGrup, lib.IsiPesan{Text: isi, Mentions: mentions}, nil)
		}
	}
	if len(evt.Join) > 0 && k.Group.Welcome != nil {
		if teks, ok := k.Group.Welcome[jidGrup].(string); ok && teks != "" {
			isi, mentions := teksSambutan(teks, evt.Join)
			_ = k.Conn.SendMessage(jidGrup, lib.IsiPesan{Text: isi, Mentions: mentions}, nil)
		}
	}
	if len(evt.Join) > 0 {
		if bufWp, err := os.ReadFile(berkasWallpaperGrup(jidGrup)); err == nil && len(bufWp) > 0 {
			botNum := k.Conn.BotNomor()
			botLid := lib.NomorSaja(k.Conn.BotLID())
			baruSelainBot := false
			for _, j := range evt.Join {
				n := lib.NomorSaja(j.String())
				if n != "" && n != botNum && n != botLid {
					baruSelainBot = true
					break
				}
			}
			if baruSelainBot && jedaKirimUlangWallpaper(jidGrup) {
				fmt.Printf("[wp-ulang] ada join baru di %s - kirim ulang wallpaper dijadwalkan\n", jidGrup)
				go func() {
					kirim := func(coba int) {
						if err := k.Conn.SetGroupWallpaper(jidGrup, bufWp); err != nil {
							fmt.Printf("[wp-ulang] GAGAL coba %d ke %s - %v\n", coba, jidGrup, err)
						} else {
							fmt.Printf("[wp-ulang] terkirim coba %d ke %s\n", coba, jidGrup)
						}
					}
					time.Sleep(5 * time.Second)
					kirim(1)
					time.Sleep(40 * time.Second)
					kirim(2)
				}()
			}
		}
	}
	return nil
}

var wpKirimUlang = struct {
	sync.Mutex
	data map[string]time.Time
}{data: map[string]time.Time{}}

func jedaKirimUlangWallpaper(grup string) bool {
	wpKirimUlang.Lock()
	defer wpKirimUlang.Unlock()
	if time.Since(wpKirimUlang.data[grup]) <= 60*time.Second {
		return false
	}
	wpKirimUlang.data[grup] = time.Now()
	return true
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "agrup.js",
		Jenis:    "group",
		Nama:     []string{"#setwelcome", "#setgoodbye", "#banmember", "#unbanmember", "#listbanmember"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(setwelcome|setgoodbye|banmember|unbanmember|listbanmember|listban|unbanaku)$"),
		Jalankan: TanganiAgrup,
	})
}
