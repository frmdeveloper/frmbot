package group

import (
	"regexp"
	"strings"

	"gow1/lib"
)

func cekAdminBanuser(k *lib.Konteks) bool {
	m := k.M
	if !m.IsGroup {
		return pengirimOwner(k)
	}
	if pengirimOwner(k) {
		return true
	}
	gc, err := k.Conn.GetGroupInfo(m.From)
	if err != nil {
		return false
	}
	return pengirimAdmin(k, gc)
}

func TanganiBanuser(k *lib.Konteks) error {
	m := k.M
	aksi := "ban"
	if strings.Contains(m.Command, "unbanuser") {
		aksi = "unban"
	} else if strings.Contains(m.Command, "listban") {
		aksi = "list"
	}
	if k.User.Banned == nil {
		k.User.Banned = []string{}
	}

	if aksi == "list" {
		if len(k.User.Banned) == 0 {
			return k.Conn.OK(m, "Belum ada yang dibanned")
		}
		baris := make([]string, len(k.User.Banned))
		mentions := make([]string, len(k.User.Banned))
		for i, n := range k.User.Banned {
			baris[i] = "- @" + n
			mentions[i] = n + "@s.whatsapp.net"
		}
		teks := "List dibanned:\n" + strings.Join(baris, "\n")
		return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: teks, Mentions: mentions}, m)
	}

	if !cekAdminBanuser(k) {
		if m.IsGroup {
			return k.Conn.OK(m, "Cuma admin sama owner")
		}
		return k.Conn.OK(m, "Cuma owner")
	}

	var target []string
	if m.Quoted != nil && m.Quoted.Sender != "" {
		target = []string{m.Quoted.Sender}
	} else if len(m.Mentions) > 0 {
		target = m.Mentions
	}
	if len(target) == 0 {
		return k.Conn.OK(m, "Reply atau tag orangnya")
	}

	hasil := []string{}
	for _, t := range target {
		num := lib.NomorSaja(t)
		if num == "" || num == k.Conn.BotNomor() || adalahOwner(k, t) {
			continue
		}
		if aksi == "ban" {
			ada := false
			for _, x := range k.User.Banned {
				if x == num {
					ada = true
					break
				}
			}
			if !ada {
				k.User.Banned = append(k.User.Banned, num)
			}
		} else {
			sisa := []string{}
			for _, x := range k.User.Banned {
				if x != num {
					sisa = append(sisa, x)
				}
			}
			k.User.Banned = sisa
		}
		hasil = append(hasil, num)
	}

	if len(hasil) == 0 {
		if aksi == "ban" {
			return k.Conn.OK(m, "Gak ada yang bisa dibanned")
		}
		return k.Conn.OK(m, "Dia gak dibanned kok")
	}
	_ = k.User.Simpan()
	baris := make([]string, len(hasil))
	mentions := make([]string, len(hasil))
	for i, n := range hasil {
		if aksi == "ban" {
			baris[i] = "@" + n + " keban, semua fitur + AI mati buat dia"
		} else {
			baris[i] = "@" + n + " dilepas dari ban"
		}
		mentions[i] = n + "@s.whatsapp.net"
	}
	return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: strings.Join(baris, "\n"), Mentions: mentions}, m)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "banuser.js",
		Jenis:    "group",
		Nama:     []string{"#banuser (reply/tag)", "#unbanuser (reply/tag)", "#listban"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(unbanuser|banuser|listban)$"),
		Jalankan: TanganiBanuser,
	})
}
