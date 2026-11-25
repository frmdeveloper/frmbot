package downloader

import (
	"strings"
	"testing"
)

func TestTeksRegionYT(t *testing.T) {
	galat := "ERROR: [youtube] QaA-33UdKsY: The uploader has not made this video available in your country\nThis video is available in Indonesia, Philippines, Thailand"
	teks, ya := teksRegionYT(galat)
	if !ya {
		t.Fatal("error region harusnya kedeteksi")
	}
	if !strings.Contains(teks, "dikunci region") || !strings.Contains(teks, "Indonesia, Philippines, Thailand") {
		t.Fatalf("teks balasan salah: %q", teks)
	}
	teks2, ya2 := teksRegionYT("ERROR: [youtube] abc: The uploader has not made this video available in your country")
	if !ya2 || !strings.Contains(teks2, "negara tertentu") {
		t.Fatalf("tanpa daftar negara harusnya fallback: %q %v", teks2, ya2)
	}
	if _, ya3 := teksRegionYT("ERROR: HTTP Error 403: Forbidden"); ya3 {
		t.Fatal("error biasa gak boleh kedeteksi sebagai region")
	}
}
