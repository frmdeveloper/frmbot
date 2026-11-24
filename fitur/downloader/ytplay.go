package downloader

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"

	"gow1/lib"
)

type videoYT struct {
	Title string
	URL   string
}

func cariRenderer(x any) *videoYT {
	switch v := x.(type) {
	case map[string]any:
		if vr, ok := v["videoRenderer"].(map[string]any); ok {
			id, _ := vr["videoId"].(string)
			judul := ""
			if t, ok := vr["title"].(map[string]any); ok {
				if runs, ok := t["runs"].([]any); ok && len(runs) > 0 {
					if r0, ok := runs[0].(map[string]any); ok {
						judul, _ = r0["text"].(string)
					}
				}
			}
			if id != "" {
				return &videoYT{Title: judul, URL: "https://www.youtube.com/watch?v=" + id}
			}
		}
		for _, val := range v {
			if hasil := cariRenderer(val); hasil != nil {
				return hasil
			}
		}
	case []any:
		for _, val := range v {
			if hasil := cariRenderer(val); hasil != nil {
				return hasil
			}
		}
	}
	return nil
}

func cariYT(judul string) *videoYT {
	body, err := lib.AmbilURL("https://www.youtube.com/results?search_query="+url.QueryEscape(judul), nil)
	if err != nil {
		return nil
	}
	teks := string(body)
	idx := strings.Index(teks, "var ytInitialData = ")
	if idx < 0 {
		return nil
	}
	sisa := teks[idx+len("var ytInitialData = "):]
	ujung := strings.Index(sisa, ";</script>")
	if ujung < 0 {
		return nil
	}
	var data any
	if json.Unmarshal([]byte(sisa[:ujung]), &data) != nil {
		return nil
	}
	return cariRenderer(data)
}

func TanganiYtplay(k *lib.Konteks) error {
	if k.M.Q == "" {
		return k.Conn.OK(k.M, "Judulnya apa njir😂")
	}
	u := k.M.Q
	if !strings.Contains(u, "youtu") {
		video := cariYT(k.M.Q)
		if video == nil {
			return fmt.Errorf("Video gak ketemu")
		}
		u = video.URL
		if err := k.Conn.OK(k.M, fmt.Sprintf("*%s*\n%s", video.Title, u)); err != nil {
			return err
		}
	}
	if err := k.Conn.OK(k.M, "sabar lagi donlod audio ny"); err != nil {
		return err
	}
	_ = os.MkdirAll("sampah", 0o755)
	output := "./sampah/" + k.M.ID + k.M.Sender
	if err := unduhYtdlp(k, u, output, "480"); err != nil {
		return err
	}
	defer bersihSampah(output + "*")
	lib.SpawnSync(fmt.Sprintf("ffmpeg -y -i %s* %s.mp3", output, output))
	b, err := os.ReadFile(output + ".mp3")
	if err != nil {
		return err
	}
	return k.Conn.SendMessage(k.M.From, lib.IsiPesan{Audio: b}, k.M)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "ytplay.js",
		Jenis:    "downloader",
		Nama:     []string{"#play"},
		Itunya:   regexp.MustCompile("(?i)^(#|!|\\/|\\.)(play|lagu)$"),
		Jalankan: TanganiYtplay,
	})
}
