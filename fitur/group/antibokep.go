package group

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"gow1/lib"

	"go.mau.fi/whatsmeow/types/events"
)

const promptVonisBokep = "Kamu filter konten grup WhatsApp. Nilai gambar ini. Balas HANYA satu kata: PORNO atau AMAN. PORNO = konten seksual eksplisit: alat kelamin terlihat, aktivitas seksual, payudara atau pantat telanjang yang ditampilkan secara seksual, termasuk gambar anime/hentai eksplisit. AMAN = selain itu, termasuk baju renang, pakaian terbuka non-eksplisit, laki-laki bertelanjang dada, seni atau medis non-seksual."

const promptDeskripsiBokep = "Deskripsikan isi gambar ini dalam Bahasa Indonesia yang santai dan singkat (1-2 kalimat saja). Langsung ke isinya, jangan pakai pembuka seperti 'gambar ini menunjukkan' atau 'foto ini adalah'."

func dasarRouterBokep() string {
	base := os.Getenv("AI_BASE_URL")
	if base == "" {
		base = "http://127.0.0.1:20128/v1"
	}
	return strings.TrimRight(base, "/")
}

func modelBokep() string {
	if v := os.Getenv("ANTIBOKEP_MODEL"); v != "" {
		return v
	}
	if v := os.Getenv("MCP_VISION_MODEL"); v != "" {
		return v
	}
	return "chatbot"
}

func kompresiBokep(buf []byte) (string, int, error) {
	dir, err := os.MkdirTemp("", "bokep-")
	if err != nil {
		return "", 0, err
	}
	defer os.RemoveAll(dir)
	masuk := filepath.Join(dir, "masuk.bin")
	keluar := filepath.Join(dir, "keluar.jpg")
	if err := os.WriteFile(masuk, buf, 0o644); err != nil {
		return "", 0, err
	}
	cmd := exec.Command("ffmpeg", "-y", "-v", "error", "-i", masuk, "-vf", "scale='min(512,iw)':'min(512,ih)':force_original_aspect_ratio=decrease", "-frames:v", "1", "-q:v", "5", keluar)
	if err := cmd.Run(); err != nil {
		return "", 0, err
	}
	jpg, err := os.ReadFile(keluar)
	if err != nil {
		return "", 0, err
	}
	return "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(jpg), len(jpg) / 1024, nil
}

func panggilVisiBokep(uri, prompt string, maksToken int, batas time.Duration) (string, bool) {
	payload := map[string]any{
		"model":      modelBokep(),
		"stream":     false,
		"max_tokens": maksToken,
		"messages": []any{map[string]any{"role": "user", "content": []any{
			map[string]any{"type": "text", "text": prompt},
			map[string]any{"type": "image_url", "image_url": map[string]any{"url": uri}},
		}}},
	}
	b, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", dasarRouterBokep()+"/chat/completions", bytes.NewReader(b))
	if err != nil {
		return "", false
	}
	req.Header.Set("content-type", "application/json")
	if kunci := os.Getenv("AI_API_KEY"); kunci != "" {
		req.Header.Set("authorization", "Bearer "+kunci)
	}
	klien := &http.Client{Timeout: batas}
	resp, err := klien.Do(req)
	if err != nil {
		fmt.Printf("[antibokep] visi error: %v\n", err)
		return "", false
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Printf("[antibokep] visi HTTP %d\n", resp.StatusCode)
		return "", false
	}
	var hasil struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&hasil); err != nil || len(hasil.Choices) == 0 {
		return "", false
	}
	return hasil.Choices[0].Message.Content, true
}

func vonisVisiBokep(uri string, batas time.Duration) string {
	teks, ok := panggilVisiBokep(uri, promptVonisBokep, 10, batas)
	if !ok {
		return "GAGAL"
	}
	atas := strings.ToUpper(teks)
	if strings.Contains(atas, "PORNO") {
		return "PORNO"
	}
	if strings.Contains(atas, "AMAN") {
		return "AMAN"
	}
	fmt.Printf("[antibokep] vonis aneh: %s\n", strings.TrimSpace(teks))
	return "GAGAL"
}

type skorLokalBokep struct {
	Nude  float64 `json:"nude"`
	Keras float64 `json:"keras"`
	Fb    float64 `json:"fb"`
	Db    float64 `json:"db"`
	Dr    float64 `json:"dr"`
	Hp    float64 `json:"hp"`
	Sexy  float64 `json:"sexy"`
}

func vonisLokalBokep(uri string) *skorLokalBokep {
	alamat := os.Getenv("ANTIBOKEP_LOCAL_URL")
	if alamat == "" {
		alamat = "http://127.0.0.1:8771/check"
	}
	idx := strings.Index(uri, ",")
	if idx < 0 {
		return nil
	}
	buf, err := base64.StdEncoding.DecodeString(uri[idx+1:])
	if err != nil || len(buf) == 0 {
		return nil
	}
	req, err := http.NewRequest("POST", alamat, bytes.NewReader(buf))
	if err != nil {
		return nil
	}
	req.Header.Set("content-type", "image/jpeg")
	klien := &http.Client{Timeout: 8 * time.Second}
	resp, err := klien.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil
	}
	var s skorLokalBokep
	if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
		return nil
	}
	return &s
}

func vonisFinalBokep(uri string, tanpaRouter bool, batasVisi time.Duration) string {
	if batasVisi <= 0 {
		batasVisi = 45 * time.Second
	}
	if os.Getenv("ANTIBOKEP_LOCAL") == "off" {
		if tanpaRouter {
			return "GAGAL"
		}
		return vonisVisiBokep(uri, batasVisi)
	}
	s := vonisLokalBokep(uri)
	if s == nil {
		if tanpaRouter {
			return "GAGAL"
		}
		return vonisVisiBokep(uri, batasVisi)
	}
	skor := fmt.Sprintf("nude=%v keras=%v fb=%v db=%v dr=%v hp=%v sexy=%v", s.Nude, s.Keras, s.Fb, s.Db, s.Dr, s.Hp, s.Sexy)
	if s.Keras >= 0.6 || s.Fb >= 0.85 || (s.Db >= 0.5 && s.Dr >= 0.5) {
		fmt.Printf("[antibokep] lokal %s -> PORNO\n", skor)
		return "PORNO"
	}
	if s.Nude < 0.35 && s.Db < 0.2 && s.Hp < 0.35 && s.Sexy < 0.6 {
		fmt.Printf("[antibokep] lokal %s -> AMAN\n", skor)
		return "AMAN"
	}
	if tanpaRouter {
		fmt.Printf("[antibokep] lokal abu-abu %s -> tanpa router (GAGAL)\n", skor)
		return "GAGAL"
	}
	fmt.Printf("[antibokep] lokal abu-abu %s -> eskalasi visi\n", skor)
	return vonisVisiBokep(uri, batasVisi)
}

func TanganiAntibokep(k *lib.Konteks) error {
	m := k.M
	if !m.IsGroup {
		return k.Conn.OK(m, "Cuma buat grup")
	}
	arg := strings.ToLower(argPertama(m))
	mode := modeGrup("antibokep", m.From)
	gc, err := k.Conn.GetGroupInfo(m.From)
	if err != nil {
		return err
	}
	if !pengirimAdmin(k, gc) && !pengirimOwner(k) {
		return k.Conn.OK(m, "Lu bukan admin")
	}
	prefix := m.Prefix
	if prefix == "" {
		prefix = "#"
	}
	switch arg {
	case "on", "kick":
		if mode == arg {
			return k.Conn.OK(m, "Udah nyala kok")
		}
		if !botAdmin(k, gc) {
			return k.Conn.OK(m, "Gw bukan admin, gak bisa hapus pesan orang")
		}
		setModeGrup("antibokep", m.From, arg)
		if arg == "kick" {
			return k.Conn.OK(m, "Antibokep nyala (mode kick). Foto/stiker porno yang kekirim di grup ini bakal gw hapus + yang ngirim langsung gw kick.")
		}
		return k.Conn.OK(m, "Antibokep nyala (mode hapus). Foto/stiker porno yang kekirim di grup ini bakal gw hapus + yang ngirim gw marahin, gak di-kick.")
	case "off":
		if mode == "" {
			return k.Conn.OK(m, "Emang udah mati")
		}
		setModeGrup("antibokep", m.From, "")
		return k.Conn.OK(m, "Antibokep mati")
	}
	status := mode
	if status == "" {
		status = "off"
	}
	return k.Conn.OK(m, fmt.Sprintf("Antibokep: %s\nPakai %santibokep on/kick/off", status, prefix))
}

type entriVonisBokep struct {
	v   string
	exp time.Time
}

var gudangVonisBokep = struct {
	sync.Mutex
	cache map[string]entriVonisBokep
	janji map[string]chan string
}{cache: map[string]entriVonisBokep{}, janji: map[string]chan string{}}

func CekAntibokep(k *lib.Konteks) (bool, error) {
	m := k.M
	if m == nil || !m.IsGroup || m.FromMe {
		return false, nil
	}
	mode := modeGrup("antibokep", m.From)
	if mode == "" {
		return false, nil
	}
	tipe := m.Type
	kunciMedia := ""
	if ev, ok := m.Full.(*events.Message); ok && ev != nil && ev.Message != nil {
		msg := ev.Message
		switch tipe {
		case "imageMessage":
			kunciMedia = string(msg.GetImageMessage().GetMediaKey())
		case "stickerMessage":
			kunciMedia = string(msg.GetStickerMessage().GetMediaKey())
		default:
			bungkus := ""
			var dalam = msg
			switch {
			case msg.ViewOnceMessage != nil:
				bungkus, dalam = "viewOnceMessage", msg.ViewOnceMessage.GetMessage()
			case msg.ViewOnceMessageV2 != nil:
				bungkus, dalam = "viewOnceMessageV2", msg.ViewOnceMessageV2.GetMessage()
			case msg.ViewOnceMessageV2Extension != nil:
				bungkus, dalam = "viewOnceMessageV2Extension", msg.ViewOnceMessageV2Extension.GetMessage()
			}
			if dalam != nil && dalam.GetImageMessage() != nil {
				fmt.Printf("[antibokep] sekali lihat (%s) dari %s di %s — dicek sebagai imageMessage\n", bungkus, lib.NomorSaja(m.Sender), m.From)
				tipe = "imageMessage"
				kunciMedia = string(dalam.GetImageMessage().GetMediaKey())
			} else {
				if bungkus != "" {
					fmt.Printf("[antibokep] skip sekali lihat bukan foto: bungkus=%s\n", bungkus)
				}
				return false, nil
			}
		}
	}
	if tipe != "imageMessage" && tipe != "stickerMessage" {
		return false, nil
	}
	num := lib.NomorSaja(m.Sender)
	if num == "" {
		return false, nil
	}
	if !m.Timestamp.IsZero() && time.Since(m.Timestamp) > 30*time.Second {
		fmt.Printf("[antibokep] skip media lama (%s)\n", tipe)
		return false, nil
	}
	if kunciMedia == "" {
		kunciMedia = m.ID
	}
	gc, err := k.Conn.GetGroupInfo(m.From)
	if err != nil {
		return false, nil
	}
	if !botAdmin(k, gc) {
		fmt.Println("[antibokep] skip: bot bukan admin di grup ini")
		return false, nil
	}
	modeKick := mode == "kick"
	pakaiVonis := func(vonis string) bool {
		if vonis != "PORNO" {
			return false
		}
		_ = k.Conn.OK(m, "GW TANDAI LU")
		if err := k.Conn.RevokeDari(m.From, m); err != nil {
			fmt.Printf("[antibokep] hapus gagal: %v\n", err)
		}
		if modeKick {
			if err := k.Conn.UpdateGroupParticipants(m.From, []string{m.Sender}, "remove"); err != nil {
				fmt.Printf("[antibokep] kick gagal: %v\n", err)
			} else {
				fmt.Printf("[antibokep] %s di-kick dari %s (porno)\n", num, m.From)
				botNum := k.Conn.BotNomor()
				sebut := []string{}
				jidAdmin := []string{}
				for _, p := range gc.Peserta {
					if !pesertaAdmin(&p) {
						continue
					}
					n := p.Nomor
					if n == "" {
						n = lib.NomorSaja(p.JID)
					}
					if n == "" || n == botNum {
						continue
					}
					sebut = append(sebut, "@"+n)
					jidAdmin = append(jidAdmin, n+"@s.whatsapp.net")
				}
				if len(jidAdmin) > 0 {
					teks := fmt.Sprintf("%s — @%s barusan gw kick, ngirim porno", strings.Join(sebut, " "), num)
					_ = k.Conn.SendMessage(m.From, lib.IsiPesan{Text: teks, Mentions: append(jidAdmin, m.Sender)}, nil)
				}
			}
		}
		return true
	}

	gudangVonisBokep.Lock()
	if c, ada := gudangVonisBokep.cache[kunciMedia]; ada && time.Now().Before(c.exp) {
		gudangVonisBokep.Unlock()
		fmt.Printf("[antibokep] %s dari %s di %s -> %s (cache)\n", tipe, num, m.From, c.v)
		return pakaiVonis(c.v), nil
	}
	if j, ada := gudangVonisBokep.janji[kunciMedia]; ada {
		gudangVonisBokep.Unlock()
		v := <-j
		fmt.Printf("[antibokep] %s dari %s di %s -> %s (ikut cek yg lagi jalan)\n", tipe, num, m.From, v)
		return pakaiVonis(v), nil
	}
	j := make(chan string, 1)
	gudangVonisBokep.janji[kunciMedia] = j
	gudangVonisBokep.Unlock()

	vonis := func() string {
		buf, _, err := k.Conn.DownloadAny(m)
		if err != nil {
			fmt.Printf("[antibokep] download gagal: %v\n", err)
			return "GAGAL"
		}
		if len(buf) == 0 {
			fmt.Printf("[antibokep] skip: buffer kosong, type=%s\n", tipe)
			return "GAGAL"
		}
		uri, kbKirim, err := kompresiBokep(buf)
		if err != nil {
			mime := m.Mime
			if mime == "" {
				mime = "image/jpeg"
			}
			uri = "data:" + mime + ";base64," + base64.StdEncoding.EncodeToString(buf)
			kbKirim = len(buf) / 1024
			fmt.Printf("[antibokep] kompresi gagal, kirim mentah (%dKB)\n", kbKirim)
		}
		v := vonisFinalBokep(uri, false, 45*time.Second)
		fmt.Printf("[antibokep] %s dari %s di %s (asli %dKB, kirim %dKB) -> %s\n", tipe, num, m.From, len(buf)/1024, kbKirim, v)
		return v
	}()

	gudangVonisBokep.Lock()
	if vonis != "GAGAL" {
		gudangVonisBokep.cache[kunciMedia] = entriVonisBokep{v: vonis, exp: time.Now().Add(10 * time.Minute)}
	}
	delete(gudangVonisBokep.janji, kunciMedia)
	gudangVonisBokep.Unlock()
	j <- vonis
	close(j)
	return pakaiVonis(vonis), nil
}

func frameVideoBokep(buf []byte) ([]byte, error) {
	dir, err := os.MkdirTemp("", "bokepvid-")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(dir)
	masuk := filepath.Join(dir, "masuk.mp4")
	keluar := filepath.Join(dir, "keluar.jpg")
	if err := os.WriteFile(masuk, buf, 0o644); err != nil {
		return nil, err
	}
	cmd := exec.Command("ffmpeg", "-y", "-v", "error", "-i", masuk, "-frames:v", "1", "-q:v", "3", keluar)
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	return os.ReadFile(keluar)
}

func VonisMedia(buf []byte, mimetype string, tanpaRouter bool) string {
	b := buf
	if strings.HasPrefix(mimetype, "video/") {
		f, err := frameVideoBokep(buf)
		if err != nil {
			fmt.Printf("[antibokep] vonisMedia gagal: %v\n", err)
			return "GAGAL"
		}
		b = f
	}
	if len(b) == 0 {
		return "GAGAL"
	}
	uri, _, err := kompresiBokep(b)
	if err != nil {
		fmt.Printf("[antibokep] vonisMedia gagal: %v\n", err)
		return "GAGAL"
	}
	return vonisFinalBokep(uri, tanpaRouter, 12*time.Second)
}

var polaTagPikir = regexp.MustCompile(`(?is)<(think|thought|thinking|reasoning)>.*?</(think|thought|thinking|reasoning)>|</?(think|thought|thinking|reasoning)>`)
var polaSpasiBokep = regexp.MustCompile(`\s+`)

func DeskripsiMedia(buf []byte, mimetype string) string {
	b := buf
	if strings.HasPrefix(mimetype, "video/") {
		f, err := frameVideoBokep(buf)
		if err != nil {
			fmt.Printf("[antibokep] deskripsiMedia gagal: %v\n", err)
			return ""
		}
		b = f
	}
	if len(b) == 0 {
		return ""
	}
	uri, _, err := kompresiBokep(b)
	if err != nil {
		fmt.Printf("[antibokep] deskripsiMedia gagal: %v\n", err)
		return ""
	}
	teks, ok := panggilVisiBokep(uri, promptDeskripsiBokep, 220, 45*time.Second)
	if !ok {
		return ""
	}
	teks = polaTagPikir.ReplaceAllString(teks, "")
	teks = polaSpasiBokep.ReplaceAllString(teks, " ")
	teks = strings.TrimSpace(teks)
	if len(teks) > 500 {
		teks = teks[:500]
	}
	return teks
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "antibokep.js",
		Jenis:    "group",
		Nama:     []string{"#antibokep on/off"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(antibokep)$"),
		Jalankan: TanganiAntibokep,
	})
}
