package maker

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	waE2E "go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"

	"gow1/lib"
)

const dirKoleksiStiker = "bantuan/stiker"
const penandaJIDMaker = "120363186235853203@g.us"
const nomorPapaMaker = "6285188261898"
const nomorOwnerMaker = "6289680237010"

const saringPrompt = `Lo penyaring stiker. Jawab HANYA satu kata: AMAN atau TOLAK.
TOLAK kalau ada: ketelanjangan, konten seksual/sugestif, alat kelamin, adegan intim,
kekerasan berdarah, atau simbol kebencian.
AMAN buat sisanya — meme, ekspresi, kartun, hewan, teks lucu, foto biasa.`

const saringTanya = "Stiker ini aman dipakai di grup umum? Jawab AMAN atau TOLAK saja."

const labelPrompt = `Lo tukang labelin stiker buat koleksi mood. Jawab HANYA satu kata (huruf kecil, tanpa spasi, maks 20 karakter).
Label harus nangkep emosi/mood utama dari stiker ini.
Contoh label: ngakak, sedih, kaget, marah, cringe, gemes, takut, bingung, santai, bucin, malu, cape, seneng, keren, galau, jijik, panik, sabar, receh, nangis`

const labelTanya = "Stiker ini cocoknya dilabelin mood apa? Jawab satu kata aja."

var polaNonLabel = regexp.MustCompile(`[^a-z0-9]+`)
var polaAkhirWebp = regexp.MustCompile(`(?i)\d*\.webp$`)

func bersihLabel(s string) string {
	return polaNonLabel.ReplaceAllString(strings.ToLower(strings.TrimSpace(s)), "")
}

func daftarKoleksiStiker() []string {
	entri, err := os.ReadDir(dirKoleksiStiker)
	if err != nil {
		return []string{}
	}
	hasil := []string{}
	for _, e := range entri {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".webp") {
			hasil = append(hasil, e.Name())
		}
	}
	return hasil
}

func moodDariNama(nama string) string {
	return polaAkhirWebp.ReplaceAllString(nama, "")
}

func kelompokKoleksiStiker() map[string]int {
	p := map[string]int{}
	for _, f := range daftarKoleksiStiker() {
		p[moodDariNama(f)]++
	}
	return p
}

func ringkasanKoleksiStiker() string {
	p := kelompokKoleksiStiker()
	nama := make([]string, 0, len(p))
	for n := range p {
		nama = append(nama, n)
	}
	sort.Strings(nama)
	if len(nama) == 0 {
		return "belum ada stiker"
	}
	bagian := make([]string, len(nama))
	for i, n := range nama {
		bagian[i] = fmt.Sprintf("%s (%d)", n, p[n])
	}
	return strings.Join(bagian, ", ")
}

func adalahOwnerMaker(k *lib.Konteks) bool {
	for _, jid := range []string{k.M.Sender, k.M.SenderAlt} {
		n := lib.NomorSaja(jid)
		if n != "" && (n == nomorOwnerMaker || n == nomorPapaMaker || n == k.Conn.BotNomor()) {
			return true
		}
	}
	return false
}

func kirimPenandaAddstiker(k *lib.Konteks, teks string, subjek string) error {
	ci := &waE2E.ContextInfo{
		GroupMentions: []*waE2E.GroupMention{
			{GroupJID: proto.String(penandaJIDMaker), GroupSubject: proto.String(subjek)},
		},
	}
	if k.M != nil {
		ci.StanzaID = proto.String(k.M.ID)
		if k.M.Sender != "" {
			ci.Participant = proto.String(k.M.Sender)
		}
		switch f := k.M.Full.(type) {
		case *events.Message:
			ci.QuotedMessage = f.Message
		case *waE2E.Message:
			ci.QuotedMessage = f
		}
	}
	pesan := &waE2E.Message{
		ExtendedTextMessage: &waE2E.ExtendedTextMessage{
			Text:        proto.String(teks),
			ContextInfo: ci,
		},
	}
	if err := k.Conn.RelayMentah(k.M.From, pesan); err != nil {
		return k.Conn.SendMessage(k.M.From, lib.IsiPesan{Text: teks}, k.M)
	}
	return nil
}

func keJpegVisi(buf []byte) ([]byte, error) {
	masuk, err := lib.TulisTemp(lib.EkstensiMime(lib.DeteksiMime(buf)), buf)
	if err != nil {
		return nil, err
	}
	keluar := strings.TrimSuffix(masuk, filepath.Ext(masuk)) + "-visi.jpg"
	defer os.Remove(masuk)
	defer os.Remove(keluar)
	if _, err := jalanPerintah("convert", masuk, "-resize", "1024x1024>", "-background", "white", "-flatten", "-quality", "88", keluar); err != nil {
		return nil, err
	}
	return os.ReadFile(keluar)
}

func buangThinkVisi(s string) string {
	if i := strings.LastIndex(s, "</think>"); i != -1 {
		return strings.TrimLeft(s[i+len("</think>"):], " \t\r\n")
	}
	if i := strings.Index(s, "<think>"); i != -1 {
		return strings.TrimLeft(s[:i], " \t\r\n")
	}
	return s
}

func panggilVisi(jpeg []byte, pertanyaan string, sistem string) (string, error) {
	base := os.Getenv("AI_BASE_URL")
	if base == "" {
		base = "http://127.0.0.1:20128/v1"
	}
	model := os.Getenv("MCP_VISION_MODEL")
	if model == "" {
		model = "ag/gemini-3.8-flash-high"
	}
	body := map[string]any{
		"model":       model,
		"stream":      true,
		"temperature": 0.2,
		"messages": []any{
			map[string]any{"role": "system", "content": sistem},
			map[string]any{
				"role": "user",
				"content": []any{
					map[string]any{"type": "text", "text": pertanyaan},
					map[string]any{"type": "image_url", "image_url": map[string]any{"url": "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(jpeg)}},
				},
			},
		},
	}
	b, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", base+"/chat/completions", bytes.NewReader(b))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	if kunci := os.Getenv("AI_API_KEY"); kunci != "" {
		req.Header.Set("Authorization", "Bearer "+kunci)
	}
	resp, err := lib.KlienHTTP.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		mentah, _ := io.ReadAll(io.LimitReader(resp.Body, 200))
		return "", fmt.Errorf("Model vision %d: %s", resp.StatusCode, string(mentah))
	}
	var isi strings.Builder
	pembaca := bufio.NewReader(resp.Body)
	for {
		baris, err := pembaca.ReadString('\n')
		baris = strings.TrimRight(baris, "\r\n")
		if strings.HasPrefix(baris, "data:") {
			d := strings.TrimSpace(baris[5:])
			if d != "" && d != "[DONE]" {
				var j struct {
					Choices []struct {
						Delta struct {
							Content string `json:"content"`
						} `json:"delta"`
					} `json:"choices"`
				}
				if json.Unmarshal([]byte(d), &j) == nil && len(j.Choices) > 0 {
					isi.WriteString(j.Choices[0].Delta.Content)
				}
			}
		}
		if err != nil {
			break
		}
	}
	teks := strings.TrimSpace(buangThinkVisi(isi.String()))
	if teks == "" {
		return "", fmt.Errorf("Model vision nggak balikin deskripsi")
	}
	return teks, nil
}

func TanganiAddstiker(k *lib.Konteks) error {
	m := k.M
	cmd := strings.ToLower(m.Command)

	if strings.Contains(cmd, "liststiker") || strings.Contains(cmd, "stikerlist") {
		return k.Conn.OK(m, "stiker mood yang ada:\n"+ringkasanKoleksiStiker())
	}

	if strings.Contains(cmd, "delstiker") || strings.Contains(cmd, "hapusstiker") {
		if !adalahOwnerMaker(k) {
			return k.Conn.OK(m, "cuma owner")
		}
		target := bersihLabel(m.Q)
		if target == "" {
			return k.Conn.OK(m, "mau hapus mood apa? yang ada: "+ringkasanKoleksiStiker())
		}
		kena := []string{}
		for _, f := range daftarKoleksiStiker() {
			if moodDariNama(f) == target {
				kena = append(kena, f)
			}
		}
		if len(kena) == 0 {
			return k.Conn.OK(m, fmt.Sprintf("mood \"%s\" nggak ada. yang ada: %s", target, ringkasanKoleksiStiker()))
		}
		for _, f := range kena {
			_ = os.Remove(filepath.Join(dirKoleksiStiker, f))
		}
		return k.Conn.OK(m, fmt.Sprintf("%d stiker mood \"%s\" kehapus", len(kena), target))
	}

	if !adalahOwnerMaker(k) {
		return k.Conn.OK(m, "cuma owner")
	}

	tunggu := bersihLabel(m.Q) == "tunggu"

	if !tunggu {
		return kirimPenandaAddstiker(k, "kirim stikernya", "cmd|#addstiker tunggu")
	}

	down := unduhMedia(k)
	if down == nil {
		return k.Conn.OK(m, "reply ke pesan tadi sambil kirim stikernya")
	}

	mime := lib.DeteksiMime(down)
	if mime == "application/octet-stream" {
		return k.Conn.OK(m, "filenya kebaca apa ya, formatnya aneh")
	}

	var webpBuf []byte
	var err error
	switch {
	case strings.Contains(mime, "webp"):
		webpBuf = down
	case strings.Contains(mime, "image"):
		webpBuf, err = Img2webp(down)
	case strings.Contains(mime, "video") || strings.Contains(mime, "gif"):
		webpBuf, err = Mp42webp(down)
	default:
		return k.Conn.OK(m, fmt.Sprintf("%s nggak bisa dijadiin stiker", mime))
	}
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dirKoleksiStiker, 0o755); err != nil {
		return err
	}

	jpeg, err := keJpegVisi(webpBuf)
	if err != nil {
		return err
	}

	putusan, err := panggilVisi(jpeg, saringTanya, saringPrompt)
	if err != nil {
		return k.Conn.OK(m, fmt.Sprintf("gagal nyaring stikernya: %s", err.Error()))
	}
	if !strings.Contains(strings.ToUpper(putusan), "AMAN") || strings.Contains(strings.ToUpper(putusan), "TOLAK") {
		return k.Conn.OK(m, "stikernya nggak lolos saringan, ganti yang lain")
	}

	moodMentah, err := panggilVisi(jpeg, labelTanya, labelPrompt)
	if err != nil {
		return k.Conn.OK(m, fmt.Sprintf("gagal labelin stikernya: %s", err.Error()))
	}
	mood := bersihLabel(moodMentah)
	if mood == "" || len(mood) > 20 {
		mood = "lain"
	}

	kembar := []string{}
	for _, f := range daftarKoleksiStiker() {
		if moodDariNama(f) == mood {
			kembar = append(kembar, f)
		}
	}
	nomor := ""
	if len(kembar) > 0 {
		nomor = fmt.Sprintf("%d", len(kembar)+1)
	}
	nama := fmt.Sprintf("%s%s.webp", mood, nomor)
	if err := os.WriteFile(filepath.Join(dirKoleksiStiker, nama), webpBuf, 0o644); err != nil {
		return err
	}

	return k.Conn.OK(m, fmt.Sprintf("kesimpen jadi %s (%dKB)\nmood \"%s\" sekarang ada %d", nama, (len(webpBuf)+512)/1024, mood, len(kembar)+1))
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "addstiker.js",
		Jenis:    "maker",
		Nama:     []string{"#addstiker"},
		Itunya:   regexp.MustCompile("(?i)^(#|!|\\/|\\.)(addstiker|tambahstiker|liststiker|stikerlist|delstiker|hapusstiker)$"),
		Jalankan: TanganiAddstiker,
	})
}
