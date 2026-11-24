package utility

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"gow1/lib"
)

var klienRedirect = &http.Client{
	Timeout: 8 * time.Second,
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

type hopRedirect struct {
	alamat string
	status int
}

func ikutiRedirect(awal string, maksHop int) ([]hopRedirect, error) {
	rantai := []hopRedirect{}
	berjalan := awal
	for i := 0; i < maksHop; i++ {
		req, err := http.NewRequest("GET", berjalan, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("User-Agent", "Mozilla/5.0")
		resp, err := klienRedirect.Do(req)
		if err != nil {
			return nil, err
		}
		resp.Body.Close()
		rantai = append(rantai, hopRedirect{alamat: berjalan, status: resp.StatusCode})
		lokasi := resp.Header.Get("Location")
		if lokasi == "" || resp.StatusCode < 300 || resp.StatusCode >= 400 {
			break
		}
		if strings.HasPrefix(lokasi, "http") {
			berjalan = lokasi
		} else {
			basis, err := url.Parse(berjalan)
			if err != nil {
				return nil, err
			}
			ref, err := url.Parse(lokasi)
			if err != nil {
				return nil, err
			}
			berjalan = basis.ResolveReference(ref).String()
		}
	}
	return rantai, nil
}

func TanganiUrlredirect(k *lib.Konteks) error {
	m := k.M
	alamat := strings.TrimSpace(m.Q)
	if alamat == "" || !regexp.MustCompile(`(?i)^https?://`).MatchString(alamat) {
		return k.Conn.SendMessage(m.From, lib.IsiPesan{
			Text: "Kirim URL-nya, contoh:\n*.cekredirect* https://bit.ly/xyz",
		}, m)
	}
	rantai, err := ikutiRedirect(alamat, 10)
	if err != nil {
		return k.Conn.SendMessage(m.From, lib.IsiPesan{
			Text: fmt.Sprintf("Gagal cek redirect:\n%s", err.Error()),
		}, m)
	}
	baris := []string{}
	for i, h := range rantai {
		panah := "↪️"
		if i == len(rantai)-1 {
			panah = "✅"
		}
		baris = append(baris, fmt.Sprintf("%s [%d] %s", panah, h.status, h.alamat))
	}
	hops := len(rantai) - 1
	ringkasan := "Tidak ada redirect. URL langsung ke tujuan."
	if hops > 0 {
		jamak := ""
		if hops > 1 {
			jamak = "s"
		}
		ringkasan = fmt.Sprintf("Total *%d* redirect%s.", hops, jamak)
	}
	teks := fmt.Sprintf("*🔗 URL Redirect Checker*\n\n%s\n\n%s", strings.Join(baris, "\n"), ringkasan)
	return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: teks}, m)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "urlredirect.js",
		Jenis:    "utility",
		Nama:     []string{"#cekredirect"},
		Itunya:   regexp.MustCompile("(?i)^(#|!|\\/|\\.)cekredirect$"),
		Jalankan: TanganiUrlredirect,
	})
}
