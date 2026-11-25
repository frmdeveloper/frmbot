package ai

import (
	"regexp"
	"strings"

	"gow1/lib"
)

func TanganiAi(k *lib.Konteks) error {
	m := k.M
	prompt := strings.TrimSpace(m.Q)
	if prompt == "" && m.Quoted != nil {
		prompt = strings.TrimSpace(m.Quoted.Text)
	}
	if prompt == "" {
		return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: "mau nanya apa jir"}, m)
	}
	return Ngobrol(k, prompt)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "ai.js",
		Jenis:    "ai",
		Nama:     nil,
		Itunya:   regexp.MustCompile("(?i)^((#|\\!|\\/|\\.)(ai|gpt|p))$"),
		Jalankan: TanganiAi,
	})
}
