package group

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unicode"

	"gow1/lib"
)

const berkasTiktok = "data/tiktok.json"
const uaTiktok = "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Mobile Safari/537.36"
const jedaCekTiktok = 3 * time.Minute
const maksKirimTiktok = 3

type recTiktok struct {
	Nick   string `json:"nick"`
	LastID string `json:"lastId"`
	Oleh   string `json:"oleh"`
	Bot    string `json:"bot"`
}

type videoTiktok struct {
	ID          string
	Desc        string
	Cover       string
	OriginCover string
	PlayAddr    string
	Width       int
	Height      int
	Private     bool
}

type profilTiktok struct {
	Nick           string
	Avatar         string
	Bio            string
	Verified       bool
	PrivateAccount bool
	Heart          any
	Follower       any
	Following      any
	Videos         []videoTiktok
}

type galatHTTP struct{ Status int }

func (e *galatHTTP) Error() string { return fmt.Sprintf("HTTP %d", e.Status) }

var muTiktok sync.Mutex
var dataTiktok map[string]map[string]*recTiktok
var tiktokTermuat bool
var connTiktok atomic.Value
var pemantauSekali sync.Once

var polaStateTiktok = regexp.MustCompile(`(?s)<script id="__FRONTITY_CONNECT_STATE__"[^>]*>(.+?)</script>`)
var polaAwalanTiktok = regexp.MustCompile(`(?i)^https?://(www\.)?tiktok\.com/`)

func muatTiktok() {
	if tiktokTermuat {
		return
	}
	tiktokTermuat = true
	dataTiktok = map[string]map[string]*recTiktok{}
	b, err := os.ReadFile(berkasTiktok)
	if err != nil {
		return
	}
	_ = json.Unmarshal(b, &dataTiktok)
	if dataTiktok == nil {
		dataTiktok = map[string]map[string]*recTiktok{}
	}
}

func simpanTiktok() {
	b, err := json.Marshal(dataTiktok)
	if err != nil {
		return
	}
	_ = os.WriteFile(berkasTiktok, b, 0o644)
}

func bersihNamaTiktok(teks string) string {
	s := strings.TrimSpace(teks)
	bagian := strings.FieldsFunc(s, func(r rune) bool {
		return r == '?' || unicode.IsSpace(r)
	})
	if len(bagian) > 0 {
		s = bagian[0]
	} else {
		s = ""
	}
	s = polaAwalanTiktok.ReplaceAllString(s, "")
	s = strings.TrimPrefix(s, "@")
	if i := strings.Index(s, "/"); i >= 0 {
		s = s[:i]
	}
	return strings.ToLower(s)
}

func stringDari(v any) string {
	switch t := v.(type) {
	case string:
		return t
	case json.Number:
		return t.String()
	}
	return ""
}

func boolDari(v any) bool {
	b, _ := v.(bool)
	return b
}

func intDari(v any) int {
	switch t := v.(type) {
	case json.Number:
		n, _ := t.Int64()
		return int(n)
	case float64:
		return int(t)
	}
	return 0
}

func formatHitung(v any) string {
	if v == nil {
		return "-"
	}
	var n int64
	switch t := v.(type) {
	case json.Number:
		x, err := t.Int64()
		if err != nil {
			return "-"
		}
		n = x
	case float64:
		n = int64(t)
	default:
		return "-"
	}
	if n == 0 {
		return "-"
	}
	s := fmt.Sprintf("%d", n)
	negatif := strings.HasPrefix(s, "-")
	if negatif {
		s = s[1:]
	}
	bagian := []string{}
	for len(s) > 3 {
		bagian = append([]string{s[len(s)-3:]}, bagian...)
		s = s[:len(s)-3]
	}
	bagian = append([]string{s}, bagian...)
	hasil := strings.Join(bagian, ".")
	if negatif {
		hasil = "-" + hasil
	}
	return hasil
}

func bandingID(a, b string) int {
	x, ok1 := new(big.Int).SetString(a, 10)
	y, ok2 := new(big.Int).SetString(b, 10)
	if !ok1 || !ok2 {
		if len(a) != len(b) {
			if len(a) > len(b) {
				return 1
			}
			return -1
		}
		return strings.Compare(a, b)
	}
	return x.Cmp(y)
}

func waktuVideoTiktok(id string) time.Time {
	n, ok := new(big.Int).SetString(id, 10)
	if !ok {
		return time.Time{}
	}
	detik := new(big.Int).Rsh(n, 32).Int64()
	return time.Unix(detik, 0)
}

func palingBaruTiktok(videos []videoTiktok) videoTiktok {
	terbaik := videos[0]
	for _, v := range videos[1:] {
		if bandingID(v.ID, terbaik.ID) > 0 {
			terbaik = v
		}
	}
	return terbaik
}

func ambilProfilTiktok(user string) (*profilTiktok, error) {
	req, err := http.NewRequest("GET", "https://www.tiktok.com/embed/@"+user, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", uaTiktok)
	klien := *lib.KlienHTTP
	klien.Timeout = 25 * time.Second
	resp, err := klien.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return nil, &galatHTTP{Status: resp.StatusCode}
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	mm := polaStateTiktok.FindSubmatch(body)
	if mm == nil {
		return nil, nil
	}
	dec := json.NewDecoder(bytes.NewReader(mm[1]))
	dec.UseNumber()
	var akar any
	if err := dec.Decode(&akar); err != nil {
		return nil, nil
	}
	root, _ := akar.(map[string]any)
	source, _ := root["source"].(map[string]any)
	data, _ := source["data"].(map[string]any)
	var halaman map[string]any
	for _, v := range data {
		obj, _ := v.(map[string]any)
		if obj == nil {
			continue
		}
		if _, ada := obj["videoList"]; ada {
			halaman = obj
			break
		}
	}
	if halaman == nil {
		return nil, nil
	}
	ui, _ := halaman["userInfo"].(map[string]any)
	p := &profilTiktok{
		Nick:           stringDari(ui["nickname"]),
		Avatar:         stringDari(ui["avatarThumbUrl"]),
		Bio:            stringDari(ui["signature"]),
		Verified:       boolDari(ui["verified"]),
		PrivateAccount: boolDari(ui["privateAccount"]),
		Heart:          ui["heartCount"],
		Follower:       ui["followerCount"],
		Following:      ui["followingCount"],
	}
	if p.Nick == "" {
		p.Nick = user
	}
	lista, _ := halaman["videoList"].([]any)
	for _, item := range lista {
		vm, _ := item.(map[string]any)
		if vm == nil {
			continue
		}
		id := stringDari(vm["id"])
		if id == "" || boolDari(vm["privateItem"]) {
			continue
		}
		p.Videos = append(p.Videos, videoTiktok{
			ID:          id,
			Desc:        stringDari(vm["desc"]),
			Cover:       stringDari(vm["coverUrl"]),
			OriginCover: stringDari(vm["originCoverUrl"]),
			PlayAddr:    stringDari(vm["playAddr"]),
			Width:       intDari(vm["width"]),
			Height:      intDari(vm["height"]),
			Private:     boolDari(vm["privateItem"]),
		})
	}
	return p, nil
}

func ambilProfilUlangi(user string) (*profilTiktok, error) {
	var hasil *profilTiktok
	var gagal error
	for i := 0; i < 4; i++ {
		if i > 0 {
			time.Sleep(3 * time.Second)
		}
		p, err := ambilProfilTiktok(user)
		if err != nil {
			hasil = nil
			gagal = err
			if gh, ok := err.(*galatHTTP); ok && gh.Status == 400 {
				break
			}
			continue
		}
		hasil = p
		gagal = nil
		if p != nil && len(p.Videos) > 0 {
			break
		}
	}
	return hasil, gagal
}

func kirimGambarURLTiktok(conn lib.Conn, chat, url, caption string, mentions []string, quoted *lib.Pesan) error {
	data, err := lib.AmbilURL(url, map[string]string{"User-Agent": uaTiktok})
	if err != nil {
		return err
	}
	return conn.SendMessage(chat, lib.IsiPesan{Image: data, Caption: caption, Mentions: mentions}, quoted)
}

func infoTiktok(k *lib.Konteks, chat, user string, quoted *lib.Pesan) {
	hasil, gagal := ambilProfilUlangi(user)
	if hasil == nil || len(hasil.Videos) == 0 {
		alasan := fmt.Sprintf("TikTok @%s lagi labil, ulangi nanti", user)
		if gh, ok := gagal.(*galatHTTP); ok && gh.Status == 400 {
			alasan = fmt.Sprintf("Akun @%s gak ada di tiktok", user)
		}
		_ = k.Conn.SendMessage(chat, lib.IsiPesan{Text: alasan}, quoted)
		return
	}
	status := "Publik 🌍"
	if hasil.PrivateAccount {
		status = "Privat 🔒"
	}
	verif := "Tidak"
	if hasil.Verified {
		verif = "Ya ✅"
	}
	teks := fmt.Sprintf("「 *INFORMASI AKUN TIKTOK* 」\n- ۶ৎ  Nickname : *%s*\n- ۶ৎ  Username : *@%s*\n- ۶ৎ  Bio : %s\n- ۶ৎ  Status : *%s*\n- ۶ৎ  Verifikasi : *%s*\n- ۶ৎ  Total Suka : *%s*\n- ۶ৎ  Pengikut : *%s*\n- ۶ৎ  Mengikuti : *%s*",
		hasil.Nick, user, hasil.Bio, status, verif,
		formatHitung(hasil.Heart), formatHitung(hasil.Follower), formatHitung(hasil.Following))
	if hasil.Bio == "" {
		teks = strings.Replace(teks, "Bio : \n", "Bio : -\n", 1)
	}
	if hasil.Avatar != "" {
		if err := kirimGambarURLTiktok(k.Conn, chat, hasil.Avatar, teks, nil, quoted); err == nil {
			return
		}
	}
	_ = k.Conn.SendMessage(chat, lib.IsiPesan{Text: teks}, quoted)
}

func mulaiPemantauTiktok() {
	pemantauSekali.Do(func() {
		go func() {
			t := time.NewTicker(jedaCekTiktok)
			defer t.Stop()
			for range t.C {
				cekSemuaTiktok()
			}
		}()
	})
}

func cekSemuaTiktok() {
	cv := connTiktok.Load()
	if cv == nil {
		return
	}
	conn, _ := cv.(lib.Conn)
	if conn == nil {
		return
	}
	type tujuan struct {
		chat string
		rec  *recTiktok
	}
	target := map[string][]tujuan{}
	muTiktok.Lock()
	muatTiktok()
	for chat, pantauan := range dataTiktok {
		if !strings.HasSuffix(chat, "@g.us") {
			continue
		}
		for user, rec := range pantauan {
			target[user] = append(target[user], tujuan{chat: chat, rec: rec})
		}
	}
	muTiktok.Unlock()
	if len(target) == 0 {
		return
	}
	users := make([]string, 0, len(target))
	for u := range target {
		users = append(users, u)
	}
	sort.Strings(users)
	for _, user := range users {
		hasil, err := ambilProfilTiktok(user)
		if err != nil || hasil == nil || len(hasil.Videos) == 0 {
			continue
		}
		for _, tj := range target[user] {
			baru := []videoTiktok{}
			for _, v := range hasil.Videos {
				if bandingID(v.ID, tj.rec.LastID) > 0 {
					baru = append(baru, v)
				}
			}
			if len(baru) == 0 {
				continue
			}
			sort.Slice(baru, func(i, j int) bool { return bandingID(baru[i].ID, baru[j].ID) < 0 })
			muTiktok.Lock()
			tj.rec.LastID = baru[len(baru)-1].ID
			tj.rec.Nick = hasil.Nick
			simpanTiktok()
			muTiktok.Unlock()
			var mentions []string
			cc := ""
			if tj.rec.Oleh != "" {
				mentions = []string{tj.rec.Oleh}
				cc = "@" + lib.NomorSaja(tj.rec.Oleh)
			}
			kirim := baru
			if len(kirim) > maksKirimTiktok {
				kirim = kirim[len(kirim)-maksKirimTiktok:]
			}
			for _, v := range kirim {
				link := "https://www.tiktok.com/@" + user
				teks := fmt.Sprintf("🎬 *%s* (@%s) upload video baru\n%s\n\n%s\n\n%s", hasil.Nick, user, cc, v.Desc, link)
				terkirim := false
				cover := v.OriginCover
				if cover == "" {
					cover = v.Cover
				}
				if cover != "" {
					if err := kirimGambarURLTiktok(conn, tj.chat, cover, teks, mentions, nil); err == nil {
						terkirim = true
					}
				}
				if !terkirim && v.Cover != "" && v.Cover != cover {
					if err := kirimGambarURLTiktok(conn, tj.chat, v.Cover, teks, mentions, nil); err == nil {
						terkirim = true
					}
				}
				if !terkirim {
					_ = conn.SendMessage(tj.chat, lib.IsiPesan{Text: teks, Mentions: mentions}, nil)
				}
				time.Sleep(2 * time.Second)
			}
		}
		time.Sleep(3 * time.Second)
	}
}

func TanganiVeriftt(k *lib.Konteks) error {
	m := k.M
	if !m.IsGroup {
		return k.Conn.OK(m, "Pantau tiktok cuma jalan di grup")
	}
	connTiktok.Store(k.Conn)
	mulaiPemantauTiktok()

	muTiktok.Lock()
	muatTiktok()
	if dataTiktok[m.From] == nil {
		dataTiktok[m.From] = map[string]*recTiktok{}
	}
	pantauan := dataTiktok[m.From]
	muTiktok.Unlock()

	cmd := m.Command

	if strings.Contains(cmd, "listtt") {
		if len(pantauan) == 0 {
			return k.Conn.OK(m, "Belum ada akun tiktok yang dipantau di sini")
		}
		users := make([]string, 0, len(pantauan))
		for u := range pantauan {
			users = append(users, u)
		}
		sort.Strings(users)
		baris := make([]string, len(users))
		for i, u := range users {
			rec := pantauan[u]
			baris[i] = fmt.Sprintf("%d. @%s — %s\n    video terakhir: %s", i+1, u, rec.Nick, waktuVideoTiktok(rec.LastID).Format("2/1/2006, 15.04.05"))
		}
		teks := "🎬 *Pantauan tiktok di sini*\n\n" + strings.Join(baris, "\n")
		return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: teks}, m)
	}

	if strings.Contains(cmd, "testtt") {
		if !pengirimOwner(k) {
			return k.Conn.OK(m, "khusus owner")
		}
		user := bersihNamaTiktok(m.Q)
		if user == "" {
			return k.Conn.OK(m, "Format: #testtt <user>")
		}
		hasil, err := ambilProfilTiktok(user)
		if err != nil {
			return k.Conn.OK(m, "gagal ambil profil: "+err.Error())
		}
		if hasil == nil || len(hasil.Videos) == 0 {
			return k.Conn.OK(m, fmt.Sprintf("@%s gak ada video / gagal parse", user))
		}
		v := palingBaruTiktok(hasil.Videos)
		link := fmt.Sprintf("https://www.tiktok.com/@%s/video/%s", user, v.ID)
		playAddr := "false"
		if v.PlayAddr != "" {
			playAddr = "true"
		}
		cover := v.Cover
		if len(cover) > 70 {
			cover = cover[:70]
		}
		teks := fmt.Sprintf("🧪 TEST @%s vid %s\nwxh: %dx%d | playAddr: %s\ncover: %s", user, v.ID, v.Width, v.Height, playAddr, cover)
		_ = k.Conn.SendMessage(m.From, lib.IsiPesan{Text: teks}, m)
		for _, varian := range []struct {
			label string
			dim   int
			q     int
		}{{"A", 320, 70}, {"B", 200, 50}} {
			teksVarian := fmt.Sprintf("🧪 TEST-%s (%dpx q%d) thumb=NULL\n%s", varian.label, varian.dim, varian.q, link)
			data, derr := lib.AmbilURL(v.PlayAddr, map[string]string{"User-Agent": uaTiktok})
			if derr != nil {
				_ = k.Conn.SendMessage(m.From, lib.IsiPesan{Text: fmt.Sprintf("TEST-%s gagal kirim video: %s", varian.label, derr.Error())}, nil)
			} else if serr := k.Conn.SendMessage(m.From, lib.IsiPesan{Video: data, Caption: teksVarian}, nil); serr != nil {
				_ = k.Conn.SendMessage(m.From, lib.IsiPesan{Text: fmt.Sprintf("TEST-%s gagal kirim video: %s", varian.label, serr.Error())}, nil)
			}
			time.Sleep(2500 * time.Millisecond)
		}
		return nil
	}

	if strings.Contains(cmd, "infott") {
		if m.Q == "" {
			punya := []string{}
			muTiktok.Lock()
			for _, p := range dataTiktok {
				for user, rec := range p {
					if rec.Oleh == m.Sender {
						ada := false
						for _, x := range punya {
							if x == user {
								ada = true
								break
							}
						}
						if !ada {
							punya = append(punya, user)
						}
					}
				}
			}
			muTiktok.Unlock()
			sort.Strings(punya)
			if len(punya) == 0 {
				return k.Conn.OK(m, "Lu belum verif akun tiktok mana pun")
			}
			for _, user := range punya {
				infoTiktok(k, m.From, user, m)
				time.Sleep(2 * time.Second)
			}
			return nil
		}
		user := bersihNamaTiktok(m.Q)
		if user == "" {
			return k.Conn.OK(m, "Username tiktoknya gak kebaca, tulis @username atau link profilnya")
		}
		infoTiktok(k, m.From, user, m)
		return nil
	}

	if m.Q == "" {
		return k.Conn.OK(m, fmt.Sprintf("Formatnya: %s <username tiktok>", m.Command))
	}
	user := bersihNamaTiktok(m.Q)
	if user == "" {
		return k.Conn.OK(m, "Username tiktoknya gak kebaca, tulis @username atau link profilnya")
	}

	if strings.Contains(cmd, "unveriftt") {
		muTiktok.Lock()
		_, ada := pantauan[user]
		if ada {
			delete(pantauan, user)
			simpanTiktok()
		}
		muTiktok.Unlock()
		if !ada {
			return k.Conn.OK(m, fmt.Sprintf("@%s gak dipantau di sini", user))
		}
		return k.Conn.OK(m, fmt.Sprintf("Berhenti mantau @%s", user))
	}

	muTiktok.Lock()
	_, sudah := pantauan[user]
	muTiktok.Unlock()
	if sudah {
		return k.Conn.OK(m, fmt.Sprintf("@%s udah dipantau di sini", user))
	}

	hasil, gagal := ambilProfilUlangi(user)
	if gh, ok := gagal.(*galatHTTP); ok && gh.Status == 400 {
		return k.Conn.OK(m, fmt.Sprintf("Akun @%s gak ada di tiktok", user))
	}
	if hasil == nil || len(hasil.Videos) == 0 {
		return k.Conn.OK(m, fmt.Sprintf("TikTok-nya lagi labil, ulangi: %s @%s", m.Command, user))
	}

	muTiktok.Lock()
	pantauan[user] = &recTiktok{
		Nick:   hasil.Nick,
		LastID: palingBaruTiktok(hasil.Videos).ID,
		Oleh:   m.Sender,
		Bot:    k.Conn.BotNomor(),
	}
	simpanTiktok()
	muTiktok.Unlock()

	teks := fmt.Sprintf("✅ Mulai mantau *%s* (@%s)\n\nTiap dia upload video baru langsung dikabarin di sini. Dicek tiap %d menit.", hasil.Nick, user, int(jedaCekTiktok/time.Minute))
	if hasil.Avatar != "" {
		if err := kirimGambarURLTiktok(k.Conn, m.From, hasil.Avatar, teks, nil, m); err == nil {
			return nil
		}
	}
	return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: teks}, m)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "veriftt.js",
		Jenis:    "group",
		Nama:     []string{"#veriftt @user", "#unveriftt @user", "#listtt", "#testtt @user"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(verif(i)?tt|unverif(i)?tt|listtt|testtt|infott)$"),
		Jalankan: TanganiVeriftt,
	})
}
