package group

import (
	"regexp"

	"gow1/lib"
)

func TanganiGetpp(k *lib.Konteks) error {
	m := k.M
	target := ""
	if len(m.Mentions) > 0 {
		target = m.Mentions[0]
	} else if m.Quoted != nil {
		target = m.Quoted.Sender
	}
	url, err := k.Conn.GetProfilePicture(target)
	if err != nil || url == "" {
		return k.Conn.OK(m, "Gabisa woilah")
	}
	data, err := lib.AmbilURL(url, nil)
	if err != nil || len(data) == 0 {
		return k.Conn.OK(m, "Gabisa woilah")
	}
	return k.Conn.SendMessage(m.From, lib.IsiPesan{Image: data}, m)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "getpp.js",
		Jenis:    "group",
		Nama:     []string{"#getpp"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(getpp)$"),
		Jalankan: TanganiGetpp,
	})
}
