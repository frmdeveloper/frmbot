package maker

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"

	"gow1/lib"
)

const ukuranMeme = 512.0
const paddingMeme = 20.0
const fontAwalMeme = 96.0
const fontMinMeme = 18.0
const fontSatuBarisMeme = 44.0

func bagiBarisTeks(u *pengukurTeks, teks string, ukuran float64, maxLebar float64) []string {
	baris := []string{}
	sekarang := ""
	for _, kata := range polaSpasi.Split(teks, -1) {
		coba := kata
		if sekarang != "" {
			coba = sekarang + " " + kata
		}
		if sekarang == "" || u.lebar(coba, ukuran) <= maxLebar {
			sekarang = coba
		} else {
			baris = append(baris, sekarang)
			sekarang = kata
		}
	}
	if sekarang != "" {
		baris = append(baris, sekarang)
	}
	return baris
}

func ekorBaris(u *pengukurTeks, baris string, ukuran float64) float64 {
	if strings.ContainsAny(baris, "gjpqy") {
		return ukuran * 0.24
	}
	if strings.ContainsAny(baris, ",;()") {
		return ukuran * 0.06
	}
	return 0
}

func coretanMeme(u *pengukurTeks, teks string, posisi string, batasTinggi float64) ([]coretanTeks, float64) {
	maxLebar := ukuranMeme - paddingMeme*2
	ukuran := fontAwalMeme
	var baris []string
	for {
		baris = bagiBarisTeks(u, teks, ukuran, maxLebar)
		muat := float64(len(baris))*ukuran*1.05 <= batasTinggi
		if muat {
			for _, b := range baris {
				if u.lebar(b, ukuran) > maxLebar {
					muat = false
					break
				}
			}
		}
		if muat && (len(baris) == 1 || ukuran <= fontSatuBarisMeme) {
			break
		}
		if ukuran <= fontMinMeme {
			break
		}
		ukuran -= 2
	}
	jarak := ukuran * 1.05
	ekor := 0.0
	if len(baris) > 0 {
		ekor = ekorBaris(u, baris[len(baris)-1], ukuran)
	}
	stroke := math.Max(4, ukuran/8)
	hasil := []coretanTeks{}
	for i, b := range baris {
		x := (ukuranMeme - u.lebar(b, ukuran)) / 2
		var y float64
		if posisi == "atas" {
			y = paddingMeme + jarak*float64(i+1)
		} else {
			y = ukuranMeme - paddingMeme - ekor - jarak*float64(len(baris)-1-i)
		}
		hasil = append(hasil, coretanTeks{teks: b, x: x, y: y, ukuran: ukuran, stroke: stroke})
	}
	return hasil, ukuran
}

func buatMeme(buffer []byte, bawah string, atas string) ([]byte, error) {
	font := "bantuan/Roboto-Bold.ttf"
	u := pengukurBaru(font)
	masuk, err := lib.TulisTemp(lib.EkstensiMime(lib.DeteksiMime(buffer)), buffer)
	if err != nil {
		return nil, err
	}
	defer os.Remove(masuk)
	batasTinggi := ukuranMeme * 0.45
	if atas != "" && bawah != "" {
		batasTinggi = ukuranMeme * 0.34
	}
	coretan := []coretanTeks{}
	if atas != "" {
		c, _ := coretanMeme(u, atas, "atas", batasTinggi)
		coretan = append(coretan, c...)
	}
	if bawah != "" {
		c, _ := coretanMeme(u, bawah, "bawah", batasTinggi)
		coretan = append(coretan, c...)
	}
	hasil, err := renderKanvasTeks(
		[]string{masuk, "-resize", "512x512^", "-gravity", "center", "-extent", "512x512"},
		font, coretan, "webp", 92, "")
	if err != nil {
		return nil, fmt.Errorf("unsupported image type: %w", err)
	}
	return hasil, nil
}

func TanganiSmeme(k *lib.Konteks) error {
	m := k.M
	down := unduhMedia(k)
	if down == nil {
		return k.Conn.OK(m, "mana gambarnya?")
	}
	bagianMentah := strings.Split(m.Q, "|")
	bagian := make([]string, len(bagianMentah))
	for i, a := range bagianMentah {
		bagian[i] = strings.TrimSpace(a)
	}
	var atas, bawah string
	if len(bagian) > 1 {
		atas, bawah = bagian[0], bagian[1]
	} else {
		atas, bawah = "", bagian[0]
	}
	if bawah == "" && atas == "" {
		return k.Conn.OK(m, fmt.Sprintf("mana teksnya?\n\ncontoh:\n%smeme ga logis\n%smeme pas gua liat|kucing ini", m.Prefix, m.Prefix))
	}
	hasil, err := buatMeme(down, bawah, atas)
	if err != nil {
		if strings.Contains(err.Error(), "unsupported image") {
			return k.Conn.OK(m, "itu bukan foto")
		}
		return err
	}
	stiker, err := AddExifBawaan(hasil)
	if err != nil {
		return err
	}
	simpanStkdb(m, stiker)
	return k.Conn.SendMessage(m.From, lib.IsiPesan{Sticker: stiker}, m)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "smeme.js",
		Jenis:    "maker",
		Nama:     []string{"#smeme"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)((s)?meme)$"),
		Jalankan: TanganiSmeme,
	})
}
