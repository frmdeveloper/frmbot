package maker

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gow1/lib"
)

func namaAcakSb() string {
	b := make([]byte, 14)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func unduhSb(k *lib.Konteks, p *lib.Pesan) ([]byte, string) {
	data, mime, err := k.Conn.DownloadAny(p)
	if err != nil {
		return nil, ""
	}
	return data, mime
}

func keOggSb(buf []byte) ([]byte, int, error) {
	ran := namaAcakSb()
	awal := "sampah/" + ran + lib.EkstensiMime(lib.DeteksiMime(buf))
	akhir := "sampah/" + ran + ".ogg"
	_ = os.MkdirAll("sampah", 0o755)
	if err := os.WriteFile(awal, buf, 0o644); err != nil {
		return nil, 0, err
	}
	defer os.Remove(awal)
	defer os.Remove(akhir)
	if keluar, err := exec.Command("ffmpeg", "-y", "-i", awal, "-vn", "-c:a", "libopus", "-b:a", "32k", "-ar", "48000", "-ac", "1", akhir).CombinedOutput(); err != nil {
		return nil, 0, fmt.Errorf("%s", strings.TrimSpace(string(keluar)))
	}
	detik := 1
	if keluar, err := exec.Command("ffprobe", "-v", "error", "-show_entries", "format=duration", "-of", "csv=p=0", akhir).Output(); err == nil {
		if f, err := strconv.ParseFloat(strings.TrimSpace(string(keluar)), 64); err == nil && f > 0 {
			detik = int(f + 0.5)
			if detik < 1 {
				detik = 1
			}
		}
	}
	data, err := os.ReadFile(akhir)
	return data, detik, err
}

func video2webpKotakSb(buf []byte) ([]byte, error) {
	ran := namaAcakSb()
	awal := "sampah/" + ran + lib.EkstensiMime(lib.DeteksiMime(buf))
	akhir := "sampah/" + ran + ".webp"
	_ = os.MkdirAll("sampah", 0o755)
	if err := os.WriteFile(awal, buf, 0o644); err != nil {
		return nil, err
	}
	defer os.Remove(awal)
	defer os.Remove(akhir)
	type langkah struct{ fps, q, s int }
	tangga := []langkah{{30, 60, 200}, {30, 45, 200}, {30, 30, 200}, {24, 30, 200}, {20, 25, 200}, {20, 25, 176}, {15, 25, 160}, {15, 20, 128}, {12, 20, 128}}
	var hasil []byte
	for _, l := range tangga {
		filter := fmt.Sprintf("crop='min(iw,ih)':'min(iw,ih)',scale=%d:%d:flags=lanczos,fps=%d,split [a][b]; [a] palettegen=reserve_transparent=on:transparency_color=ffffff [p]; [b][p] paletteuse", l.s, l.s, l.fps)
		if err := exec.Command("ffmpeg", "-y", "-i", awal, "-vcodec", "libwebp", "-vf", filter,
			"-loop", "0", "-ss", "00:00:00", "-t", "30", "-preset", "default", "-an", "-vsync", "0",
			"-quality", strconv.Itoa(l.q), akhir).Run(); err != nil {
			continue
		}
		b, err := os.ReadFile(akhir)
		if err != nil {
			continue
		}
		if hasil == nil || len(b) < len(hasil) {
			hasil = b
		}
		if len(b) <= 900*1024 {
			break
		}
	}
	if hasil == nil {
		return nil, fmt.Errorf("konversi video gagal")
	}
	return hasil, nil
}

func buatStikerSb(buf []byte, mime string) ([]byte, error) {
	switch {
	case strings.Contains(mime, "webp"):
		return buf, nil
	case strings.Contains(mime, "image"):
		return BuatStiker(buf, mime)
	case strings.Contains(mime, "video"):
		return video2webpKotakSb(buf)
	}
	return nil, fmt.Errorf("tipe %s gak didukung", mime)
}

func promptSb(k *lib.Konteks, teks, cmd string) error {
	return k.Conn.SendMessage(k.M.From, lib.IsiPesan{
		Text:          teks,
		GroupMentions: []lib.GrupMention{{JID: "120363186235853203@g.us", Subjek: "cmd|" + cmd}},
	}, k.M)
}

func rakitKirimSb(k *lib.Konteks, bufStiker, ogg []byte, detik int) error {
	m := k.M
	pack, author := "Elfa Chan", strconv.Itoa(time.Now().Year())
	if bagian := strings.Split(m.Q, "|"); len(bagian) == 2 && !strings.HasPrefix(strings.TrimSpace(bagian[0]), "suara") {
		pack = strings.TrimSpace(bagian[0])
		author = strings.TrimSpace(bagian[1])
	}
	stiker, err := AddExif(bufStiker, pack, author)
	if err != nil {
		stiker = bufStiker
	}
	target := m.Quoted
	if target == nil {
		target = m
	}
	return k.Conn.KirimStikerBersuara(m.From, stiker, ogg, detik, target)
}

func cocokGambarSb(tipe string) bool {
	return tipe == "imageMessage" || tipe == "stickerMessage" || tipe == "videoMessage"
}

func TanganiStikerBersuara(k *lib.Konteks) error {
	m := k.M
	args := strings.Fields(m.Q)

	if len(args) >= 2 && args[0] == "suara" {
		path := args[1]
		var bufSuara []byte
		if m.Type == "audioMessage" || m.Type == "videoMessage" {
			bufSuara, _ = unduhSb(k, m)
		}
		if bufSuara == nil && m.Quoted != nil && (m.Quoted.Type == "audioMessage" || m.Quoted.Type == "videoMessage") {
			bufSuara, _ = unduhSb(k, m.Quoted)
		}
		if bufSuara == nil {
			return promptSb(k, "audionya mana jir, kirim VN/audio coba", "#stikerbersuara suara "+path)
		}
		nomor := lib.NomorSaja(m.Sender)
		if !strings.HasPrefix(path, "sampah/") || !strings.HasPrefix(filepath.Base(path), "sb"+nomor) {
			return k.Conn.OK(m, "fotonya udah ilang jir, ulang dari awal pake #stikerbersuara")
		}
		bufGambar, err := os.ReadFile(path)
		if err != nil {
			return k.Conn.OK(m, "fotonya udah ilang jir, ulang dari awal pake #stikerbersuara")
		}
		os.Remove(path)
		mimeG := lib.DeteksiMime(bufGambar)
		if strings.Contains(mimeG, "video") {
			_ = k.Conn.OK(m, "bentar ya, videonya lagi gue proses...")
		}
		ogg, detik, err := keOggSb(bufSuara)
		if err != nil {
			return k.Conn.OK(m, "suaranya gak bisa diproses jir")
		}
		stik, err := buatStikerSb(bufGambar, mimeG)
		if err != nil {
			return k.Conn.OK(m, "medianya gak bisa dibaca jir")
		}
		return rakitKirimSb(k, stik, ogg, detik)
	}

	var bufGambar []byte
	var mimeGambar, tipeSumber string
	if cocokGambarSb(m.Type) {
		if b, mime := unduhSb(k, m); b != nil {
			bufGambar, mimeGambar, tipeSumber = b, mime, m.Type
		}
	}
	if bufGambar == nil && m.Quoted != nil && cocokGambarSb(m.Quoted.Type) {
		if b, mime := unduhSb(k, m.Quoted); b != nil {
			bufGambar, mimeGambar, tipeSumber = b, mime, m.Quoted.Type
		}
	}

	if bufGambar != nil {
		var bufSuara []byte
		if m.Type == "audioMessage" {
			bufSuara, _ = unduhSb(k, m)
		}
		if bufSuara == nil && m.Quoted != nil && (m.Quoted.Type == "audioMessage" || m.Quoted.Type == "videoMessage") {
			bufSuara, _ = unduhSb(k, m.Quoted)
		}
		if bufSuara != nil {
			if tipeSumber == "videoMessage" {
				_ = k.Conn.OK(m, "bentar ya, videonya lagi gue proses...")
			}
			ogg, detik, err := keOggSb(bufSuara)
			if err != nil {
				return k.Conn.OK(m, "suaranya gak bisa diproses jir")
			}
			stik, err := buatStikerSb(bufGambar, mimeGambar)
			if err != nil {
				return k.Conn.OK(m, "medianya gak bisa dibaca jir")
			}
			return rakitKirimSb(k, stik, ogg, detik)
		}
		if tipeSumber == "videoMessage" {
			_ = k.Conn.OK(m, "bentar ya, videonya lagi gue proses...")
			if ogg, detik, err := keOggSb(bufGambar); err == nil {
				if stikV, err := buatStikerSb(bufGambar, mimeGambar); err == nil {
					return rakitKirimSb(k, stikV, ogg, detik)
				}
			}
		}
		nomor := lib.NomorSaja(m.Sender)
		_ = os.MkdirAll("sampah", 0o755)
		if entri, err := os.ReadDir("sampah"); err == nil {
			for _, e := range entri {
				if strings.HasPrefix(e.Name(), "sb"+nomor) {
					os.Remove("sampah/" + e.Name())
				}
			}
		}
		berkas := fmt.Sprintf("sampah/sb%s%d.bin", nomor, time.Now().UnixMilli())
		if err := os.WriteFile(berkas, bufGambar, 0o644); err != nil {
			return err
		}
		return promptSb(k, "Sekarang kirim audionya coba\nreply dulu pesan ini", "#stikerbersuara suara "+berkas)
	}

	return promptSb(k, "Kirim fotonya coba", "#stikerbersuara foto")
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "stikerbersuara.js",
		Jenis:    "maker",
		Nama:     []string{"#stikerbersuara"},
		Itunya:   regexp.MustCompile(`(?i)^(#|!|/|\.)(s(tic?ker)?(bersuara|suara)|stiker2|s2)$`),
		Jalankan: TanganiStikerBersuara,
	})
}
