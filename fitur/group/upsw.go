package group

import (
	"fmt"
	"regexp"
	"strings"

	waE2E "go.mau.fi/whatsmeow/proto/waE2E"
	"google.golang.org/protobuf/proto"

	"gow1/lib"
)

func TanganiUpsw(k *lib.Konteks) error {
	m := k.M
	if !strings.HasSuffix(m.From, "@g.us") {
		return k.Conn.OK(m, "Emng ini grup?")
	}
	if m.Quoted == nil {
		return k.Conn.OK(m, "Bales pesannya yg mau di bikin sw grup")
	}
	sumber := m.Sender
	if m.SenderAlt != "" {
		sumber = m.SenderAlt
	}
	tanda := fmt.Sprintf("\n_diupload oleh +%s_", lib.NomorSaja(sumber))

	asli, ok := m.Quoted.Full.(*waE2E.Message)
	if !ok || asli == nil {
		if m.Quoted.Text == "" {
			return fmt.Errorf("upsw: pesan yang dibalas gak punya isi yang bisa diteruskan")
		}
		bungkus := &waE2E.Message{
			GroupStatusMessageV2: &waE2E.FutureProofMessage{
				Message: &waE2E.Message{Conversation: proto.String(m.Quoted.Text + tanda)},
			},
		}
		return k.Conn.RelayMentah(m.From, bungkus)
	}

	q := proto.Clone(asli).(*waE2E.Message)
	switch {
	case q.Conversation != nil:
		q.Conversation = proto.String(q.GetConversation() + tanda)
	case q.ExtendedTextMessage != nil:
		q.ExtendedTextMessage.Text = proto.String(q.ExtendedTextMessage.GetText() + tanda)
	case q.ImageMessage != nil:
		q.ImageMessage.Caption = proto.String(q.ImageMessage.GetCaption() + tanda)
	case q.VideoMessage != nil:
		q.VideoMessage.Caption = proto.String(q.VideoMessage.GetCaption() + tanda)
	case q.DocumentMessage != nil:
		q.DocumentMessage.Caption = proto.String(q.DocumentMessage.GetCaption() + tanda)
	}
	bungkus := &waE2E.Message{
		GroupStatusMessageV2: &waE2E.FutureProofMessage{Message: q},
	}
	return k.Conn.RelayMentah(m.From, bungkus)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "upsw.js",
		Jenis:    "group",
		Nama:     []string{"#upsw"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(upsw(gc)?)$"),
		Jalankan: TanganiUpsw,
	})
}
