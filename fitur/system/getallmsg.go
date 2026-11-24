package system

import (
	"regexp"

	"gow1/lib"
)

func TanganiGetallmsg(k *lib.Konteks) error {
	m := k.M
	if m.Quoted == nil || m.Quoted.Sender == "" {
		return k.Conn.OK(m, "Reply dulu pesan usernya")
	}
	t := lib.NomorSaja(m.Quoted.Sender)
	alt := PasanganLID(t)
	cocok := func(v string) bool {
		n := lib.NomorSaja(v)
		return n != "" && (n == t || (alt != "" && n == alt))
	}
	data := []CatatanPesan{}
	for _, c := range SemuaCatatan() {
		if cocok(c.Sender) || cocok(c.LID) {
			data = append(data, c)
		}
	}
	if len(data) == 0 {
		return k.Conn.OK(m, "Belum ada catatan pesan dari user itu")
	}
	return kirimJSON(k, data, "allmsg.json")
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "getallmsg.js",
		Jenis:    "system",
		Nama:     nil,
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)getallmsg$"),
		Jalankan: TanganiGetallmsg,
	})
}
