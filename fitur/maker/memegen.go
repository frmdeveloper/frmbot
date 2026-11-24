package maker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"

	"gow1/lib"
)

func mp4KeGif(buffer []byte) ([]byte, error) {
	awal, err := lib.TulisTemp(".mp4", buffer)
	if err != nil {
		return nil, err
	}
	akhir := strings.TrimSuffix(awal, ".mp4") + ".gif"
	defer os.Remove(awal)
	defer os.Remove(akhir)
	lib.SpawnSync(fmt.Sprintf("ffmpeg -i %s %s", awal, akhir))
	return os.ReadFile(akhir)
}

func uguuUpload(buffer []byte) (string, error) {
	ext := lib.EkstensiMime(lib.DeteksiMime(buffer))
	var body bytes.Buffer
	w := multipart.NewWriter(&body)
	part, err := w.CreateFormFile("files[]", "."+strings.TrimPrefix(ext, "."))
	if err != nil {
		return "", err
	}
	if _, err := part.Write(buffer); err != nil {
		return "", err
	}
	if err := w.Close(); err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", "https://uguu.se/upload", &body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, err := lib.KlienHTTP.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	mentah, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var data struct {
		Files []struct {
			URL string `json:"url"`
		} `json:"files"`
	}
	if err := json.Unmarshal(mentah, &data); err != nil {
		return "", err
	}
	if len(data.Files) == 0 {
		return "", fmt.Errorf("uguu nggak ngasih url")
	}
	return data.Files[0].URL, nil
}

func TanganiMemegen(k *lib.Konteks) error {
	m := k.M
	down := unduhMedia(k)
	bagian := strings.Split(m.Q, "|")
	var atas, bawah string
	if len(bagian) > 0 {
		atas = bagian[0]
	}
	if len(bagian) > 1 {
		bawah = bagian[1]
	}
	if down == nil || (atas == "" && bawah == "") {
		return k.Conn.OK(m, "Gak gitu caranya")
	}
	mime := lib.DeteksiMime(down)
	ext := strings.TrimPrefix(lib.EkstensiMime(mime), ".")
	if strings.Contains(mime, "image") {
		ext = "jpg"
	}
	if strings.Contains(mime, "webp") {
		ext = "gif"
	}
	if strings.Contains(mime, "video") {
		gif, err := mp4KeGif(down)
		if err != nil {
			return err
		}
		down = gif
		ext = "gif"
	}
	up, err := uguuUpload(down)
	if err != nil {
		return err
	}
	if atas == "" {
		atas = "_"
	}
	if bawah == "" {
		bawah = "_"
	}
	alamat := fmt.Sprintf("https://api.memegen.link/images/custom/%s/%s.%s?background=%s",
		url.PathEscape(atas), url.PathEscape(bawah), ext, url.QueryEscape(up))
	data, err := lib.AmbilURL(alamat, nil)
	if err != nil {
		return err
	}
	var sticker []byte
	if ext == "gif" {
		sticker, err = Mp42webp(data)
	} else {
		sticker, err = Img2webp(data)
	}
	if err != nil {
		return err
	}
	sss, err := AddExifBawaan(sticker)
	if err != nil {
		return err
	}
	simpanStkdb(m, sss)
	return k.Conn.SendMessage(m.From, lib.IsiPesan{Sticker: sss}, m)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "memegen.js",
		Jenis:    "maker",
		Nama:     nil,
		Itunya:   nil,
		Jalankan: TanganiMemegen,
	})
}
