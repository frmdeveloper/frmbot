package converter

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"gow1/lib"
)

const layananRmbg = "http://127.0.0.1:8773/rmbg"

func TanganiRmbg(k *lib.Konteks) error {
	adalahURL := strings.HasPrefix(k.M.Q, "http://") || strings.HasPrefix(k.M.Q, "https://")
	if k.M.Quoted == nil && k.M.Full == nil && !adalahURL {
		return k.Conn.OK(k.M, "Reply foto/stikernya dulu, atau #rmbg <url gambar>")
	}
	var down []byte
	if adalahURL {
		if b, err := lib.AmbilURL(k.M.Q, nil); err == nil {
			down = b
		}
	}
	if len(down) == 0 {
		target := k.M
		if k.M.Quoted != nil {
			target = k.M.Quoted
		}
		if b, _, err := k.Conn.DownloadAny(target); err == nil {
			down = b
		}
	}
	if len(down) == 0 {
		return k.Conn.OK(k.M, "Gagal ngunduh medianya, coba reply ulang")
	}
	if err := k.Conn.OK(k.M, "Sabar, lagi ngilangin background-nya — modelnya berat, bisa semenit lebih"); err != nil {
		return err
	}
	ran := acakHex(8)
	if err := os.MkdirAll(filepath.Join("data", "tmp"), 0o755); err != nil {
		return err
	}
	tmp, err := os.MkdirTemp(filepath.Join("data", "tmp"), "rmbg*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp)
	keluar := filepath.Join(tmp, "keluar.png")
	mime := lib.DeteksiMime(down)
	jadiStiker := strings.Contains(mime, "webp") || (k.M.Quoted != nil && k.M.Quoted.Type == "stickerMessage") || k.M.Type == "stickerMessage"
	klien := &http.Client{Timeout: 300 * time.Second}
	resp, err := klien.Post(layananRmbg, "application/octet-stream", bytes.NewReader(down))
	if err != nil {
		return k.Conn.OK(k.M, "Gagal remove background\n\n"+potong(err.Error(), 200))
	}
	defer resp.Body.Close()
	hasilRmbg, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return k.Conn.OK(k.M, "Gagal remove background\n\n"+potong(string(hasilRmbg), 200))
	}
	if err := os.WriteFile(keluar, hasilRmbg, 0o644); err != nil {
		return err
	}
	if jadiStiker {
		webp := filepath.Join(tmp, "keluar.webp")
		lib.SpawnSync(fmt.Sprintf(`ffmpeg -y -i "%s" -vf "scale=512:512:flags=lanczos:force_original_aspect_ratio=decrease,format=rgba,pad=512:512:(ow-iw)/2:(oh-ih)/2:color=#00000000,setsar=1" "%s"`, keluar, webp))
		if !jadiBerkas(webp) {
			return k.Conn.OK(k.M, "Gagal konversi balik ke stiker")
		}
		data, err := os.ReadFile(webp)
		if err != nil {
			return err
		}
		buf, err := addExif(data)
		if err != nil {
			return err
		}
		return k.Conn.SendMessage(k.M.From, lib.IsiPesan{Sticker: buf}, k.M)
	}
	data, err := os.ReadFile(keluar)
	if err != nil {
		return err
	}
	return k.Conn.SendMessage(k.M.From, lib.IsiPesan{Document: data, Mimetype: "image/png", FileName: "rmbg-" + ran + ".png"}, k.M)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "rmbg.js",
		Jenis:    "converter",
		Nama:     []string{"#rmbg"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(rmbg|removebg)$"),
		Jalankan: TanganiRmbg,
	})
}
