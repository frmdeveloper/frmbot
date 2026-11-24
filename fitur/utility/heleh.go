package utility

import (
	"regexp"
	"strings"

	"gow1/lib"
)

var (
	polaVokalKecil = regexp.MustCompile(`[aiueo]`)
	polaVokalBesar = regexp.MustCompile(`[AIUEO]`)
)

func TanganiHeleh(k *lib.Konteks) error {
	m := k.M
	r := []rune(m.Command)
	if len(r) < 3 {
		return nil
	}
	c := string(r[2])
	teks := m.Q
	if m.Quoted != nil && m.Quoted.Text != "" {
		teks = m.Quoted.Text
	}
	if teks == "" {
		return nil
	}
	hasil := polaVokalKecil.ReplaceAllString(teks, c)
	hasil = polaVokalBesar.ReplaceAllString(hasil, strings.ToUpper(c))
	return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: hasil}, m)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "heleh.js",
		Jenis:    "utility",
		Nama:     nil,
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(h[aiueo]l[aiueo]h)$"),
		Jalankan: TanganiHeleh,
	})
}
