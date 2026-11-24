package audio

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"gow1/lib"
)

const dirAudio = "sampah"
const srAudio = 48000

var polaAngkaJS = regexp.MustCompile(`^[+-]?(\d+\.?\d*|\.\d+)([eE][+-]?\d+)?`)

func parseFloatJS(s string) float64 {
	s = strings.TrimSpace(s)
	cocok := polaAngkaJS.FindString(s)
	if cocok == "" {
		return math.NaN()
	}
	n, err := strconv.ParseFloat(cocok, 64)
	if err != nil {
		return math.NaN()
	}
	return n
}

func angka(nilai string, min, max, bawaan float64) float64 {
	n := parseFloatJS(nilai)
	if math.IsNaN(n) || math.IsInf(n, 0) {
		return bawaan
	}
	return math.Min(max, math.Max(min, n))
}

func formatJS(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func kecepatanAudio(x float64) string {
	return fmt.Sprintf("aresample=%d,asetrate=%d*%s,aresample=%d", srAudio, srAudio, formatJS(x), srAudio)
}

var polaWaktu = regexp.MustCompile(`^([0-9]+:)?[0-9]+(:[0-9]+)?(\.[0-9]+)?$`)

func waktuOk(t string) bool {
	return polaWaktu.MatchString(strings.TrimSpace(t))
}

type efekAudio struct {
	contoh string
	buat   func(a string, durasi float64) (string, string, string)
}

var efek = map[string]efekAudio{
	"bass": {contoh: "#bass 12", buat: func(a string, durasi float64) (string, string, string) {
		return "bass=g=" + formatJS(angka(a, -30, 30, 12)), "", ""
	}},
	"treble": {contoh: "#treble 10", buat: func(a string, durasi float64) (string, string, string) {
		return "treble=g=" + formatJS(angka(a, -30, 30, 10)), "", ""
	}},
	"volume": {contoh: "#volume 8", buat: func(a string, durasi float64) (string, string, string) {
		return "volume=" + formatJS(angka(a, -30, 30, 8)) + "dB", "", ""
	}},
	"reverb": {contoh: "#reverb", buat: func(a string, durasi float64) (string, string, string) {
		return "aecho=0.8:0.88:60|85|110:0.4|0.3|0.2", "", ""
	}},
	"slow": {contoh: "#slow", buat: func(a string, durasi float64) (string, string, string) { return kecepatanAudio(0.85), "", "" }},
	"fast": {contoh: "#fast", buat: func(a string, durasi float64) (string, string, string) { return kecepatanAudio(1.25), "", "" }},
	"speed": {contoh: "#speed 1.5", buat: func(a string, durasi float64) (string, string, string) {
		return kecepatanAudio(angka(a, 0.5, 2.5, 1.5)), "", ""
	}},
	"vibra": {contoh: "#vibra 6", buat: func(a string, durasi float64) (string, string, string) {
		return "vibrato=f=" + formatJS(angka(a, 0.1, 20, 6)) + ":d=0.6", "", ""
	}},
	"fade": {contoh: "#fade 3", buat: func(a string, durasi float64) (string, string, string) {
		maks := math.Max(0.5, durasi/2)
		d := angka(a, 0.5, maks, math.Min(2, maks))
		st := math.Max(0, durasi-d)
		return fmt.Sprintf("afade=t=in:st=0:d=%.2f,afade=t=out:st=%.2f:d=%.2f", d, st, d), "", ""
	}},
	"cut": {contoh: "#cut 0:10-0:25", buat: func(a string, durasi float64) (string, string, string) {
		bagian := strings.SplitN(a, "-", 2)
		var awal, akhir string
		if len(bagian) > 0 {
			awal = bagian[0]
		}
		if len(bagian) > 1 {
			akhir = bagian[1]
		}
		if !waktuOk(awal) || !waktuOk(akhir) {
			return "", "", "Formatnya *mulai-selesai*, contoh #cut 0:10-0:25"
		}
		return "", fmt.Sprintf("-ss %s -to %s", strings.TrimSpace(awal), strings.TrimSpace(akhir)), ""
	}},
}

func durasiAudio(file string) float64 {
	out, _ := lib.SpawnSync(fmt.Sprintf(`ffprobe -v error -select_streams a:0 -show_entries format=duration -of csv=p=0 "%s"`, file))
	n := parseFloatJS(out)
	if math.IsNaN(n) {
		return 0
	}
	return n
}

func TanganiAudio(k *lib.Konteks) error {
	m := k.M
	nama := m.Command
	if r := []rune(nama); len(r) > 0 && strings.ContainsRune("!#%./\\", r[0]) {
		nama = string(r[1:])
	}
	efekDipilih, ada := efek[strings.ToLower(nama)]
	if !ada {
		return nil
	}

	var target *lib.Pesan
	if m.Quoted != nil && m.Quoted.Type == "audioMessage" {
		target = m.Quoted
	} else if m.Type == "audioMessage" {
		target = m
	}
	if target == nil {
		return k.Conn.OK(m, "Reply audionya dulu, contoh "+efekDipilih.contoh)
	}

	down, _, err := k.Conn.DownloadAny(target)
	if err != nil || len(down) == 0 {
		return k.Conn.OK(m, "Gagal ngunduh audionya, coba reply ulang")
	}

	if err := os.MkdirAll(dirAudio, 0o755); err != nil {
		return err
	}
	acak := make([]byte, 8)
	if _, err := rand.Read(acak); err != nil {
		return err
	}
	ran := hex.EncodeToString(acak)
	masuk := filepath.Join(dirAudio, ran+".masuk")
	keluar := filepath.Join(dirAudio, ran+".mp3")
	if err := os.WriteFile(masuk, down, 0o644); err != nil {
		return err
	}
	defer os.Remove(masuk)
	defer os.Remove(keluar)

	durasi := durasiAudio(masuk)
	if durasi == 0 {
		return k.Conn.OK(m, "Audionya rusak atau nggak kebaca")
	}

	args := strings.Fields(m.Q)
	arg0 := ""
	if len(args) > 0 {
		arg0 = args[0]
	}
	af, pre, salah := efekDipilih.buat(arg0, durasi)
	if salah != "" {
		return k.Conn.OK(m, salah)
	}

	perintah := fmt.Sprintf(`ffmpeg -hide_banner -y %s -i "%s" %s -c:a libmp3lame -b:a 128k "%s"`, pre, masuk, bagianAF(af), keluar)
	_, stderr := lib.SpawnSync(perintah)

	info, statErr := os.Stat(keluar)
	if statErr != nil || info.Size() == 0 {
		var baris []string
		for _, b := range strings.Split(strings.TrimSpace(stderr), "\n") {
			if strings.TrimSpace(b) != "" {
				baris = append(baris, b)
			}
		}
		alasan := ""
		if len(baris) > 2 {
			alasan = strings.Join(baris[len(baris)-2:], "\n")
		} else {
			alasan = strings.Join(baris, "\n")
		}
		if alasan == "" {
			alasan = "ffmpeg nggak ngeluarin apa-apa"
		}
		return k.Conn.OK(m, "Gagal ngolah audionya\n\n"+alasan)
	}

	data, err := os.ReadFile(keluar)
	if err != nil {
		return err
	}
	return k.Conn.SendMessage(m.From, lib.IsiPesan{Audio: data, Mimetype: "audio/mpeg"}, m)
}

func bagianAF(af string) string {
	if af == "" {
		return ""
	}
	return fmt.Sprintf(`-af "%s"`, af)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "audio.js",
		Jenis:    "audio",
		Nama:     []string{"#bass", "#reverb", "#slow", "#cut", "#fast", "#speed", "#vibra", "#fade", "#treble", "#volume"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(bass|reverb|slow|cut|fast|speed|vibra|fade|treble|volume)$"),
		Jalankan: TanganiAudio,
	})
}
