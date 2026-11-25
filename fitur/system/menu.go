package system

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"gow1/lib"
)

func ramTerpakaiMB() int {
	b, err := os.ReadFile("/proc/self/status")
	if err != nil {
		return 0
	}
	for _, baris := range strings.Split(string(b), "\n") {
		if strings.HasPrefix(baris, "VmRSS:") {
			var kb int
			_, _ = fmt.Sscanf(baris, "VmRSS: %d kB", &kb)
			return kb / 1024
		}
	}
	return 0
}

func TanganiMenu(k *lib.Konteks) error {
	m := k.M
	if strings.Contains(m.Command, "mani") {
		_ = k.Conn.OK(m, "Mani loh ya")
	}

	semuaFitur := lib.Semua()
	perFile := []string{}
	for file := range semuaFitur {
		perFile = append(perFile, file)
	}
	sort.Strings(perFile)

	byJenis := map[string][]string{}
	jenisList := []string{}
	for _, file := range perFile {
		f := semuaFitur[file]
		if len(f.Nama) == 0 || f.Jenis == "" {
			continue
		}
		if _, ada := byJenis[f.Jenis]; !ada {
			jenisList = append(jenisList, f.Jenis)
		}
		byJenis[f.Jenis] = append(byJenis[f.Jenis], f.Nama...)
	}
	sort.Strings(jenisList)

	mau := strings.ToLower(strings.TrimSpace(m.Q))
	if mau == "grup" {
		mau = "group"
	}
	ketemu := ""
	for _, j := range jenisList {
		if strings.ToLower(j) == mau {
			ketemu = j
		}
	}
	if mau != "" && ketemu == "" {
		return k.Conn.OK(m, m.Q)
	}
	jenisPakai := []string{}
	if ketemu != "" {
		jenisPakai = []string{ketemu}
	} else {
		for _, j := range jenisList {
			if strings.ToLower(j) != "group" {
				jenisPakai = append(jenisPakai, j)
			}
		}
	}

	prefix := m.Prefix
	if prefix == "" {
		prefix = "#"
	}
	blok := []string{}
	for _, j := range jenisPakai {
		var sb strings.Builder
		sb.WriteString("╭─「 *" + j + "* 」")
		for _, n := range byJenis[j] {
			sb.WriteString("\n│  ︎  " + strings.ReplaceAll(n, "#", prefix))
		}
		blok = append(blok, sb.String())
	}
	isi := strings.Join(blok, "\n╰──⊷\n\n")

	teks := "✦ *frm chan* ✦\n" +
		"lib: whatsmeow • golang\n" +
		"github.com/frmdeveloper/frmbot\n\n" +
		"📌 Jangan lupa follow\n" +
		"    tiktok.com/@rifki_mekanik\n\n" +
		"⚙️ RAM: " + fmt.Sprint(ramTerpakaiMB()) + " MB terpakai\n\n" +
		"🪡 daftar perintah 🪡\n" +
		strings.Repeat("\u200e", 8000) + "\n" +
		isi + "\n╰──⊷"
	return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: teks}, m)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "menu.js",
		Jenis:    "system",
		Nama:     nil,
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(help|(all)?m[e\\w]n[u\\w](all)?)$"),
		Jalankan: TanganiMenu,
	})
}
