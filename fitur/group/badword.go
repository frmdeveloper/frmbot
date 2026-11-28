package group

import (
	"fmt"
	"regexp"
	"strings"

	"gow1/lib"
)

var kataKasar = []string{
	"anjing", "bangsat", "bajingan", "kontol", "memek", "ngentot",
	"jancok", "jancuk", "kampret", "keparat", "brengsek", "sialan",
	"goblok", "tolol",
	"fuck", "fucking", "shit", "bitch", "asshole", "bastard", "dick", "pussy",
}

var polaKataValid = regexp.MustCompile(`^[a-z]{2,}$`)

func polaBadword(grup string) *regexp.Regexp {
	semua := append(append([]string{}, kataKasar...), kataBadwordGrup(grup)...)
	bagian := make([]string, 0, len(semua))
	for _, w := range semua {
		huruf := make([]string, 0, len(w))
		for _, c := range w {
			huruf = append(huruf, string(c)+"+")
		}
		bagian = append(bagian, strings.Join(huruf, ""))
	}
	return regexp.MustCompile(`(?i)\b(` + strings.Join(bagian, "|") + `)\b`)
}

func perintahBadword(m *lib.Pesan) string {
	return strings.ToLower(strings.TrimLeft(m.Command, "#!/."))
}

func TanganiBadword(k *lib.Konteks) error {
	m := k.M
	if !m.IsGroup {
		return k.Conn.OK(m, "Cuma buat grup")
	}
	cmd := perintahBadword(m)
	arg := strings.ToLower(argPertama(m))
	mode := modeGrup("badword", m.From)
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
	if cmd == "listbadword" {
		tambahan := kataBadwordGrup(m.From)
		isi := "-"
		if len(tambahan) > 0 {
			isi = strings.Join(tambahan, ", ")
		}
		return k.Conn.OK(m, fmt.Sprintf("Kata kasar bawaan: %s\n\nTambahan grup ini: %s", strings.Join(kataKasar, ", "), isi))
	}
	if cmd == "addbadword" {
		mentah := strings.FieldsFunc(strings.ToLower(m.Q), func(r rune) bool { return r == ',' || r == ' ' || r == '\t' || r == '\n' })
		valid := []string{}
		udah := map[string]bool{}
		for _, w := range mentah {
			if polaKataValid.MatchString(w) && !udah[w] {
				udah[w] = true
				valid = append(valid, w)
			}
		}
		if len(valid) == 0 {
			return k.Conn.OK(m, fmt.Sprintf("Pakai %saddbadword <kata> — bisa banyak kata sekalian, pisahin pakai spasi/koma", prefix))
		}
		daftar := kataBadwordGrup(m.From)
		ada := map[string]bool{}
		for _, w := range daftar {
			ada[w] = true
		}
		for _, w := range kataKasar {
			ada[w] = true
		}
		masuk, udahAda := []string{}, []string{}
		for _, w := range valid {
			if ada[w] {
				udahAda = append(udahAda, w)
				continue
			}
			if len(daftar) >= 100 {
				break
			}
			daftar = append(daftar, w)
			masuk = append(masuk, w)
		}
		if len(masuk) > 0 {
			simpanKataBadwordGrup(m.From, daftar)
		}
		bagian := []string{}
		if len(masuk) > 0 {
			bagian = append(bagian, "Ditambahin buat grup ini: "+strings.Join(masuk, ", "))
		}
		if len(udahAda) > 0 {
			bagian = append(bagian, "Udah ada: "+strings.Join(udahAda, ", "))
		}
		if len(bagian) == 0 {
			bagian = append(bagian, "Gak ada kata baru yang masuk (daftar tambahan grup ini udah penuh, maks 100)")
		}
		return k.Conn.OK(m, strings.Join(bagian, "\n"))
	}
	switch arg {
	case "on", "kick":
		if mode == arg {
			return k.Conn.OK(m, "Udah nyala kok")
		}
		if !botAdmin(k, gc) {
			return k.Conn.OK(m, "Gw bukan admin, gak bisa hapus pesan orang")
		}
		setModeGrup("badword", m.From, arg)
		if arg == "kick" {
			return k.Conn.OK(m, "Badword nyala (mode kick). Yang ngomong kasar pesannya gw hapus + orangnya langsung gw kick.")
		}
		return k.Conn.OK(m, "Badword nyala (mode hapus). Yang ngomong kasar pesannya langsung gw hapus, gak di-kick.")
	case "off":
		if mode == "" {
			return k.Conn.OK(m, "Emang udah mati")
		}
		setModeGrup("badword", m.From, "")
		return k.Conn.OK(m, "Badword mati")
	}
	status := mode
	if status == "" {
		status = "off"
	}
	return k.Conn.OK(m, fmt.Sprintf("Badword: %s\nPakai %s%s on/kick/off", status, prefix, cmd))
}

func CekBadword(k *lib.Konteks) (bool, error) {
	m := k.M
	if m == nil || !m.IsGroup || m.FromMe {
		return false, nil
	}
	mode := modeGrup("badword", m.From)
	if mode == "" {
		return false, nil
	}
	if m.Text == "" || !polaBadword(m.From).MatchString(m.Text) {
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
		fmt.Println("[badword] skip: bot bukan admin di grup ini")
		return false, nil
	}
	num := lib.NomorSaja(m.Sender)
	if err := k.Conn.RevokeDari(m.From, m); err != nil {
		fmt.Printf("[badword] hapus gagal: %v\n", err)
		return false, nil
	}
	fmt.Printf("[badword] pesan %s di %s dihapus (mode %s)\n", num, m.From, mode)
	if mode == "kick" {
		if err := k.Conn.UpdateGroupParticipants(m.From, []string{m.Sender}, "remove"); err != nil {
			fmt.Printf("[badword] kick gagal: %v\n", err)
		} else {
			fmt.Printf("[badword] %s di-kick dari %s\n", num, m.From)
		}
	}
	return true, nil
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "badword.js",
		Jenis:    "group",
		Nama:     []string{"#antibadword on/kick/off", "#listbadword", "#addbadword"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(badword|antibadword|listbadword|addbadword)$"),
		Jalankan: TanganiBadword,
	})
}
