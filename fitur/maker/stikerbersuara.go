package maker

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"gow1/lib"

	waE2E "go.mau.fi/whatsmeow/proto/waE2E"
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

func durasiSb(buf []byte) float64 {
	ran := namaAcakSb()
	berkas := "sampah/" + ran + lib.EkstensiMime(lib.DeteksiMime(buf))
	_ = os.MkdirAll("sampah", 0o755)
	if err := os.WriteFile(berkas, buf, 0o644); err != nil {
		return 0
	}
	defer os.Remove(berkas)
	keluar, err := exec.Command("ffprobe", "-v", "error", "-show_entries", "format=duration", "-of", "csv=p=0", berkas).Output()
	if err != nil {
		return 0
	}
	f, _ := strconv.ParseFloat(strings.TrimSpace(string(keluar)), 64)
	return f
}

func video2webpKotakSb(buf []byte) ([]byte, error) {
	tAwal := time.Now()
	ran := namaAcakSb()
	awal := "sampah/" + ran + lib.EkstensiMime(lib.DeteksiMime(buf))
	_ = os.MkdirAll("sampah", 0o755)
	if err := os.WriteFile(awal, buf, 0o644); err != nil {
		return nil, err
	}
	defer os.Remove(awal)
	type langkah struct{ fps, q, s int }
	tangga := []langkah{{30, 60, 200}, {30, 45, 200}, {30, 30, 200}, {24, 30, 200}, {20, 25, 200}, {20, 25, 176}, {15, 25, 160}, {15, 20, 128}, {12, 20, 128}}
	berkasLangkah := func(i int) string { return fmt.Sprintf("sampah/%s-%d.webp", ran, i) }
	berkasContoh := "sampah/" + ran + "-contoh.webp"
	per := map[int][]byte{}
	var mu sync.Mutex
	encode := func(i int, potong float64) ([]byte, error) {
		l := tangga[i]
		tujuan := berkasLangkah(i)
		tArg := "30"
		if potong > 0 {
			tujuan = berkasContoh
			tArg = strconv.FormatFloat(potong, 'f', -1, 64)
		}
		filter := fmt.Sprintf("crop='min(iw,ih)':'min(iw,ih)',scale=%d:%d:flags=lanczos,fps=%d", l.s, l.s, l.fps)
		if err := exec.Command("ffmpeg", "-y", "-i", awal, "-vcodec", "libwebp", "-vf", filter,
			"-loop", "0", "-ss", "00:00:00", "-t", tArg, "-preset", "default", "-an", "-vsync", "0",
			"-quality", strconv.Itoa(l.q), tujuan).Run(); err != nil {
			return nil, err
		}
		b, err := os.ReadFile(tujuan)
		os.Remove(tujuan)
		if err != nil {
			return nil, err
		}
		if potong == 0 {
			mu.Lock()
			per[i] = b
			mu.Unlock()
		}
		return b, nil
	}
	cocok := func(b []byte) bool { return b != nil && len(b) <= 900*1024 }
	ambil := func(i int) []byte {
		mu.Lock()
		defer mu.Unlock()
		return per[i]
	}
	encodeParalel := func(daftar []int) {
		var wg sync.WaitGroup
		for _, i := range daftar {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				_, _ = encode(i, 0)
			}(i)
		}
		wg.Wait()
	}
	fasePasangan := func(dari int) bool {
		for p := dari; p < len(tangga); p += 2 {
			grup := []int{p}
			if p+1 < len(tangga) {
				grup = append(grup, p+1)
			}
			encodeParalel(grup)
			for _, i := range grup {
				if cocok(ambil(i)) {
					return true
				}
			}
		}
		return false
	}
	dur := durasiSb(buf)
	if dur > 8 && len(buf) > 4*1024*1024 {
		target := 0
		if contoh, err := encode(0, 0.5); err == nil {
			perDetik := float64(len(contoh)) / math.Min(0.5, dur)
			target = -1
			for i, l := range tangga {
				pred := perDetik * dur * (float64(l.fps) / 30) * math.Pow(float64(l.s)/200, 2) * math.Pow(float64(l.q)/60, 0.6)
				if pred <= 850*1024 {
					target = i
					break
				}
			}
			if target < 0 {
				target = len(tangga) - 1
			}
		}
		if target == 0 {
			_, _ = encode(0, 0)
			if !cocok(ambil(0)) {
				fasePasangan(1)
			}
		} else {
			pasangan := []int{target}
			if target < len(tangga)-1 {
				pasangan = append(pasangan, target+1)
			}
			encodeParalel(pasangan)
			adaCocok := false
			for _, i := range pasangan {
				if cocok(ambil(i)) {
					adaCocok = true
				}
			}
			if !adaCocok {
				for i := target + 2; i < len(tangga); i++ {
					_, _ = encode(i, 0)
					if cocok(ambil(i)) {
						break
					}
				}
			} else {
				for _, i := range pasangan {
					if b := ambil(i); cocok(b) && len(b) < 400*1024 {
						_, _ = encode(target-1, 0)
						break
					}
				}
			}
		}
	} else {
		_, _ = encode(0, 0)
		if !cocok(ambil(0)) {
			fasePasangan(1)
		}
	}
	urut := make([]int, 0, len(per))
	for i := range per {
		urut = append(urut, i)
	}
	sort.Ints(urut)
	var keluaran []byte
	for _, i := range urut {
		if cocok(per[i]) {
			keluaran = per[i]
			break
		}
	}
	if keluaran == nil {
		for _, i := range urut {
			if keluaran == nil || len(per[i]) < len(keluaran) {
				keluaran = per[i]
			}
		}
	}
	if keluaran == nil {
		return nil, fmt.Errorf("konversi video gagal")
	}
	fmt.Printf("[s2time] webp %.1fs dur=%.0f sumber=%db hasil=%db\n", time.Since(tAwal).Seconds(), dur, len(buf), len(keluaran))
	return keluaran, nil
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
	pack, author := "FRM Chan", strconv.Itoa(time.Now().Year())
	if bagian := strings.Split(m.Q, "|"); len(bagian) == 2 && !strings.HasPrefix(strings.TrimSpace(bagian[0]), "suara") {
		pack = strings.TrimSpace(bagian[0])
		author = strings.TrimSpace(bagian[1])
	}
	stiker, err := AddExif(bufStiker, pack, author)
	if err != nil {
		stiker = bufStiker
	}
	var target *lib.Pesan
	if m.Quoted != nil {
		target = m.Quoted
	} else if len(strings.Fields(m.Q)) > 0 {
		target = m
	}
	percobaan := 0
	for {
		err := k.Conn.KirimStikerBersuara(m.From, stiker, ogg, detik, target, m.Sender)
		if err == nil {
			return nil
		}
		percobaan++
		kenaLimit := strings.Contains(err.Error(), "429")
		if !kenaLimit || percobaan > 3 {
			if kenaLimit {
				return k.Conn.OK(m, "upload stikernya ditolak WA mulu jir (rate limit), coba lagi bentar ya")
			}
			return err
		}
		time.Sleep(time.Duration(5*percobaan) * time.Second)
	}
}

func cocokGambarSb(tipe string) bool {
	return tipe == "imageMessage" || tipe == "stickerMessage" || tipe == "videoMessage"
}

func TanganiStikerBersuara(k *lib.Konteks) error {
	m := k.M
	args := strings.Fields(m.Q)

	var stikMentah *waE2E.StickerMessage
	if s := StikerAudioDari(mentahDariPesan(m)); s != nil {
		stikMentah = s
	}
	if stikMentah == nil && m.Quoted != nil {
		stikMentah = StikerAudioDari(mentahDariFull(m.Quoted.Full))
	}
	if stikMentah == nil && m.Quoted != nil {
		if st, _, ok := AmbilCacheStikerBersuara(m.Quoted.ID); ok {
			stikMentah = st
		}
	}
	if stikMentah != nil && AudioDariStiker(stikMentah) != nil {
		aud := AudioDariStiker(stikMentah)
		pesanStik := &lib.Pesan{Full: &waE2E.Message{StickerMessage: stikerDariFld(fldDariStiker(stikMentah))}}
		pesanAud := &lib.Pesan{Full: &waE2E.Message{AudioMessage: audioDariFld(fldDariAudio(aud))}}
		bufStik, _ := unduhSb(k, pesanStik)
		bufAud, _ := unduhSb(k, pesanAud)
		if bufStik == nil || bufAud == nil {
			return k.Conn.OK(m, "stiker bersuaranya gak bisa gue ambil jir")
		}
		ogg, detik, err := keOggSb(bufAud)
		if err != nil {
			return k.Conn.OK(m, "suaranya gak bisa gue ambil jir")
		}
		return rakitKirimSb(k, bufStik, ogg, detik)
	}

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
