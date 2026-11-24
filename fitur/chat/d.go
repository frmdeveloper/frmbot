package chat

import (
	"regexp"

	"gow1/lib"
)

func TanganiD(k *lib.Konteks) error {
	m := k.M
	if m.Quoted == nil || m.Quoted.ID == "" {
		return k.Conn.OK(m, "Reply pesan yang mau didelete")
	}

	quotedDariBot := lib.NomorSaja(m.Quoted.Sender) == k.Conn.BotNomor() ||
		(m.Quoted.Sender != "" && m.Quoted.Sender == k.Conn.BotLID())
	if quotedDariBot {
		if err := k.Conn.RevokeMessage(m.From, m.Quoted); err != nil {
			return k.Conn.OK(m, "Gagal delete: "+err.Error())
		}
		return nil
	}

	if !m.IsGroup {
		return k.Conn.OK(m, "Di chat pribadi cuma bisa delete pesan bot sendiri")
	}

	gc, err := k.Conn.GetGroupInfo(m.From)
	if err != nil {
		return k.Conn.OK(m, "Gagal delete: "+err.Error())
	}
	cari := func(jid string) *lib.Peserta {
		num := lib.NomorSaja(jid)
		if num == "" {
			return nil
		}
		for i := range gc.Peserta {
			p := &gc.Peserta[i]
			for _, x := range []string{p.Nomor, p.JID, p.LID} {
				if lib.NomorSaja(x) == num {
					return p
				}
			}
		}
		return nil
	}
	isAdmin := func(p *lib.Peserta) bool {
		return p != nil && (p.IsAdmin || p.IsSuperAdmin)
	}
	if !isAdmin(cari(m.Sender)) && !isAdmin(cari(m.SenderAlt)) {
		return k.Conn.OK(m, "Gue bukan admin di grup ini")
	}

	if err := k.Conn.RevokeMessage(m.From, m.Quoted); err != nil {
		return k.Conn.OK(m, "Gagal delete: "+err.Error())
	}
	return nil
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "d.js",
		Jenis:    "chat",
		Nama:     []string{"#d"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)d$"),
		Jalankan: TanganiD,
	})
}
