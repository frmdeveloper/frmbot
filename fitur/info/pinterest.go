package info

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gow1/lib"
)

const pinBase = "https://www.pinterest.com"

var pinHeaders = map[string]string{
	"accept":                  "application/json, text/javascript, */*, q=0.01",
	"referer":                 "https://www.pinterest.com/",
	"user-agent":              "Postify/1.0.0",
	"x-app-version":           "a9522f",
	"x-pinterest-appstate":    "active",
	"x-pinterest-pws-handler": "www/[username]/[slug].js",
	"x-pinterest-source-url":  "/search/pins/?rs=typed&q=kucing%20anggora/",
	"x-requested-with":        "XMLHttpRequest",
}

var reLinkPin = regexp.MustCompile(`(?i)https?://(?:www\.)?pinterest\.[a-z.]+/\S+|https?://pin\.it/\S+`)

var klienWARP = &http.Client{
	Timeout: 60 * time.Second,
	Transport: &http.Transport{
		Proxy: func(r *http.Request) (*url.URL, error) {
			return url.Parse("socks5://127.0.0.1:40000")
		},
	},
}

func pinAmbil(alamat string, header map[string]string, params url.Values) ([]byte, http.Header, error) {
	if params != nil {
		alamat = alamat + "?" + params.Encode()
	}
	var errAkhir error
	for _, klien := range []*http.Client{klienWARP, lib.KlienHTTP} {
		req, err := http.NewRequest("GET", alamat, nil)
		if err != nil {
			return nil, nil, err
		}
		for k, v := range header {
			req.Header.Set(k, v)
		}
		resp, err := klien.Do(req)
		if err != nil {
			errAkhir = err
			continue
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			errAkhir = err
			continue
		}
		return body, resp.Header, nil
	}
	return nil, nil, errAkhir
}

func pinCookies() string {
	_, header, err := pinAmbil(pinBase, nil, nil)
	if err != nil || header == nil {
		return ""
	}
	setCookies := header.Values("Set-Cookie")
	if len(setCookies) == 0 {
		return ""
	}
	var bagian []string
	for _, c := range setCookies {
		bagian = append(bagian, strings.TrimSpace(strings.Split(c, ";")[0]))
	}
	return strings.Join(bagian, "; ")
}

type pinGambar struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type pinHasil struct {
	ID    string
	Title string
	URL   string
}

func pinSearch(query string, limit int) []pinHasil {
	if query == "" {
		return nil
	}
	cookies := pinCookies()
	if cookies == "" {
		return nil
	}
	dataJSON, _ := json.Marshal(map[string]any{
		"options": map[string]any{
			"isPrefetch":                   false,
			"query":                        query,
			"scope":                        "pins",
			"bookmarks":                    []string{""},
			"no_fetch_context_on_resource": false,
			"page_size":                    limit,
		},
		"context": map[string]any{},
	})
	params := url.Values{}
	params.Set("source_url", "/search/pins/?q="+query)
	params.Set("data", string(dataJSON))
	params.Set("_", strconv.FormatInt(time.Now().UnixMilli(), 10))
	header := map[string]string{}
	for k, v := range pinHeaders {
		header[k] = v
	}
	header["cookie"] = cookies
	b, _, err := pinAmbil(pinBase+"/resource/BaseSearchResource/get/", header, params)
	if err != nil {
		return nil
	}
	var resp struct {
		ResourceResponse struct {
			Data struct {
				Results []struct {
					ID          string `json:"id"`
					Title       string `json:"title"`
					Description string `json:"description"`
					Images      struct {
						Orig *pinGambar `json:"orig"`
					} `json:"images"`
				} `json:"results"`
			} `json:"data"`
		} `json:"resource_response"`
	}
	if err := json.Unmarshal(b, &resp); err != nil {
		return nil
	}
	var hasil []pinHasil
	for _, r := range resp.ResourceResponse.Data.Results {
		if r.Images.Orig == nil || r.Images.Orig.URL == "" {
			continue
		}
		hasil = append(hasil, pinHasil{ID: r.ID, Title: r.Title, URL: r.Images.Orig.URL})
	}
	if len(hasil) == 0 {
		return nil
	}
	return hasil
}

type kandidatToolPin struct {
	bin string
	pre []string
}

func cobaToolPin(kandidat []kandidatToolPin, args []string) bool {
	for _, knd := range kandidat {
		ctx, batal := context.WithTimeout(context.Background(), 120*time.Second)
		cmd := exec.CommandContext(ctx, knd.bin, append(append([]string{}, knd.pre...), args...)...)
		err := cmd.Run()
		batal()
		if err == nil {
			return true
		}
		var errExec *exec.Error
		if errors.As(err, &errExec) {
			continue
		}
		return false
	}
	return false
}

func galeriDlPin(link, dir string) bool {
	home, _ := os.UserHomeDir()
	return cobaToolPin([]kandidatToolPin{
		{bin: "gallery-dl"},
		{bin: filepath.Join(home, ".local/bin/gallery-dl")},
		{bin: "python3", pre: []string{"-m", "gallery_dl"}},
	}, []string{"--no-part", "-D", dir, link})
}

func ytdlpDlPin(link, dir string) bool {
	home, _ := os.UserHomeDir()
	return cobaToolPin([]kandidatToolPin{
		{bin: "yt-dlp"},
		{bin: filepath.Join(home, ".local/bin/yt-dlp")},
		{bin: "python3", pre: []string{"-m", "yt_dlp"}},
	}, []string{"-o", filepath.Join(dir, "%(id)s.%(ext)s"), "--no-playlist", "-q", link})
}

func unduhPin(k *lib.Konteks, link string) error {
	m := k.M
	dir, err := os.MkdirTemp("", "pin-*")
	if err != nil {
		return k.Conn.OK(m, "Error pas download, coba lagi ntar")
	}
	defer os.RemoveAll(dir)
	if err := k.Conn.OK(m, "⏳ Lagi download dari pinterest..."); err != nil {
		return err
	}
	sukses := galeriDlPin(link, dir)
	if !sukses {
		sukses = ytdlpDlPin(link, dir)
	}
	entri, _ := os.ReadDir(dir)
	var files []string
	for _, e := range entri {
		if !e.IsDir() {
			files = append(files, filepath.Join(dir, e.Name()))
		}
	}
	if len(files) == 0 {
		return k.Conn.OK(m, "Gagal download 😔 Pin nya mungkin udah dihapus, private, atau link nya salah")
	}
	jumlahFile := len(files)
	if len(files) > 10 {
		files = files[:10]
	}
	for i, f := range files {
		buf, err := os.ReadFile(f)
		if err != nil {
			continue
		}
		mime := lib.DeteksiMime(buf)
		caption := ""
		if i == 0 {
			caption = fmt.Sprintf("📌 Pinterest\n%d file terunduh", jumlahFile)
		}
		var errKirim error
		switch {
		case strings.HasPrefix(mime, "video"):
			errKirim = k.Conn.SendMessage(m.From, lib.IsiPesan{Video: buf, Mimetype: mime, Caption: caption}, m)
		case strings.HasPrefix(mime, "image"):
			errKirim = k.Conn.SendMessage(m.From, lib.IsiPesan{Image: buf, Caption: caption}, m)
		default:
			if mime == "" || mime == "application/octet-stream" {
				mime = "application/octet-stream"
			}
			errKirim = k.Conn.SendMessage(m.From, lib.IsiPesan{Document: buf, Mimetype: mime, FileName: filepath.Base(f)}, m)
		}
		if errKirim == nil {
			time.Sleep(700 * time.Millisecond)
		}
	}
	return nil
}

func TanganiPinterest(k *lib.Konteks) error {
	m := k.M
	if m.Q == "" {
		return k.Conn.OK(m, "Kirim query buat cari, atau link pinterest buat download\nContoh:\n#pinterest kucing anggora\n#pinterest https://pin.it/5d9yKmyG3")
	}
	if link := reLinkPin.FindString(m.Q); link != "" {
		return unduhPin(k, link)
	}
	pins := pinSearch(m.Q, 1)
	if len(pins) == 0 {
		return k.Conn.OK(m, "Malas nyari")
	}
	for _, pin := range pins {
		b, err := lib.AmbilURL(pin.URL, nil)
		if err != nil {
			continue
		}
		_ = k.Conn.SendMessage(m.From, lib.IsiPesan{Image: b}, m)
	}
	return nil
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "pinterest.js",
		Jenis:    "info",
		Nama:     []string{"#pinterest"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(pin)(terest)?$"),
		Jalankan: TanganiPinterest,
	})
}
