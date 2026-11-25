package group

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gow1/lib"
)

type connGrupPalsu struct {
	lib.Conn
	terkirim       []lib.IsiPesan
	quotedTerkirim []*lib.Pesan
}

func (c *connGrupPalsu) SendMessage(jid string, isi lib.IsiPesan, quoted *lib.Pesan) error {
	c.terkirim = append(c.terkirim, isi)
	c.quotedTerkirim = append(c.quotedTerkirim, quoted)
	return nil
}
func (c *connGrupPalsu) OK(m *lib.Pesan, teks string) error {
	c.terkirim = append(c.terkirim, lib.IsiPesan{Text: teks})
	c.quotedTerkirim = append(c.quotedTerkirim, m)
	return nil
}
func (c *connGrupPalsu) BotNomor() string { return "6283169480682" }
func (c *connGrupPalsu) BotJID() string   { return "6283169480682@s.whatsapp.net" }
func (c *connGrupPalsu) BotLID() string   { return "" }
func (c *connGrupPalsu) GetGroupInfo(jid string) (*lib.GrupInfo, error) {
	return &lib.GrupInfo{
		JID: jid,
		Peserta: []lib.Peserta{
			{JID: "6283169480682@s.whatsapp.net", Nomor: "6283169480682@s.whatsapp.net", IsAdmin: true},
			{JID: "6285188261898@s.whatsapp.net", Nomor: "6285188261898@s.whatsapp.net", IsAdmin: true},
			{JID: "6281111111111@s.whatsapp.net", Nomor: "6281111111111@s.whatsapp.net"},
		},
	}, nil
}

func konteksGrup(t *testing.T, c *connGrupPalsu, perintah, q, pengirim string, mentions []string) *lib.Konteks {
	t.Helper()
	path := filepath.Join(t.TempDir(), "group.json")
	if err := os.WriteFile(path, []byte("{}"), 0o644); err != nil {
		t.Fatal(err)
	}
	g := lib.MuatGroup(path)
	m := &lib.Pesan{
		From:     "120363000000000000@g.us",
		Sender:   pengirim,
		Command:  perintah,
		Q:        q,
		IsGroup:  true,
		Mentions: mentions,
	}
	return &lib.Konteks{Conn: c, M: m, Group: g}
}

func TestBanmemberDurasiDanListbanDanUnbanaku(t *testing.T) {
	c := &connGrupPalsu{}
	k := konteksGrup(t, c, "#banmember", "@6281111111111 500h", "6285188261898@s.whatsapp.net", []string{"6281111111111@s.whatsapp.net"})
	if err := TanganiAgrup(k); err != nil {
		t.Fatal(err)
	}
	if len(c.terkirim) != 1 || !strings.Contains(c.terkirim[0].Text, "dibanned dari chat (500 hari)") {
		t.Fatalf("balasan banmember salah: %+v", c.terkirim)
	}
	daftar := lib.DaftarBan(k.Group, k.M.From)
	if len(daftar) != 1 {
		t.Fatalf("daftar ban harusnya 1: %v", daftar)
	}
	entri, _ := daftar[0].(map[string]any)
	if entri["d"] != float64(500*86400000) {
		t.Fatalf("field d salah: %v", entri)
	}

	c.terkirim = nil
	k.M.Command = "#listbanmember"
	k.M.Q = ""
	k.M.Mentions = nil
	if err := TanganiAgrup(k); err != nil {
		t.Fatal(err)
	}
	if len(c.terkirim) != 1 || !strings.Contains(c.terkirim[0].Text, "@6281111111111 (") || !strings.Contains(c.terkirim[0].Text, "lagi)") {
		t.Fatalf("listbanmember salah: %+v", c.terkirim)
	}
	if !strings.Contains(c.terkirim[0].Text, "11999j") && !strings.Contains(c.terkirim[0].Text, "12000j") {
		t.Fatalf("sisa waktu listban harusnya ~12000j: %+v", c.terkirim[0].Text)
	}

	c.terkirim = nil
	c.quotedTerkirim = nil
	k.M.Command = "#unbanaku"
	k.M.Sender = "6281111111111@s.whatsapp.net"
	if err := TanganiAgrup(k); err != nil {
		t.Fatal(err)
	}
	if len(c.terkirim) != 1 || !strings.Contains(c.terkirim[0].Text, "hoki lu nemu fitur rahasia ini jir") {
		t.Fatalf("unbanaku salah: %+v", c.terkirim)
	}
	if len(c.quotedTerkirim) != 1 || c.quotedTerkirim[0] != nil {
		t.Fatalf("unbanaku gak boleh nge-quote: %+v", c.quotedTerkirim)
	}
	if len(lib.DaftarBan(k.Group, k.M.From)) != 0 {
		t.Fatal("ban harusnya sudah lepas")
	}

	c.terkirim = nil
	c.quotedTerkirim = nil
	if err := TanganiAgrup(k); err != nil {
		t.Fatal(err)
	}
	if len(c.quotedTerkirim) != 1 || c.quotedTerkirim[0] != nil {
		t.Fatalf("unbanaku kedua gak boleh nge-quote: %+v", c.quotedTerkirim)
	}
	if len(c.terkirim) != 1 || !strings.Contains(c.terkirim[0].Text, "lu gak lagi ke-ban jir") {
		t.Fatalf("unbanaku kedua salah: %+v", c.terkirim)
	}
}

func TestBanmemberPapaKebal(t *testing.T) {
	c := &connGrupPalsu{}
	k := konteksGrup(t, c, "#banmember", "@6285188261898", "6281111111111@s.whatsapp.net", []string{"6285188261898@s.whatsapp.net"})
	k.M.Sender = "6285188261898@s.whatsapp.net"
	if err := TanganiAgrup(k); err != nil {
		t.Fatal(err)
	}
	ketemu := false
	for _, isi := range c.terkirim {
		if strings.Contains(isi.Text, "itu papa gw anjir, aku gamau jadi anak durhaka") {
			ketemu = true
		}
	}
	if !ketemu {
		t.Fatalf("ban papa harusnya ditolak: %+v", c.terkirim)
	}
	if len(lib.DaftarBan(k.Group, k.M.From)) != 0 {
		t.Fatal("papa gak boleh masuk daftar ban")
	}
}
