package group

import (
	"regexp"
	"strings"

	"gow1/lib"
)

func TanganiBot(k *lib.Konteks) error {
	m := k.M
	if !m.IsGroup {
		return k.Conn.OK(m, "Fitur ini cuma buat grup")
	}
	arg := strings.ToLower(argPertama(m))
	nyala := toggleGrupAktif("bot", m.From)
	switch arg {
	case "on":
		if nyala {
			return k.Conn.OK(m, "Bot udah nyala kok")
		}
		setToggleGrup("bot", m.From, true)
		return k.Conn.OK(m, "Bot nyala, semua fitur bisa dipakai di grup ini")
	case "off":
		if !nyala {
			return k.Conn.OK(m, "Emang udah mati")
		}
		setToggleGrup("bot", m.From, false)
		return k.Conn.OK(m, "Bot mati, cuma AI doang yang jawab")
	}
	return nil
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "bot.js",
		Jenis:    "group",
		Nama:     []string{"#bot on/off"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(bot)$"),
		Jalankan: TanganiBot,
	})
}
