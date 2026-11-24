package maker

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"gow1/lib"
)

func TanganiReadqr(k *lib.Konteks) error {
	m := k.M
	down := unduhMedia(k)
	if down == nil {
		return k.Conn.OK(m, "Kirim atau reply gambar yang ada QR code-nya")
	}
	mime := lib.DeteksiMime(down)
	if !strings.Contains(mime, "image") {
		return k.Conn.OK(m, "Kirim gambar yang berisi QR code")
	}
	path, err := lib.TulisTemp(lib.EkstensiMime(mime), down)
	if err != nil {
		return k.Conn.OK(m, fmt.Sprintf("Error: %s", err.Error()))
	}
	defer os.Remove(path)
	out, err := jalanPerintah("zbarimg", "--quiet", "--raw", path)
	if err != nil {
		return k.Conn.OK(m, fmt.Sprintf("Error: %s", err.Error()))
	}
	hasil := strings.TrimRight(out, "\n")
	if strings.HasPrefix(hasil, "QR-Code:") {
		hasil = strings.TrimPrefix(hasil, "QR-Code:")
	}
	if hasil == "" {
		return k.Conn.OK(m, "QR code tidak ditemukan dalam gambar")
	}
	return k.Conn.OK(m, fmt.Sprintf("*QR Code terdeteksi:*\n\n%s", hasil))
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "readqr.js",
		Jenis:    "maker",
		Nama:     []string{"#readqr"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(readqr|qrread|scanqr|qrscan)$"),
		Jalankan: TanganiReadqr,
	})
}
