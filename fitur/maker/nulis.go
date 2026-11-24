package maker

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gow1/lib"
)

var polaBungkusNulis = regexp.MustCompile(`([^\s]+[ \t]*){1,9}`)

func TanganiNulis(k *lib.Konteks) error {
	m := k.M
	teks := m.Q
	if teks == "" && m.Quoted != nil {
		teks = m.Quoted.Text
	}
	if teks == "" {
		return k.Conn.OK(m, "Gaada teksnya tuh")
	}
	baris := strings.Split(teks, "\n")
	for i, b := range baris {
		if b != "" {
			baris[i] = strings.TrimSpace(polaBungkusNulis.ReplaceAllString(b, "$&\n"))
		}
	}
	splitText := strings.Join(baris, "\n")
	potong := strings.Split(splitText, "\n")
	if len(potong) > 31 {
		potong = potong[:31]
	}
	fixHeight := strings.Join(potong, "\n")
	ran := namaAcak()
	keluar := filepath.Join("sampah", ran+".jpg")
	if err := os.MkdirAll("sampah", 0o755); err != nil {
		return nil
	}
	_, _ = jalanPerintah("convert",
		"./bantuan/sebelumkiri.jpg",
		"-font", "./bantuan/Indie-Flower.ttf",
		"-size", "960x1280",
		"-pointsize", "22",
		"-interline-spacing", "2",
		"-annotate", "+140+153",
		fixHeight,
		"./"+keluar)
	data, err := os.ReadFile(keluar)
	if err != nil {
		return nil
	}
	return k.Conn.SendMessage(m.From, lib.IsiPesan{Image: data}, m)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "nulis.js",
		Jenis:    "maker",
		Nama:     nil,
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(nulis)$"),
		Jalankan: TanganiNulis,
	})
}
