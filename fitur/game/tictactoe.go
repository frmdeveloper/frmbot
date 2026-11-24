package game

import (
	"encoding/json"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"sync"

	waE2E "go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"

	"gow1/lib"
)

const penandaJIDTTT = "120363186235853203@g.us"

var angkaTTT = []string{"1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣", "6️⃣", "7️⃣", "8️⃣", "9️⃣"}

var ikonTTT = map[byte]string{'X': "❌", 'O': "⭕"}

var garisTTT = [][3]int{
	{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
	{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
	{0, 4, 8}, {2, 4, 6},
}

type keadaanTTT struct {
	Papan string
	X     string
	O     string
}

var tttMu sync.Mutex
var tttSesi = map[string]*keadaanTTT{}

func hitungTTT(papan string, c byte) int {
	n := 0
	for i := 0; i < len(papan); i++ {
		if papan[i] == c {
			n++
		}
	}
	return n
}

func giliranTTT(papan string) byte {
	if hitungTTT(papan, 'X') <= hitungTTT(papan, 'O') {
		return 'X'
	}
	return 'O'
}

func menangTTT(papan string) *[3]int {
	for i := range garisTTT {
		g := garisTTT[i]
		if papan[g[0]] != '.' && papan[g[0]] == papan[g[1]] && papan[g[1]] == papan[g[2]] {
			return &g
		}
	}
	return nil
}

func gambarTTT(papan string) string {
	baris := make([]string, 0, 3)
	for _, i := range []int{0, 3, 6} {
		var sb strings.Builder
		for j := 0; j < 3; j++ {
			c := papan[i+j]
			if c == '.' {
				sb.WriteString(angkaTTT[i+j])
			} else {
				sb.WriteString(ikonTTT[c])
			}
		}
		baris = append(baris, sb.String())
	}
	return strings.Join(baris, "\n")
}

func pasangTTT(papan string, i int, c byte) string {
	b := []byte(papan)
	b[i] = c
	return string(b)
}

func skorTTT(papan string, dalam int) int {
	if g := menangTTT(papan); g != nil {
		if papan[g[0]] == 'O' {
			return 10 - dalam
		}
		return dalam - 10
	}
	if !strings.Contains(papan, ".") {
		return 0
	}
	gil := giliranTTT(papan)
	pertama := true
	terbaik := 0
	for i := 0; i < 9; i++ {
		if papan[i] != '.' {
			continue
		}
		n := skorTTT(pasangTTT(papan, i, gil), dalam+1)
		if pertama || (gil == 'O' && n > terbaik) || (gil == 'X' && n < terbaik) {
			terbaik = n
			pertama = false
		}
	}
	return terbaik
}

func langkahBotTTT(papan string) int {
	terbaik := -1 << 30
	var pilihan []int
	for i := 0; i < 9; i++ {
		if papan[i] != '.' {
			continue
		}
		n := skorTTT(pasangTTT(papan, i, 'O'), 1)
		if n > terbaik {
			terbaik = n
			pilihan = []int{i}
		} else if n == terbaik {
			pilihan = append(pilihan, i)
		}
	}
	return pilihan[rand.Intn(len(pilihan))]
}

func petaTTT(papan, x, o string) string {
	j := map[string]string{}
	for i := 0; i < 9; i++ {
		if papan[i] == '.' {
			j[strconv.Itoa(i+1)] = "#ttt " + strconv.Itoa(i+1) + "|" + papan + "|" + x + "|" + o
		}
	}
	b, _ := json.Marshal(j)
	return "json|" + string(b)
}

func kirimTTT(k *lib.Konteks, papan, x, o string) error {
	m := k.M
	bot := k.Conn.BotNomor()
	nama := func(a string) string {
		switch {
		case a == "-":
			return "_menunggu lawan_"
		case a == bot:
			return "*bot*"
		default:
			return "@" + a
		}
	}
	garis := menangTTT(papan)
	var juara byte
	if garis != nil {
		juara = papan[garis[0]]
	}
	seri := garis == nil && !strings.Contains(papan, ".")
	usai := garis != nil || seri
	gil := giliranTTT(papan)

	var sb strings.Builder
	sb.WriteString("*TicTacToe*\n")
	sb.WriteString(ikonTTT['X'] + " " + nama(x) + "\n")
	sb.WriteString(ikonTTT['O'] + " " + nama(o) + "\n\n")
	sb.WriteString(gambarTTT(papan) + "\n\n")
	switch {
	case garis != nil:
		pemenang := x
		if juara == 'O' {
			pemenang = o
		}
		sb.WriteString("🏆 " + nama(pemenang) + " menang!")
	case seri:
		sb.WriteString("Seri!")
	default:
		giliranNama := ""
		if gil == 'X' {
			giliranNama = nama(x)
		} else if o == "-" {
			giliranNama = "_siapa saja_"
		} else {
			giliranNama = nama(o)
		}
		sb.WriteString("Giliran " + ikonTTT[gil] + " " + giliranNama + "\n_Balas pesan ini dengan angka kotak._")
	}
	teks := sb.String()

	var mentions []string
	for _, a := range []string{x, o} {
		if a != "-" && a != bot {
			mentions = append(mentions, a+"@s.whatsapp.net")
		}
	}

	tttMu.Lock()
	if usai {
		delete(tttSesi, m.From)
	} else {
		tttSesi[m.From] = &keadaanTTT{Papan: papan, X: x, O: o}
	}
	tttMu.Unlock()

	if !usai {
		ci := &waE2E.ContextInfo{
			MentionedJID: mentions,
			GroupMentions: []*waE2E.GroupMention{
				{GroupJID: proto.String(penandaJIDTTT), GroupSubject: proto.String(petaTTT(papan, x, o))},
			},
		}
		ci.StanzaID = proto.String(m.ID)
		if m.Sender != "" {
			ci.Participant = proto.String(m.Sender)
		}
		switch f := m.Full.(type) {
		case *events.Message:
			ci.QuotedMessage = f.Message
		case *waE2E.Message:
			ci.QuotedMessage = f
		}
		pesan := &waE2E.Message{
			ExtendedTextMessage: &waE2E.ExtendedTextMessage{
				Text:        proto.String(teks),
				ContextInfo: ci,
			},
		}
		if err := k.Conn.RelayMentah(m.From, pesan); err == nil {
			return nil
		}
	}
	return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: teks, Mentions: mentions}, m)
}

func parseIntJS(s string) (int, bool) {
	s = strings.TrimSpace(s)
	i := 0
	if i < len(s) && (s[i] == '+' || s[i] == '-') {
		i++
	}
	mulai := i
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		i++
	}
	if i == mulai {
		return 0, false
	}
	n, err := strconv.Atoi(s[:i])
	if err != nil {
		return 0, false
	}
	return n, true
}

func TanganiTictactoe(k *lib.Konteks) error {
	m := k.M
	var aku []string
	for _, j := range []string{m.Sender, m.SenderAlt} {
		if n := lib.NomorSaja(j); n != "" {
			aku = append(aku, n)
		}
	}
	bot := k.Conn.BotNomor()
	termasuk := func(n string) bool {
		for _, a := range aku {
			if a == n {
				return true
			}
		}
		return false
	}

	if !strings.Contains(m.Q, "|") {
		lawan := bot
		if m.IsGroup {
			lawan = "-"
		}
		x := ""
		if len(aku) > 0 {
			x = aku[0]
		}
		return kirimTTT(k, ".........", x, lawan)
	}

	bagian := strings.Split(m.Q, "|")
	if len(bagian) < 4 {
		return k.Conn.OK(m, "Papannya udah basi")
	}
	sel, papan, x, o := bagian[0], bagian[1], bagian[2], bagian[3]
	angka, ok := parseIntJS(sel)
	i := angka - 1
	if !ok || i < 0 || i > 8 || len(papan) != 9 || papan[i] != '.' {
		return k.Conn.OK(m, "Papannya udah basi")
	}

	gil := giliranTTT(papan)
	lawan := o
	if m.IsGroup {
		if gil == 'X' && !termasuk(x) {
			return k.Conn.OK(m, "Bukan giliran kamu")
		}
		if gil == 'O' {
			if o == "-" && termasuk(x) {
				return k.Conn.OK(m, "Tunggu lawan dulu")
			}
			if o != "-" && !termasuk(o) {
				return k.Conn.OK(m, "Bukan giliran kamu")
			}
		}
	}
	if gil == 'O' && lawan == "-" {
		if len(aku) > 0 {
			lawan = aku[0]
		}
	}

	baru := pasangTTT(papan, i, gil)
	if lawan == bot && menangTTT(baru) == nil && strings.Contains(baru, ".") {
		j := langkahBotTTT(baru)
		baru = pasangTTT(baru, j, 'O')
	}

	return kirimTTT(k, baru, x, lawan)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "tictactoe.js",
		Jenis:    "game",
		Nama:     []string{"#ttt"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(ttt|tictactoe|xo)$"),
		Jalankan: TanganiTictactoe,
	})
}
