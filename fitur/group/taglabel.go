package group

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"gow1/lib"

	waE2E "go.mau.fi/whatsmeow/proto/waE2E"
	"google.golang.org/protobuf/proto"
)

func kirimLabel(k *lib.Konteks, grup, label string) error {
	pesan := &waE2E.Message{ProtocolMessage: &waE2E.ProtocolMessage{
		Type:        waE2E.ProtocolMessage_GROUP_MEMBER_LABEL_CHANGE.Enum(),
		MemberLabel: &waE2E.MemberLabel{Label: proto.String(label), LabelTimestamp: proto.Int64(time.Now().UnixMilli())},
	}}
	return k.Conn.RelayMentah(grup, pesan)
}

func TanganiTaglabel(k *lib.Konteks) error {
	m := k.M
	if !pengirimOwner(k) {
		return k.Conn.OK(m, "khusus owner")
	}
	cmd := strings.ToLower(strings.TrimLeft(m.Command, "#!/."))
	label := strings.TrimSpace(m.Q)
	prefix := m.Prefix
	if prefix == "" {
		prefix = "#"
	}
	caraPakai := fmt.Sprintf("Pakai %staglabel <teks> (grup ini) atau %staglabelall <teks> (semua grup)", prefix, prefix)
	if label == "" {
		return k.Conn.OK(m, caraPakai)
	}
	if utf8.RuneCountInString(label) > 30 {
		return k.Conn.OK(m, fmt.Sprintf("Labelnya kepanjangan: %d/30 karakter", utf8.RuneCountInString(label)))
	}
	if cmd == "taglabelall" {
		grup, err := k.Conn.GetJoinedGroups()
		if err != nil {
			return k.Conn.OK(m, fmt.Sprintf("gagal: %v", err))
		}
		ok := 0
		gagal := []string{}
		for _, g := range grup {
			if g == nil || g.JID == "" {
				continue
			}
			if err := kirimLabel(k, g.JID, label); err != nil {
				nama := g.Nama
				if nama == "" {
					nama = g.JID
				}
				gagal = append(gagal, nama)
			} else {
				ok++
			}
		}
		teks := fmt.Sprintf("tag \"%s\" kepasang di %d/%d grup", label, ok, len(grup))
		if len(gagal) > 0 {
			teks += " — gagal: " + strings.Join(gagal, ", ")
		}
		return k.Conn.OK(m, teks)
	}
	if !m.IsGroup {
		return k.Conn.OK(m, caraPakai)
	}
	if err := kirimLabel(k, m.From, label); err != nil {
		return k.Conn.OK(m, fmt.Sprintf("gagal: %v", err))
	}
	return k.Conn.OK(m, fmt.Sprintf("tag kepasang: %s", label))
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "taglabel.js",
		Jenis:    "group",
		Nama:     nil,
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(taglabel|taglabelall)$"),
		Jalankan: TanganiTaglabel,
	})
}
