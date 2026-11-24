package info

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"gow1/lib"

	waE2E "go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

type videoYT struct {
	URL    string
	Image  string
	Title  string
	Author string
	Tim    string
}

var (
	reYts2     = regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(yts2)$")
	reVidID    = regexp.MustCompile(`"videoId":"([A-Za-z0-9_-]{11})"`)
	reVidTitle = regexp.MustCompile(`"title":\{"runs":\[\{"text":"((?:[^"\\]|\\.)*)"`)
	reVidAuth  = regexp.MustCompile(`"longBylineText":\{"runs":\[\{"text":"((?:[^"\\]|\\.)*)"`)
	reVidLen   = regexp.MustCompile(`"lengthText":\{(?:[^{}]|\{[^{}]*\})*?"simpleText":"([^"]+)"`)
	reVidThumb = regexp.MustCompile(`"url":"(https://i\.ytimg\.com[^"]+)"`)
)

func lepasEscapeJSON(s string) string {
	if u, err := strconv.Unquote("\"" + s + "\""); err == nil {
		return u
	}
	return s
}

func cariYT(kata string) ([]videoYT, error) {
	alamat := "https://www.youtube.com/results?search_query=" + url.QueryEscape(kata)
	b, err := lib.AmbilURL(alamat, map[string]string{
		"Accept-Language": "en-US,en;q=0.9",
	})
	if err != nil {
		return nil, err
	}
	html := string(b)
	bagian := strings.Split(html, "\"videoRenderer\":{")
	var hasil []videoYT
	udah := map[string]bool{}
	for _, bgn := range bagian[1:] {
		if len(bgn) > 8000 {
			bgn = bgn[:8000]
		}
		mv := reVidID.FindStringSubmatch(bgn)
		if mv == nil {
			continue
		}
		id := mv[1]
		if udah[id] {
			continue
		}
		udah[id] = true
		v := videoYT{URL: "https://www.youtube.com/watch?v=" + id, Image: "https://i.ytimg.com/vi/" + id + "/hqdefault.jpg"}
		if mt := reVidTitle.FindStringSubmatch(bgn); mt != nil {
			v.Title = lepasEscapeJSON(mt[1])
		}
		if ma := reVidAuth.FindStringSubmatch(bgn); ma != nil {
			v.Author = lepasEscapeJSON(ma[1])
		}
		if ml := reVidLen.FindStringSubmatch(bgn); ml != nil {
			v.Tim = ml[1]
		}
		thumbs := reVidThumb.FindAllStringSubmatch(bgn, -1)
		if len(thumbs) > 0 {
			v.Image = strings.ReplaceAll(thumbs[len(thumbs)-1][1], "\\u0026", "&")
		}
		hasil = append(hasil, v)
		if len(hasil) >= 20 {
			break
		}
	}
	return hasil, nil
}

func kirimGrupMentionYT(k *lib.Konteks, teks, subjek string) error {
	ci := &waE2E.ContextInfo{
		GroupMentions: []*waE2E.GroupMention{
			{GroupJID: proto.String("120363186235853203@g.us"), GroupSubject: proto.String(subjek)},
		},
	}
	if ev, ok := k.M.Full.(*events.Message); ok && ev != nil && ev.Message != nil {
		ci.StanzaID = proto.String(k.M.ID)
		if k.M.Sender != "" {
			ci.Participant = proto.String(k.M.Sender)
		}
		ci.QuotedMessage = ev.Message
	}
	pesan := &waE2E.Message{ExtendedTextMessage: &waE2E.ExtendedTextMessage{
		Text:        proto.String(teks),
		ContextInfo: ci,
	}}
	return k.Conn.RelayMentah(k.M.From, pesan)
}

func TanganiYtsearch(k *lib.Konteks) error {
	m := k.M
	if reYts2.MatchString(m.Command) {
		bagian := strings.Split(m.Q, "|")
		linkyt := bagian[0]
		linkthumb := ""
		if len(bagian) > 1 {
			linkthumb = bagian[1]
		}
		jsp := map[string]string{"1": "#ytmp4 " + linkyt, "2": "#ytmp3 " + linkyt}
		jspJSON, _ := json.Marshal(jsp)
		caption := fmt.Sprintf("Download %s\n\n    1. Video\n    2. Audio", linkyt)
		if thumb, err := lib.AmbilURL(linkthumb, nil); err == nil && len(thumb) > 0 {
			if err := k.Conn.SendMessage(m.From, lib.IsiPesan{Image: thumb, Caption: caption}, m); err == nil {
				return nil
			}
		}
		return kirimGrupMentionYT(k, caption, "json|"+string(jspJSON))
	}
	if m.Q == "" {
		return k.Conn.OK(m, "Judulnya mana?")
	}
	videos, err := cariYT(m.Q)
	if err != nil {
		return err
	}
	if len(videos) == 0 {
		return k.Conn.OK(m, "Gak ada njir")
	}
	text := "Usap kekanan dan balas angka!\n\n"
	jsonPeta := map[string]string{}
	for i, b := range videos {
		no := i + 1
		jsonPeta[strconv.Itoa(no)] = "#yts2 " + b.URL + "|" + b.Image
		text += fmt.Sprintf("%d. %s\n*%s* Durasi: %s\n\n", no, b.Title, b.Author, b.Tim)
	}
	jsonStr, _ := json.Marshal(jsonPeta)
	if err := kirimGrupMentionYT(k, text, "json|"+string(jsonStr)); err != nil {
		return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: text}, m)
	}
	return nil
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "ytsearch.js",
		Jenis:    "info",
		Nama:     []string{"#ytsearch"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(yts|ytsearch|ytcari|yts2)$"),
		Jalankan: TanganiYtsearch,
	})
}
