package downloader

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"golang.org/x/net/html"

	"gow1/lib"
)

var polaIG = regexp.MustCompile(`(?i)^https?://(?:www\.)?instagram\.com/(?:.*/)?(?:p|tv|reel|reels|stories)/([^/?#&]+)`)
var polaCekIG = regexp.MustCompile(`(?i)(www)?(.instagram.com)`)
var polaSC = regexp.MustCompile(`/(p|reel|reels|tv)/([A-Za-z0-9_-]+)`)
var polaStory = regexp.MustCompile(`/stories/[^/]+/(\d+)`)

const uaIG = "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Mobile Safari/537.36"

const abcIG = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

func shortcodeKeMediaID(sc string) string {
	n := big.NewInt(0)
	for _, c := range sc {
		n.Mul(n, big.NewInt(64))
		n.Add(n, big.NewInt(int64(strings.IndexRune(abcIG, c))))
	}
	return n.String()
}

type mediaIG struct {
	URL string
}

type hasilIG struct {
	Photos []mediaIG
	Videos []mediaIG
}

func kosongIG() hasilIG {
	return hasilIG{Photos: []mediaIG{}, Videos: []mediaIG{}}
}

func isiIG(h hasilIG) bool {
	return len(h.Photos) > 0 || len(h.Videos) > 0
}

func panenIG(doc *html.Node, cocok func(*html.Node) bool, ambilTipe func(*html.Node, *html.Node) string) hasilIG {
	hasil := kosongIG()
	for _, e := range cariNode(doc, cocok) {
		var a *html.Node
		for _, x := range cariNode(e, func(n *html.Node) bool { return n.Data == "a" }) {
			a = x
			break
		}
		if a == nil {
			continue
		}
		href := attrNode(a, "href")
		if href == "" {
			continue
		}
		if ambilTipe(e, a) == "image" {
			hasil.Photos = append(hasil.Photos, mediaIG{URL: href})
		} else {
			hasil.Videos = append(hasil.Videos, mediaIG{URL: href})
		}
	}
	return hasil
}

func snapinstaIG(u string) hasilIG {
	hasil := kosongIG()
	kepala := map[string]string{"User-Agent": uaIG, "Referer": "https://snapinsta.to/en2"}
	body, err := lib.KirimForm("https://snapinsta.to/api/userverify", url.Values{"url": {u}}, kepala)
	if err != nil {
		return hasil
	}
	var rv struct {
		Token string `json:"token"`
	}
	if json.Unmarshal(body, &rv) != nil || rv.Token == "" {
		return hasil
	}
	body, err = lib.KirimForm("https://snapinsta.to/api/ajaxSearch", url.Values{
		"q": {u}, "t": {"media"}, "v": {"v2"}, "lang": {"en"}, "cftoken": {rv.Token},
	}, kepala)
	if err != nil {
		return hasil
	}
	var rs struct {
		Data string `json:"data"`
	}
	if json.Unmarshal(body, &rs) != nil {
		return hasil
	}
	htmlMentah := rs.Data
	if strings.HasPrefix(htmlMentah, "var") {
		htmlMentah = innerHTMLApaAja(htmlMentah)
	}
	if strings.TrimSpace(htmlMentah) == "" {
		return hasil
	}
	doc := parseHTML(htmlMentah)
	if doc == nil {
		return hasil
	}
	return panenIG(doc, func(n *html.Node) bool {
		return kelasAda(n, "download-items__btn")
	}, func(e, a *html.Node) string {
		if strings.Contains(strings.ToLower(attrNode(a, "title")), "photo") {
			return "image"
		}
		return "video"
	})
}

func snapsaveIG(u string) hasilIG {
	hasil := kosongIG()
	body, err := lib.KirimForm("https://snapsave.app/action.php?lang=en", url.Values{"url": {u}}, map[string]string{
		"User-Agent": uaIG, "Referer": "https://snapsave.app/",
	})
	if err != nil {
		return hasil
	}
	skrip := string(body)
	if !strings.HasPrefix(skrip, "var ") {
		return hasil
	}
	args, ok := getEncodedSnapApp(skrip)
	if !ok {
		return hasil
	}
	js := decodeSnapApp(args)
	htmlMentah := ambilInnerHTML(js, "download-section")
	if strings.TrimSpace(htmlMentah) == "" {
		return hasil
	}
	doc := parseHTML(htmlMentah)
	if doc == nil {
		return hasil
	}
	return panenIG(doc, func(n *html.Node) bool {
		return kelasAda(n, "download-items__btn")
	}, func(e, a *html.Node) string {
		teks := ""
		for _, s := range cariNode(e, func(n *html.Node) bool { return n.Data == "span" }) {
			teks = teksNode(s)
			break
		}
		if strings.Contains(strings.ToLower(teks), "photo") {
			return "image"
		}
		return "video"
	})
}

func downloadgramIG(u string) hasilIG {
	hasil := kosongIG()
	body, err := lib.KirimForm("https://api.downloadgram.org/media", url.Values{
		"url": {u}, "v": {"3"}, "lang": {"en"},
	}, map[string]string{"User-Agent": uaIG})
	if err != nil {
		return hasil
	}
	htmlMentah := string(body)
	if strings.Contains(htmlMentah, "getElementById") {
		htmlMentah = innerHTMLApaAja(htmlMentah)
	}
	if strings.TrimSpace(htmlMentah) == "" {
		return hasil
	}
	doc := parseHTML(htmlMentah)
	if doc == nil {
		return hasil
	}
	for _, e := range cariNode(doc, func(n *html.Node) bool {
		return n.Data == "video" || n.Data == "img"
	}) {
		isVideo := e.Data == "video"
		src := ""
		if isVideo {
			for _, s := range cariNode(e, func(n *html.Node) bool { return n.Data == "source" }) {
				src = attrNode(s, "src")
				break
			}
			if src == "" {
				src = attrNode(e, "poster")
			}
		} else {
			src = attrNode(e, "src")
		}
		src = strings.NewReplacer("\\", "", "\"", "").Replace(src)
		if src == "" {
			continue
		}
		if isVideo {
			hasil.Videos = append(hasil.Videos, mediaIG{URL: src})
		} else {
			hasil.Photos = append(hasil.Photos, mediaIG{URL: src})
		}
	}
	return hasil
}

type itemIG struct {
	CarouselMedia []itemIG `json:"carousel_media"`
	VideoVersions []struct {
		URL string `json:"url"`
	} `json:"video_versions"`
	ImageVersions2 struct {
		Candidates []struct {
			URL string `json:"url"`
		} `json:"candidates"`
	} `json:"image_versions2"`
}

func apiLangsungIG(u, sc, story string) hasilIG {
	hasil := kosongIG()
	sid := os.Getenv("IG_SESSION_ID")
	if sid == "" {
		return hasil
	}
	uid := os.Getenv("IG_DS_USER_ID")
	if uid == "" {
		uid = regexp.MustCompile(`%3A|:`).Split(sid, 2)[0]
	}
	mid := story
	if mid == "" {
		mid = shortcodeKeMediaID(sc)
	}
	var resp struct {
		Items []itemIG `json:"items"`
	}
	err := lib.AmbilJSON(fmt.Sprintf("https://www.instagram.com/api/v1/media/%s/info/", mid), map[string]string{
		"User-Agent":  uaIG,
		"X-IG-App-ID": "936619743392459",
		"Cookie":      fmt.Sprintf("sessionid=%s; ds_user_id=%s", sid, uid),
		"Referer":     "https://www.instagram.com/",
	}, &resp)
	if err != nil || len(resp.Items) == 0 {
		return hasil
	}
	item := resp.Items[0]
	carousel := item.CarouselMedia
	if len(carousel) == 0 {
		carousel = []itemIG{item}
	}
	for _, c := range carousel {
		if len(c.VideoVersions) > 0 {
			hasil.Videos = append(hasil.Videos, mediaIG{URL: c.VideoVersions[0].URL})
		} else if len(c.ImageVersions2.Candidates) > 0 {
			hasil.Photos = append(hasil.Photos, mediaIG{URL: c.ImageVersions2.Candidates[0].URL})
		}
	}
	return hasil
}

func apiFallbackIG(u string) hasilIG {
	hasil := kosongIG()
	var vkr struct {
		Status string `json:"status"`
		Data   []struct {
			URL  string `json:"url"`
			Type string `json:"type"`
		} `json:"data"`
	}
	if lib.AmbilJSON("https://api.vkrdown.com/api/v1/dl?url="+url.QueryEscape(u), nil, &vkr) == nil {
		for _, d := range vkr.Data {
			if strings.Contains(d.Type, "video") || strings.Contains(d.URL, ".mp4") {
				hasil.Videos = append(hasil.Videos, mediaIG{URL: d.URL})
			} else if d.URL != "" {
				hasil.Photos = append(hasil.Photos, mediaIG{URL: d.URL})
			}
		}
	}
	return hasil
}

type galatIG struct {
	Pesan string
}

func (e *galatIG) Error() string { return e.Pesan }

func Igdl(u string) (hasilIG, error) {
	if u == "" {
		return kosongIG(), &galatIG{Pesan: "URL Required"}
	}
	if strings.Contains(u, "/share/") {
		req, err := http.NewRequest("HEAD", u, nil)
		if err == nil {
			if resp, err := lib.KlienHTTP.Do(req); err == nil {
				u = resp.Request.URL.String()
				resp.Body.Close()
			}
		}
	}
	if !polaIG.MatchString(u) {
		return kosongIG(), &galatIG{Pesan: "Invalid URL"}
	}
	sc := ""
	if cocok := polaSC.FindStringSubmatch(u); cocok != nil {
		sc = cocok[2]
	}
	story := ""
	if cocok := polaStory.FindStringSubmatch(u); cocok != nil {
		story = cocok[1]
	}
	if sc == "" && story == "" {
		return kosongIG(), &galatIG{Pesan: "Shortcode nggak ketemu di URL"}
	}
	for _, jalur := range []func(string, string, string) hasilIG{
		func(a, b, c string) hasilIG { return snapinstaIG(a) },
		func(a, b, c string) hasilIG { return snapsaveIG(a) },
		func(a, b, c string) hasilIG { return downloadgramIG(a) },
		func(a, b, c string) hasilIG { return apiFallbackIG(a) },
		apiLangsungIG,
	} {
		hasil := jalur(u, sc, story)
		if isiIG(hasil) {
			return hasil, nil
		}
	}
	return kosongIG(), &galatIG{Pesan: "Semua jalur IG mati, postingannya private atau IG_SESSION_ID udah expired"}
}

func TanganiInstagram(k *lib.Konteks) error {
	if !polaCekIG.MatchString(k.M.Q) {
		return k.Conn.OK(k.M, "Format link Instagram tidak sesuai")
	}
	idl, err := Igdl(k.M.Q)
	if err != nil {
		return err
	}
	for _, p := range idl.Photos {
		if err := kirimGambarURL(k, p.URL, ""); err != nil {
			return err
		}
	}
	for _, p := range idl.Videos {
		if err := kirimVideoURL(k, p.URL, ""); err != nil {
			return err
		}
	}
	return nil
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "instagram.js",
		Jenis:    "downloader",
		Nama:     []string{"#instagram"},
		Itunya:   regexp.MustCompile("(?i)^(#|!|\\/|\\.)(instagram|ig|insta)(dl)?$"),
		Jalankan: TanganiInstagram,
	})
}
