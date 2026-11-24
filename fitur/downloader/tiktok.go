package downloader

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html"

	"gow1/lib"
)

type dataTT struct {
	Region   string
	Nickname string
	Title    string
	Music    string
	MusicURL string
	Play     string
	Images   []string
}

func formatRibuan(n int64, ada bool) string {
	if !ada {
		return "-"
	}
	s := strconv.FormatInt(n, 10)
	negatif := strings.HasPrefix(s, "-")
	if negatif {
		s = s[1:]
	}
	var bagian []string
	for len(s) > 3 {
		bagian = append([]string{s[len(s)-3:]}, bagian...)
		s = s[:len(s)-3]
	}
	bagian = append([]string{s}, bagian...)
	hasil := strings.Join(bagian, ".")
	if negatif {
		return "-" + hasil
	}
	return hasil
}

type profilTT struct {
	Nick     string
	Avatar   string
	Bio      string
	Verified bool
	Private  bool
	Heart    int64
	HeartAda bool
	Follower int64
	FollAda  bool
	Follwing int64
	FingAda  bool
	Video    []any
}

func ambilProfilTT(user string) *profilTT {
	body, err := lib.AmbilURL("https://www.tiktok.com/embed/@"+user, map[string]string{
		"User-Agent": "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Mobile Safari/537.36",
	})
	if err != nil {
		return nil
	}
	re := regexp.MustCompile(`(?s)<script id="__FRONTITY_CONNECT_STATE__"[^>]*>(.+?)</script>`)
	cocok := re.FindSubmatch(body)
	if cocok == nil {
		return nil
	}
	var state struct {
		Source struct {
			Data map[string]json.RawMessage `json:"data"`
		} `json:"source"`
	}
	if json.Unmarshal(cocok[1], &state) != nil {
		return nil
	}
	for _, mentah := range state.Source.Data {
		var halaman struct {
			VideoList []struct {
				ID          string `json:"id"`
				PrivateItem bool   `json:"privateItem"`
			} `json:"videoList"`
			UserInfo *struct {
				Nickname       string `json:"nickname"`
				AvatarThumbURL string `json:"avatarThumbUrl"`
				Signature      string `json:"signature"`
				Verified       bool   `json:"verified"`
				PrivateAccount bool   `json:"privateAccount"`
				HeartCount     *int64 `json:"heartCount"`
				FollowerCount  *int64 `json:"followerCount"`
				FollowingCount *int64 `json:"followingCount"`
			} `json:"userInfo"`
		}
		if json.Unmarshal(mentah, &halaman) != nil || halaman.UserInfo == nil {
			continue
		}
		if !strings.Contains(string(mentah), "\"videoList\"") {
			continue
		}
		p := &profilTT{
			Nick:     halaman.UserInfo.Nickname,
			Avatar:   halaman.UserInfo.AvatarThumbURL,
			Bio:      halaman.UserInfo.Signature,
			Verified: halaman.UserInfo.Verified,
			Private:  halaman.UserInfo.PrivateAccount,
		}
		if p.Nick == "" {
			p.Nick = user
		}
		if halaman.UserInfo.HeartCount != nil {
			p.Heart, p.HeartAda = *halaman.UserInfo.HeartCount, true
		}
		if halaman.UserInfo.FollowerCount != nil {
			p.Follower, p.FollAda = *halaman.UserInfo.FollowerCount, true
		}
		if halaman.UserInfo.FollowingCount != nil {
			p.Follwing, p.FingAda = *halaman.UserInfo.FollowingCount, true
		}
		for _, v := range halaman.VideoList {
			if v.ID != "" && !v.PrivateItem {
				p.Video = append(p.Video, v)
			}
		}
		return p
	}
	return nil
}

func infoTiktok(k *lib.Konteks, user string) error {
	var hasil *profilTT
	var gagal any
	for i := 0; i < 4 && (hasil == nil || len(hasil.Video) == 0); i++ {
		if i > 0 {
			time.Sleep(3 * time.Second)
		}
		hasil = ambilProfilTT(user)
		if hasil == nil {
			gagal = "gagal"
		} else {
			gagal = nil
		}
	}
	if hasil == nil || len(hasil.Video) == 0 {
		alasan := fmt.Sprintf("TikTok @%s lagi labil, ulangi nanti", user)
		if gagal == 400 {
			alasan = fmt.Sprintf("Akun @%s gak ada di tiktok", user)
		}
		return k.Conn.SendMessage(k.M.From, lib.IsiPesan{Text: alasan}, k.M)
	}
	status := "Publik 🌍"
	if hasil.Private {
		status = "Privat 🔒"
	}
	verif := "Tidak"
	if hasil.Verified {
		verif = "Ya ✅"
	}
	teks := fmt.Sprintf("「 *INFORMASI AKUN TIKTOK* 」\n- ۶ৎ  Nickname : *%s*\n- ۶ৎ  Username : *@%s*\n- ۶ৎ  Bio : %s\n- ۶ৎ  Status : *%s*\n- ۶ৎ  Verifikasi : *%s*\n- ۶ৎ  Total Suka : *%s*\n- ۶ৎ  Pengikut : *%s*\n- ۶ৎ  Mengikuti : *%s*",
		hasil.Nick, user, hasil.Bio, status, verif,
		formatRibuan(hasil.Heart, hasil.HeartAda),
		formatRibuan(hasil.Follower, hasil.FollAda),
		formatRibuan(hasil.Follwing, hasil.FingAda))
	if hasil.Avatar != "" {
		b, err := lib.AmbilURL(hasil.Avatar, nil)
		if err == nil {
			return k.Conn.SendMessage(k.M.From, lib.IsiPesan{Image: b, Caption: teks}, k.M)
		}
	}
	return k.Conn.SendMessage(k.M.From, lib.IsiPesan{Text: teks}, k.M)
}

type tokenTT struct {
	Type     string `json:"type"`
	Filename string `json:"filename"`
}

func bacaTokenTT(u string) *tokenTT {
	parsed, err := url.Parse(u)
	if err != nil {
		return nil
	}
	bagian := strings.Split(parsed.Query().Get("token"), ".")
	if len(bagian) < 2 {
		return nil
	}
	mentah, err := base64.RawURLEncoding.DecodeString(bagian[1])
	if err != nil {
		mentah, err = base64.URLEncoding.DecodeString(bagian[1])
		if err != nil {
			return nil
		}
	}
	var tok tokenTT
	if json.Unmarshal(mentah, &tok) != nil {
		return nil
	}
	return &tok
}

func kanonikTT(u string) string {
	asli := u
	if !strings.Contains(u, "tiktok.com/@") {
		req, err := http.NewRequest("HEAD", u, nil)
		if err == nil {
			if resp, err := lib.KlienHTTP.Do(req); err == nil {
				asli = resp.Request.URL.String()
				resp.Body.Close()
			}
		}
	}
	parsed, err := url.Parse(asli)
	if err != nil {
		return asli
	}
	return parsed.Scheme + "://" + parsed.Host + strings.Replace(parsed.Path, "/photo/", "/video/", 1)
}

type hasilMusical struct {
	Username   string
	Desc       string
	TitleAudio string
	Thumbnail  string
	Video      []string
	Photo      []string
	Audio      []string
}

func musicalTT(linktt string) (hasilMusical, error) {
	hasil := hasilMusical{Video: []string{}, Photo: []string{}, Audio: []string{}}
	ua := "Mozilla/5.0 (Linux Android 9 CPH1923) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.62 Mobile Safari/537.36"
	req, err := http.NewRequest("GET", "https://musicaldown.com/id", nil)
	if err != nil {
		return hasil, err
	}
	req.Header.Set("User-Agent", ua)
	resp, err := lib.KlienHTTP.Do(req)
	if err != nil {
		return hasil, err
	}
	mentah, _ := io.ReadAll(resp.Body)
	cookie := strings.Join(resp.Header.Values("Set-Cookie"), "; ")
	resp.Body.Close()
	doc := parseHTML(string(mentah))
	if doc == nil {
		return hasil, fmt.Errorf("gagal parse halaman musicaldown")
	}
	type inputForm struct{ nama, nilai string }
	var inputs []inputForm
	for _, form := range cariNode(doc, func(n *html.Node) bool { return n.Data == "form" }) {
		for _, in := range cariNode(form, func(n *html.Node) bool { return n.Data == "input" }) {
			inputs = append(inputs, inputForm{attrNode(in, "name"), attrNode(in, "value")})
		}
		break
	}
	if len(inputs) < 3 {
		return hasil, fmt.Errorf("form musicaldown berubah")
	}
	link, err := url.Parse(linktt)
	if err != nil {
		return hasil, err
	}
	dataKirim := fmt.Sprintf("%s=%s&%s=%s&%s=%s",
		inputs[0].nama, url.QueryEscape(link.Scheme+"://"+link.Host+link.Path),
		inputs[1].nama, inputs[1].nilai,
		inputs[2].nama, inputs[2].nilai)
	reqPost, err := http.NewRequest("POST", "https://musicaldown.com/id/download", strings.NewReader(dataKirim))
	if err != nil {
		return hasil, err
	}
	reqPost.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	reqPost.Header.Set("User-Agent", ua)
	reqPost.Header.Set("Accept", "text/html,application/xhtml+xml,application/xmlq=0.9,image/avif,image/webp,image/apng,*/*q=0.8,application/signed-exchangev=b3q=0.9")
	reqPost.Header.Set("Origin", "https://musicaldown.com")
	reqPost.Header.Set("Referer", "https://musicaldown.com/id")
	reqPost.Header.Set("Cookie", cookie)
	respPost, err := lib.KlienHTTP.Do(reqPost)
	if err != nil {
		return hasil, err
	}
	mentahPost, _ := io.ReadAll(respPost.Body)
	respPost.Body.Close()
	reqMp3, err := http.NewRequest("POST", "https://musicaldown.com/id/mp3", nil)
	if err != nil {
		return hasil, err
	}
	reqMp3.Header.Set("User-Agent", ua)
	reqMp3.Header.Set("Cookie", cookie)
	respMp3, err := lib.KlienHTTP.Do(reqMp3)
	if err != nil {
		return hasil, err
	}
	mentahMp3, _ := io.ReadAll(respMp3.Body)
	respMp3.Body.Close()
	dok := parseHTML(string(mentahPost))
	dokMp3 := parseHTML(string(mentahMp3))
	if dok == nil || dokMp3 == nil {
		return hasil, fmt.Errorf("gagal parse hasil musicaldown")
	}
	for _, a := range cariNode(dok, func(n *html.Node) bool {
		return n.Data == "a" && kelasAda(n, "btn", "waves-effect", "waves-light", "orange", "download")
	}) {
		if href := attrNode(a, "href"); href != "" {
			hasil.Video = append(hasil.Video, href)
		}
	}
	for _, img := range cariNode(dok, func(n *html.Node) bool {
		return n.Data == "img"
	}) {
		induk := img.Parent
		diCard := false
		for x := induk; x != nil; x = x.Parent {
			if x.Type == html.ElementNode && x.Data == "div" && kelasAda(x, "card-image") {
				diCard = true
				break
			}
		}
		if diCard {
			if src := attrNode(img, "src"); src != "" {
				hasil.Photo = append(hasil.Photo, src)
			}
		}
		if kelasAda(img, "responsive-img") && hasil.Thumbnail == "" {
			hasil.Thumbnail = attrNode(img, "src")
		}
	}
	for _, a := range cariNode(dokMp3, func(n *html.Node) bool {
		return n.Data == "a" && kelasAda(n, "waves-effect")
	}) {
		if href := attrNode(a, "href"); href != "" {
			hasil.Audio = append(hasil.Audio, href)
		}
	}
	if len(hasil.Photo) > 0 {
		for _, a := range cariNode(dok, func(n *html.Node) bool {
			return n.Data == "a" && kelasAda(n, "btn", "waves-effect", "waves-light", "orange", "download")
		}) {
			diWelcome := false
			for x := a.Parent; x != nil; x = x.Parent {
				if x.Type == html.ElementNode && x.Data == "div" && kelasAda(x, "welcome", "section") {
					diWelcome = true
					break
				}
			}
			if diWelcome {
				if href := attrNode(a, "href"); href != "" {
					hasil.Audio = append(hasil.Audio, href)
				}
				break
			}
		}
	}
	for _, h := range cariNode(dok, func(n *html.Node) bool { return n.Data == "h2" }) {
		teks := teksNode(h)
		if hasil.Desc == "" {
			hasil.Desc = strings.TrimSpace(teks)
		}
		for _, b := range cariNode(h, func(n *html.Node) bool { return n.Data == "b" }) {
			if hasil.Username == "" {
				hasil.Username = strings.TrimSpace(teksNode(b))
			}
			break
		}
		break
	}
	for _, t := range cariNode(dokMp3, func(n *html.Node) bool { return n.Data == "title" }) {
		hasil.TitleAudio = strings.Split(teksNode(t), " |")[0]
		break
	}
	var videoBersih []string
	for _, v := range hasil.Video {
		if strings.Contains(v, "type=mp3") {
			continue
		}
		videoBersih = append(videoBersih, v)
	}
	hasil.Video = videoBersih
	var audioBersih []string
	for _, a := range hasil.Audio {
		if strings.HasPrefix(a, "https://") {
			audioBersih = append(audioBersih, a)
		}
	}
	hasil.Audio = audioBersih
	return hasil, nil
}

func AmbilDataTT(u string) dataTT {
	var tikwm struct {
		Msg  string `json:"msg"`
		Data struct {
			Region string   `json:"region"`
			Title  string   `json:"title"`
			Play   string   `json:"play"`
			Music  string   `json:"music"`
			Images []string `json:"images"`
			Author struct {
				Nickname string `json:"nickname"`
			} `json:"author"`
			MusicInfo struct {
				Title string `json:"title"`
			} `json:"music_info"`
		} `json:"data"`
	}
	if lib.AmbilJSON("https://tikwm.com/api/?url="+u, nil, &tikwm) == nil && tikwm.Msg == "success" {
		d := tikwm.Data
		return dataTT{
			Region:   d.Region,
			Nickname: d.Author.Nickname,
			Title:    d.Title,
			Music:    d.MusicInfo.Title,
			MusicURL: d.Music,
			Play:     d.Play,
			Images:   d.Images,
		}
	}
	md, _ := musicalTT(u)
	var ombed struct {
		AuthorName string `json:"author_name"`
		Title      string `json:"title"`
	}
	_ = lib.AmbilJSON("https://www.tiktok.com/oembed?url="+url.QueryEscape(kanonikTT(u)), nil, &ombed)
	type isiVideo struct {
		u   string
		tok *tokenTT
	}
	var isi []isiVideo
	for _, v := range md.Video {
		isi = append(isi, isiVideo{v, bacaTokenTT(v)})
	}
	var video []isiVideo
	var lagu *isiVideo
	for i := range isi {
		if isi[i].tok != nil && isi[i].tok.Type == "mp3" {
			if lagu == nil {
				klon := isi[i]
				lagu = &klon
			}
			continue
		}
		video = append(video, isi[i])
	}
	hasil := dataTT{
		Nickname: ombed.AuthorName,
		Title:    ombed.Title,
		Images:   md.Photo,
	}
	if hasil.Nickname == "" {
		hasil.Nickname = md.Username
	}
	if hasil.Title == "" {
		hasil.Title = md.Desc
	}
	if lagu != nil {
		hasil.MusicURL = lagu.u
		if lagu.tok != nil {
			hasil.Music = lagu.tok.Filename
		}
	}
	var terpilih *isiVideo
	for i := range video {
		if video[i].tok != nil && strings.Contains(video[i].tok.Filename, "[HD]") {
			terpilih = &video[i]
			break
		}
	}
	if terpilih == nil && len(video) > 0 {
		terpilih = &video[0]
	}
	if terpilih != nil {
		hasil.Play = terpilih.u
	}
	return hasil
}

func TanganiTiktok(k *lib.Konteks) error {
	if k.M.Q == "" {
		return k.Conn.OK(k.M, "Link tiktoknya mana?")
	}
	if strings.HasPrefix(k.M.Q, "@") {
		user := strings.ToLower(strings.TrimPrefix(strings.Fields(k.M.Q)[0], "@"))
		return infoTiktok(k, user)
	}
	link := polaTiktok.FindString(k.M.Q)
	if link == "" {
		if f := lib.Semua()["ttsearch.js"]; f != nil && f.Jalankan != nil {
			return f.Jalankan(k)
		}
		return nil
	}
	if !strings.HasPrefix(link, "http") {
		link = "https://" + link
	}
	tii := AmbilDataTT(link)
	if tii.Play == "" && len(tii.Images) == 0 {
		return k.Conn.OK(k.M, "Gagal ambil videonya, coba lagi atau cek linknya")
	}
	caption := ""
	if tii.Region != "" {
		caption += fmt.Sprintf("Di upload dari %s %s\n\n", tii.Region, emojiFlag(tii.Region))
	}
	caption += tii.Nickname + "\n" + tii.Title
	if tii.Music != "" {
		caption += "\n----------\n🎵 " + tii.Music
	}
	if len(tii.Images) > 0 {
		if err := k.Conn.SendMessage(k.M.From, lib.IsiPesan{Text: caption}, k.M); err != nil {
			return err
		}
		for _, img := range tii.Images {
			time.Sleep(500 * time.Millisecond)
			if err := kirimGambarURL(k, img, ""); err != nil {
				return err
			}
		}
		return nil
	}
	return kirimVideoURL(k, tii.Play, caption)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "tiktok.js",
		Jenis:    "downloader",
		Nama:     []string{"#tiktok"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)((tikok|tt|tiktok)(mp3|audio|mp4)?)$"),
		Jalankan: TanganiTiktok,
	})
}
