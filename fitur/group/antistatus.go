package group

import (
	"fmt"
	"regexp"
	"strings"

	"gow1/lib"

	"go.mau.fi/whatsmeow/types/events"
)

func TanganiAntistatus(k *lib.Konteks) error {
	m := k.M
	if !m.IsGroup {
		return k.Conn.OK(m, "Cuma buat grup")
	}
	arg := strings.ToLower(argPertama(m))
	nyala := toggleGrupAktif("antistatus", m.From)
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
			return k.Conn.OK(m, "Gw bukan admin, gak bisa hapus pesan orang")
		}
		setToggleGrup("antistatus", m.From, true)
		return k.Conn.OK(m, "Antistatus nyala. Status yang dikirim langsung ke grup ini bakal gw hapus.")
	case "off":
		if !nyala {
			return k.Conn.OK(m, "Emang udah mati")
		}
		setToggleGrup("antistatus", m.From, false)
		return k.Conn.OK(m, "Antistatus mati")
	}
	status := "off"
	if nyala {
		status = "on"
	}
	return k.Conn.OK(m, fmt.Sprintf("Antistatus: %s\nPakai %santistatus on/off", status, prefix))
}

func CekAntistatus(k *lib.Konteks) (bool, error) {
	m := k.M
	if m == nil || !m.IsGroup || m.FromMe {
		return false, nil
	}
	if !toggleGrupAktif("antistatus", m.From) {
		return false, nil
	}
	ev, ok := m.Full.(*events.Message)
	if !ok || ev == nil || ev.Message == nil {
		return false, nil
	}
	if ev.Message.GroupStatusMessageV2 == nil && ev.Message.GroupStatusMessage == nil {
		return false, nil
	}
	gc, err := k.Conn.GetGroupInfo(m.From)
	if err != nil {
		return false, nil
	}
	if !botAdmin(k, gc) {
		fmt.Println("[antistatus] skip: bot bukan admin di grup ini")
		return false, nil
	}
	if err := k.Conn.RevokeDari(m.From, m); err != nil {
		fmt.Printf("[antistatus] hapus gagal: %v\n", err)
		return false, nil
	}
	fmt.Printf("[antistatus] status grup dari %s di %s dihapus\n", lib.NomorSaja(m.Sender), m.From)
	return true, nil
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "antistatus.js",
		Jenis:    "group",
		Nama:     []string{"#antistatus"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(antistatus)$"),
		Jalankan: TanganiAntistatus,
	})
}
