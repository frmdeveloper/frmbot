package converter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"regexp"
	"strings"
	"time"

	"gow1/lib"
)

func kirimMultipart(alamat, field, namaBerkas string, data []byte, headers map[string]string) ([]byte, error) {
	var buf bytes.Buffer
	mw := multipart.NewWriter(&buf)
	part, err := mw.CreateFormFile(field, namaBerkas)
	if err != nil {
		return nil, err
	}
	if _, err := part.Write(data); err != nil {
		return nil, err
	}
	if err := mw.Close(); err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", alamat, &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", mw.FormDataContentType())
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	klien := &http.Client{Timeout: 300 * time.Second}
	resp, err := klien.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

type berkasUguu struct {
	URL string `json:"url"`
}

type hasilUguu struct {
	Files []berkasUguu `json:"files"`
}

func UploadFileUgu(buffer []byte, ext string) (*hasilUguu, error) {
	if ext == "" {
		ext = lib.EkstensiMime(lib.DeteksiMime(buffer))
	}
	body, err := kirimMultipart("https://uguu.se/upload", "files[]", ext, buffer, map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36",
	})
	if err != nil {
		return nil, err
	}
	var hasil hasilUguu
	if err := json.Unmarshal(body, &hasil); err != nil {
		return nil, err
	}
	return &hasil, nil
}

func Uguu(buffer []byte) (string, error) {
	hasil, err := UploadFileUgu(buffer, "")
	if err != nil {
		return "", err
	}
	if len(hasil.Files) == 0 {
		return "", fmt.Errorf("uguu nggak ngasih hasil")
	}
	return hasil.Files[0].URL, nil
}

func TelegraPh(buffer []byte, ext string) (string, error) {
	if ext == "" {
		ext = lib.EkstensiMime(lib.DeteksiMime(buffer))
	}
	body, err := kirimMultipart("https://telegra.ph/upload", "file", ext, buffer, nil)
	if err != nil {
		return "", err
	}
	var hasil []struct {
		Src string `json:"src"`
	}
	if err := json.Unmarshal(body, &hasil); err != nil {
		return "", err
	}
	if len(hasil) == 0 {
		return "", fmt.Errorf("telegra.ph nggak ngasih hasil")
	}
	return "https://telegra.ph" + hasil[0].Src, nil
}

type hasilTmpfile struct {
	URLOriginal string
	URL         string
}

func Tmpfile(buffer []byte, ext string) (*hasilTmpfile, error) {
	if ext == "" {
		ext = strings.TrimPrefix(lib.EkstensiMime(lib.DeteksiMime(buffer)), ".")
	}
	nama := fmt.Sprintf("file-%d.%s", time.Now().UnixMilli(), ext)
	body, err := kirimMultipart("https://tmpfiles.org/api/v1/upload", "file", nama, buffer, map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0 Safari/537.36",
	})
	if err != nil {
		return nil, err
	}
	var mentah struct {
		Data struct {
			URL string `json:"url"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &mentah); err != nil {
		return nil, err
	}
	return &hasilTmpfile{
		URLOriginal: mentah.Data.URL,
		URL:         strings.Replace(mentah.Data.URL, "tmpfiles.org/", "tmpfiles.org/dl/", 1),
	}, nil
}

func TanganiTourl(k *lib.Konteks) error {
	var down []byte
	if b, _, err := k.Conn.DownloadAny(k.M); err == nil {
		down = b
	} else if k.M.Quoted != nil {
		if b2, _, err2 := k.Conn.DownloadAny(k.M.Quoted); err2 == nil {
			down = b2
		}
	}
	if len(down) == 0 {
		return k.Conn.OK(k.M, "mana filenya")
	}
	up, err := UploadFileUgu(down, "")
	if err != nil {
		return err
	}
	urls := make([]string, 0, len(up.Files))
	for _, f := range up.Files {
		urls = append(urls, f.URL)
	}
	return k.Conn.SendMessage(k.M.From, lib.IsiPesan{Text: strings.Join(urls, " ")}, k.M)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "tourl.js",
		Jenis:    "converter",
		Nama:     []string{"#tourl"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(tourl|upload)$"),
		Jalankan: TanganiTourl,
	})
}
