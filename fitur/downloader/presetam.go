package downloader

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"sync"

	"gow1/lib"
)

const uaPreset = "Mozilla/5.0 (Linux; Android 13; SM-A175F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0 Mobile Safari/537.36"

var polaPreset = regexp.MustCompile(`(?i)(?:https?://)?(?:(?:www\.)?alight\.link|(?:www\.)?alightcreative\.com|alight\.page\.link|drive\.google\.com)/[^\s"'<>)\]]+`)
var polaIDVideo = regexp.MustCompile(`/(?:video|photo)/([0-9]{15,})`)
var polaDriveID = regexp.MustCompile(`/d/([\w-]{20,})|[?&]id=([\w-]{20,})`)
var polaNamaDrive = regexp.MustCompile(`filename="?([^"]+)"?`)
var polaEkorLink = regexp.MustCompile(`[.,)\]]+$`)

type komentarTT struct {
	CID               string       `json:"cid"`
	Text              string       `json:"text"`
	DiggCount         int          `json:"digg_count"`
	ReplyCommentTotal int          `json:"reply_comment_total"`
	ReplyComment      []komentarTT `json:"reply_comment"`
	User              struct {
		UniqueID string `json:"unique_id"`
	} `json:"user"`
}

type responKomentar struct {
	Comments []komentarTT `json:"comments"`
	Cursor   int          `json:"cursor"`
	HasMore  int          `json:"has_more"`
}

type tautanPreset struct {
	Link string
	Dari string
	Suka int
}

func ambilJSONPreset(u string, target any) error {
	return lib.AmbilJSON(u, map[string]string{"User-Agent": uaPreset}, target)
}

func idVideoPreset(u string) string {
	if cocok := polaIDVideo.FindStringSubmatch(u); cocok != nil {
		return cocok[1]
	}
	req, err := http.NewRequest("HEAD", u, nil)
	if err != nil {
		return ""
	}
	req.Header.Set("User-Agent", uaPreset)
	resp, err := lib.KlienHTTP.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	jadi := resp.Request.URL.String()
	if cocok := polaIDVideo.FindStringSubmatch(jadi); cocok != nil {
		return cocok[1]
	}
	return ""
}

func bersihLinkPreset(u string) string {
	p := u
	if !strings.HasPrefix(p, "http") {
		p = "https://" + p
	}
	return polaEkorLink.ReplaceAllString(p, "")
}

func CariPreset(u string) (string, []tautanPreset) {
	id := idVideoPreset(u)
	if id == "" {
		return "", nil
	}
	var komentar []komentarTT
	cursor := 0
	for i := 0; i < 3; i++ {
		var j responKomentar
		err := ambilJSONPreset(fmt.Sprintf("https://www.tiktok.com/api/comment/list/?aweme_id=%s&count=50&cursor=%d&aid=1988", id, cursor), &j)
		if err != nil || len(j.Comments) == 0 {
			break
		}
		komentar = append(komentar, j.Comments...)
		cursor = j.Cursor
		if j.HasMore == 0 {
			break
		}
	}
	ketemu := map[string]*tautanPreset{}
	var urutan []string
	catat := func(km komentarTT) {
		for _, t := range polaPreset.FindAllString(km.Text, -1) {
			link := bersihLinkPreset(t)
			if _, ada := ketemu[link]; !ada {
				ketemu[link] = &tautanPreset{Link: link, Dari: km.User.UniqueID, Suka: km.DiggCount}
				urutan = append(urutan, link)
			}
		}
	}
	for _, km := range komentar {
		catat(km)
		for _, b := range km.ReplyComment {
			catat(b)
		}
	}
	var perluCek []komentarTT
	for _, km := range komentar {
		if km.ReplyCommentTotal > 0 && len(km.ReplyComment) == 0 {
			perluCek = append(perluCek, km)
		}
	}
	if len(perluCek) > 30 {
		perluCek = perluCek[:30]
	}
	for i := 0; i < len(perluCek); i += 5 {
		ujung := i + 5
		if ujung > len(perluCek) {
			ujung = len(perluCek)
		}
		bagian := perluCek[i:ujung]
		hasil := make([]*responKomentar, len(bagian))
		var wg sync.WaitGroup
		for j, km := range bagian {
			wg.Add(1)
			go func(j int, km komentarTT) {
				defer wg.Done()
				var r responKomentar
				if ambilJSONPreset(fmt.Sprintf("https://www.tiktok.com/api/comment/list/reply/?item_id=%s&comment_id=%s&count=50&cursor=0&aid=1988", id, km.CID), &r) == nil {
					hasil[j] = &r
				}
			}(j, km)
		}
		wg.Wait()
		for _, r := range hasil {
			if r == nil {
				continue
			}
			for _, b := range r.Comments {
				catat(b)
			}
		}
	}
	daftar := make([]tautanPreset, 0, len(urutan))
	for _, link := range urutan {
		daftar = append(daftar, *ketemu[link])
	}
	sort.SliceStable(daftar, func(a, b int) bool { return daftar[a].Suka > daftar[b].Suka })
	return id, daftar
}

type berkasDrive struct {
	Isi  []byte
	Nama string
}

func unduhDrive(link string) *berkasDrive {
	cocok := polaDriveID.FindStringSubmatch(link)
	if cocok == nil {
		return nil
	}
	berkas := cocok[1]
	if berkas == "" {
		berkas = cocok[2]
	}
	if berkas == "" {
		return nil
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://drive.usercontent.google.com/download?id=%s&export=download&confirm=t", berkas), nil)
	if err != nil {
		return nil
	}
	req.Header.Set("User-Agent", uaPreset)
	resp, err := lib.KlienHTTP.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	if strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		return nil
	}
	isi, err := io.ReadAll(io.LimitReader(resp.Body, 100*1024*1024))
	if err != nil {
		return nil
	}
	nama := ""
	if m := polaNamaDrive.FindStringSubmatch(resp.Header.Get("Content-Disposition")); m != nil {
		mentah := []byte(m[1])
		runes := make([]rune, len(mentah))
		for i, b := range mentah {
			runes[i] = rune(b)
		}
		nama = string(runes)
	}
	if nama == "" {
		nama = "preset.xml"
	}
	return &berkasDrive{Isi: isi, Nama: nama}
}

func TanganiPresetam(k *lib.Konteks) error {
	if k.M.Q == "" {
		return k.Conn.OK(k.M, "Link tiktoknya mana? Contoh:\n#presetam https://www.tiktok.com/@user/video/123")
	}
	u := polaTiktok.FindString(k.M.Q)
	if u == "" {
		u = regexp.MustCompile(`(?i)(?:https?://)?(?:[\w-]+\.)?vt\.tiktok\.com/\S+`).FindString(k.M.Q)
	}
	if u == "" {
		return k.Conn.OK(k.M, "Itu bukan link tiktok njir")
	}
	if !strings.HasPrefix(u, "http") {
		u = "https://" + u
	}
	id, tautan := CariPreset(u)
	if id == "" {
		return k.Conn.OK(k.M, "Gagal baca ID videonya, coba link versi panjang")
	}
	if len(tautan) == 0 {
		return k.Conn.OK(k.M, "Gak nemu link preset di komentar. Biasanya link-nya ada di balesan komentar si pengupload, mungkin videonya belum dibales atau presetnya di bio")
	}
	var baris []string
	for i, t := range tautan {
		s := fmt.Sprintf("%d. %s\n   dari @%s", i+1, t.Link, t.Dari)
		if t.Suka > 0 {
			s += fmt.Sprintf(" (%d suka)", t.Suka)
		}
		baris = append(baris, s)
	}
	teks := fmt.Sprintf("Nemu %d link preset:\n\n%s", len(tautan), strings.Join(baris, "\n\n"))
	if err := k.Conn.SendMessage(k.M.From, lib.IsiPesan{Text: teks}, k.M); err != nil {
		return err
	}
	jumlah := 0
	for _, t := range tautan {
		if !strings.Contains(t.Link, "drive.google.com") {
			continue
		}
		if jumlah >= 3 {
			break
		}
		jumlah++
		berkas := unduhDrive(t.Link)
		if berkas == nil {
			continue
		}
		mime := "application/octet-stream"
		if strings.HasSuffix(strings.ToLower(berkas.Nama), ".xml") {
			mime = "text/xml"
		}
		_ = k.Conn.SendMessage(k.M.From, lib.IsiPesan{Document: berkas.Isi, FileName: berkas.Nama, Mimetype: mime}, k.M)
	}
	return nil
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "presetam.js",
		Jenis:    "downloader",
		Nama:     []string{"#presetam"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(presetam|preset|amreset|cekpreset)$"),
		Jalankan: TanganiPresetam,
	})
}
