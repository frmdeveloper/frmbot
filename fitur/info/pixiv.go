package info

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"gow1/lib"

	"github.com/coder/websocket"
	"golang.org/x/net/proxy"
)

const uaPixiv = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0 Safari/537.36"

var (
	errTagTerlarangPixiv = errors.New("TAG_TERLARANG")
	errR18Pixiv          = errors.New("R18")
	errUgoiraPixiv       = errors.New("UGOIRA")
	polaTagTerlarang     = regexp.MustCompile(`(?i)ロリ|ロリコン|ショタ|ショタコン|幼女|幼男|児童|ランドセル|小学生|中学生|女子高生|高校生|loli|lolicon|shota|shotacon|underage`)
	setTagTerlarang      = map[string]bool{"js": true, "jc": true, "jk": true}
	reLinkPixiv          = regexp.MustCompile(`(?i)pixiv\.net/(?:[a-z-]+/)?artworks/(\d+)`)
	reIDPixiv            = regexp.MustCompile(`^\d{6,}$`)
	reDomainPixiv        = regexp.MustCompile(`(^|\.)pixiv\.net$`)
)

const maksHalamanPixiv = 4
const maksCobaCariPixiv = 6
const riwayatTTLPixiv = 6 * 60 * 60 * 1000

type idPixiv string

func (i *idPixiv) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	*i = idPixiv(s)
	return nil
}

func fileCookiePixiv() string {
	if v := os.Getenv("PIXIV_COOKIE_FILE"); v != "" {
		return v
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "pixiv-cookie.txt")
}

var (
	muCookiePixiv  sync.Mutex
	cookiePixiv    string
	cookiePixivAda bool
	cookieMtime    int64
	sesiValidSd    int64
	muSegarPixiv   sync.Mutex
)

func ambilCookiePixiv() string {
	muCookiePixiv.Lock()
	defer muCookiePixiv.Unlock()
	st, err := os.Stat(fileCookiePixiv())
	if err != nil {
		cookiePixiv, cookiePixivAda = "", true
		return ""
	}
	if cookiePixivAda && st.ModTime().UnixMilli() == cookieMtime {
		return cookiePixiv
	}
	b, err := os.ReadFile(fileCookiePixiv())
	if err != nil {
		cookiePixiv = ""
	} else {
		cookiePixiv = strings.TrimSpace(string(b))
	}
	cookiePixivAda = true
	cookieMtime = st.ModTime().UnixMilli()
	return cookiePixiv
}

func headJSONPixiv() map[string]string {
	h := map[string]string{"User-Agent": uaPixiv, "Referer": "https://www.pixiv.net/", "Accept": "application/json"}
	if c := ambilCookiePixiv(); c != "" {
		h["Cookie"] = c
	}
	return h
}

var klienPixivLangsung = &http.Client{Timeout: 30 * time.Second}

var klienPixivWarp = func() *http.Client {
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:40000", nil, proxy.Direct)
	if err != nil {
		return klienPixivLangsung
	}
	tr := &http.Transport{}
	if cd, ok := dialer.(proxy.ContextDialer); ok {
		tr.DialContext = cd.DialContext
	} else {
		tr.Dial = dialer.Dial
	}
	return &http.Client{Timeout: 30 * time.Second, Transport: tr}
}()

func pixivGetSekali(klien *http.Client, url string, headers map[string]string) ([]byte, string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, "", err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := klien.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}
	if resp.StatusCode >= 400 {
		return nil, "", fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	return body, resp.Header.Get("Content-Type"), nil
}

func pixivGet(url string, headers map[string]string) ([]byte, string, error) {
	body, ct, err := pixivGetSekali(klienPixivLangsung, url, headers)
	if err == nil {
		return body, ct, nil
	}
	body2, ct2, err2 := pixivGetSekali(klienPixivWarp, url, headers)
	if err2 != nil {
		return nil, "", err
	}
	return body2, ct2, nil
}

func sesiLoginAktifPixiv(cookie string) bool {
	if cookie == "" {
		return false
	}
	body, _, err := pixivGetSekali(klienPixivLangsung, "https://www.pixiv.net/ajax/user/self", map[string]string{"User-Agent": uaPixiv, "Referer": "https://www.pixiv.net/", "Accept": "application/json", "Cookie": cookie})
	if err != nil {
		return false
	}
	var hasil struct {
		UserData struct {
			ID json.RawMessage `json:"id"`
		} `json:"userData"`
	}
	if json.Unmarshal(body, &hasil) != nil {
		return false
	}
	id := strings.Trim(string(hasil.UserData.ID), `"`)
	return id != "" && id != "0" && id != "null"
}

func cdpJSONPixiv(path string, out any) error {
	klien := &http.Client{Timeout: 5 * time.Second}
	resp, err := klien.Get("http://127.0.0.1:9223" + path)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("CDP HTTP %d", resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

func tungguCdpPixiv() error {
	akhir := time.Now().Add(20 * time.Second)
	for time.Now().Before(akhir) {
		var v map[string]any
		if err := cdpJSONPixiv("/json/version", &v); err == nil {
			return nil
		}
		time.Sleep(time.Second)
	}
	return fmt.Errorf("CDP chrome pixiv gak nyala")
}

func nyalakanChromePixiv() error {
	var v map[string]any
	if err := cdpJSONPixiv("/json/version", &v); err == nil {
		return nil
	}
	home, _ := os.UserHomeDir()
	profil := filepath.Join(home, ".config", "chrome-pixiv")
	display := os.Getenv("DISPLAY")
	if display == "" {
		display = ":1"
	}
	cmd := exec.Command("google-chrome", "--headless=new", "--user-data-dir="+profil, "--remote-debugging-address=127.0.0.1", "--remote-debugging-port=9223", "--no-first-run", "--no-default-browser-check", "https://www.pixiv.net/en/")
	env := os.Environ()
	adaDisplay := false
	for _, e := range env {
		if strings.HasPrefix(e, "DISPLAY=") {
			adaDisplay = true
		}
	}
	if !adaDisplay {
		env = append(env, "DISPLAY="+display)
	}
	cmd.Env = env
	_ = cmd.Start()
	return tungguCdpPixiv()
}

func ambilCookieDariChromePixiv() (string, error) {
	if err := nyalakanChromePixiv(); err != nil {
		return "", err
	}
	var targets []struct {
		Type                 string `json:"type"`
		URL                  string `json:"url"`
		WebSocketDebuggerUrl string `json:"webSocketDebuggerUrl"`
	}
	if err := cdpJSONPixiv("/json/list", &targets); err != nil {
		return "", err
	}
	wsURL := ""
	for _, t := range targets {
		if t.Type == "page" && strings.Contains(t.URL, "pixiv.net") && t.WebSocketDebuggerUrl != "" {
			wsURL = t.WebSocketDebuggerUrl
			break
		}
	}
	if wsURL == "" {
		for _, t := range targets {
			if t.Type == "page" && t.WebSocketDebuggerUrl != "" {
				wsURL = t.WebSocketDebuggerUrl
				break
			}
		}
	}
	if wsURL == "" {
		return "", fmt.Errorf("tab chrome pixiv gak ada")
	}
	ctx, batal := context.WithTimeout(context.Background(), 20*time.Second)
	defer batal()
	ws, _, err := websocket.Dial(ctx, wsURL, nil)
	if err != nil {
		return "", err
	}
	defer ws.Close(websocket.StatusNormalClosure, "")
	ws.SetReadLimit(4 << 20)
	if err := ws.Write(ctx, websocket.MessageText, []byte(`{"id":1,"method":"Storage.getCookies","params":{}}`)); err != nil {
		return "", err
	}
	for {
		_, data, err := ws.Read(ctx)
		if err != nil {
			return "", err
		}
		var j struct {
			ID     int `json:"id"`
			Result struct {
				Cookies []struct {
					Name   string `json:"name"`
					Value  string `json:"value"`
					Domain string `json:"domain"`
				} `json:"cookies"`
			} `json:"result"`
		}
		if json.Unmarshal(data, &j) != nil || j.ID != 1 {
			continue
		}
		var px []string
		adaSesi := false
		for _, c := range j.Result.Cookies {
			if !reDomainPixiv.MatchString(c.Domain) {
				continue
			}
			if c.Name == "PHPSESSID" {
				adaSesi = true
			}
			px = append(px, c.Name+"="+c.Value)
		}
		if !adaSesi {
			return "", fmt.Errorf("PHPSESSID gak ketemu di chrome")
		}
		return strings.Join(px, "; "), nil
	}
}

func segarkanCookiePixiv() bool {
	muSegarPixiv.Lock()
	defer muSegarPixiv.Unlock()
	header, err := ambilCookieDariChromePixiv()
	if err != nil {
		fmt.Println("[pixiv] gagal segarkan cookie:", err)
		return false
	}
	if !sesiLoginAktifPixiv(header) {
		fmt.Println("[pixiv] profil chrome pixiv belum login")
		return false
	}
	tujuan := fileCookiePixiv()
	if err := os.WriteFile(tujuan, []byte(header+"\n"), 0o600); err != nil {
		fmt.Println("[pixiv] gagal segarkan cookie:", err)
		return false
	}
	_ = os.Chmod(tujuan, 0o600)
	muCookiePixiv.Lock()
	cookiePixiv = header
	cookiePixivAda = true
	if st, err := os.Stat(tujuan); err == nil {
		cookieMtime = st.ModTime().UnixMilli()
	}
	muCookiePixiv.Unlock()
	sesiValidSd = time.Now().UnixMilli() + 30*60*1000
	fmt.Println("[pixiv] cookie sesi disegarkan dari profil chrome")
	return true
}

func pastikanLoginPixiv() bool {
	if time.Now().UnixMilli() < sesiValidSd && ambilCookiePixiv() != "" {
		return true
	}
	c := ambilCookiePixiv()
	if c != "" && sesiLoginAktifPixiv(c) {
		sesiValidSd = time.Now().UnixMilli() + 30*60*1000
		return true
	}
	return segarkanCookiePixiv()
}

func adaTagTerlarangPixiv(tags []string) bool {
	for _, t := range tags {
		if polaTagTerlarang.MatchString(t) || setTagTerlarang[strings.ToLower(t)] {
			return true
		}
	}
	return false
}

type itemCariPixiv struct {
	ID         idPixiv  `json:"id"`
	IllustType int      `json:"illustType"`
	XRestrict  int      `json:"xRestrict"`
	AiType     int      `json:"aiType"`
	Tags       []string `json:"tags"`
}

type detailPixiv struct {
	ID          idPixiv `json:"id"`
	IllustID    idPixiv `json:"illustId"`
	Title       string  `json:"title"`
	IllustTitle string  `json:"illustTitle"`
	UserName    string  `json:"userName"`
	XRestrict   int     `json:"xRestrict"`
	IllustType  int     `json:"illustType"`
	PageCount   int     `json:"pageCount"`
	AiType      int     `json:"aiType"`
	URLs        struct {
		Regular  string `json:"regular"`
		Original string `json:"original"`
	} `json:"urls"`
	Tags struct {
		Tags []struct {
			Tag string `json:"tag"`
		} `json:"tags"`
	} `json:"tags"`
}

func cariPixiv(kata string) ([]itemCariPixiv, error) {
	login := pastikanLoginPixiv()
	var mentah []itemCariPixiv
	sudah := map[string]bool{}
	for hal := 1; hal <= 3; hal++ {
		params := url.Values{}
		params.Set("word", kata)
		params.Set("order", "date_d")
		if login {
			params.Set("mode", "all")
		} else {
			params.Set("mode", "safe")
		}
		params.Set("p", strconv.Itoa(hal))
		params.Set("s_mode", "s_tag")
		params.Set("type", "all")
		params.Set("lang", "en")
		u := "https://www.pixiv.net/ajax/search/artworks/" + url.PathEscape(kata) + "?" + params.Encode()
		body, _, err := pixivGet(u, headJSONPixiv())
		if err != nil {
			return nil, err
		}
		var hasil struct {
			Error   bool   `json:"error"`
			Message string `json:"message"`
			Body    struct {
				IllustManga struct {
					Data []itemCariPixiv `json:"data"`
				} `json:"illustManga"`
			} `json:"body"`
		}
		if err := json.Unmarshal(body, &hasil); err != nil {
			return nil, err
		}
		if hasil.Error {
			return nil, fmt.Errorf("%s", hasil.Message)
		}
		batch := hasil.Body.IllustManga.Data
		for _, x := range batch {
			id := string(x.ID)
			if id != "" && !sudah[id] {
				sudah[id] = true
				mentah = append(mentah, x)
			}
		}
		if len(batch) < 60 {
			break
		}
	}
	var lolos []itemCariPixiv
	for _, x := range mentah {
		if x.IllustType == 2 {
			continue
		}
		if login {
			if x.XRestrict != 0 && x.XRestrict != 1 && x.XRestrict != 2 {
				continue
			}
		} else if x.XRestrict != 0 {
			continue
		}
		if adaTagTerlarangPixiv(x.Tags) {
			continue
		}
		lolos = append(lolos, x)
	}
	return lolos, nil
}

func detailPixivAmbil(id string) (*detailPixiv, error) {
	body, _, err := pixivGet("https://www.pixiv.net/ajax/illust/"+id+"?lang=en", headJSONPixiv())
	if err != nil {
		return nil, err
	}
	var hasil struct {
		Error bool        `json:"error"`
		Body  detailPixiv `json:"body"`
	}
	if err := json.Unmarshal(body, &hasil); err != nil {
		return nil, err
	}
	if hasil.Error {
		return nil, fmt.Errorf("karya gak ketemu")
	}
	return &hasil.Body, nil
}

type berkasPixiv struct {
	buf  []byte
	mime string
}

func unduhGambarPixiv(u string) (berkasPixiv, error) {
	body, ct, err := pixivGet(u, map[string]string{"User-Agent": uaPixiv, "Referer": "https://www.pixiv.net/"})
	if err != nil {
		return berkasPixiv{}, err
	}
	mime := strings.Split(ct, ";")[0]
	if mime == "" {
		mime = "image/jpeg"
	}
	return berkasPixiv{buf: body, mime: mime}, nil
}

func ambilKaryaPixiv(id string, maksHalaman int) (*detailPixiv, []berkasPixiv, []int, error) {
	login := pastikanLoginPixiv()
	d, err := detailPixivAmbil(id)
	if err != nil {
		return nil, nil, nil, err
	}
	var tagsDetail []string
	for _, t := range d.Tags.Tags {
		if t.Tag != "" {
			tagsDetail = append(tagsDetail, t.Tag)
		}
	}
	if adaTagTerlarangPixiv(tagsDetail) {
		return nil, nil, nil, errTagTerlarangPixiv
	}
	if d.XRestrict != 0 && !login {
		return nil, nil, nil, errR18Pixiv
	}
	if d.IllustType == 2 {
		return nil, nil, nil, errUgoiraPixiv
	}
	var urls []string
	if u := d.URLs.Regular; u != "" {
		urls = append(urls, u)
	} else if u := d.URLs.Original; u != "" {
		urls = append(urls, u)
	}
	if d.PageCount > 1 {
		body, _, err := pixivGet("https://www.pixiv.net/ajax/illust/"+id+"/pages?lang=en", headJSONPixiv())
		if err == nil {
			var hasil struct {
				Body []struct {
					URLs struct {
						Regular  string `json:"regular"`
						Original string `json:"original"`
					} `json:"urls"`
				} `json:"body"`
			}
			if json.Unmarshal(body, &hasil) == nil && len(hasil.Body) > 0 {
				var hal []string
				for _, p := range hasil.Body {
					u := p.URLs.Regular
					if u == "" {
						u = p.URLs.Original
					}
					if u != "" {
						hal = append(hal, u)
					}
				}
				if len(hal) > 0 {
					urls = hal
				}
			}
		}
	}
	idx := make([]int, len(urls))
	for i := range idx {
		idx[i] = i
	}
	if len(idx) > maksHalaman {
		rand.Shuffle(len(idx), func(a, b int) { idx[a], idx[b] = idx[b], idx[a] })
		idx = idx[:maksHalaman]
		sort.Ints(idx)
	}
	var berkas []berkasPixiv
	for _, i := range idx {
		b, err := unduhGambarPixiv(urls[i])
		if err != nil {
			return nil, nil, nil, err
		}
		berkas = append(berkas, b)
	}
	if len(berkas) == 0 {
		return nil, nil, nil, fmt.Errorf("gambar kosong")
	}
	halaman := make([]int, len(idx))
	for i, v := range idx {
		halaman[i] = v + 1
	}
	return d, berkas, halaman, nil
}

type entriRiwayatPixiv struct {
	ids []string
	exp int64
}

var (
	muRiwayatPixiv sync.Mutex
	riwayatPixiv   = map[string]*entriRiwayatPixiv{}
)

func kirimKaryaPixiv(k *lib.Konteks, d *detailPixiv, berkas []berkasPixiv, konteks string, halaman []int) error {
	m := k.M
	idD := string(d.ID)
	if idD == "" {
		idD = string(d.IllustID)
	}
	judul := d.Title
	if judul == "" {
		judul = d.IllustTitle
	}
	fmt.Printf("[pixiv] kirim id=%s R%d ai=%d %.40s (%s)\n", idD, d.XRestrict, d.AiType, judul, konteks)
	if judul == "" {
		judul = "-"
	}
	artis := d.UserName
	if artis == "" {
		artis = "-"
	}
	caption := "🎨 " + judul + "\n👤 " + artis + "\nhttps://www.pixiv.net/en/artworks/" + idD
	if len(halaman) > 0 && d.PageCount > 1 {
		bagian := make([]string, len(halaman))
		for i, h := range halaman {
			bagian[i] = strconv.Itoa(h)
		}
		caption += "\n📄 halaman " + strings.Join(bagian, ", ") + " dari " + strconv.Itoa(d.PageCount)
	}
	for i, b := range berkas {
		isi := lib.IsiPesan{Image: b.buf, Mimetype: b.mime}
		if i == 0 {
			isi.Caption = caption
		}
		if err := k.Conn.SendMessage(m.From, isi, m); err != nil {
			return err
		}
	}
	return nil
}

func TanganiPixiv(k *lib.Konteks) error {
	m := k.M
	if strings.TrimSpace(m.Q) == "" {
		return k.Conn.OK(m, "Cari gambar Pixiv atau kirim link/ID karyanya\nContoh:\n#pixiv elaina\n#pixiv https://www.pixiv.net/en/artworks/151511113")
	}
	id := ""
	if cocok := reLinkPixiv.FindStringSubmatch(m.Q); cocok != nil {
		id = cocok[1]
	} else if reIDPixiv.MatchString(strings.TrimSpace(m.Q)) {
		id = strings.TrimSpace(m.Q)
	}
	if id != "" {
		d, berkas, halaman, err := ambilKaryaPixiv(id, maksHalamanPixiv)
		if err != nil {
			fmt.Println("[pixiv] ambil karya gagal:", err)
			switch {
			case errors.Is(err, errTagTerlarangPixiv):
				return k.Conn.OK(m, "Karya itu gak bisa gue ambil — ada tag yg gue blokir permanen")
			case errors.Is(err, errR18Pixiv):
				return k.Conn.OK(m, "Itu karya 18+ — sesi login Pixiv bot lagi gak aktif")
			case errors.Is(err, errUgoiraPixiv):
				return k.Conn.OK(m, "Itu ugoira (animasi), belum didukung")
			}
			return k.Conn.OK(m, "Gagal ngambil karyanya 😔 Link/ID-nya bener gak?")
		}
		return kirimKaryaPixiv(k, d, berkas, "link/id", halaman)
	}
	hasil, err := cariPixiv(m.Q)
	if err != nil {
		fmt.Println("[pixiv] search gagal:", err)
		return k.Conn.OK(m, "Search Pixiv lagi error, coba lagi ntar")
	}
	if len(hasil) == 0 {
		return k.Conn.OK(m, fmt.Sprintf("Gak nemu gambar buat \"%s\" di Pixiv", m.Q))
	}
	susun := func(arr []itemCariPixiv) []itemCariPixiv {
		var r18, aman []itemCariPixiv
		for _, x := range arr {
			if x.XRestrict > 0 {
				r18 = append(r18, x)
			} else {
				aman = append(aman, x)
			}
		}
		rand.Shuffle(len(r18), func(a, b int) { r18[a], r18[b] = r18[b], r18[a] })
		rand.Shuffle(len(aman), func(a, b int) { aman[a], aman[b] = aman[b], aman[a] })
		return append(r18, aman...)
	}
	var manusia, ai []itemCariPixiv
	for _, x := range hasil {
		if x.AiType == 2 {
			ai = append(ai, x)
		} else {
			manusia = append(manusia, x)
		}
	}
	gabungan := append(susun(manusia), susun(ai)...)
	kunciRiwayat := m.From + "|" + strings.ToLower(strings.TrimSpace(m.Q))
	muRiwayatPixiv.Lock()
	rw := riwayatPixiv[kunciRiwayat]
	var idBarusan []string
	if rw != nil && rw.exp > time.Now().UnixMilli() {
		idBarusan = rw.ids
	}
	muRiwayatPixiv.Unlock()
	baruSet := map[string]bool{}
	for _, s := range idBarusan {
		baruSet[s] = true
	}
	var urutan []itemCariPixiv
	for _, x := range gabungan {
		if !baruSet[string(x.ID)] {
			urutan = append(urutan, x)
		}
	}
	for _, x := range gabungan {
		if baruSet[string(x.ID)] {
			urutan = append(urutan, x)
		}
	}
	batas := maksCobaCariPixiv
	if len(urutan) < batas {
		batas = len(urutan)
	}
	for _, item := range urutan[:batas] {
		d, berkas, halaman, err := ambilKaryaPixiv(string(item.ID), 1)
		if err != nil {
			fmt.Println("[pixiv] kandidat gagal:", item.ID, err)
			continue
		}
		var idsBaru []string
		for _, s := range idBarusan {
			if s != string(item.ID) {
				idsBaru = append(idsBaru, s)
			}
		}
		idsBaru = append(idsBaru, string(item.ID))
		if len(idsBaru) > 12 {
			idsBaru = idsBaru[len(idsBaru)-12:]
		}
		muRiwayatPixiv.Lock()
		riwayatPixiv[kunciRiwayat] = &entriRiwayatPixiv{ids: idsBaru, exp: time.Now().UnixMilli() + riwayatTTLPixiv}
		muRiwayatPixiv.Unlock()
		return kirimKaryaPixiv(k, d, berkas, fmt.Sprintf("search \"%s\"", m.Q), halaman)
	}
	return k.Conn.OK(m, "Gagal ngambil gambarnya 😔 coba kata kunci lain")
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "pixiv.js",
		Jenis:    "info",
		Nama:     []string{"#pixiv"},
		Itunya:   regexp.MustCompile(`(?i)^(#|!|/|\.)pixiv$`),
		Jalankan: TanganiPixiv,
	})
}
