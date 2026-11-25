package maker

import (
	crand "crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	mrand "math/rand"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gow1/lib"
)

func namaAcak() string {
	b := make([]byte, 14)
	_, _ = crand.Read(b)
	return hex.EncodeToString(b)
}

func jalanPerintah(nama string, args ...string) (string, error) {
	cmd := exec.Command(nama, args...)
	out, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return string(out), fmt.Errorf("%s: %s", nama, strings.TrimSpace(string(ee.Stderr)))
		}
		return string(out), err
	}
	return string(out), nil
}

func unduhMedia(k *lib.Konteks) []byte {
	if d, _, err := k.Conn.DownloadAny(k.M); err == nil && len(d) > 0 {
		return d
	}
	if k.M.Quoted != nil {
		if d, _, err := k.Conn.DownloadAny(k.M.Quoted); err == nil && len(d) > 0 {
			return d
		}
	}
	return nil
}

func simpanStkdb(m *lib.Pesan, data []byte) {
	dir := filepath.Join("data", "stiker", lib.NomorSaja(m.Sender))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return
	}
	_ = os.WriteFile(filepath.Join(dir, m.ID+".webp"), data, 0o644)
}

func tglStiker() string {
	s := time.Now()
	return fmt.Sprintf("%d:%d %d/%d/%d", s.Hour(), s.Minute(), s.Day(), int(s.Month()), s.Year())
}

func encodeURIComponentJS(s string) string {
	return strings.ReplaceAll(url.QueryEscape(s), "+", "%20")
}

var polaTenor = regexp.MustCompile(`https?://media[0-9]*\.tenor\.com/[^"\\]+\.gif`)
var polaSpasi = regexp.MustCompile(`\s+`)

func cariTenor(kata string) []byte {
	q := encodeURIComponentJS(polaSpasi.ReplaceAllString(strings.TrimSpace(kata), "-") + "-gifs")
	html, err := lib.AmbilURL("https://tenor.com/search/"+q, map[string]string{
		"User-Agent": "Mozilla/5.0 (Linux; Android 10) AppleWebKit/537.36",
	})
	if err != nil {
		return nil
	}
	mentah := polaTenor.FindAllString(string(html), -1)
	unik := []string{}
	udah := map[string]bool{}
	for _, u := range mentah {
		if !udah[u] {
			udah[u] = true
			unik = append(unik, u)
		}
	}
	if len(unik) == 0 {
		return nil
	}
	data, err := lib.AmbilURL(unik[mrand.Intn(len(unik))], nil)
	if err != nil {
		return nil
	}
	return data
}

func Img2webp(buffer []byte) ([]byte, error) {
	ext := lib.EkstensiMime(lib.DeteksiMime(buffer))
	awal, err := lib.TulisTemp(ext, buffer)
	if err != nil {
		return nil, err
	}
	akhir := strings.TrimSuffix(awal, ext) + ".webp"
	defer os.Remove(awal)
	defer os.Remove(akhir)
	filter := "scale=512:512:flags=lanczos:force_original_aspect_ratio=decrease,format=rgba,pad=512:512:(ow-iw)/2:(oh-ih)/2:color=#00000000,setsar=1"
	_ = exec.Command("ffmpeg", "-y", "-i", awal, "-vf", filter, akhir).Run()
	return os.ReadFile(akhir)
}

func Mp42webp(buffer []byte) ([]byte, error) {
	ext := lib.EkstensiMime(lib.DeteksiMime(buffer))
	awal, err := lib.TulisTemp(ext, buffer)
	if err != nil {
		return nil, err
	}
	akhir := strings.TrimSuffix(awal, ext) + ".webp"
	defer os.Remove(awal)
	defer os.Remove(akhir)
	filter := "scale='min(320,512)':min'(320,512)':force_original_aspect_ratio=decrease,fps=15, pad=320:320:-1:-1:color=white@0.0, split [a][b]; [a] palettegen=reserve_transparent=on:transparency_color=ffffff [p]; [b][p] paletteuse"
	_ = exec.Command("ffmpeg", "-y", "-i", awal, "-vcodec", "libwebp", "-vf", filter, "-loop", "0", "-ss", "00:00:00", "-t", "00:00:05", "-preset", "default", "-an", akhir).Run()
	return os.ReadFile(akhir)
}

func BuatStiker(data []byte, mime string) ([]byte, error) {
	switch {
	case strings.Contains(mime, "webp"):
		return data, nil
	case strings.Contains(mime, "image"):
		return Img2webp(data)
	case strings.Contains(mime, "video"):
		return Mp42webp(data)
	default:
		return nil, fmt.Errorf("tipe %s nggak bisa dijadiin stiker", mime)
	}
}

type exifStiker struct {
	PackID      string   `json:"sticker-pack-id"`
	PackName    string   `json:"sticker-pack-name"`
	Publisher   string   `json:"sticker-pack-publisher"`
	AndroidLink string   `json:"android-app-store-link"`
	IosLink     string   `json:"Ios-app-store-link"`
	Emojis      []string `json:"emojis"`
}

func bangunExifJSON(pack string, author string) []byte {
	isi := exifStiker{
		PackID:      "com.frmdeveloper.sticker",
		PackName:    pack,
		Publisher:   author,
		AndroidLink: "https://play.google.com/store/search?c=apps&q=" + encodeURIComponentJS("Dibuat oleh FRM pada "+tglStiker()),
		IosLink:     "https://itunes.apple.com/app/sticker-maker-studio/id1443326857",
		Emojis:      []string{""},
	}
	b, _ := json.Marshal(isi)
	return b
}

func bungkusExifAttr(jsonBuf []byte) []byte {
	attr := []byte{0x49, 0x49, 0x2A, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01, 0x00, 0x41, 0x57, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00}
	binary.LittleEndian.PutUint32(attr[14:18], uint32(len(jsonBuf)))
	return append(attr, jsonBuf...)
}

func muxExifWebp(data []byte, exif []byte) ([]byte, error) {
	inPath, err := lib.TulisTemp(".webp", data)
	if err != nil {
		return nil, err
	}
	exifPath, err := lib.TulisTemp(".exif", exif)
	if err != nil {
		os.Remove(inPath)
		return nil, err
	}
	defer os.Remove(inPath)
	defer os.Remove(exifPath)
	if _, err := jalanPerintah("webpmux", "-set", "exif", exifPath, inPath, "-o", inPath); err == nil {
		return os.ReadFile(inPath)
	}
	return sisipExifWebp(data, exif)
}

func sisipExifWebp(data []byte, exif []byte) ([]byte, error) {
	if len(data) < 20 || string(data[0:4]) != "RIFF" || string(data[8:12]) != "WEBP" {
		return nil, fmt.Errorf("bukan berkas webp")
	}
	type potong struct {
		tag string
		isi []byte
	}
	potongan := []potong{}
	pos := 12
	for pos+8 <= len(data) {
		tag := string(data[pos : pos+4])
		ukuran := int(binary.LittleEndian.Uint32(data[pos+4 : pos+8]))
		if pos+8+ukuran > len(data) {
			return nil, fmt.Errorf("chunk webp rusak")
		}
		potongan = append(potongan, potong{tag, data[pos+8 : pos+8+ukuran]})
		pos += 8 + ukuran
		if ukuran%2 == 1 {
			pos++
		}
	}
	if len(potongan) == 0 {
		return nil, fmt.Errorf("webp kosong")
	}
	keluaran := []potong{}
	if potongan[0].tag == "VP8X" {
		vp8x := append([]byte{}, potongan[0].isi...)
		if len(vp8x) < 10 {
			return nil, fmt.Errorf("vp8x rusak")
		}
		vp8x[0] |= 0x08
		keluaran = append(keluaran, potong{"VP8X", vp8x})
		keluaran = append(keluaran, potongan[1:]...)
	} else {
		var lebar, tinggi int
		switch potongan[0].tag {
		case "VP8 ":
			d := potongan[0].isi
			if len(d) < 10 {
				return nil, fmt.Errorf("vp8 rusak")
			}
			lebar = int(binary.LittleEndian.Uint16(d[6:8]) & 0x3FFF)
			tinggi = int(binary.LittleEndian.Uint16(d[8:10]) & 0x3FFF)
		case "VP8L":
			d := potongan[0].isi
			if len(d) < 5 {
				return nil, fmt.Errorf("vp8l rusak")
			}
			lebar = int(d[1]) | (int(d[2]&0x3F) << 8) + 1
			tinggi = (int(d[2]) >> 6) | (int(d[3]) << 2) | (int(d[4]&0x0F) << 10) + 1
		default:
			return nil, fmt.Errorf("chunk pertama webp tidak dikenal: %s", potongan[0].tag)
		}
		vp8x := make([]byte, 10)
		vp8x[0] = 0x08 | 0x10
		w := lebar - 1
		h := tinggi - 1
		vp8x[4] = byte(w)
		vp8x[5] = byte(w >> 8)
		vp8x[6] = byte(w >> 16)
		vp8x[7] = byte(h)
		vp8x[8] = byte(h >> 8)
		vp8x[9] = byte(h >> 16)
		keluaran = append(keluaran, potong{"VP8X", vp8x})
		keluaran = append(keluaran, potongan...)
	}
	bersih := keluaran[:0]
	for _, p := range keluaran {
		if p.tag != "EXIF" {
			bersih = append(bersih, p)
		}
	}
	bersih = append(bersih, potong{"EXIF", exif})
	total := 4
	for _, p := range bersih {
		total += 8 + len(p.isi)
		if len(p.isi)%2 == 1 {
			total++
		}
	}
	hasil := make([]byte, 0, total+8)
	hasil = append(hasil, []byte("RIFF")...)
	var uk [4]byte
	binary.LittleEndian.PutUint32(uk[:], uint32(total))
	hasil = append(hasil, uk[:]...)
	hasil = append(hasil, []byte("WEBP")...)
	for _, p := range bersih {
		hasil = append(hasil, []byte(p.tag)...)
		binary.LittleEndian.PutUint32(uk[:], uint32(len(p.isi)))
		hasil = append(hasil, uk[:]...)
		hasil = append(hasil, p.isi...)
		if len(p.isi)%2 == 1 {
			hasil = append(hasil, 0)
		}
	}
	return hasil, nil
}

func AddExif(data []byte, pack string, author string) ([]byte, error) {
	return muxExifWebp(data, bungkusExifAttr(bangunExifJSON(pack, author)))
}

func AddExifBawaan(data []byte) ([]byte, error) {
	return AddExif(data, "Stiker Kita", strconv.Itoa(time.Now().Year()))
}

type exifSet struct {
	PackID      string `json:"sticker-pack-id"`
	PackName    string `json:"sticker-pack-name"`
	Publisher   string `json:"sticker-pack-publisher"`
	AndroidLink string `json:"android-app-store-link"`
	IosLink     string `json:"ios-app-store-link"`
}

func CreateExif(pack string, auth string, pakage string) []byte {
	b, _ := json.Marshal(exifSet{
		PackID:      "com.frmdeveloper.sticker",
		PackName:    pack,
		Publisher:   auth,
		AndroidLink: "https://play.google.com/store/search?c=apps&q=" + encodeURIComponentJS(pakage),
		IosLink:     "",
	})
	panjang := len(b)
	kode := []byte{0x00, 0x00, 0x16, 0x00, 0x00, 0x00}
	if panjang > 256 {
		panjang -= 256
		kode = append([]byte{0x01}, kode...)
	} else {
		kode = append([]byte{0x00}, kode...)
	}
	hasil := []byte{0x49, 0x49, 0x2A, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01, 0x00, 0x41, 0x57, 0x07, 0x00}
	hasil = append(hasil, byte(panjang))
	hasil = append(hasil, kode...)
	hasil = append(hasil, b...)
	return hasil
}

func SetExif(data []byte, pack string, auth string) ([]byte, error) {
	return muxExifWebp(data, CreateExif(pack, auth, "Dibuat oleh FRM pada "+tglStiker()))
}

func EmojiToEscapeSequence(emoji string) string {
	for _, c := range emoji {
		return fmt.Sprintf("\\u{%x}", c)
	}
	return ""
}

var reSPersis = regexp.MustCompile(`^[#!\/\.]s$`)

func TanganiStiker(k *lib.Konteks) error {
	m := k.M
	adaSuara := StikerAudioDari(mentahDariPesan(m)) != nil
	if m.Quoted != nil {
		if StikerAudioDari(mentahDariFull(m.Quoted.Full)) != nil {
			adaSuara = true
		}
		if _, _, ok := AmbilCacheStikerBersuara(m.Quoted.ID); ok {
			adaSuara = true
		}
	}
	if reSPersis.MatchString(m.Command) && (m.Type == "videoMessage" || (m.Quoted != nil && m.Quoted.Type == "videoMessage") || m.Type == "audioStickerMessage" || (m.Quoted != nil && m.Quoted.Type == "audioStickerMessage") || adaSuara) {
		return TanganiStikerBersuara(k)
	}
	down := unduhMedia(k)
	if down == nil && m.Q != "" {
		down = cariTenor(m.Q)
		if down == nil {
			return k.Conn.OK(m, "gif-nya ga ketemu 😔")
		}
	}
	if down == nil {
		return k.Conn.OK(m, "mana gambarnya?")
	}
	mime := lib.DeteksiMime(down)
	kiri := "FRM Chan"
	kanan := strconv.Itoa(time.Now().Year())
	if m.Q != "" {
		h := strings.Split(m.Q, "|")
		kiri = h[0]
		if len(h) > 1 {
			kanan = h[1]
		}
	}
	var webpBuf []byte
	var err error
	switch {
	case strings.Contains(mime, "webp"):
		webpBuf = down
	case strings.Contains(mime, "image"):
		webpBuf, err = Img2webp(down)
	case strings.Contains(mime, "video"):
		webpBuf, err = Mp42webp(down)
	}
	if err != nil {
		return err
	}
	if webpBuf == nil {
		return k.Conn.OK(m, "mana gambarnya?")
	}
	stiker, err := AddExif(webpBuf, kiri, kanan)
	if err != nil {
		return err
	}
	if stiker == nil {
		return k.Conn.OK(m, "mana gambarnya?")
	}
	simpanStkdb(m, stiker)
	return k.Conn.SendMessage(m.From, lib.IsiPesan{Sticker: stiker}, m)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "stiker.js",
		Jenis:    "maker",
		Nama:     []string{"#stiker"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(wm|s(tic?ker)?(gif)?(wm)?)$"),
		Jalankan: TanganiStiker,
	})
}
