package maker

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	mrand "math/rand"
	"net/http"
	"regexp"

	"gow1/lib"
)

type quotePhoto struct {
	URL string `json:"url"`
}

type quoteFrom struct {
	ID    int        `json:"id"`
	Name  string     `json:"name"`
	Photo quotePhoto `json:"photo"`
}

type quotePesan struct {
	Entities     []any          `json:"entities"`
	Avatar       bool           `json:"avatar"`
	From         quoteFrom      `json:"from"`
	Text         string         `json:"text"`
	ReplyMessage map[string]any `json:"replyMessage"`
	Media        *quotePhoto    `json:"media,omitempty"`
}

type quoteBody struct {
	Type            string       `json:"type"`
	Format          string       `json:"format"`
	BackgroundColor string       `json:"backgroundColor"`
	Width           int          `json:"width"`
	Height          int          `json:"height"`
	Scale           int          `json:"scale"`
	Messages        []quotePesan `json:"messages"`
}

func BuatQuote(avatar string, name string, media string, text string) ([]byte, error) {
	pesan := quotePesan{
		Entities: []any{},
		Avatar:   true,
		From: quoteFrom{
			ID:    mrand.Intn(10) + 1,
			Name:  name,
			Photo: quotePhoto{URL: avatar},
		},
		Text:         text,
		ReplyMessage: map[string]any{},
	}
	if media != "" {
		pesan.Media = &quotePhoto{URL: media}
	}
	body := quoteBody{
		Type:            "",
		Format:          "png",
		BackgroundColor: "#FFFFFF",
		Width:           512,
		Height:          768,
		Scale:           2,
		Messages:        []quotePesan{pesan},
	}
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "https://quote.yuri.ly/generate", bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := lib.KlienHTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var hasil struct {
		OK     bool `json:"ok"`
		Result struct {
			Image string `json:"image"`
		} `json:"result"`
		Error *struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&hasil); err != nil {
		return nil, err
	}
	if !hasil.OK {
		if hasil.Error != nil && hasil.Error.Message != "" {
			return nil, fmt.Errorf("%s", hasil.Error.Message)
		}
		return nil, fmt.Errorf("API quote gagal dibales")
	}
	return base64.StdEncoding.DecodeString(hasil.Result.Image)
}

func TanganiQuote(k *lib.Konteks) error {
	m := k.M
	if m.Q == "" {
		return k.Conn.OK(m, "Gak gitu kocak 🤭")
	}
	avatar, err := k.Conn.GetProfilePicture(m.Sender)
	if err != nil || avatar == "" {
		avatar = "https://telegra.ph/file/9932abc07a99e4938a4e7.png"
	}
	nomor := lib.NomorSaja(m.SenderAlt)
	if nomor == "" {
		nomor = lib.NomorSaja(m.Sender)
	}
	buat, err := BuatQuote(avatar, "+"+nomor, "", m.Q)
	if err != nil {
		return err
	}
	webp, err := Img2webp(buat)
	if err != nil {
		return err
	}
	qcc, err := AddExifBawaan(webp)
	if err != nil {
		return err
	}
	simpanStkdb(m, qcc)
	return k.Conn.SendMessage(m.From, lib.IsiPesan{Sticker: qcc}, m)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "quote.js",
		Jenis:    "maker",
		Nama:     []string{"#qc"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(qc|quote|quickchat)$"),
		Jalankan: TanganiQuote,
	})
}
