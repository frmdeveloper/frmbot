package info

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"gow1/lib"
)

const uaCariTT = "Mozilla/5.0 (Linux; Android 13; SM-A175F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0 Mobile Safari/537.36"
const uaMusical = "Mozilla/5.0 (Linux Android 9 CPH1923) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.62 Mobile Safari/537.36"

var reURLTT = regexp.MustCompile(`https://www\.tiktok\.com/@[A-Za-z0-9._-]{2,30}/(video|photo)/[0-9]{15,}`)

func cariTiktok(kata string) ([]string, error) {
	target := "https://search.brave.com/search?q=" + url.QueryEscape("site:tiktok.com "+kata)
	cmd := exec.Command("curl", "-sL", "--http2", "-A", uaCariTT, "-H", "Accept: */*", target)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	mentah := reURLTT.FindAllString(string(out), -1)
	var hasil []string
	udah := map[string]bool{}
	for _, u := range mentah {
		if !udah[u] {
			udah[u] = true
			hasil = append(hasil, u)
		}
	}
	return hasil, nil
}

type dataTT struct {
	Region   string
	Nickname string
	Title    string
	Music    string
	MusicURL string
	Play     string
	Images   []string
}

type tikwmResp struct {
	Msg  string `json:"msg"`
	Data struct {
		Region string `json:"region"`
		Author struct {
			Nickname string `json:"nickname"`
		} `json:"author"`
		Title     string `json:"title"`
		MusicInfo struct {
			Title string `json:"title"`
		} `json:"music_info"`
		Music  string   `json:"music"`
		Play   string   `json:"play"`
		Images []string `json:"images"`
	} `json:"data"`
}

func ambilDataTT(link string) dataTT {
	var tw tikwmResp
	if err := lib.AmbilJSON("https://tikwm.com/api/?url="+link, nil, &tw); err == nil && tw.Msg == "success" {
		return dataTT{
			Region:   tw.Data.Region,
			Nickname: tw.Data.Author.Nickname,
			Title:    tw.Data.Title,
			Music:    tw.Data.MusicInfo.Title,
			MusicURL: tw.Data.Music,
			Play:     tw.Data.Play,
			Images:   tw.Data.Images,
		}
	}
	md, _ := musicalTT(link)
	var ombed struct {
		AuthorName string `json:"author_name"`
		Title      string `json:"title"`
	}
	if u, err := kanonikTT(link); err == nil {
		_ = lib.AmbilJSON("https://www.tiktok.com/oembed?url="+url.QueryEscape(u), nil, &ombed)
	}
	type isiVideo struct {
		u   string
		tok map[string]any
	}
	var isi []isiVideo
	for _, u := range md.Video {
		isi = append(isi, isiVideo{u, bacaTokenTT(u)})
	}
	var video []isiVideo
	var lagu *isiVideo
	for i := range isi {
		if isi[i].tok["type"] == "mp3" {
			if lagu == nil {
				lagu = &isi[i]
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
		if fn, ok := lagu.tok["filename"].(string); ok {
			hasil.Music = fn
		}
		hasil.MusicURL = lagu.u
	}
	if len(video) > 0 {
		pilih := video[0]
		for _, v := range video {
			if fn, ok := v.tok["filename"].(string); ok && strings.Contains(fn, "[HD]") {
				pilih = v
				break
			}
		}
		hasil.Play = pilih.u
	}
	return hasil
}

func bacaTokenTT(u string) map[string]any {
	parsed, err := url.Parse(u)
	if err != nil {
		return nil
	}
	bagian := strings.Split(parsed.Query().Get("token"), ".")
	if len(bagian) < 2 {
		return nil
	}
	s := bagian[1]
	for len(s)%4 != 0 {
		s += "="
	}
	dec, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return nil
	}
	var hasil map[string]any
	if err := json.Unmarshal(dec, &hasil); err != nil {
		return nil
	}
	return hasil
}

func kanonikTT(link string) (string, error) {
	if strings.Contains(link, "tiktok.com/@") {
		return link, nil
	}
	req, err := http.NewRequest("HEAD", link, nil)
	if err != nil {
		return "", err
	}
	resp, err := lib.KlienHTTP.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	akhir := resp.Request.URL
	return akhir.Scheme + "://" + akhir.Host + strings.Replace(akhir.Path, "/photo/", "/video/", 1), nil
}

type musicalHasil struct {
	Username string
	Desc     string
	Video    []string
	Audio    []string
	Photo    []string
}

var (
	reInputTag  = regexp.MustCompile(`<input[^>]*>`)
	reNamaAttr  = regexp.MustCompile(`name="([^"]*)"`)
	reValueAttr = regexp.MustCompile(`value="([^"]*)"`)
	reATag      = regexp.MustCompile(`<a\s[^>]*>`)
	reHrefAttr  = regexp.MustCompile(`href="([^"]+)"`)
	reClassAttr = regexp.MustCompile(`class="([^"]*)"`)
	reFotoCard  = regexp.MustCompile(`(?s)<div class="card-image">\s*<img[^>]+src="([^"]+)"`)
	reUserH2    = regexp.MustCompile(`(?s)<div class="col s12">\s*<h2>\s*<b>(.*?)</b>`)
	reDescH2    = regexp.MustCompile(`(?s)<div class="col s12">\s*<h2>(.*?)</h2>`)
	reTagHTML   = regexp.MustCompile(`<[^>]+>`)
)

func stripTagHTML(s string) string {
	return strings.TrimSpace(reTagHTML.ReplaceAllString(s, ""))
}

func musicalTT(linktt string) (musicalHasil, error) {
	var hasil musicalHasil
	parsed, err := url.Parse(linktt)
	if err != nil {
		return hasil, err
	}
	reqAwal, err := http.NewRequest("GET", "https://musicaldown.com/id", nil)
	if err != nil {
		return hasil, err
	}
	reqAwal.Header.Set("User-Agent", uaMusical)
	respAwal, err := lib.KlienHTTP.Do(reqAwal)
	if err != nil {
		return hasil, err
	}
	bAwal, _ := io.ReadAll(respAwal.Body)
	respAwal.Body.Close()
	var kue []string
	for _, c := range respAwal.Header.Values("Set-Cookie") {
		kue = append(kue, strings.Split(c, ";")[0])
	}
	cookie := strings.Join(kue, "; ")
	type inputForm struct{ nama, nilai string }
	var inputs []inputForm
	for _, tag := range reInputTag.FindAllString(string(bAwal), -1) {
		nm := reNamaAttr.FindStringSubmatch(tag)
		if nm == nil {
			continue
		}
		in := inputForm{nama: nm[1]}
		if vl := reValueAttr.FindStringSubmatch(tag); vl != nil {
			in.nilai = vl[1]
		}
		inputs = append(inputs, in)
	}
	if len(inputs) < 3 {
		return hasil, errMusicalKurangInput
	}
	isiForm := inputs[0].nama + "=" + url.QueryEscape(parsed.Scheme+"://"+parsed.Host+parsed.Path) +
		"&" + inputs[1].nama + "=" + inputs[1].nilai +
		"&" + inputs[2].nama + "=" + inputs[2].nilai
	reqDl, err := http.NewRequest("POST", "https://musicaldown.com/id/download", strings.NewReader(isiForm))
	if err != nil {
		return hasil, err
	}
	reqDl.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	reqDl.Header.Set("User-Agent", uaMusical)
	reqDl.Header.Set("Accept", "text/html,application/xhtml+xml,application/xmlq=0.9,image/avif,image/webp,image/apng,*/*q=0.8,application/signed-exchangev=b3q=0.9")
	reqDl.Header.Set("Origin", "https://musicaldown.com")
	reqDl.Header.Set("Referer", "https://musicaldown.com/id")
	reqDl.Header.Set("Cookie", cookie)
	respDl, err := lib.KlienHTTP.Do(reqDl)
	if err != nil {
		return hasil, err
	}
	bDl, _ := io.ReadAll(respDl.Body)
	respDl.Body.Close()
	htmlDl := string(bDl)

	reqMp3, err := http.NewRequest("POST", "https://musicaldown.com/id/mp3", nil)
	if err != nil {
		return hasil, err
	}
	reqMp3.Header.Set("User-Agent", uaMusical)
	reqMp3.Header.Set("Cookie", cookie)
	respMp3, err := lib.KlienHTTP.Do(reqMp3)
	if err != nil {
		return hasil, err
	}
	bMp3, _ := io.ReadAll(respMp3.Body)
	respMp3.Body.Close()

	for _, tag := range reATag.FindAllString(htmlDl, -1) {
		cls := reClassAttr.FindStringSubmatch(tag)
		if cls == nil || cls[1] != "btn waves-effect waves-light orange download" {
			continue
		}
		if hr := reHrefAttr.FindStringSubmatch(tag); hr != nil {
			hasil.Video = append(hasil.Video, hr[1])
		}
	}
	for _, f := range reFotoCard.FindAllStringSubmatch(htmlDl, -1) {
		hasil.Photo = append(hasil.Photo, f[1])
	}
	for _, tag := range reATag.FindAllString(string(bMp3), -1) {
		cls := reClassAttr.FindStringSubmatch(tag)
		if cls == nil || !strings.Contains(cls[1], "waves-effect") {
			continue
		}
		if hr := reHrefAttr.FindStringSubmatch(tag); hr != nil {
			hasil.Audio = append(hasil.Audio, hr[1])
		}
	}
	if mu := reUserH2.FindStringSubmatch(htmlDl); mu != nil {
		hasil.Username = stripTagHTML(mu[1])
	}
	if md := reDescH2.FindStringSubmatch(htmlDl); md != nil {
		hasil.Desc = stripTagHTML(md[1])
	}
	return hasil, nil
}

var errMusicalKurangInput = &musicalError{"input form musicaldown kurang dari 3"}

type musicalError struct{ s string }

func (e *musicalError) Error() string { return e.s }

func TanganiTtsearch(k *lib.Konteks) error {
	m := k.M
	if m.Q == "" {
		return k.Conn.OK(m, "Mau nyari apaan?")
	}
	hasil, err := cariTiktok(m.Q)
	if err != nil {
		return err
	}
	if len(hasil) == 0 {
		return k.Conn.OK(m, "Gak nemu videonya njir")
	}
	teratas := hasil
	if len(teratas) > 10 {
		teratas = teratas[:10]
	}
	link := teratas[rand.Intn(len(teratas))]
	tii := ambilDataTT(link)
	if tii.Play == "" && len(tii.Images) == 0 {
		return k.Conn.OK(m, "Nemu videonya tapi gagal diambil, coba lagi")
	}
	caption := ""
	if tii.Region != "" {
		caption += "Di upload dari " + tii.Region + " " + emojiFlag(tii.Region) + "\n\n"
	}
	caption += tii.Nickname + "\n" + tii.Title
	if tii.Music != "" {
		caption += "\n----------\n🎵 " + tii.Music
	}
	caption += "\n\n" + link
	if len(tii.Images) > 0 {
		if err := k.Conn.SendMessage(m.From, lib.IsiPesan{Text: caption}, m); err != nil {
			return err
		}
		for _, u := range tii.Images {
			time.Sleep(500 * time.Millisecond)
			b, err := lib.AmbilURL(u, nil)
			if err != nil {
				continue
			}
			_ = k.Conn.SendMessage(m.From, lib.IsiPesan{Image: b}, m)
		}
		return nil
	}
	b, err := lib.AmbilURL(tii.Play, nil)
	if err != nil {
		return k.Conn.OK(m, "Nemu videonya tapi gagal diambil, coba lagi")
	}
	return k.Conn.SendMessage(m.From, lib.IsiPesan{Video: b, Caption: caption}, m)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "ttsearch.js",
		Jenis:    "info",
		Nama:     []string{"#ttsearch"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(ttsearch|tiktoksearch|tiktokcari|caritiktok)$"),
		Jalankan: TanganiTtsearch,
	})
}
