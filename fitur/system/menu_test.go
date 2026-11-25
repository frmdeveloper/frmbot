package system

import (
	"strings"
	"testing"

	"gow1/lib"

	_ "gow1/fitur/audio"
)

type mockConn struct {
	terakhirTeks string
}

func (m *mockConn) SendMessage(jid string, isi lib.IsiPesan, quoted *lib.Pesan) error {
	m.terakhirTeks = isi.Text
	return nil
}
func (m *mockConn) OK(msg *lib.Pesan, teks string) error                                              { return nil }
func (m *mockConn) DownloadAny(msg *lib.Pesan) ([]byte, string, error)                               { return nil, "", nil }
func (m *mockConn) GetGroupInfo(jid string) (*lib.GrupInfo, error)                                   { return nil, nil }
func (m *mockConn) GetJoinedGroups() ([]*lib.GrupInfo, error)                                        { return nil, nil }
func (m *mockConn) UpdateGroupParticipants(jid string, p []string, a string) error                   { return nil }
func (m *mockConn) GetProfilePicture(jid string) (string, error)                                     { return "", nil }
func (m *mockConn) GetGroupInviteLink(jid string) (string, error)                                    { return "", nil }
func (m *mockConn) GetGroupInviteInfo(link string) (*lib.GrupInfo, error)                             { return nil, nil }
func (m *mockConn) SetGroupName(jid string, nama string) error                                        { return nil }
func (m *mockConn) SetGroupAnnounce(jid string, announce bool) error                                 { return nil }
func (m *mockConn) IsOnWhatsApp(nomor string) (bool, string, error)                                   { return true, "", nil }
func (m *mockConn) KirimStikerBersuara(jid string, s, a []byte, d int, q *lib.Pesan, mt string) error { return nil }
func (m *mockConn) SendPresence(state string) error                                                  { return nil }
func (m *mockConn) SendChatPresence(jid string, state string) error                                  { return nil }
func (m *mockConn) RevokeMessage(jid string, msg *lib.Pesan) error                                   { return nil }
func (m *mockConn) RelayMentah(jid string, mentah any) error                                         { return nil }
func (m *mockConn) React(msg *lib.Pesan, emoji string) error                                         { return nil }
func (m *mockConn) SendButtons(jid string, teks string, footer string, buttons []lib.Tombol, quoted *lib.Pesan) error {
	return nil
}
func (m *mockConn) SendList(jid string, judul, teks, footer, teksTombol string, bagian []lib.BagianList, quoted *lib.Pesan) error {
	return nil
}
func (m *mockConn) BotJID() string   { return "" }
func (m *mockConn) BotLID() string   { return "" }
func (m *mockConn) BotNomor() string { return "" }
func (m *mockConn) Terhubung() bool  { return true }

func TestMenuFormat(t *testing.T) {
	conn := &mockConn{}
	k := &lib.Konteks{
		Conn: conn,
		M: &lib.Pesan{
			From:    "123@s.whatsapp.net",
			Command: ".menu",
			Prefix:  ".",
			Q:       "",
		},
	}

	err := TanganiMenu(k)
	if err != nil {
		t.Fatalf("TanganiMenu error: %v", err)
	}

	if !strings.Contains(conn.terakhirTeks, "╭─「 *daftar kategori* 」") {
		t.Errorf("Header kategori tidak sesuai: %s", conn.terakhirTeks)
	}
	if !strings.Contains(conn.terakhirTeks, "╰──⊷") {
		t.Errorf("Footer box tidak sesuai: %s", conn.terakhirTeks)
	}

	// Larangan 1: Tidak boleh ada emoji
	emojiList := []string{"🤖", "🎧", "💬", "🔄", "📥", "🎮", "👥", "🔍", "🎨", "⚙️", "🧰", "✦", "💾", "⚡", "🎵", "🪡", "💡"}
	for _, em := range emojiList {
		if strings.Contains(conn.terakhirTeks, em) {
			t.Errorf("Ditemukan emoji terlarang %s di teks menu", em)
		}
	}

	// Larangan 2: Tidak boleh ada tanda kurung sebagai penjelasan
	if strings.Contains(conn.terakhirTeks, "(") || strings.Contains(conn.terakhirTeks, ")") {
		t.Errorf("Ditemukan tanda kurung terlarang di teks menu: %s", conn.terakhirTeks)
	}

	// Larangan 3: Tidak boleh ada iklan whatsmeow/golang/tiktok
	iklan := []string{"whatsmeow", "golang", "tiktok.com", "github.com/frmdeveloper"}
	for _, ik := range iklan {
		if strings.Contains(strings.ToLower(conn.terakhirTeks), ik) {
			t.Errorf("Ditemukan iklan terlarang %s di teks menu", ik)
		}
	}

	// Test kategori audio
	kAudio := &lib.Konteks{
		Conn: conn,
		M: &lib.Pesan{
			From:    "123@s.whatsapp.net",
			Command: ".menu",
			Prefix:  ".",
			Q:       "audio",
		},
	}
	err = TanganiMenu(kAudio)
	if err != nil {
		t.Fatalf("TanganiMenu audio error: %v", err)
	}
	if !strings.Contains(conn.terakhirTeks, "╭─「 *audio* 」") {
		t.Errorf("Header audio tidak sesuai: %s", conn.terakhirTeks)
	}
	if !strings.Contains(conn.terakhirTeks, "╰──⊷") {
		t.Errorf("Footer audio tidak sesuai: %s", conn.terakhirTeks)
	}
	if strings.Contains(conn.terakhirTeks, "(") || strings.Contains(conn.terakhirTeks, ")") {
		t.Errorf("Ditemukan tanda kurung di menu audio: %s", conn.terakhirTeks)
	}
}
