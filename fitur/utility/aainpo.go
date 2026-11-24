package utility

import (
	"regexp"
	"strings"

	"gow1/lib"
)

func TanganiAainpo(k *lib.Konteks) error {
	m := k.M
	c := ""
	if len(m.Command) > 1 {
		c = m.Command[1:]
	}
	balas := func(teks string) error {
		return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: teks}, m)
	}
	switch {
	case strings.Contains(c, "sewa"):
		return balas("Culik aja terus chat wa.me/frmchan biar di acc")
	case strings.Contains(c, "script") || strings.Contains(c, "sc"):
		return balas("Wah ngelunjak")
	case strings.Contains(c, "donasi"):
		return balas("Gausah sih, chat wa.me/frmchan coba")
	case strings.Contains(c, "owner"):
		return balas("wa.me/rifkimekanik\nwww.tiktok.com/@rifki_mekanik")
	}
	return nil
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "aainpo.js",
		Jenis:    "utility",
		Nama:     nil,
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(sewa|owner|script|sc|donasi)$"),
		Jalankan: TanganiAainpo,
	})
}
