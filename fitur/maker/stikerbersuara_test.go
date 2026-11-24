package maker

import (
	"encoding/base64"
	"os"
	"os/exec"
	"strings"
	"testing"

	"gow1/lib"
)

type connSb struct {
	lib.Conn
	png     []byte
	mp3     []byte
	vid     []byte
	terkirim []lib.IsiPesan
	oks     []string
	sbJid   string
	sbStik  []byte
	sbAudio []byte
	sbDetik int
	sbCalls int
}

func (c *connSb) DownloadAny(p *lib.Pesan) ([]byte, string, error) {
	switch p.Type {
	case "imageMessage":
		return c.png, "image/png", nil
	case "audioMessage":
		return c.mp3, "audio/mpeg", nil
	case "videoMessage":
		return c.vid, "video/mp4", nil
	}
	return nil, "", os.ErrNotExist
}

func (c *connSb) SendMessage(jid string, isi lib.IsiPesan, quoted *lib.Pesan) error {
	c.terkirim = append(c.terkirim, isi)
	return nil
}

func (c *connSb) OK(m *lib.Pesan, teks string) error {
	c.oks = append(c.oks, teks)
	return nil
}

func (c *connSb) KirimStikerBersuara(jid string, stiker, audio []byte, detik int, quoted *lib.Pesan) error {
	c.sbCalls++
	c.sbJid, c.sbStik, c.sbAudio, c.sbDetik = jid, stiker, audio, detik
	return nil
}

func TestStikerBersuaraAlur(t *testing.T) {
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		t.Skip("ffmpeg gak ada")
	}
	png, _ := base64.StdEncoding.DecodeString("iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z8BQDwAEhQGAhKmMIQAAAABJRU5ErkJggg==")
	tmpMp3 := t.TempDir() + "/nada.mp3"
	if err := exec.Command("ffmpeg", "-y", "-v", "error", "-f", "lavfi", "-i", "sine=frequency=440:duration=1", tmpMp3).Run(); err != nil {
		t.Skip("bikin mp3 tes gagal")
	}
	mp3, err := os.ReadFile(tmpMp3)
	if err != nil {
		t.Fatal(err)
	}
	pengirim := "628111111111@s.whatsapp.net"
	grup := "120363000000000000@g.us"

	c := &connSb{png: png, mp3: mp3}
	mA := &lib.Pesan{From: grup, Sender: pengirim, IsGroup: true, Type: "textMessage", Text: "#s2", Command: "#s2"}
	if err := TanganiStikerBersuara(&lib.Konteks{Conn: c, M: mA}); err != nil {
		t.Fatal(err)
	}
	if len(c.terkirim) != 1 || c.terkirim[0].Text != "Kirim fotonya coba" ||
		len(c.terkirim[0].GroupMentions) != 1 || c.terkirim[0].GroupMentions[0].Subjek != "cmd|#stikerbersuara foto" {
		t.Fatalf("prompt foto salah: %+v", c.terkirim)
	}

	c.terkirim = nil
	mB := &lib.Pesan{From: grup, Sender: pengirim, IsGroup: true, Type: "imageMessage", Command: "#stikerbersuara", Q: "foto"}
	if err := TanganiStikerBersuara(&lib.Konteks{Conn: c, M: mB}); err != nil {
		t.Fatal(err)
	}
	if len(c.terkirim) != 1 || !strings.HasPrefix(c.terkirim[0].Text, "Sekarang kirim audionya coba") {
		t.Fatalf("prompt audio salah: %+v", c.terkirim)
	}
	subjek := c.terkirim[0].GroupMentions[0].Subjek
	if !strings.HasPrefix(subjek, "cmd|#stikerbersuara suara sampah/sb628111111111") {
		t.Fatalf("subjek tahap suara salah: %s", subjek)
	}
	path := strings.TrimPrefix(subjek, "cmd|#stikerbersuara suara ")
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("berkas gambar tahap gak kesimpen: %v", err)
	}
	defer os.Remove(path)

	mC := &lib.Pesan{From: grup, Sender: pengirim, IsGroup: true, Type: "audioMessage", Command: "#stikerbersuara", Q: "suara " + path}
	if err := TanganiStikerBersuara(&lib.Konteks{Conn: c, M: mC}); err != nil {
		t.Fatal(err)
	}
	if c.sbCalls != 1 || len(c.sbStik) == 0 || len(c.sbAudio) == 0 || c.sbDetik < 1 {
		t.Fatalf("kirim stiker bersuara salah: calls=%d stik=%d audio=%d detik=%d", c.sbCalls, len(c.sbStik), len(c.sbAudio), c.sbDetik)
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatalf("berkas tahap harusnya kehapus")
	}

	c.sbCalls = 0
	mD := &lib.Pesan{From: grup, Sender: pengirim, IsGroup: true, Type: "imageMessage", Command: "#s2",
		Quoted: &lib.Pesan{Type: "audioMessage", Sender: pengirim}}
	if err := TanganiStikerBersuara(&lib.Konteks{Conn: c, M: mD}); err != nil {
		t.Fatal(err)
	}
	if c.sbCalls != 1 {
		t.Fatalf("one-shot harusnya langsung kirim, calls=%d", c.sbCalls)
	}
}
