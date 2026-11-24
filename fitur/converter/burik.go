package converter

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gow1/lib"
)

func acakHex(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func jadiBerkas(path string) bool {
	st, err := os.Stat(path)
	return err == nil && st.Size() > 0
}

func ekorBaris(s string, n int) string {
	bagian := []string{}
	for _, baris := range strings.Split(strings.TrimSpace(s), "\n") {
		if strings.TrimSpace(baris) != "" {
			bagian = append(bagian, baris)
		}
	}
	if len(bagian) > n {
		bagian = bagian[len(bagian)-n:]
	}
	return strings.Join(bagian, "\n")
}

func potong(s string, n int) string {
	r := []rune(s)
	if len(r) > n {
		return string(r[:n])
	}
	return s
}

func encodeURIComponent(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || strings.IndexByte("-_.!~*'()", c) >= 0 {
			b.WriteByte(c)
		} else {
			fmt.Fprintf(&b, "%%%02X", c)
		}
	}
	return b.String()
}

func jsonKutip(s string) string {
	b, _ := json.Marshal(s)
	return string(b)
}

func tglSekarang() string {
	t := time.Now()
	return fmt.Sprintf("%d:%d %d/%d/%d", t.Hour(), t.Minute(), t.Day(), int(t.Month()), t.Year())
}

func createExif(pack, auth, pakage string) []byte {
	jsonStr := fmt.Sprintf(`{"sticker-pack-id":"com.frmdeveloper.sticker","sticker-pack-name":%s,"sticker-pack-publisher":%s,"android-app-store-link":%s,"ios-app-store-link":""}`,
		jsonKutip(pack), jsonKutip(auth), jsonKutip("https://play.google.com/store/search?c=apps&q="+encodeURIComponent(pakage)))
	code := []byte{0x00, 0x00, 0x16, 0x00, 0x00, 0x00}
	ln := len(jsonStr)
	if ln > 256 {
		ln -= 256
		code = append([]byte{0x01}, code...)
	} else {
		code = append([]byte{0x00}, code...)
	}
	hexLen := strconv.FormatInt(int64(ln), 16)
	if ln < 16 {
		hexLen = "0" + hexLen
	}
	lenBytes, _ := hex.DecodeString(hexLen)
	head := []byte{0x49, 0x49, 0x2A, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01, 0x00, 0x41, 0x57, 0x07, 0x00}
	out := append([]byte{}, head...)
	out = append(out, lenBytes...)
	out = append(out, code...)
	out = append(out, jsonStr...)
	return out
}

func addExif(webp []byte) ([]byte, error) {
	now := time.Now()
	jsonStr := fmt.Sprintf(`{"sticker-pack-id":"com.frmdeveloper.sticker","sticker-pack-name":"Stiker Kita","sticker-pack-publisher":%d,"android-app-store-link":%s,"Ios-app-store-link":"https://itunes.apple.com/app/sticker-maker-studio/id1443326857","emojis":[""]}`,
		now.Year(), jsonKutip("https://play.google.com/store/search?c=apps&q="+encodeURIComponent("Dibuat oleh FRM pada "+tglSekarang())))
	jsonBuffer := []byte(jsonStr)
	exif := []byte{0x49, 0x49, 0x2A, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01, 0x00, 0x41, 0x57, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00}
	binary.LittleEndian.PutUint32(exif[14:18], uint32(len(jsonBuffer)))
	exif = append(exif, jsonBuffer...)
	fWebp, err := lib.TulisTemp(".webp", webp)
	if err != nil {
		return nil, err
	}
	defer os.Remove(fWebp)
	fExif, err := lib.TulisTemp(".exif", exif)
	if err != nil {
		return nil, err
	}
	defer os.Remove(fExif)
	keluar := strings.TrimSuffix(fWebp, ".webp") + "-exif.webp"
	defer os.Remove(keluar)
	_, stderr := lib.SpawnSync(fmt.Sprintf(`webpmux -set exif "%s" "%s" -o "%s"`, fExif, fWebp, keluar))
	if !jadiBerkas(keluar) {
		return nil, fmt.Errorf("webpmux gagal pasang exif: %s", ekorBaris(stderr, 2))
	}
	return os.ReadFile(keluar)
}

func genapF(n float64) int {
	v := int(math.Trunc(n/2) * 2)
	if v < 2 {
		v = 2
	}
	return v
}

func kecilin(n int) int {
	v := genapF(float64(n) / 16)
	if v < 16 {
		v = 16
	}
	return v
}

func blurStr(w, h int) string {
	return fmt.Sprintf("scale=%d:%d:flags=bilinear,scale=%d:%d:flags=bilinear,gblur=sigma=2.4", kecilin(w), kecilin(h), genapF(float64(w)), genapF(float64(h)))
}

func ukuranMedia(file string) (int, int, bool) {
	output, _ := lib.SpawnSync(fmt.Sprintf(`ffprobe -v error -select_streams v:0 -show_entries stream=width,height -of csv=p=0 "%s"`, file))
	baris := strings.Split(strings.TrimSpace(output), "\n")[0]
	bagian := strings.Split(baris, ",")
	if len(bagian) < 2 {
		return 0, 0, false
	}
	w, err1 := strconv.Atoi(strings.TrimSpace(bagian[0]))
	h, err2 := strconv.Atoi(strings.TrimSpace(bagian[1]))
	if err1 != nil || err2 != nil || w == 0 || h == 0 {
		return 0, 0, false
	}
	return genapF(float64(w)), genapF(float64(h)), true
}

func capBerkas(file string) string {
	data, _ := os.ReadFile(file)
	head := data
	if len(head) > 16 {
		head = head[:16]
	}
	var sb strings.Builder
	for _, c := range head {
		if c >= 0x20 && c <= 0x7e {
			sb.WriteByte(c)
		} else {
			sb.WriteByte('.')
		}
	}
	return fmt.Sprintf("%db %s %s", len(data), sb.String(), hex.EncodeToString(head))
}

type infoStikerHasil struct {
	w     int
	h     int
	jeda  []int
	salah string
}

var polaKanvas = regexp.MustCompile(`Canvas size:\s*(\d+)\s*x\s*(\d+)`)
var polaFrame = regexp.MustCompile(`(?m)^\s*\d+:\s+\d+\s+\d+\s+\S+\s+\S+\s+\S+\s+(\d+)`)

func infoStiker(file string) infoStikerHasil {
	output, stderr := lib.SpawnSync(fmt.Sprintf(`webpmux -info "%s"`, file))
	kanvas := polaKanvas.FindStringSubmatch(output)
	if kanvas == nil {
		sumber := stderr
		if strings.TrimSpace(sumber) == "" {
			sumber = output
		}
		kata := strings.ReplaceAll(ekorBaris(sumber, 2), "\n", " | ")
		if kata == "" {
			kata = "diem aja"
		}
		return infoStikerHasil{salah: fmt.Sprintf("webpmux: %s\nberkas: %s", kata, capBerkas(file))}
	}
	w, _ := strconv.Atoi(kanvas[1])
	h, _ := strconv.Atoi(kanvas[2])
	jeda := []int{}
	for _, m := range polaFrame.FindAllStringSubmatch(output, -1) {
		d, _ := strconv.Atoi(m[1])
		jeda = append(jeda, d)
	}
	return infoStikerHasil{w: genapF(float64(w)), h: genapF(float64(h)), jeda: jeda}
}

func burikStiker(masuk, tmp, keluar string) (string, string) {
	info := infoStiker(masuk)
	if info.salah != "" {
		return "", info.salah
	}
	vf := "format=rgba," + blurStr(info.w, info.h)
	if len(info.jeda) < 2 {
		_, stderr := lib.SpawnSync(fmt.Sprintf(`ffmpeg -hide_banner -y -i "%s" -vf "%s" -c:v libwebp -q:v 8 "%s"`, masuk, vf, keluar))
		if jadiBerkas(keluar) {
			return "gambar", ""
		}
		return "", stderr
	}
	langkah := int(math.Ceil(float64(len(info.jeda)) / 24))
	if langkah < 1 {
		langkah = 1
	}
	pakai := []int{}
	gagalFrame := ""
	for i := 0; i < len(info.jeda); i += langkah {
		n := len(pakai) + 1
		fWebp := filepath.Join(tmp, fmt.Sprintf("f%d.webp", n))
		fPng := filepath.Join(tmp, fmt.Sprintf("f%d.png", n))
		_, stderr := lib.SpawnSync(fmt.Sprintf(`webpmux -get frame %d "%s" -o "%s" && dwebp -quiet "%s" -o "%s"`, i+1, masuk, fWebp, fWebp, fPng))
		if !jadiBerkas(fPng) {
			alasan := strings.TrimSpace(stderr)
			if alasan == "" {
				alasan = "png nggak kebikin"
			}
			gagalFrame = fmt.Sprintf("frame %d gagal diambil: %s", i+1, alasan)
			break
		}
		jumlah := 0
		for j := i; j < i+langkah && j < len(info.jeda); j++ {
			jumlah += info.jeda[j]
		}
		if jumlah < 20 {
			jumlah = 20
		}
		pakai = append(pakai, jumlah)
	}
	if len(pakai) == 0 {
		if gagalFrame != "" {
			return "", gagalFrame
		}
		return "", "nggak ada frame kepake"
	}
	if len(pakai) < 2 {
		_, stderr := lib.SpawnSync(fmt.Sprintf(`ffmpeg -hide_banner -y -i "%s" -vf "%s" -c:v libwebp -q:v 8 "%s"`, filepath.Join(tmp, "f1.png"), vf, keluar))
		if jadiBerkas(keluar) {
			return "gambar", ""
		}
		if strings.TrimSpace(stderr) != "" {
			return "", stderr
		}
		return "", gagalFrame
	}
	_, stderrOlah := lib.SpawnSync(fmt.Sprintf(`ffmpeg -hide_banner -y -i "%s" -vf "%s" "%s"`, filepath.Join(tmp, "f%d.png"), vf, filepath.Join(tmp, "b%d.png")))
	if !jadiBerkas(filepath.Join(tmp, "b1.png")) {
		return "", "ffmpeg gagal ngeblur frame: " + stderrOlah
	}
	var daftar strings.Builder
	for i, d := range pakai {
		fmt.Fprintf(&daftar, "-lossy -q 6 -d %d \"%s\" ", d, filepath.Join(tmp, fmt.Sprintf("b%d.png", i+1)))
	}
	_, stderrSusun := lib.SpawnSync(fmt.Sprintf(`img2webp -loop 0 %s-o "%s"`, daftar.String(), keluar))
	if !jadiBerkas(keluar) {
		return "", "img2webp gagal nyusun: " + stderrSusun
	}
	return "animasi", ""
}

func cocokMedia(tipe string) bool {
	return tipe == "videoMessage" || tipe == "imageMessage" || tipe == "stickerMessage"
}

func TanganiBurik(k *lib.Konteks) error {
	tipe := ""
	if k.M.Quoted != nil && cocokMedia(k.M.Quoted.Type) {
		tipe = k.M.Quoted.Type
	} else if cocokMedia(k.M.Type) {
		tipe = k.M.Type
	}
	if tipe == "" {
		return k.Conn.OK(k.M, "Reply video, foto, atau stikernya dulu pakai #burik")
	}
	target := k.M
	if k.M.Quoted != nil && k.M.Quoted.Type == tipe {
		target = k.M.Quoted
	}
	down, _, _ := k.Conn.DownloadAny(target)
	if len(down) == 0 {
		return k.Conn.OK(k.M, "Gagal ngunduh medianya, coba reply ulang")
	}
	if err := os.MkdirAll(filepath.Join("data", "tmp"), 0o755); err != nil {
		return err
	}
	tmp, err := os.MkdirTemp(filepath.Join("data", "tmp"), "burik*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp)
	masuk := filepath.Join(tmp, "masuk")
	stiker := tipe == "stickerMessage"
	gambar := tipe == "imageMessage"
	ext := ".mp4"
	if stiker {
		ext = ".webp"
	} else if gambar {
		ext = ".jpg"
	}
	keluar := filepath.Join(tmp, "hasil"+ext)
	if err := os.WriteFile(masuk, down, 0o644); err != nil {
		return err
	}
	var ok, salah string
	if stiker {
		ok, salah = burikStiker(masuk, tmp, keluar)
	} else {
		w, h, ada := ukuranMedia(masuk)
		if !ada {
			return k.Conn.OK(k.M, "Medianya rusak atau nggak kebaca")
		}
		vf := blurStr(w, h) + ",format=yuv420p"
		warna := "format=yuv410p,format=yuv420p,"
		var cmd string
		if gambar {
			cmd = fmt.Sprintf(`ffmpeg -hide_banner -y -i "%s" -vf "%s%s" -q:v 28 "%s"`, masuk, warna, vf, keluar)
		} else {
			cmd = fmt.Sprintf(`ffmpeg -hide_banner -y -i "%s" -vf "fps=12,%s%s" -c:v libx264 -crf 48 -maxrate 70k -bufsize 140k -preset veryfast -movflags +faststart -c:a aac -b:a 16k -ac 1 -ar 16000 "%s"`, masuk, warna, vf, keluar)
		}
		_, stderr := lib.SpawnSync(cmd)
		salah = stderr
		if jadiBerkas(keluar) {
			if gambar {
				ok = "gambar"
			} else {
				ok = "video"
			}
		}
	}
	if ok == "" {
		alasan := ekorBaris(salah, 3)
		if alasan == "" {
			alasan = "nggak ada output"
		}
		return k.Conn.OK(k.M, "Gagal ngeburikin\n\n"+alasan)
	}
	if stiker {
		if ok == "animasi" {
			berkas := filepath.Join(tmp, "pack.exif")
			_ = os.WriteFile(berkas, createExif("Stiker Kita", strconv.Itoa(time.Now().Year()), "sticker whatsapp"), 0o644)
			lib.SpawnSync(fmt.Sprintf(`webpmux -set exif "%s" "%s" -o "%s"`, berkas, keluar, keluar))
			data, err := os.ReadFile(keluar)
			if err != nil {
				return err
			}
			return k.Conn.SendMessage(k.M.From, lib.IsiPesan{Sticker: data}, k.M)
		}
		data, err := os.ReadFile(keluar)
		if err != nil {
			return err
		}
		webp, err := addExif(data)
		if err != nil {
			return err
		}
		return k.Conn.SendMessage(k.M.From, lib.IsiPesan{Sticker: webp}, k.M)
	}
	data, err := os.ReadFile(keluar)
	if err != nil {
		return err
	}
	if gambar {
		return k.Conn.SendMessage(k.M.From, lib.IsiPesan{Image: data}, k.M)
	}
	return k.Conn.SendMessage(k.M.From, lib.IsiPesan{Video: data}, k.M)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "burik.js",
		Jenis:    "converter",
		Nama:     []string{"#burik"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(burik)$"),
		Jalankan: TanganiBurik,
	})
}
