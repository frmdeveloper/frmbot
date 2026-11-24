package info

import (
	"regexp"
	"strings"

	"gow1/lib"
)

func TanganiTranslate(k *lib.Konteks) error {
	if k.M.Q == "" {
		return nil
	}
	isi := strings.Split(k.M.Q, " ")
	teks := strings.Join(isi[1:], " ")
	hasil, err := lib.Gtr(isi[0], teks)
	if err != nil {
		return err
	}
	return k.Conn.SendMessage(k.M.From, lib.IsiPesan{Text: hasil}, k.M)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "translate.js",
		Jenis:    "info",
		Nama:     []string{"#translate"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(tr|translate)$"),
		Jalankan: TanganiTranslate,
	})
}
