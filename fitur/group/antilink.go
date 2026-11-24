package group

import (
	"fmt"
	"regexp"
	"strings"

	"gow1/lib"
)

var polaLinkGrup = regexp.MustCompile(`(chat\.whatsapp\.com/(?:invite/)?[0-9A-Za-z]{15,})`)

func TanganiAntilink(k *lib.Konteks) error {
	m := k.M
	if !m.IsGroup {
		return k.Conn.OK(m, "Cuma buat grup")
	}
	arg := strings.ToLower(argPertama(m))
	nyala := toggleGrupAktif("antilink", m.From)

	gc, err := k.Conn.GetGroupInfo(m.From)
	if err != nil {
		return err
	}
	if !pengirimAdmin(k, gc) && !pengirimOwner(k) {
		return k.Conn.OK(m, "Lu bukan admin")
	}
	prefix := m.Prefix
	if prefix == "" {
		prefix = "#"
	}
	switch arg {
	case "on":
		if nyala {
			return k.Conn.OK(m, "Udah nyala kok")
		}
		if !botAdmin(k, gc) {
			return k.Conn.OK(m, "Gw bukan admin, gak bisa hapus pesan")
		}
		setToggleGrup("antilink", m.From, true)
		return k.Conn.OK(m, "Antilink nyala, yang ngirim link pesannya gw hapus")
	case "off":
		if !nyala {
			return k.Conn.OK(m, "Emang udah mati")
		}
		setToggleGrup("antilink", m.From, false)
		return k.Conn.OK(m, "Antilink mati")
	}
	status := "off"
	if nyala {
		status = "on"
	}
	return k.Conn.OK(m, fmt.Sprintf("Antilink: %s\nPakai %santilink on/off", status, prefix))
}

func CekAntilink(k *lib.Konteks) (bool, error) {
	m := k.M
	if m == nil || !m.IsGroup || m.FromMe {
		return false, nil
	}
	if !toggleGrupAktif("antilink", m.From) {
		return false, nil
	}
	if !polaLinkGrup.MatchString(m.Text) {
		return false, nil
	}
	if pengirimOwner(k) {
		return false, nil
	}
	gc, err := k.Conn.GetGroupInfo(m.From)
	if err != nil {
		return false, nil
	}
	if pengirimAdmin(k, gc) {
		return false, nil
	}
	if !botAdmin(k, gc) {
		return false, nil
	}
	if err := k.Conn.RevokeMessage(m.From, m); err != nil {
		return false, nil
	}
	return true, nil
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "antilink.js",
		Jenis:    "group",
		Nama:     []string{"#antilink on/off"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(antilink)$"),
		Jalankan: TanganiAntilink,
	})
}
