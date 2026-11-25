package ai

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gow1/lib"
)

type connPalsu struct {
	lib.Conn
	terkirim []lib.IsiPesan
	presensi []string
}

func (c *connPalsu) SendChatPresence(jid string, state string) error {
	c.presensi = append(c.presensi, state)
	return nil
}

func (c *connPalsu) SendMessage(jid string, isi lib.IsiPesan, quoted *lib.Pesan) error {
	c.terkirim = append(c.terkirim, isi)
	return nil
}
func (c *connPalsu) BotNomor() string { return "6283169480682" }
func (c *connPalsu) BotJID() string   { return "6283169480682@s.whatsapp.net" }
func (c *connPalsu) BotLID() string   { return "" }

type tangkapanReq struct {
	Model    string         `json:"model"`
	Messages []pesanObrolan `json:"messages"`
}

func TestNgobrolGroupMentions(t *testing.T) {
	dir := t.TempDir()
	PathSifat = filepath.Join(dir, "sifat.md")
	if err := os.WriteFile(PathSifat, []byte("SIFAT-UNIK-123: lo Elfa"), 0o644); err != nil {
		t.Fatal(err)
	}
	var reqs []tangkapanReq
	var auth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth = r.Header.Get("Authorization")
		var tr tangkapanReq
		_ = json.NewDecoder(r.Body).Decode(&tr)
		reqs = append(reqs, tr)
		jumlahTool := 0
		for _, p := range tr.Messages {
			if p.Role == "tool" {
				jumlahTool++
			}
		}
		w.Header().Set("Content-Type", "application/json")
		switch jumlahTool {
		case 0:
			_, _ = w.Write([]byte(`{"choices":[{"message":{"content":"","tool_calls":[{"id":"c1","type":"function","function":{"name":"listFitur","arguments":"{}"}}]}}]}`))
		case 1:
			_, _ = w.Write([]byte(`{"choices":[{"message":{"content":"","tool_calls":[{"id":"c2","type":"function","function":{"name":"jalankanKode","arguments":"{\"kode\":\"1+1\"}"}}]}}]}`))
		default:
			_, _ = w.Write([]byte(`{"choices":[{"message":{"content":"halo juga njir"}}]}`))
		}
	}))
	defer srv.Close()
	t.Setenv("AI_BASE_URL", srv.URL)
	t.Setenv("AI_API_KEY", "tes-kunci")
	t.Setenv("AI_MODEL", "")

	c := &connPalsu{}
	buat := func(teks string, quoted *lib.Pesan) *lib.Konteks {
		return &lib.Konteks{Conn: c, M: &lib.Pesan{
			From:   "6281111111111@s.whatsapp.net",
			Sender: "6281111111111@s.whatsapp.net",
			Text:   teks,
			Quoted: quoted,
		}}
	}
	if err := Ngobrol(buat("halo elfa", nil), "halo elfa"); err != nil {
		t.Fatal(err)
	}
	if len(c.terkirim) != 1 {
		t.Fatalf("terkirim: %d", len(c.terkirim))
	}
	if len(c.presensi) != 2 || c.presensi[0] != "composing" || c.presensi[1] != "paused" {
		t.Fatalf("presensi salah: %v", c.presensi)
	}
	gms := c.terkirim[0].GroupMentions
	if len(gms) != 2 || gms[0].Subjek != "ask|halo elfa" || gms[1].Subjek != "halo juga njir" {
		t.Fatalf("groupMentions salah: %+v", gms)
	}
	quoted := &lib.Pesan{
		Sender:        "6283169480682@s.whatsapp.net",
		Text:          "halo juga njir",
		GroupMentions: gms,
	}
	if err := Ngobrol(buat("apa kabar", quoted), "apa kabar"); err != nil {
		t.Fatal(err)
	}
	if auth != "Bearer tes-kunci" {
		t.Fatalf("auth salah: %q", auth)
	}
	if len(reqs) != 6 || reqs[0].Model != "chatbot" {
		t.Fatalf("request/model salah: %d", len(reqs))
	}
	if !strings.Contains(reqs[0].Messages[0].Content, "SIFAT-UNIK-123") {
		t.Fatalf("system prompt gak dari sifat.md: %q", reqs[0].Messages[0].Content)
	}
	hasilList := reqs[1].Messages[len(reqs[1].Messages)-1]
	if hasilList.Role != "tool" || !strings.Contains(hasilList.Content, "mcp.js") {
		t.Fatalf("hasil listFitur salah: %+v", hasilList)
	}
	hasilKode := reqs[2].Messages[len(reqs[2].Messages)-1]
	if hasilKode.Role != "tool" || hasilKode.Content != "2" {
		t.Fatalf("hasil jalankanKode salah: %+v", hasilKode)
	}
	m2 := reqs[3].Messages
	if len(m2) != 4 {
		t.Fatalf("konteks quote gak kepakai, pesan: %d", len(m2))
	}
	if m2[1].Role != "user" || m2[1].Content != "halo elfa" ||
		m2[2].Role != "assistant" || m2[2].Content != "halo juga njir" ||
		m2[3].Content != "apa kabar" {
		t.Fatalf("isi konteks salah: %+v", m2)
	}
}
