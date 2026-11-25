package system

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"gow1/lib"
)

func TanganiMenu(k *lib.Konteks) error {
	m := k.M
	prefix := m.Prefix
	if prefix == "" {
		prefix = "."
	}
	q := strings.ToLower(strings.TrimSpace(m.Q))

	semuaFitur := lib.Semua()
	perFile := make([]string, 0, len(semuaFitur))
	for file := range semuaFitur {
		perFile = append(perFile, file)
	}
	sort.Strings(perFile)

	byJenis := map[string][]string{}
	for _, file := range perFile {
		f := semuaFitur[file]
		if len(f.Nama) == 0 || f.Jenis == "" {
			continue
		}
		j := strings.ToLower(f.Jenis)
		for _, nm := range f.Nama {
			cleanName := nm
			if strings.HasPrefix(cleanName, "#") || strings.HasPrefix(cleanName, ".") || strings.HasPrefix(cleanName, "!") || strings.HasPrefix(cleanName, "/") {
				cleanName = cleanName[1:]
			}
			byJenis[j] = append(byJenis[j], prefix+cleanName)
		}
	}

	jenisList := make([]string, 0, len(byJenis))
	for j := range byJenis {
		jenisList = append(jenisList, j)
	}
	sort.Strings(jenisList)

	if q == "" {
		var sb strings.Builder
		sb.WriteString("╭─「 *daftar kategori* 」\n")
		for _, j := range jenisList {
			sb.WriteString(fmt.Sprintf("│    %smenu %s\n", prefix, j))
		}
		sb.WriteString(fmt.Sprintf("│    %smenu all\n", prefix))
		sb.WriteString("╰──⊷\n\n")
		sb.WriteString(fmt.Sprintf("Ketik %smenu nama kategori untuk melihat daftar perintah\nContoh %smenu audio", prefix, prefix))
		return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: sb.String()}, m)
	}

	targetJenis := []string{}
	if q == "all" {
		targetJenis = jenisList
	} else {
		for _, j := range jenisList {
			if strings.EqualFold(j, q) {
				targetJenis = append(targetJenis, j)
			}
		}
		if len(targetJenis) == 0 {
			pesan := fmt.Sprintf("Kategori %s tidak ditemukan\nKetik %smenu untuk melihat daftar kategori yang tersedia", q, prefix)
			return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: pesan}, m)
		}
	}

	blok := []string{}
	for _, j := range targetJenis {
		var sb strings.Builder
		sb.WriteString("╭─「 *" + j + "* 」")
		if strings.EqualFold(j, "group") && !m.IsGroup {
			sb.WriteString("\n│    Khusus grup")
			sb.WriteString("\n╰──⊷")
			blok = append(blok, sb.String())
			continue
		}
		for _, cmd := range byJenis[j] {
			sb.WriteString("\n│    " + cmd)
		}
		sb.WriteString("\n╰──⊷")
		blok = append(blok, sb.String())
	}

	hasil := strings.Join(blok, "\n\n")
	return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: hasil}, m)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "menu.js",
		Jenis:    "system",
		Nama:     []string{".menu"},
		Itunya:   regexp.MustCompile(`(?i)^(#|!|/|\.)(help|(all)?m[e\w]n[u\w](all)?)$`),
		Jalankan: TanganiMenu,
	})
}
