package downloader

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"gow1/lib"
)

func berkasSampah(pola string) []string {
	hasil, _ := filepath.Glob(pola)
	return hasil
}

func bersihSampah(pola string) {
	for _, f := range berkasSampah(pola) {
		_ = os.RemoveAll(f)
	}
}

func unduhYtdlp(k *lib.Konteks, urlBerkas, output, resolusi string) error {
	warp := "--proxy \"socks5://127.0.0.1:40000\""
	android := "--extractor-args \"youtube:player_client=android\""
	ios := "--extractor-args \"youtube:player_client=ios\""
	mweb := "--extractor-args \"youtube:player_client=mweb\""
	percobaan := []string{
		android + " " + warp,
		warp,
		android,
		ios,
		mweb + " --no-check-certificates",
		"--no-check-certificates",
		"",
	}
	galat := ""
	for _, c := range percobaan {
		perintah := fmt.Sprintf("yt-dlp \"%s\" --cookies \"./bantuan/cookie-yt.txt\" --js-runtimes node --no-playlist --output \"%s\" -S \"res:%s,ext:mp4:m4a\" %s", urlBerkas, output, resolusi, c)
		_, stderr := lib.SpawnSync(perintah)
		if len(berkasSampah(output+"*")) > 0 && !strings.Contains(stderr, "ERROR:") {
			return nil
		}
		galat = stderr
		bersihSampah(output + "*")
	}
	if strings.TrimSpace(galat) == "" {
		galat = "yt-dlp gagal tanpa pesan error"
	}
	return fmt.Errorf("%s", galat)
}

var reRegionYT = regexp.MustCompile(`(?i)not made this video available in your country`)
var reNegaraYT = regexp.MustCompile(`This video is available in ([^\n]+)`)

func teksRegionYT(galat string) (string, bool) {
	if !reRegionYT.MatchString(galat) {
		return "", false
	}
	negara := "negara tertentu"
	if m := reNegaraYT.FindStringSubmatch(galat); m != nil && strings.TrimSpace(m[1]) != "" {
		negara = strings.TrimSpace(m[1])
	}
	return "videonya dikunci region njir — cuma bisa ditonton di " + negara + ", server bot gak di sana jadi gak bisa donlod. coba cari upload lain atau link lain aja", true
}

func TanganiYtdl(k *lib.Konteks) error {
	if !strings.Contains(k.M.Q, "youtu") {
		return k.Conn.OK(k.M, "Link YouTube tidak valid")
	}
	if err := k.Conn.OK(k.M, "Sedang mengunduh video"); err != nil {
		return err
	}
	args := strings.Fields(k.M.Q)
	if len(args) == 0 {
		return fmt.Errorf("url kosong")
	}
	_ = os.MkdirAll("sampah", 0o755)
	output := "./sampah/" + k.M.ID + k.M.Sender
	if strings.ContainsAny(k.M.Command, "a3") {
		if err := unduhYtdlp(k, args[0], output, "480"); err != nil {
			if teks, ya := teksRegionYT(err.Error()); ya {
				return k.Conn.OK(k.M, teks)
			}
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
	resolusi := "360"
	if len(args) > 1 {
		resolusi = args[1]
	}
	if err := unduhYtdlp(k, args[0], output, resolusi); err != nil {
		if teks, ya := teksRegionYT(err.Error()); ya {
			return k.Conn.OK(k.M, teks)
		}
		return err
	}
	defer bersihSampah(output + "*")
	berkas := berkasSampah(output + "*")
	if len(berkas) == 0 {
		return fmt.Errorf("hasil yt-dlp gak ketemu")
	}
	b, err := os.ReadFile(berkas[0])
	if err != nil {
		return err
	}
	if strings.HasSuffix(berkas[0], ".mp4") {
		return k.Conn.SendMessage(k.M.From, lib.IsiPesan{
			Video: b,
		}, k.M)
	}
	return k.Conn.SendMessage(k.M.From, lib.IsiPesan{
		Document: b,
		FileName: fmt.Sprintf("%d.mp4", time.Now().UnixMilli()),
		Mimetype: "video/mp4",
	}, k.M)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "ytdl.js",
		Jenis:    "downloader",
		Nama:     []string{".ytmp4", ".ytmp3"},
		Itunya:   regexp.MustCompile(`(?i)^(#|!|/|\.)(yt|youtube|ytmp3|ytmp4)(a|v|mp3|mp4)?$`),
		Jalankan: TanganiYtdl,
	})
}
