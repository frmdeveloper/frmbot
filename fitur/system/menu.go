package system

import (
	"fmt"
	"regexp"
	"runtime"
	"sort"
	"strings"

	"gow1/lib"
)

func TanganiMenu(k *lib.Konteks) error {
	m := k.M
	if strings.Contains(m.Command, "mani") {
		_ = k.Conn.OK(m, "Mani loh ya")
	}
	q := strings.ToLower(strings.TrimSpace(m.Q))
	if q == "hehe" {
		return k.Conn.OK(m, "hehe")
	}

	semuaFitur := lib.Semua()
	perFile := []string{}
	for file := range semuaFitur {
		perFile = append(perFile, file)
	}
	sort.Strings(perFile)

	byJenis := map[string][][]string{}
	for _, file := range perFile {
		f := semuaFitur[file]
		if len(f.Nama) == 0 || f.Jenis == "" {
			continue
		}
		byJenis[f.Jenis] = append(byJenis[f.Jenis], f.Nama)
	}
	jenisList := []string{}
	for j := range byJenis {
		jenisList = append(jenisList, j)
	}
	sort.Strings(jenisList)

	if q != "" {
		filter := []string{}
		for _, j := range jenisList {
			if strings.EqualFold(j, q) {
				filter = append(filter, j)
			}
		}
		if len(filter) == 0 {
			return k.Conn.OK(m, m.Q)
		}
		jenisList = filter
	}

	emojiKategori := map[string]string{
		"ai": "🤖", "audio": "🎧", "chat": "💬", "converter": "🔄",
		"downloader": "📥", "game": "🎮", "group": "👥", "info": "🔍",
		"maker": "🎨", "system": "⚙️", "utility": "🧰",
	}
	blok := []string{}
	for _, j := range jenisList {
		emoji := emojiKategori[strings.ToLower(j)]
		if emoji == "" {
			emoji = "✨"
		}
		var sb strings.Builder
		sb.WriteString(emoji + " *" + strings.ToUpper(j) + "*")
		if strings.EqualFold(j, "group") && !m.IsGroup {
			sb.WriteString("\n🔒 disembunyikan, khusus grup")
			blok = append(blok, sb.String())
			continue
		}
		perintah := []string{}
		for _, nm := range byJenis[j] {
			perintah = append(perintah, nm...)
		}
		for _, n := range perintah {
			sb.WriteString("\n• " + n)
		}
		blok = append(blok, sb.String())
	}
	isi := strings.Join(blok, "\n\n")

	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	ramMB := float64(ms.Sys) / 1024 / 1024
	ram := fmt.Sprintf("%.1f MB", ramMB)
	if ramMB >= 1024 {
		ram = fmt.Sprintf("%.2f GB", ramMB/1024)
	}

	kepala := "✦ *FRM Bot* ✦\n" +
		"github.com/frmdeveloper/frmbot\n" +
		"_whatsmeow • golang_\n\n" +
		"💾 *RAM:* " + ram + "\n" +
		"⚡ *" + fmt.Sprint(lib.Jumlah()) + " fitur* siap dipakai\n" +
		"🎵 *follow:* tiktok.com/@rifki_mekanik\n\n" +
		"🪡 *daftar perintah* 🪡\n\n"
	kaki := "\n\n💡 ketik *#menu <kategori>* buat fokus ke satu kategori\n" +
		"contoh: #menu maker"

	teks := kepala + isi + kaki
	prefix := m.Prefix
	if prefix == "" {
		prefix = "#"
	}
	teks = strings.ReplaceAll(teks, "#", prefix)
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
