package downloader

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html"

	"gow1/lib"
)

var polaFB = regexp.MustCompile(`(?i)https?://(?:[\w-]+\.)*(?:facebook\.com|fb\.watch|fb\.me|fb\.gg)/\S+`)

var daftarUA = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/125.0.0.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
}

func uaAcak() string {
	return daftarUA[rand.Intn(len(daftarUA))]
}

func parseHTML(s string) *html.Node {
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		return nil
	}
	return doc
}

func jalanNode(n *html.Node, fn func(*html.Node)) {
	if n == nil {
		return
	}
	fn(n)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		jalanNode(c, fn)
	}
}

func cariNode(akar *html.Node, cocok func(*html.Node) bool) []*html.Node {
	var hasil []*html.Node
	jalanNode(akar, func(n *html.Node) {
		if n.Type == html.ElementNode && cocok(n) {
			hasil = append(hasil, n)
		}
	})
	return hasil
}

func attrNode(n *html.Node, nama string) string {
	for _, a := range n.Attr {
		if a.Key == nama {
			return a.Val
		}
	}
	return ""
}

func kelasAda(n *html.Node, kelas ...string) bool {
	punya := map[string]bool{}
	for _, k := range strings.Fields(attrNode(n, "class")) {
		punya[k] = true
	}
	for _, k := range kelas {
		if !punya[k] {
			return false
		}
	}
	return true
}

func teksNode(n *html.Node) string {
	var sb strings.Builder
	var rek func(*html.Node)
	rek = func(x *html.Node) {
		if x.Type == html.TextNode {
			sb.WriteString(x.Data)
		}
		for c := x.FirstChild; c != nil; c = c.NextSibling {
			rek(c)
		}
	}
	rek(n)
	return sb.String()
}

type hasilFB struct {
	Videos []string
	Photos []string
	Pesan  string
}

func kosongFB() hasilFB {
	return hasilFB{Videos: []string{}, Photos: []string{}}
}

var polaTag = regexp.MustCompile(`<[^>]*>`)
var polaSpasi = regexp.MustCompile(`\s+`)

func bersihTeksFB(s string) string {
	return strings.TrimSpace(polaSpasi.ReplaceAllString(polaTag.ReplaceAllString(s, " "), " "))
}

func Fbdl(u string) hasilFB {
	hasil := getmyfb(u)
	if len(hasil.Videos) > 0 || len(hasil.Photos) > 0 {
		return hasil
	}
	alt := snapsaveFB(u)
	if len(alt.Videos) > 0 || len(alt.Photos) > 0 {
		return alt
	}
	pesan := hasil.Pesan
	if pesan == "" {
		pesan = alt.Pesan
	}
	return hasilFB{Videos: []string{}, Photos: []string{}, Pesan: pesan}
}

func getmyfb(u string) hasilFB {
	hasil := kosongFB()
	form := url.Values{"id": {u}, "locale": {"en"}}
	body, err := lib.KirimForm("https://getmyfb.com/process", form, map[string]string{
		"Origin":         "https://getmyfb.com",
		"Referer":        "https://getmyfb.com/",
		"Hx-Current-Url": "https://getmyfb.com",
		"Hx-Request":     "true",
		"Hx-Target":      "target",
		"Hx-Trigger":     "form",
		"User-Agent":     uaAcak(),
	})
	if err != nil {
		return hasil
	}
	doc := parseHTML(string(body))
	if doc == nil {
		return hasil
	}
	for _, a := range cariNode(doc, func(n *html.Node) bool {
		return n.Data == "a" && kelasAda(n, "photo-button")
	}) {
		if href := attrNode(a, "href"); href != "" {
			hasil.Photos = append(hasil.Photos, href)
		}
	}
	video := ""
	for _, a := range cariNode(doc, func(n *html.Node) bool {
		return n.Data == "a" && kelasAda(n, "hd-button")
	}) {
		video = attrNode(a, "href")
		break
	}
	if video == "" {
		for _, a := range cariNode(doc, func(n *html.Node) bool {
			return n.Data == "a" && kelasAda(n, "sd-button") && !kelasAda(n, "mp3")
		}) {
			video = attrNode(a, "href")
			break
		}
	}
	if video != "" {
		hasil.Videos = append(hasil.Videos, video)
	}
	if len(hasil.Videos) == 0 && len(hasil.Photos) == 0 {
		for _, h := range cariNode(doc, func(n *html.Node) bool { return n.Data == "h4" }) {
			hasil.Pesan = bersihTeksFB(teksNode(h))
			break
		}
	}
	return hasil
}

func snapsaveFB(u string) hasilFB {
	hasil := kosongFB()
	var buf bytes.Buffer
	mw := multipart.NewWriter(&buf)
	_ = mw.WriteField("url", u)
	_ = mw.Close()
	tanggal := strconv.FormatInt(time.Now().UnixMilli(), 10)
	if len(tanggal) > 10 {
		tanggal = tanggal[:10]
	}
	req, err := http.NewRequest("POST", "https://snapsave.app/action.php?lang=en", &buf)
	if err != nil {
		return hasil
	}
	req.Header.Set("Content-Type", mw.FormDataContentType())
	req.Header.Set("Origin", "https://snapsave.app")
	req.Header.Set("Referer", "https://snapsave.app/")
	req.Header.Set("Cookie", fmt.Sprintf("_ga=GA1.1.2035716653.%s; _ga_WNPZGVDWE9=GS1.1.%s.2.1.%s.49.0.0", tanggal, tanggal, tanggal))
	req.Header.Set("User-Agent", uaAcak())
	resp, err := lib.KlienHTTP.Do(req)
	if err != nil {
		return hasil
	}
	defer resp.Body.Close()
	var sb strings.Builder
	bufResp := make([]byte, 65536)
	for {
		n, err := resp.Body.Read(bufResp)
		if n > 0 {
			sb.Write(bufResp[:n])
		}
		if err != nil {
			break
		}
	}
	data := sb.String()
	args, ok := getEncodedSnapApp(data)
	if !ok {
		return hasil
	}
	js := decodeSnapApp(args)
	isi := ambilInnerHTML(js, "download-section")
	if isi == "" {
		hasil.Pesan = bersihTeksFB(ambilInnerHTML(js, "alert"))
		return hasil
	}
	doc := parseHTML(isi)
	if doc == nil {
		return hasil
	}
	type pilihan struct {
		mutu int
		href string
	}
	var daftar []pilihan
	for _, tr := range cariNode(doc, func(n *html.Node) bool { return n.Data == "tr" }) {
		var href string
		for _, a := range cariNode(tr, func(n *html.Node) bool { return n.Data == "a" }) {
			href = attrNode(a, "href")
			break
		}
		if href == "" || strings.Contains(href, "render.php") {
			continue
		}
		mutu := 0
		for _, td := range cariNode(tr, func(n *html.Node) bool {
			return n.Data == "td" && kelasAda(n, "video-quality")
		}) {
			mutu, _ = strconv.Atoi(strings.TrimSpace(teksNode(td)))
			break
		}
		daftar = append(daftar, pilihan{mutu: mutu, href: href})
	}
	for i := 0; i < len(daftar); i++ {
		for j := i + 1; j < len(daftar); j++ {
			if daftar[j].mutu > daftar[i].mutu {
				daftar[i], daftar[j] = daftar[j], daftar[i]
			}
		}
	}
	if len(daftar) > 0 {
		hasil.Videos = append(hasil.Videos, daftar[0].href)
	}
	return hasil
}

var polaInner = regexp.MustCompile(`(?s)innerHTML\s*=\s*"((?:[^"\\]|\\.)*)"`)
var polaLolos = regexp.MustCompile(`\\(.)`)

func ambilInnerHTML(js, id string) string {
	re := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(id) + `"\)\.innerHTML\s*=\s*"((?:[^"\\]|\\.)*)"`)
	cocok := re.FindStringSubmatch(js)
	if cocok == nil {
		return ""
	}
	return polaLolos.ReplaceAllString(cocok[1], "$1")
}

func innerHTMLApaAja(js string) string {
	cocok := polaInner.FindStringSubmatch(js)
	if cocok == nil {
		return ""
	}
	return polaLolos.ReplaceAllString(cocok[1], "$1")
}

func getEncodedSnapApp(data string) ([]string, bool) {
	bagian := strings.Split(data, "decodeURIComponent(escape(r))}(")
	if len(bagian) < 2 {
		return nil, false
	}
	sisa := strings.Split(bagian[1], "))")[0]
	mentah := strings.Split(sisa, ",")
	hasil := make([]string, len(mentah))
	for i, v := range mentah {
		hasil[i] = strings.TrimSpace(strings.ReplaceAll(v, "\"", ""))
	}
	return hasil, true
}

func decodeSnapApp(args []string) string {
	if len(args) < 6 {
		return ""
	}
	h, n := args[0], args[2]
	t, _ := strconv.Atoi(args[3])
	e, _ := strconv.Atoi(args[4])
	decode := func(d string, ee, f int) string {
		g := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/"
		hh := ""
		ii := ""
		if ee >= 0 && ee <= len(g) {
			hh = g[:ee]
		}
		if f >= 0 && f <= len(g) {
			ii = g[:f]
		}
		j := 0.0
		huruf := []rune(d)
		for c := 0; c < len(huruf); c++ {
			b := huruf[len(huruf)-1-c]
			if idx := strings.IndexRune(hh, b); idx != -1 {
				j += float64(idx) * math.Pow(float64(ee), float64(c))
			}
		}
		k := ""
		jj := int64(j)
		for jj > 0 {
			k = string(ii[jj%int64(f)]) + k
			jj = (jj - (jj % int64(f))) / int64(f)
		}
		if k == "" {
			k = "0"
		}
		return k
	}
	hr := []rune(h)
	nr := []rune(n)
	var r strings.Builder
	for i := 0; i < len(hr); i++ {
		s := ""
		for i < len(hr) && hr[i] != nr[e] {
			s += string(hr[i])
			i++
		}
		for j := 0; j < len(nr); j++ {
			s = strings.ReplaceAll(s, string(nr[j]), strconv.Itoa(j))
		}
		angka, _ := strconv.Atoi(decode(s, e, 10))
		r.WriteRune(rune(angka - t))
	}
	return r.String()
}

func TanganiFacebook(k *lib.Konteks) error {
	link := polaFB.FindString(k.M.Q)
	if link == "" {
		return k.Conn.OK(k.M, "Mau download FB malah dikasih apaan")
	}
	hasil := Fbdl(link)
	if len(hasil.Videos) == 0 && len(hasil.Photos) == 0 {
		pesan := hasil.Pesan
		if pesan == "" {
			pesan = "Medianya nggak ketemu, postingannya private atau linknya salah"
		}
		return k.Conn.OK(k.M, pesan)
	}
	for _, u := range hasil.Videos {
		if err := kirimVideoURL(k, u, ""); err != nil {
			return err
		}
	}
	for i, u := range hasil.Photos {
		if i > 0 {
			time.Sleep(500 * time.Millisecond)
		}
		if err := kirimGambarURL(k, u, ""); err != nil {
			return err
		}
	}
	return nil
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "facebook.js",
		Jenis:    "downloader",
		Nama:     []string{"#facebook"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(fb|facebook)$"),
		Jalankan: TanganiFacebook,
	})
}
