package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var KlienHTTP = &http.Client{Timeout: 60 * time.Second}

func NomorSaja(jid string) string {
	jid = strings.Split(jid, "@")[0]
	jid = strings.Split(jid, ":")[0]
	return jid
}

func SpawnSync(perintah string) (string, string) {
	cmd := exec.Command("sh", "-c", perintah)
	var out, errb bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errb
	_ = cmd.Run()
	return out.String(), errb.String()
}

func AmbilURL(alamat string, header map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", alamat, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0 Safari/537.36")
	for k, v := range header {
		req.Header.Set(k, v)
	}
	resp, err := KlienHTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP %d dari %s", resp.StatusCode, alamat)
	}
	return io.ReadAll(resp.Body)
}

func AmbilJSON(alamat string, header map[string]string, target any) error {
	b, err := AmbilURL(alamat, header)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, target)
}

func KirimForm(alamat string, form url.Values, header map[string]string) ([]byte, error) {
	req, err := http.NewRequest("POST", alamat, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0 Safari/537.36")
	for k, v := range header {
		req.Header.Set(k, v)
	}
	resp, err := KlienHTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func Gtr(bahasa, teks string) (string, error) {
	u := "https://translate.googleapis.com/translate_a/single?client=gtx&sl=auto&tl=" +
		url.QueryEscape(bahasa) + "&dt=t&q=" + url.QueryEscape(teks)
	b, err := AmbilURL(u, nil)
	if err != nil {
		return "", err
	}
	var hasil []any
	if err := json.Unmarshal(b, &hasil); err != nil {
		return "", err
	}
	if len(hasil) == 0 {
		return "", fmt.Errorf("hasil translate kosong")
	}
	segmen, _ := hasil[0].([]any)
	var sb strings.Builder
	for _, s := range segmen {
		bagian, _ := s.([]any)
		if len(bagian) > 0 {
			t, _ := bagian[0].(string)
			sb.WriteString(t)
		}
	}
	return sb.String(), nil
}

func DeteksiMime(data []byte) string {
	switch {
	case len(data) >= 3 && data[0] == 0xFF && data[1] == 0xD8 && data[2] == 0xFF:
		return "image/jpeg"
	case len(data) >= 8 && string(data[1:4]) == "PNG":
		return "image/png"
	case len(data) >= 12 && string(data[0:4]) == "RIFF" && string(data[8:12]) == "WEBP":
		return "image/webp"
	case len(data) >= 6 && (string(data[0:6]) == "GIF89a" || string(data[0:6]) == "GIF87a"):
		return "image/gif"
	case len(data) >= 8 && string(data[4:8]) == "ftyp":
		return "video/mp4"
	case len(data) >= 4 && string(data[0:4]) == "OggS":
		return "audio/ogg"
	case len(data) >= 3 && string(data[0:3]) == "ID3":
		return "audio/mpeg"
	case len(data) >= 2 && data[0] == 0xFF && data[1]&0xE0 == 0xE0:
		return "audio/mpeg"
	case len(data) >= 4 && string(data[0:4]) == "%PDF":
		return "application/pdf"
	default:
		return "application/octet-stream"
	}
}

func EkstensiMime(mime string) string {
	switch mime {
	case "image/jpeg":
		return ".jpg"
	case "image/png":
		return ".png"
	case "image/webp":
		return ".webp"
	case "image/gif":
		return ".gif"
	case "video/mp4":
		return ".mp4"
	case "audio/ogg":
		return ".ogg"
	case "audio/mpeg":
		return ".mp3"
	case "application/pdf":
		return ".pdf"
	default:
		return ".bin"
	}
}

func TulisTemp(ext string, data []byte) (string, error) {
	if err := os.MkdirAll(filepath.Join("data", "tmp"), 0o755); err != nil {
		return "", err
	}
	f, err := os.CreateTemp(filepath.Join("data", "tmp"), "gow1-*"+ext)
	if err != nil {
		return "", err
	}
	defer f.Close()
	if _, err := f.Write(data); err != nil {
		return "", err
	}
	return f.Name(), nil
}

func BacaBerkas(path string) ([]byte, error) {
	return os.ReadFile(path)
}
