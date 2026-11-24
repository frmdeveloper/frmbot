package maker

import (
	"os"
	"os/exec"
	"testing"
)

func TestVideoKotakMu900(t *testing.T) {
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		t.Skip("ffmpeg gak ada")
	}
	vid := t.TempDir() + "/noise.mp4"
	err := exec.Command("ffmpeg", "-y", "-v", "error", "-f", "lavfi",
		"-i", "nullsrc=size=640x640:duration=25:rate=30",
		"-vf", "geq=r='random(1)*255':g='random(1)*255':b='random(1)*255'",
		"-c:v", "libx264", "-pix_fmt", "yuv420p", vid).Run()
	if err != nil {
		t.Skip("bikin video noise gagal")
	}
	buf, err := os.ReadFile(vid)
	if err != nil {
		t.Fatal(err)
	}
	hasil, err := video2webpKotakSb(buf)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("hasil: %d byte", len(hasil))
	if len(hasil) > 900*1024 {
		t.Fatalf("hasil %d byte masih di atas 900KB", len(hasil))
	}
}
