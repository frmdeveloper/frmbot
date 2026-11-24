package lib

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestParseDurasiBan(t *testing.T) {
	kasus := map[string]int64{
		"30m":  30 * 60000,
		"2j":   2 * 3600000,
		"1h":   86400000,
		"500h": 500 * 86400000,
		"1m":   60000,
	}
	for s, mau := range kasus {
		d, ok := ParseDurasiBan(s)
		if !ok || d != mau {
			t.Fatalf("ParseDurasiBan(%q) = %d,%v mau %d", s, d, ok, mau)
		}
	}
	for _, s := range []string{"0m", "10", "abc", "1d", "", "m30"} {
		if _, ok := ParseDurasiBan(s); ok {
			t.Fatalf("ParseDurasiBan(%q) harusnya gagal", s)
		}
	}
}

func TestFormatDurasiBan(t *testing.T) {
	if got := FormatDurasiBan(500 * 86400000); got != "500 hari" {
		t.Fatalf("got %q", got)
	}
	if got := FormatDurasiBan(86400000 + 3600000); got != "1 hari 1j" {
		t.Fatalf("got %q", got)
	}
	if got := FormatDurasiBan(2*3600000 + 30*60000); got != "2j 30m" {
		t.Fatalf("got %q", got)
	}
	if got := FormatDurasiBan(45 * 60000); got != "45m" {
		t.Fatalf("got %q", got)
	}
}

func grupUji(t *testing.T) *DataGroup {
	t.Helper()
	path := filepath.Join(t.TempDir(), "group.json")
	if err := os.WriteFile(path, []byte("{}"), 0o644); err != nil {
		t.Fatal(err)
	}
	return MuatGroup(path)
}

func TestPruneDanAktif(t *testing.T) {
	g := grupUji(t)
	now := time.Now().UnixMilli()
	grup := "123@g.us"
	g.BanMember[grup] = []any{
		map[string]any{"j": "628111@s.whatsapp.net", "t": float64(now - 1000), "d": float64(500 * 86400000)},
		map[string]any{"j": "628222@s.whatsapp.net", "t": float64(now - 25*3600*1000)},
		"628333@s.whatsapp.net",
	}
	PruneBanMember(g, grup)
	daftar := DaftarBan(g, grup)
	if len(daftar) != 2 {
		t.Fatalf("setelah prune harusnya 2 entri, dapat %d: %v", len(daftar), daftar)
	}
	if !BanAktif(g, grup, "628111") {
		t.Fatal("628111 harusnya masih aktif (500 hari)")
	}
	if BanAktif(g, grup, "628222") {
		t.Fatal("628222 harusnya sudah kedaluwarsa")
	}
	if !BanAktif(g, grup, "628333") {
		t.Fatal("entri string legacy harusnya aktif")
	}
	sisa := BanSisaMs(daftar[0], now)
	if sisa <= 499*86400000 || sisa > 500*86400000 {
		t.Fatalf("sisa 500h salah: %d", sisa)
	}
}

func TestLepasBan(t *testing.T) {
	g := grupUji(t)
	grup := "123@g.us"
	g.BanMember[grup] = []any{
		map[string]any{"j": "628111@s.whatsapp.net", "t": float64(time.Now().UnixMilli())},
		map[string]any{"j": "628222@lid", "t": float64(time.Now().UnixMilli())},
	}
	if !LepasBan(g, grup, "628222") {
		t.Fatal("LepasBan via nomor LID harusnya berhasil")
	}
	if LepasBan(g, grup, "628999") {
		t.Fatal("LepasBan nomor tak dikenal harusnya false")
	}
	if len(DaftarBan(g, grup)) != 1 {
		t.Fatal("sisa daftar harusnya 1")
	}
}
