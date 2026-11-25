package maker

import (
	"bytes"
	"image"
	_ "image/png"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"gow1/lib"
)

func TestWebpKePngHD(t *testing.T) {
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		t.Skip("ffmpeg gak ada")
	}
	dir := t.TempDir()
	anim := filepath.Join(dir, "anim.webp")
	if err := exec.Command("ffmpeg", "-y", "-v", "error", "-f", "lavfi", "-i", "testsrc=duration=1:size=64x64:rate=10", "-vcodec", "libwebp", "-loop", "0", "-an", anim).Run(); err != nil {
		t.Fatalf("bikin webp animasi: %v", err)
	}
	mentah, err := os.ReadFile(anim)
	if err != nil {
		t.Fatal(err)
	}
	if mime := lib.DeteksiMime(mentah); mime != "image/webp" {
		t.Fatalf("mime sumber %q, mau image/webp", mime)
	}
	png, err := webpKePngHD(mentah)
	if err != nil {
		t.Fatalf("webpKePngHD: %v", err)
	}
	if mime := lib.DeteksiMime(png); mime != "image/png" {
		t.Fatalf("mime hasil %q, mau image/png", mime)
	}
	cfg, _, err := image.DecodeConfig(bytes.NewReader(png))
	if err != nil || cfg.Width != 64 || cfg.Height != 64 {
		t.Fatalf("DecodeConfig hasil: %+v, %v", cfg, err)
	}
	if _, err := webpKePngHD([]byte("ini bukan gambar sama sekali")); err == nil {
		t.Fatal("berkas rusak harusnya error")
	}
}
