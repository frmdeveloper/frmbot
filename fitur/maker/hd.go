package maker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gow1/lib"
)

func dimensiGambar(data []byte) (int, int, error) {
	cfg, _, err := image.DecodeConfig(bytes.NewReader(data))
	if err == nil && cfg.Width > 0 {
		return cfg.Width, cfg.Height, nil
	}
	path, err := lib.TulisTemp(lib.EkstensiMime(lib.DeteksiMime(data)), data)
	if err != nil {
		return 0, 0, err
	}
	defer os.Remove(path)
	out, err := jalanPerintah("ffprobe", "-v", "error", "-select_streams", "v:0", "-show_entries", "stream=width,height", "-of", "csv=s=x:p=0", path)
	if err != nil {
		return 0, 0, err
	}
	bagian := strings.Split(strings.TrimSpace(out), "x")
	if len(bagian) != 2 {
		return 0, 0, fmt.Errorf("dimensi gambar tidak kebaca")
	}
	w, err1 := strconv.Atoi(bagian[0])
	h, err2 := strconv.Atoi(bagian[1])
	if err1 != nil || err2 != nil {
		return 0, 0, fmt.Errorf("dimensi gambar tidak kebaca")
	}
	return w, h, nil
}

func upscaleHD(buffer []byte, size int, anime bool) (string, error) {
	if len(buffer) == 0 {
		return "", fmt.Errorf("undefined buffer input!")
	}
	if size != 2 && size != 4 && size != 6 && size != 8 && size != 16 {
		return "", fmt.Errorf("invalid upscale size!")
	}
	width, height, err := dimensiGambar(buffer)
	if err != nil {
		return "", err
	}
	stempel := strconv.FormatInt(time.Now().UnixMilli(), 10)
	var body bytes.Buffer
	w := multipart.NewWriter(&body)
	_ = w.WriteField("name", "upscale-"+stempel)
	_ = w.WriteField("imageName", "upscale-"+stempel)
	_ = w.WriteField("desiredHeight", strconv.Itoa(height*size))
	_ = w.WriteField("desiredWidth", strconv.Itoa(width*size))
	_ = w.WriteField("outputFormat", "png")
	_ = w.WriteField("compressionLevel", "none")
	_ = w.WriteField("anime", strconv.FormatBool(anime))
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="image_file"; filename="upscale-%s.png"`, stempel))
	h.Set("Content-Type", "image/png")
	part, err := w.CreatePart(h)
	if err != nil {
		return "", err
	}
	if _, err := part.Write(buffer); err != nil {
		return "", err
	}
	if err := w.Close(); err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", "https://api.upscalepics.com/upscale-to-size", &body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Set("Origin", "https://upscalepics.com")
	req.Header.Set("Referer", "https://upscalepics.com")
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
		BgRemoved string          `json:"bgRemoved"`
		Error     json.RawMessage `json:"error"`
	}
	if err := json.Unmarshal(mentah, &data); err != nil {
		return "", err
	}
	if len(data.Error) > 0 && string(data.Error) != "null" && string(data.Error) != `""` && string(data.Error) != "false" {
		return "", fmt.Errorf("something error from upscaler api!")
	}
	return data.BgRemoved, nil
}

func TanganiHd(k *lib.Konteks) error {
	m := k.M
	down := unduhMedia(k)
	if down == nil {
		return k.Conn.OK(m, "mana gambarnya?")
	}
	mime := lib.DeteksiMime(down)
	if !strings.Contains(mime, "image") {
		return k.Conn.OK(m, "Gw cuma bisa foto🙂")
	}
	hddd, err := upscaleHD(down, 2, strings.Contains(m.Q, "anime"))
	if err != nil {
		return err
	}
	hasil, err := lib.AmbilURL(hddd, nil)
	if err != nil {
		return err
	}
	return k.Conn.SendMessage(m.From, lib.IsiPesan{Image: hasil}, m)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "hd.js",
		Jenis:    "maker",
		Nama:     []string{"#hd"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(hd|remini|enhance)$"),
		Jalankan: TanganiHd,
	})
}
