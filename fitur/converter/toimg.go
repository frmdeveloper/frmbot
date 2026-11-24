package converter

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"gow1/lib"
)

func TanganiToimg(k *lib.Konteks) error {
	var down []byte
	if k.M.Quoted != nil {
		if b, _, err := k.Conn.DownloadAny(k.M.Quoted); err == nil {
			down = b
		}
	}
	ada := len(down) > 0 && lib.DeteksiMime(down) == "image/webp"
	if !ada {
		return k.Conn.OK(k.M, "Hadeh, stikernya mana 😂")
	}
	png, err := Webp2png(down)
	if err != nil {
		return err
	}
	return k.Conn.SendMessage(k.M.From, lib.IsiPesan{Image: png}, k.M)
}

func Webp2png(buffer []byte) ([]byte, error) {
	awal, err := lib.TulisTemp(".webp", buffer)
	if err != nil {
		return nil, err
	}
	defer os.Remove(awal)
	akhir := strings.TrimSuffix(awal, ".webp") + ".png"
	defer os.Remove(akhir)
	lib.SpawnSync(fmt.Sprintf(`convert "%s[0]" "%s"`, awal, akhir))
	if !jadiBerkas(akhir) {
		lib.SpawnSync(fmt.Sprintf(`ffmpeg -i "%s" -frames:v 1 "%s"`, awal, akhir))
	}
	return os.ReadFile(akhir)
}

func Any2mp4(buffer []byte) ([]byte, error) {
	awal, err := lib.TulisTemp("", buffer)
	if err != nil {
		return nil, err
	}
	defer os.Remove(awal)
	akhir := awal + ".mp4"
	defer os.Remove(akhir)
	_, stderr := lib.SpawnSync(fmt.Sprintf(`ffmpeg -i "%s" "%s"`, awal, akhir))
	if !jadiBerkas(akhir) {
		return nil, fmt.Errorf("%s", ekorBaris(stderr, 2))
	}
	return os.ReadFile(akhir)
}

func Mp42gif(buffer []byte) ([]byte, error) {
	awal, err := lib.TulisTemp(".mp4", buffer)
	if err != nil {
		return nil, err
	}
	defer os.Remove(awal)
	akhir := strings.TrimSuffix(awal, ".mp4") + ".gif"
	defer os.Remove(akhir)
	_, stderr := lib.SpawnSync(fmt.Sprintf(`ffmpeg -i "%s" "%s"`, awal, akhir))
	if !jadiBerkas(akhir) {
		return nil, fmt.Errorf("%s", ekorBaris(stderr, 2))
	}
	return os.ReadFile(akhir)
}

var polaFileEzgif = regexp.MustCompile(`<input[^>]*name="file"[^>]*value="([^"]*)"`)
var polaFileEzgifAlt = regexp.MustCompile(`<input[^>]*value="([^"]*)"[^>]*name="file"`)
var polaOutfile = regexp.MustCompile(`(?s)outfile.*?<source[^>]*src="([^"]*)"`)

func Webp2mp4(url string) (string, error) {
	html, err := lib.AmbilURL("https://ezgif.com/webp-to-mp4?url="+url, nil)
	if err != nil {
		return "", err
	}
	file := ""
	if m := polaFileEzgif.FindSubmatch(html); m != nil {
		file = string(m[1])
	} else if m := polaFileEzgifAlt.FindSubmatch(html); m != nil {
		file = string(m[1])
	}
	form := map[string][]string{
		"file":    {file},
		"convert": {"Convert WebP to MP4!"},
	}
	hasil, err := lib.KirimForm("https://ezgif.com/webp-to-mp4/"+file, form, nil)
	if err != nil {
		return "", err
	}
	m := polaOutfile.FindSubmatch(hasil)
	if m == nil {
		return "", fmt.Errorf("link hasil ezgif nggak ketemu")
	}
	return "https:" + string(m[1]), nil
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "toimg.js",
		Jenis:    "converter",
		Nama:     []string{"#toimg"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)to(img|image)$"),
		Jalankan: TanganiToimg,
	})
}
