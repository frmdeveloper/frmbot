package group

import (
	"regexp"

	"gow1/lib"
)

func TanganiHidetag(k *lib.Konteks) error {
	m := k.M
	gc, err := k.Conn.GetGroupInfo(m.From)
	if err != nil {
		return err
	}
	if !pengirimAdmin(k, gc) {
		return k.Conn.OK(m, "lu bukan admin")
	}
	mem := make([]string, 0, len(gc.Peserta))
	for _, p := range gc.Peserta {
		jid := p.JID
		if jid == "" && p.Nomor != "" {
			jid = p.Nomor + "@s.whatsapp.net"
		}
		if jid != "" {
			mem = append(mem, jid)
		}
	}
	return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: m.Q, Mentions: mem}, m)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "hidetag.js",
		Jenis:    "group",
		Nama:     nil,
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(hidetag|ht)$"),
		Jalankan: TanganiHidetag,
	})
}
