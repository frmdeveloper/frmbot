package utility

import (
	"strings"
	"testing"

	"gow1/lib"
)

type connPalsu struct {
	terkirim []lib.IsiPesan
}

func (c *connPalsu) SendMessage(jid string, isi lib.IsiPesan, quoted *lib.Pesan) error {
	c.terkirim = append(c.terkirim, isi)
	return nil
}
func (c *connPalsu) OK(m *lib.Pesan, teks string) error {
	c.terkirim = append(c.terkirim, lib.IsiPesan{Text: teks})
	return nil
}
func (c *connPalsu) DownloadAny(m *lib.Pesan) ([]byte, string, error) { return nil, "", nil }
func (c *connPalsu) GetGroupInfo(jid string) (*lib.GrupInfo, error)  { return nil, nil }
func (c *connPalsu) GetJoinedGroups() ([]*lib.GrupInfo, error)       { return nil, nil }
func (c *connPalsu) UpdateGroupParticipants(jid string, p []string, aksi string) error {
	return nil
}
func (c *connPalsu) GetProfilePicture(jid string) (string, error) { return "", nil }
func (c *connPalsu) GetGroupInviteLink(jid string) (string, error) {
	return "", nil
}
func (c *connPalsu) GetGroupInviteInfo(link string) (*lib.GrupInfo, error) { return nil, nil }
func (c *connPalsu) SetGroupName(jid string, nama string) error            { return nil }
func (c *connPalsu) SetGroupAnnounce(jid string, a bool) error             { return nil }
func (c *connPalsu) IsOnWhatsApp(nomor string) (bool, string, error)       { return false, "", nil }
func (c *connPalsu) KirimStikerBersuara(jid string, s, a []byte, d int, q *lib.Pesan) error {
	return nil
}
func (c *connPalsu) SendPresence(state string) error { return nil }
func (c *connPalsu) SendChatPresence(jid string, state string) error       { return nil }
func (c *connPalsu) RevokeMessage(jid string, m *lib.Pesan) error          { return nil }
func (c *connPalsu) RelayMentah(jid string, mentah any) error              { return nil }
func (c *connPalsu) BotJID() string                                        { return "6283169480682@s.whatsapp.net" }
func (c *connPalsu) BotLID() string                                        { return "" }
func (c *connPalsu) BotNomor() string                                      { return "6283169480682" }
func (c *connPalsu) Terhubung() bool                                       { return true }

func konteksUji(perintah, q string) (*lib.Konteks, *connPalsu) {
	c := &connPalsu{}
	m := &lib.Pesan{
		From:      "6285188261898@s.whatsapp.net",
		Sender:    "36572367954136@lid",
		SenderAlt: "6285188261898@s.whatsapp.net",
		Command:   perintah,
		Q:         q,
		Text:      perintah + " " + q,
	}
	return &lib.Konteks{Conn: c, M: m}, c
}

func TestExShell(t *testing.T) {
	k, c := konteksUji(".ex", "echo halo-$from")
	if err := TanganiEval(k); err != nil {
		t.Fatal(err)
	}
	if len(c.terkirim) != 1 || !strings.Contains(c.terkirim[0].Text, "halo-6285188261898@s.whatsapp.net") {
		t.Fatalf("hasil .ex salah: %+v", c.terkirim)
	}
}

func TestEEkspresi(t *testing.T) {
	k, c := konteksUji(".e", "return 1 + 1")
	if err := TanganiEval(k); err != nil {
		t.Fatal(err)
	}
	if len(c.terkirim) != 1 || c.terkirim[0].Text != "2" {
		t.Fatalf("hasil .e salah: %+v", c.terkirim)
	}
}

func TestEObjekDanConn(t *testing.T) {
	k, c := konteksUji(".e", "conn.ok('tes dari eval'); return { nomor: conn.number, dari: m.from }")
	if err := TanganiEval(k); err != nil {
		t.Fatal(err)
	}
	if len(c.terkirim) != 2 {
		t.Fatalf("jumlah kiriman salah: %+v", c.terkirim)
	}
	if c.terkirim[0].Text != "tes dari eval" {
		t.Fatalf("conn.ok salah: %+v", c.terkirim[0])
	}
	if !strings.Contains(c.terkirim[1].Text, "6283169480682") {
		t.Fatalf("hasil objek salah: %+v", c.terkirim[1])
	}
}

func TestEBukanOwnerDiam(t *testing.T) {
	k, c := konteksUji(".ex", "echo halo")
	k.M.SenderAlt = "6289999999999@s.whatsapp.net"
	if err := TanganiEval(k); err != nil {
		t.Fatal(err)
	}
	if len(c.terkirim) != 0 {
		t.Fatalf("non-owner harus diam: %+v", c.terkirim)
	}
}

func TestEKosong(t *testing.T) {
	k, c := konteksUji(".e", "")
	if err := TanganiEval(k); err != nil {
		t.Fatal(err)
	}
	if len(c.terkirim) != 1 || !strings.Contains(c.terkirim[0].Text, "astagfirullahaladzim") {
		t.Fatalf("hasil .e kosong salah: %+v", c.terkirim)
	}
}
