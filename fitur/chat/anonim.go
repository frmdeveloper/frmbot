package chat

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"gow1/lib"
)

type sesiAnon struct {
	Partner string
	Bot     string
}

type antriAnon struct {
	JID string
	Bot string
}

var anonMu sync.Mutex
var anonAntri *antriAnon
var anonPasangan = map[string]*sesiAnon{}

func kabarAnon(k *lib.Konteks, botNum, jid, teks string) {
	if botNum != "" && botNum != k.Conn.BotNomor() {
		return
	}
	_ = k.Conn.SendMessage(jid, lib.IsiPesan{Text: teks}, nil)
}

func SesiAnonAktif(jid string) bool {
	anonMu.Lock()
	defer anonMu.Unlock()
	return anonPasangan[jid] != nil
}

func bersihinAnon(jid string) *sesiAnon {
	sesi := anonPasangan[jid]
	if sesi != nil {
		delete(anonPasangan, sesi.Partner)
		delete(anonPasangan, jid)
	}
	if anonAntri != nil && anonAntri.JID == jid {
		anonAntri = nil
	}
	return sesi
}

func mulaiAnon(k *lib.Konteks) error {
	m := k.M
	me := m.From
	anonMu.Lock()
	if anonPasangan[me] != nil {
		anonMu.Unlock()
		return k.Conn.OK(m, "kamu udah nyambung sama seseorang 😏\n_.next_ ganti · _.stop_ udahan")
	}
	if anonAntri != nil && anonAntri.JID == me {
		anonMu.Unlock()
		return k.Conn.OK(m, "sabar, lagi dicariin partner... 🔎")
	}
	if anonAntri == nil {
		anonAntri = &antriAnon{JID: me, Bot: k.Conn.BotNomor()}
		anonMu.Unlock()
		return k.Conn.OK(m, "🔎 lagi nyari partner... tunggu ya\n_.stop_ buat batal")
	}
	lawan := anonAntri
	anonAntri = nil
	anonPasangan[me] = &sesiAnon{Partner: lawan.JID, Bot: lawan.Bot}
	anonPasangan[lawan.JID] = &sesiAnon{Partner: me, Bot: k.Conn.BotNomor()}
	anonMu.Unlock()
	sapa := "✅ ketemu partner! ngobrol aja, semua pesanmu diteruskan anonim.\n_.next_ ganti · _.stop_ udahan"
	if err := k.Conn.OK(m, sapa); err != nil {
		return err
	}
	kabarAnon(k, lawan.Bot, lawan.JID, sapa)
	return nil
}

func nextAnon(k *lib.Konteks) error {
	m := k.M
	me := m.From
	anonMu.Lock()
	sesi := anonPasangan[me]
	if sesi != nil {
		bersihinAnon(me)
	}
	sudahAntri := anonAntri != nil && anonAntri.JID == me
	anonMu.Unlock()
	if sesi != nil {
		kabarAnon(k, sesi.Bot, sesi.Partner, "😶 partner pergi nyari yang lain.\nketik _.anon_ buat cari partner baru")
	}
	if sudahAntri {
		return k.Conn.OK(m, "kamu udah di antrian, sabar ya 🔎")
	}
	return mulaiAnon(k)
}

func stopAnon(k *lib.Konteks) error {
	m := k.M
	me := m.From
	anonMu.Lock()
	if anonAntri != nil && anonAntri.JID == me {
		anonAntri = nil
		anonMu.Unlock()
		return k.Conn.OK(m, "❌ berhenti nyari partner.")
	}
	sesi := anonPasangan[me]
	if sesi == nil {
		anonMu.Unlock()
		return k.Conn.OK(m, "kamu lagi gak di sesi anonim 🤔\nketik _.anon_ buat mulai")
	}
	bersihinAnon(me)
	anonMu.Unlock()
	kabarAnon(k, sesi.Bot, sesi.Partner, "👋 partner ngakhirin sesi.\nketik _.anon_ buat cari lagi")
	return k.Conn.OK(m, "👋 sesi anonim diakhiri.")
}

func TanganiAnonim(k *lib.Konteks) error {
	m := k.M
	if m.IsGroup {
		return k.Conn.OK(m, "anonim chat cuma bisa di chat pribadi ya 🙏")
	}
	cmd := m.Command
	if r := []rune(cmd); len(r) > 0 && strings.ContainsRune("!#%./\\", r[0]) {
		cmd = string(r[1:])
	}
	if strings.EqualFold(cmd, "next") {
		return nextAnon(k)
	}
	if strings.EqualFold(cmd, "stop") || strings.EqualFold(cmd, "keluar") || strings.EqualFold(cmd, "leave") {
		return stopAnon(k)
	}
	return mulaiAnon(k)
}

func TeruskanAnonim(k *lib.Konteks) error {
	m := k.M
	if m.FromMe || m.IsGroup || m.Prefix != "" {
		return nil
	}
	anonMu.Lock()
	sesi := anonPasangan[m.From]
	anonMu.Unlock()
	if sesi == nil {
		return nil
	}
	if sesi.Bot != "" && sesi.Bot != k.Conn.BotNomor() {
		anonMu.Lock()
		bersihinAnon(m.From)
		anonMu.Unlock()
		return k.Conn.OK(m, "⚠️ partner udah gak keconek, sesi diputus.\nketik _.anon_ buat cari lagi")
	}

	gagal := func() error {
		fmt.Println("[anonim] forward gagal")
		anonMu.Lock()
		bersihinAnon(m.From)
		anonMu.Unlock()
		return k.Conn.OK(m, "⚠️ gagal nerusin ke partner, sesi diputus.\nketik _.anon_ buat cari lagi")
	}

	jenis := m.Media
	if jenis == "" {
		if m.Type == "conversation" || m.Type == "extendedTextMessage" {
			if err := k.Conn.SendMessage(sesi.Partner, lib.IsiPesan{Text: m.Text}, nil); err != nil {
				return gagal()
			}
			return nil
		}
		return k.Conn.OK(m, "🚫 jenis pesan itu belum bisa diteruskan")
	}

	data, mime, err := k.Conn.DownloadAny(m)
	if err != nil {
		return gagal()
	}
	isi := lib.IsiPesan{Mimetype: mime}
	switch jenis {
	case "image":
		isi.Image = data
	case "video":
		isi.Video = data
	case "sticker":
		isi.Sticker = data
	case "audio":
		isi.Audio = data
	case "document":
		isi.Document = data
	default:
		return k.Conn.OK(m, "🚫 jenis pesan itu belum bisa diteruskan")
	}
	if m.Text != "" && jenis != "sticker" {
		isi.Caption = m.Text
	}
	if err := k.Conn.SendMessage(sesi.Partner, isi, nil); err != nil {
		return gagal()
	}
	return nil
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "anonim.js",
		Jenis:    "chat",
		Nama:     []string{"#anon", "#next", "#stop"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(anon(im)?|acs|cari|next|stop|keluar|leave)$"),
		Jalankan: TanganiAnonim,
	})
}
