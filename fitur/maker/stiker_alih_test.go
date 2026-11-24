package maker

import (
	"encoding/base64"
	"os"
	"os/exec"
	"testing"

	"gow1/lib"
)

func TestSPersisVideoDialihkan(t *testing.T) {
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		t.Skip("ffmpeg gak ada")
	}
	vidPath := t.TempDir() + "/v.mp4"
	err := exec.Command("ffmpeg", "-y", "-v", "error",
		"-f", "lavfi", "-i", "testsrc=duration=1:size=128x128:rate=10",
		"-f", "lavfi", "-i", "sine=frequency=440:duration=1",
		"-c:v", "libx264", "-pix_fmt", "yuv420p", "-shortest", vidPath).Run()
	if err != nil {
		t.Skip("bikin video tes gagal")
	}
	vid, err := os.ReadFile(vidPath)
	if err != nil {
		t.Fatal(err)
	}
	png, _ := base64.StdEncoding.DecodeString("iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z8BQDwAEhQGAhKmMIQAAAABJRU5ErkJggg==")
	pengirim := "628111111111@s.whatsapp.net"
	grup := "120363000000000000@g.us"

	c := &connSb{vid: vid, png: png}
	mS := &lib.Pesan{From: grup, Sender: pengirim, IsGroup: true, Type: "videoMessage", Command: "#s"}
	if err := TanganiStiker(&lib.Konteks{Conn: c, M: mS}); err != nil {
		t.Fatal(err)
	}
	if c.sbCalls != 1 {
		t.Fatalf("#s + video harusnya dialihkan ke stiker bersuara, sbCalls=%d", c.sbCalls)
	}

	c2 := &connSb{vid: vid, png: png}
	mStiker := &lib.Pesan{From: grup, Sender: pengirim, IsGroup: true, Type: "videoMessage", Command: "#stiker"}
	if err := TanganiStiker(&lib.Konteks{Conn: c2, M: mStiker}); err != nil {
		t.Fatal(err)
	}
	if c2.sbCalls != 0 {
		t.Fatalf("#stiker + video gak boleh dialihkan")
	}
	adaStikerBiasa := false
	for _, isi := range c2.terkirim {
		if len(isi.Sticker) > 0 {
			adaStikerBiasa = true
		}
	}
	if !adaStikerBiasa {
		t.Fatalf("#stiker + video harusnya kirim stiker animasi biasa")
	}

	c3 := &connSb{vid: vid, png: png}
	mSGambar := &lib.Pesan{From: grup, Sender: pengirim, IsGroup: true, Type: "imageMessage", Command: "#s"}
	if err := TanganiStiker(&lib.Konteks{Conn: c3, M: mSGambar}); err != nil {
		t.Fatal(err)
	}
	if c3.sbCalls != 0 {
		t.Fatalf("#s + gambar gak boleh dialihkan")
	}
}
