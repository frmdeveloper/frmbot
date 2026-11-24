package group

import (
	"fmt"
	"regexp"
	"strings"

	"gow1/lib"
)

func TanganiChatai(k *lib.Konteks) error {
	m := k.M
	if !m.IsGroup {
		return k.Conn.OK(m, "Fitur ini cuma buat grup")
	}
	arg := strings.ToLower(argPertama(m))
	if k.Group.Chatai == nil {
		k.Group.Chatai = map[string]any{}
	}
	nyala := true
	if v, ada := k.Group.Chatai[m.From]; ada {
		if f, ok := v.(bool); ok && !f {
			nyala = false
		}
	}
	prefix := m.Prefix
	if prefix == "" {
		prefix = "#"
	}
	switch arg {
	case "on":
		if nyala {
			return k.Conn.OK(m, "Udah nyala dari tadi")
		}
		delete(k.Group.Chatai, m.From)
		_ = k.Group.Simpan()
		return k.Conn.OK(m, "Oke, chat AI nyala lagi di grup ini")
	case "off":
		if !nyala {
			return k.Conn.OK(m, "Emang udah mati")
		}
		k.Group.Chatai[m.From] = false
		_ = k.Group.Simpan()
		return k.Conn.OK(m, "Yaudah, chat AI gue matiin di grup ini")
	}
	status := "off"
	if nyala {
		status = "on"
	}
	return k.Conn.OK(m, fmt.Sprintf("Chat AI: %s\nPakai %schatai on/off", status, prefix))
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "chatai.js",
		Jenis:    "group",
		Nama:     []string{"#chatai on/off"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)chatai$"),
		Jalankan: TanganiChatai,
	})
}
