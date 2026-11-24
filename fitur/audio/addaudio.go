package audio

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	waE2E "go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"

	"gow1/lib"
)

const penandaJIDAddaudio = "120363186235853203@g.us"

func kirimPenandaAddaudio(k *lib.Konteks, teks, subjek string) error {
	ci := &waE2E.ContextInfo{
		GroupMentions: []*waE2E.GroupMention{
			{GroupJID: proto.String(penandaJIDAddaudio), GroupSubject: proto.String(subjek)},
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

func TanganiAddaudio(k *lib.Konteks) error {
	m := k.M
	args := strings.Fields(m.Q)

	if len(args) == 0 {
		if m.Type == "imageMessage" {
			data, _, err := k.Conn.DownloadAny(m)
			if err != nil {
				return err
			}
			img := "./sampah/" + m.ID + m.Sender + ".jpeg"
			if err := os.MkdirAll(filepath.Dir(img), 0o755); err != nil {
				return err
			}
			if err := os.WriteFile(img, data, 0o644); err != nil {
				return err
			}
			return kirimPenandaAddaudio(k, "Sekarang kirim audionya coba", "cmd|#addaudio "+img)
		}
		return kirimPenandaAddaudio(k, "Kirim fotonya coba", "cmd|#addaudio ")
	}

	if m.Type == "audioMessage" {
		data, _, err := k.Conn.DownloadAny(m)
		if err != nil {
			return err
		}
		aud := "./sampah/" + m.ID + m.Sender + ".mp3"
		if err := os.MkdirAll(filepath.Dir(aud), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(aud, data, 0o644); err != nil {
			return err
		}
		return kirimPenandaAddaudio(k, "Yakin dek?", fmt.Sprintf("cmd|#addaudio %s %s", args[0], aud))
	}

	if len(args) >= 2 {
		if err := k.Conn.OK(m, fmt.Sprintf("Bentar lagi manggil\n\n%s %s %s", m.Command, args[0], args[1])); err != nil {
			return err
		}
		out := "./sampah/" + m.ID + m.Sender + ".mp4"
		perintah := fmt.Sprintf(`ffmpeg -loop 1 -i %s -i %s -c:v libx264 -tune stillimage -crf 23 -preset medium -c:a aac -b:a 192k -pix_fmt yuv420p -vf "pad=ceil(iw/2)*2:ceil(ih/2)*2:(ow-iw)/2:(oh-ih)/2" -shortest -movflags +faststart %s`, args[0], args[1], out)
		_, stderr := lib.SpawnSync(perintah)
		if strings.Contains(stderr, "Conversion failed!") {
			return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: "ERROR JIR\n\n" + stderr}, m)
		}
		data, err := os.ReadFile(out)
		if err != nil {
			return err
		}
		return k.Conn.SendMessage(m.From, lib.IsiPesan{Video: data, Mimetype: "video/mp4"}, m)
	}

	return nil
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "addaudio.js",
		Jenis:    "audio",
		Nama:     []string{"#addaudio"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(addaudio)$"),
		Jalankan: TanganiAddaudio,
	})
}
