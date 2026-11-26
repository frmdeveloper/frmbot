package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"gow1/lib"
)

var PathSifat = "bantuan/sifat.md"

type panggilanTool struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Function struct {
		Name      string `json:"name"`
		Arguments string `json:"arguments"`
	} `json:"function"`
}

type pesanObrolan struct {
	Role       string          `json:"role"`
	Content    string          `json:"content,omitempty"`
	ToolCalls  []panggilanTool `json:"tool_calls,omitempty"`
	ToolCallID string          `json:"tool_call_id,omitempty"`
}

func muatSifat(nama string) string {
	if nama == "" {
		nama = "FRM Chan"
	}
	b, err := os.ReadFile(PathSifat)
	if err == nil && strings.TrimSpace(string(b)) != "" {
		return strings.ReplaceAll(strings.TrimSpace(string(b)), "{nama}", nama)
	}
	return "Nama lo " + nama + ". Lo nyebelin, gaul, balasan lo sesingkat mungkin, kalo bisa 1 kata doang. Kosakata lo: njir, jir, anjir — dilarang bilang anj dan cok."
}

func nomorPengirim(m *lib.Pesan) string {
	if n := lib.NomorSaja(m.SenderAlt); n != "" {
		return n
	}
	return lib.NomorSaja(m.Sender)
}

func konteksPengirim(k *lib.Konteks) string {
	m := k.M
	wib, _ := time.LoadLocation("Asia/Jakarta")
	tempat := "chat pribadi (DM)"
	if m.IsGroup {
		tempat = "grup WhatsApp"
	}
	return fmt.Sprintf("\n\nKonteks obrolan:\n- Sekarang: %s WIB\n- Yang lagi ngomong sama lo: nomor %s\n- Tempat: %s\n- Kenalin orang dari nomornya.",
		time.Now().In(wib).Format("Monday, 2 January 2006 15:04"), nomorPengirim(m), tempat)
}

type balasanModel struct {
	Choices []struct {
		Message struct {
			Content   string          `json:"content"`
			ToolCalls []panggilanTool `json:"tool_calls"`
		} `json:"message"`
	} `json:"choices"`
}

func panggilModel(msgs []pesanObrolan) (string, []panggilanTool, error) {
	base := os.Getenv("AI_BASE_URL")
	if base == "" {
		base = "http://127.0.0.1:20128/v1"
	}
	model := os.Getenv("AI_MODEL")
	if model == "" {
		model = "chatbot"
	}
	isi, err := json.Marshal(map[string]any{
		"model":             model,
		"messages":          msgs,
		"tools":             skemaTools,
		"temperature":       0.95,
		"frequency_penalty": 0.5,
		"presence_penalty":  0.4,
		"stream":            false,
	})
	if err != nil {
		return "", nil, err
	}
	req, err := http.NewRequest(http.MethodPost, strings.TrimSuffix(base, "/")+"/chat/completions", bytes.NewReader(isi))
	if err != nil {
		return "", nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if kunci := os.Getenv("AI_API_KEY"); kunci != "" {
		req.Header.Set("Authorization", "Bearer "+kunci)
	}
	klien := &http.Client{Timeout: 120 * time.Second}
	resp, err := klien.Do(req)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()
	mentah, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", nil, fmt.Errorf("model balas HTTP %d: %s", resp.StatusCode, strings.TrimSpace(string(mentah))[:min(200, len(mentah))])
	}
	var hasil balasanModel
	if err := json.Unmarshal(mentah, &hasil); err != nil {
		return "", nil, err
	}
	if len(hasil.Choices) == 0 {
		return "", nil, fmt.Errorf("model gak ngasih pilihan jawaban")
	}
	return strings.TrimSpace(hasil.Choices[0].Message.Content), hasil.Choices[0].Message.ToolCalls, nil
}

func jalanObrolan(k *lib.Konteks, msgs []pesanObrolan) (string, error) {
	maks := 15
	if v := os.Getenv("MCP_MAX_STEPS"); v != "" {
		var n int
		if _, err := fmt.Sscanf(v, "%d", &n); err == nil && n > 0 {
			maks = n
		}
	}
	for langkah := 0; langkah < maks; langkah++ {
		isi, panggilan, err := panggilModel(msgs)
		if err != nil {
			return "", err
		}
		if len(panggilan) == 0 {
			return isi, nil
		}
		msgs = append(msgs, pesanObrolan{Role: "assistant", Content: isi, ToolCalls: panggilan})
		for _, p := range panggilan {
			hasil := jalankanTool(k, p.Function.Name, p.Function.Arguments)
			msgs = append(msgs, pesanObrolan{Role: "tool", ToolCallID: p.ID, Content: hasil})
		}
	}
	return "", fmt.Errorf("kebanyakan langkah tool (maks %d)", maks)
}

var reKataMcp = regexp.MustCompile(`(?i)#?mcp\b`)

func konteksDariQuote(m *lib.Pesan) (pesanObrolan, pesanObrolan, bool) {
	if m.Quoted == nil {
		return pesanObrolan{}, pesanObrolan{}, false
	}
	for _, gm := range m.Quoted.GroupMentions {
		if strings.HasPrefix(gm.Subjek, "ask|") {
			sebelum := strings.TrimSpace(strings.TrimPrefix(gm.Subjek, "ask|"))
			jawaban := strings.TrimSpace(m.Quoted.Text)
			if sebelum != "" && jawaban != "" {
				return pesanObrolan{Role: "user", Content: sebelum},
					pesanObrolan{Role: "assistant", Content: jawaban}, true
			}
		}
	}
	return pesanObrolan{}, pesanObrolan{}, false
}

func potongRunes(s string, maks int) string {
	r := []rune(s)
	if len(r) <= maks {
		return s
	}
	return string(r[:maks])
}

func Ngobrol(k *lib.Konteks, prompt string) error {
	m := k.M
	prompt = strings.TrimSpace(prompt)
	if prompt == "" {
		return nil
	}
	_ = k.Conn.SendChatPresence(m.From, "composing")
	defer func() { _ = k.Conn.SendChatPresence(m.From, "paused") }()
	msgs := []pesanObrolan{{Role: "system", Content: muatSifat(k.Conn.BotNama()) + konteksPengirim(k)}}
	if sebelum, jawaban, ok := konteksDariQuote(m); ok {
		msgs = append(msgs, sebelum, jawaban)
	} else if m.Quoted != nil && strings.TrimSpace(m.Quoted.Text) != "" {
		prompt = "[ngebalas pesan: \"" + strings.TrimSpace(m.Quoted.Text) + "\"]\n" + prompt
	}
	msgs = append(msgs, pesanObrolan{Role: "user", Content: prompt})
	jawab, err := jalanObrolan(k, msgs)
	if err != nil {
		return err
	}
	jawab = strings.TrimSpace(reKataMcp.ReplaceAllString(jawab, ""))
	if jawab == "" {
		return nil
	}
	gms := []lib.GrupMention{
		{JID: "120363186235853203@g.us", Subjek: "ask|" + potongRunes(prompt, 500)},
		{JID: "1234" + lib.RandNum(14) + "@g.us", Subjek: jawab},
	}
	return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: jawab, GroupMentions: gms}, m)
}

func TanganiMcp(k *lib.Konteks) error {
	return Ngobrol(k, k.M.Text)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "mcp.js",
		Jenis:    "ai",
		Nama:     nil,
		Itunya:   regexp.MustCompile("(?i)^(assalamualaikum|woy|woi|halo|hi|hai|p|oi|oy|frm|oagi|pagi|siang|sore|malam|selamat|bangun|(#|\\!|\\/|\\.)(mcp|ai|jembot))$"),
		Jalankan: TanganiMcp,
	})
}
