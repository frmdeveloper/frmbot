package group

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"gow1/lib"
)

func keJpegWallpaper(buf []byte) ([]byte, error) {
	dir, err := os.MkdirTemp("", "wp-")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(dir)
	masuk := filepath.Join(dir, "masuk.bin")
	keluar := filepath.Join(dir, "keluar.jpg")
	if err := os.WriteFile(masuk, buf, 0o644); err != nil {
		return nil, err
	}
	cmd := exec.Command("ffmpeg", "-y", "-v", "error", "-i", masuk, "-vf", "scale='min(1440,iw)':'min(1440,ih)':force_original_aspect_ratio=decrease", "-frames:v", "1", "-q:v", "3", keluar)
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	return os.ReadFile(keluar)
}

func TanganiSetwallpaper(k *lib.Konteks) error {
	m := k.M
	if !m.IsGroup {
		return k.Conn.OK(m, "ini cuma buat grup jir")
	}
	gc, err := k.Conn.GetGroupInfo(m.From)
	if err != nil {
		return err
	}
	if !pengirimOwner(k) && !pengirimAdmin(k, gc) {
		return k.Conn.OK(m, "lu bukan admin jir")
	}
	if !botAdmin(k, gc) {
		return k.Conn.OK(m, "gw bukan admin di grup ini — jadiin admin dulu biar bisa ganti wallpaper buat semua orang")
	}
	var sumber *lib.Pesan
	if m.Type == "imageMessage" {
		sumber = m
	} else if m.Quoted != nil && m.Quoted.Type == "imageMessage" {
		sumber = m.Quoted
	}
	if sumber == nil {
		return k.Conn.OK(m, "reply foto yang mau dijadiin wallpaper, atau kirim fotonya pake caption #setwallpaper")
	}
	buf, _, err := k.Conn.DownloadAny(sumber)
	if err != nil || len(buf) == 0 {
		return k.Conn.OK(m, "fotonya gak bisa diunduh jir, coba kirim ulang")
	}
	jpg, err := keJpegWallpaper(buf)
	if err != nil {
		return k.Conn.OK(m, "gambarnya gak bisa dibaca jir, coba kirim ulang gambarnya")
	}
	if err := k.Conn.SetGroupWallpaper(m.From, jpg); err != nil {
		return k.Conn.OK(m, fmt.Sprintf("gagal ganti wallpapernya jir: %v", err))
	}
	simpanWallpaperGrup(m.From, jpg)
	return k.Conn.OK(m, "wallpaper grup udah diganti — semua orang di grup ini bakal liat wallpaper baru ini")
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "setwallpaper.js",
		Jenis:    "group",
		Nama:     []string{"#setwallpaper"},
		Itunya:   regexp.MustCompile(`(?i)^(#|!|/|\.)(set|ganti)?wallpaper$`),
		Jalankan: TanganiSetwallpaper,
	})
}

var polaNamaBerkasWallpaper = regexp.MustCompile(`[^0-9a-zA-Z.-]`)

func berkasWallpaperGrup(grup string) string {
	return filepath.Join("data", "wallpaper", polaNamaBerkasWallpaper.ReplaceAllString(grup, "_")+".jpg")
}

func simpanWallpaperGrup(grup string, jpg []byte) {
	if err := os.MkdirAll(filepath.Dir(berkasWallpaperGrup(grup)), 0o755); err != nil {
		return
	}
	_ = os.WriteFile(berkasWallpaperGrup(grup), jpg, 0o644)
}

func hapusWallpaperGrup(grup string) {
	_ = os.Remove(berkasWallpaperGrup(grup))
}
