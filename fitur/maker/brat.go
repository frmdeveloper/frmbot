package maker

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"gow1/lib"
)

type pengukurTeks struct {
	font  string
	cache map[string]float64
}

func pengukurBaru(font string) *pengukurTeks {
	return &pengukurTeks{font: font, cache: map[string]float64{}}
}

func (p *pengukurTeks) lebar(teks string, ukuran float64) float64 {
	if teks == "" {
		return 0
	}
	w100, ok := p.cache[teks]
	if !ok {
		out, err := jalanPerintah("convert", "-font", p.font, "-pointsize", "100", "label:"+teks, "-format", "%w", "info:")
		if err != nil {
			p.cache[teks] = 0
			return 0
		}
		w100, _ = strconv.ParseFloat(strings.TrimSpace(out), 64)
		p.cache[teks] = w100
	}
	return w100 * ukuran / 100
}

const naikFontRoboto = 0.9277

type coretanTeks struct {
	teks   string
	x      float64
	y      float64
	ukuran float64
	stroke float64
}

func renderKanvasTeks(baseArgs []string, font string, coretan []coretanTeks, keluarWebp string, kualitas int, ubahUkuran string) ([]byte, error) {
	tmpBase, err := lib.TulisTemp(".png", nil)
	if err != nil {
		return nil, err
	}
	os.Remove(tmpBase)
	args := append([]string{}, baseArgs...)
	args = append(args, tmpBase)
	if _, err := jalanPerintah("convert", args...); err != nil {
		return nil, err
	}
	defer os.Remove(tmpBase)
	args = []string{tmpBase, "-font", font, "-gravity", "NorthWest"}
	berStroke := len(coretan) > 0 && coretan[0].stroke > 0
	if berStroke {
		args = append(args, "-stroke", "black", "-fill", "white")
	} else {
		args = append(args, "-fill", "black")
	}
	for _, c := range coretan {
		y := c.y - naikFontRoboto*c.ukuran
		args = append(args, "-pointsize", strconv.Itoa(int(math.Round(c.ukuran))))
		if c.stroke > 0 {
			args = append(args, "-strokewidth", strconv.Itoa(int(math.Round(c.stroke))))
		}
		args = append(args, "-annotate", fmt.Sprintf("+%d+%d", int(math.Round(c.x)), int(math.Round(y))), c.teks)
	}
	tmpHasil, err := lib.TulisTemp(".png", nil)
	if err != nil {
		return nil, err
	}
	os.Remove(tmpHasil)
	if ubahUkuran != "" {
		args = append(args, "-resize", ubahUkuran)
	}
	args = append(args, tmpHasil)
	if _, err := jalanPerintah("convert", args...); err != nil {
		return nil, err
	}
	defer os.Remove(tmpHasil)
	if keluarWebp == "" {
		return os.ReadFile(tmpHasil)
	}
	tmpWebp, err := lib.TulisTemp(".webp", nil)
	if err != nil {
		return nil, err
	}
	os.Remove(tmpWebp)
	defer os.Remove(tmpWebp)
	if _, err := jalanPerintah("cwebp", "-q", strconv.Itoa(kualitas), tmpHasil, "-o", tmpWebp); err != nil {
		return nil, err
	}
	return os.ReadFile(tmpWebp)
}

func buatBrat(teks string) ([]byte, error) {
	font := "bantuan/RobotoCondensed-Regular.ttf"
	u := pengukurBaru(font)
	const skala = 4.0
	const sisi = 512 * skala
	const padding = 40 * skala
	maxLebar := sisi - padding*2
	maxTinggi := sisi - padding*2
	kataKata := strings.Split(teks, " ")
	ukuranFont := 100 * skala
	var barisBaris [][]string
	for ukuranFont > 0 {
		barisBaris = [][]string{}
		barisSekarang := []string{}
		lebarSekarang := 0.0
		for _, kata := range kataKata {
			lebarKata := u.lebar(kata+" ", ukuranFont)
			if lebarSekarang+lebarKata <= maxLebar {
				barisSekarang = append(barisSekarang, kata)
				lebarSekarang += lebarKata
			} else {
				if len(barisSekarang) > 0 {
					barisBaris = append(barisBaris, barisSekarang)
				}
				barisSekarang = []string{kata}
				lebarSekarang = lebarKata
			}
		}
		if len(barisSekarang) > 0 {
			barisBaris = append(barisBaris, barisSekarang)
		}
		totalTinggi := float64(len(barisBaris)) * (ukuranFont + 10*skala)
		if totalTinggi <= maxTinggi {
			break
		}
		ukuranFont -= 2 * skala
	}
	if ukuranFont <= 0 {
		ukuranFont = 2 * skala
	}
	tinggiBaris := ukuranFont + 10*skala
	totalTinggi := float64(len(barisBaris)) * tinggiBaris
	startY := (sisi-totalTinggi)/2 + ukuranFont/2
	coretan := []coretanTeks{}
	for i, baris := range barisBaris {
		y := startY + float64(i)*tinggiBaris
		if len(baris) == 1 {
			coretan = append(coretan, coretanTeks{teks: baris[0], x: padding, y: y, ukuran: ukuranFont})
			continue
		}
		jumlahLebar := 0.0
		for _, kata := range baris {
			jumlahLebar += u.lebar(kata, ukuranFont)
		}
		totalSpasi := maxLebar - jumlahLebar
		spasiAntar := totalSpasi / float64(len(baris)-1)
		x := padding
		for _, kata := range baris {
			coretan = append(coretan, coretanTeks{teks: kata, x: x, y: y, ukuran: ukuranFont})
			x += u.lebar(kata, ukuranFont) + spasiAntar
		}
	}
	return renderKanvasTeks(
		[]string{"-size", "2048x2048", "xc:white"},
		font, coretan, "webp", 100, "512x512")
}

func TanganiBrat(k *lib.Konteks) error {
	m := k.M
	if m.Q == "" {
		return k.Conn.OK(m, "mana teksnya?")
	}
	mentah, err := buatBrat(m.Q)
	if err != nil {
		return err
	}
	bratt, err := AddExifBawaan(mentah)
	if err != nil {
		return err
	}
	simpanStkdb(m, bratt)
	return k.Conn.SendMessage(m.From, lib.IsiPesan{Sticker: bratt}, m)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "brat.js",
		Jenis:    "maker",
		Nama:     []string{"#brat"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(brat)$"),
		Jalankan: TanganiBrat,
	})
}
