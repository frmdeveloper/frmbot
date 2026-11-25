package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync/atomic"
	"time"

	"go.mau.fi/whatsmeow"
	waCompanionReg "go.mau.fi/whatsmeow/proto/waCompanionReg"
	waE2E "go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"

	_ "github.com/mattn/go-sqlite3"

	fitchat "gow1/fitur/chat"
	fiturgrup "gow1/fitur/group"
	fiturai "gow1/fitur/ai"
	fiturmaker "gow1/fitur/maker"
	fitursystem "gow1/fitur/system"
	"gow1/lib"
)

var rePerintahAntilink = regexp.MustCompile(`(?i)^(#|!|/|\.)antilink$`)

var HitungPesan = struct {
	Masuk    atomic.Int64
	Diproses atomic.Int64
	Perintah atomic.Int64
	Mulai    time.Time
}{Mulai: time.Now()}

type Bot struct {
	user      *lib.DataUser
	group     *lib.DataGroup
	nomor     string
	client    *whatsmeow.Client
	container *sqlstore.Container
	logout    atomic.Bool
	cobaLagi  atomic.Int64
}

func BotBaru(user *lib.DataUser, group *lib.DataGroup, nomor string) *Bot {
	return &Bot{user: user, group: group, nomor: nomor}
}

func (b *Bot) Nomor() string { return b.nomor }

func bukaStore() (*sqlstore.Container, error) {
	return sqlstore.New(context.Background(), "sqlite3",
		"file:data/auth.db?_foreign_keys=on&_pragma=journal_mode(WAL)&_pragma=busy_timeout(10000)", nil)
}

func (b *Bot) Start() error {
	ctx := context.Background()
	container, err := bukaStore()
	if err != nil {
		return fmt.Errorf("buka store sesi: %w", err)
	}
	b.container = container

	devices, err := container.GetAllDevices(ctx)
	if err != nil {
		return fmt.Errorf("baca daftar device: %w", err)
	}
	var device = container.NewDevice()
	for _, d := range devices {
		if d.ID != nil && strings.Split(d.ID.User, "@")[0] == b.nomor {
			device = d
			break
		}
		if d.ID != nil && device.ID == nil {
			device = d
		}
	}

	store.DeviceProps.PlatformType = waCompanionReg.DeviceProps_CHROME.Enum()
	store.DeviceProps.Os = proto.String("Chrome")

	clientLog := waLog.Stdout("Client", "INFO", true)
	b.client = whatsmeow.NewClient(device, clientLog)
	b.client.AddEventHandler(b.tanganiEvent)

	if err := b.client.Connect(); err != nil {
		return fmt.Errorf("connect: %w", err)
	}

	if b.client.Store.ID == nil {
		kode, err := b.client.PairPhone(ctx, b.nomor, true, whatsmeow.PairClientChrome, "Chrome (Linux)")
		if err != nil {
			return fmt.Errorf("pairphone: %w", err)
		}
		fmt.Println("==================================================")
		fmt.Printf("[pairing] KODE PAIRING buat %s: %s\n", b.nomor, kode)
		fmt.Println("[pairing] Di HP nomor itu: WhatsApp > Perangkat tertaut >")
		fmt.Println("[pairing] Tautkan perangkat > Tautkan dengan nomor telepon >")
		fmt.Println("[pairing] masukkan kode di atas.")
		fmt.Println("==================================================")
	} else {
		fmt.Printf("[conn] sesi ketemu buat %s — nyambung tanpa pairing\n", b.client.Store.ID.User)
	}
	return nil
}

func (b *Bot) Stop() {
	if b.client != nil {
		b.client.Disconnect()
	}
}

func LogoutMain() {
	ctx := context.Background()
	container, err := bukaStore()
	if err != nil {
		fmt.Println("[logout] buka store gagal:", err)
		return
	}
	devices, err := container.GetAllDevices(ctx)
	if err != nil || len(devices) == 0 {
		fmt.Println("[logout] tidak ada device di store:", err)
		return
	}
	for _, d := range devices {
		if d.ID == nil {
			_ = container.DeleteDevice(ctx, d)
			fmt.Println("[logout] device belum pairing dibuang dari store")
			continue
		}
		client := whatsmeow.NewClient(d, nil)
		if err := client.Connect(); err != nil {
			fmt.Println("[logout] connect gagal:", err)
			continue
		}
		for i := 0; i < 50 && !client.IsLoggedIn(); i++ {
			time.Sleep(200 * time.Millisecond)
		}
		if err := client.Logout(ctx); err != nil {
			fmt.Println("[logout] logout server gagal:", err)
		} else {
			fmt.Println("[logout] tautan diputus dari server")
		}
		client.Disconnect()
		_ = container.DeleteDevice(ctx, d)
		fmt.Println("[logout] device dihapus dari store lokal")
	}
}

func (b *Bot) tanganiEvent(evt any) {
	switch e := evt.(type) {
	case *events.Connected:
		b.cobaLagi.Store(0)
		fmt.Printf("[conn] Connected — login sebagai %s\n", b.client.Store.ID)
		_ = b.client.SendPresence(context.Background(), types.PresenceAvailable)
	case *events.Disconnected:
		fmt.Println("[conn] Disconnected")
		if !b.logout.Load() {
			go b.sambungUlang()
		}
	case *events.LoggedOut:
		b.logout.Store(true)
		fmt.Printf("[conn] LoggedOut (%v) — device dihapus, restart gow1 buat pairing ulang\n", e.Reason)
		if b.client != nil && b.client.Store != nil {
			_ = b.container.DeleteDevice(context.Background(), b.client.Store)
		}
	case *events.GroupInfo:
		k := &lib.Konteks{Conn: b, M: &lib.Pesan{From: e.JID.String(), IsGroup: true}, User: b.user, Group: b.group}
		if err := fiturgrup.AksiGrup(k, e); err != nil {
			fmt.Println("[agrup] action error:", err)
		}
	case *events.Message:
		m := konversiPesan(e)
		if m == nil || m.From == "status@broadcast" {
			return
		}
		b.TanganiPesan(m)
	}
}

func (b *Bot) sambungUlang() {
	coba := b.cobaLagi.Add(1)
	jeda := time.Duration(min(5*coba, 60)) * time.Second
	fmt.Printf("[reconnect] percobaan ke-%d dalam %s\n", coba, jeda)
	time.Sleep(jeda)
	if b.logout.Load() || b.client.IsConnected() {
		return
	}
	if err := b.client.Connect(); err != nil {
		fmt.Println("[reconnect] gagal:", err)
		go b.sambungUlang()
	}
}

func konversiPesan(e *events.Message) *lib.Pesan {
	msg := e.Message
	if msg == nil {
		return nil
	}
	var teks, tipe, media, mime string
	var ctxInfo *waE2E.ContextInfo
	switch {
	case msg.Conversation != nil:
		teks, tipe = msg.GetConversation(), "conversation"
	case msg.ExtendedTextMessage != nil:
		teks, tipe = msg.ExtendedTextMessage.GetText(), "extendedTextMessage"
		ctxInfo = msg.ExtendedTextMessage.GetContextInfo()
	case msg.ImageMessage != nil:
		teks, tipe = msg.ImageMessage.GetCaption(), "imageMessage"
		media, mime = "image", msg.ImageMessage.GetMimetype()
		ctxInfo = msg.ImageMessage.GetContextInfo()
	case msg.VideoMessage != nil:
		teks, tipe = msg.VideoMessage.GetCaption(), "videoMessage"
		media, mime = "video", msg.VideoMessage.GetMimetype()
		ctxInfo = msg.VideoMessage.GetContextInfo()
	case msg.DocumentMessage != nil:
		teks, tipe = msg.DocumentMessage.GetCaption(), "documentMessage"
		media, mime = "document", msg.DocumentMessage.GetMimetype()
		ctxInfo = msg.DocumentMessage.GetContextInfo()
	case msg.AudioMessage != nil:
		tipe, media, mime = "audioMessage", "audio", msg.AudioMessage.GetMimetype()
		ctxInfo = msg.AudioMessage.GetContextInfo()
	case msg.StickerMessage != nil:
		tipe, media, mime = "stickerMessage", "sticker", msg.StickerMessage.GetMimetype()
		ctxInfo = msg.StickerMessage.GetContextInfo()
	case msg.AudioStickerMessage != nil:
		tipe, media = "audioStickerMessage", "sticker"
		if dalam := msg.AudioStickerMessage.GetMessage().GetStickerMessage(); dalam != nil {
			mime = dalam.GetMimetype()
			ctxInfo = dalam.GetContextInfo()
		}
	default:
		tipe = "other"
	}

	m := &lib.Pesan{
		ID:        e.Info.ID,
		From:      e.Info.Chat.String(),
		Sender:    e.Info.Sender.String(),
		SenderAlt: e.Info.SenderAlt.String(),
		Text:      teks,
		IsGroup:   e.Info.Chat.Server == types.GroupServer,
		FromMe:    e.Info.IsFromMe,
		Type:      tipe,
		Media:     media,
		Mime:      mime,
		Full:      e,
	}
	if ctxInfo != nil {
		m.Mentions = ctxInfo.GetMentionedJID()
		m.GroupMentions = grupMentionsDariCI(ctxInfo)
		if ctxInfo.GetStanzaID() != "" {
			q := &lib.Pesan{
				ID:     ctxInfo.GetStanzaID(),
				From:   m.From,
				Sender: ctxInfo.GetParticipant(),
				Full:   ctxInfo.GetQuotedMessage(),
			}
			if qm := ctxInfo.GetQuotedMessage(); qm != nil {
				q.GroupMentions = grupMentionsDariPesan(qm)
				switch {
				case qm.Conversation != nil:
					q.Text, q.Type = qm.GetConversation(), "conversation"
				case qm.ExtendedTextMessage != nil:
					q.Text, q.Type = qm.ExtendedTextMessage.GetText(), "extendedTextMessage"
				case qm.ImageMessage != nil:
					q.Text, q.Type, q.Media, q.Mime = qm.ImageMessage.GetCaption(), "imageMessage", "image", qm.ImageMessage.GetMimetype()
				case qm.VideoMessage != nil:
					q.Text, q.Type, q.Media, q.Mime = qm.VideoMessage.GetCaption(), "videoMessage", "video", qm.VideoMessage.GetMimetype()
				case qm.AudioMessage != nil:
					q.Type, q.Media, q.Mime = "audioMessage", "audio", qm.AudioMessage.GetMimetype()
				case qm.StickerMessage != nil:
					q.Type, q.Media, q.Mime = "stickerMessage", "sticker", qm.StickerMessage.GetMimetype()
				case qm.AudioStickerMessage != nil:
					q.Type, q.Media = "audioStickerMessage", "sticker"
					if dalam := qm.AudioStickerMessage.GetMessage().GetStickerMessage(); dalam != nil {
						q.Mime = dalam.GetMimetype()
					}
				case qm.DocumentMessage != nil:
					q.Text, q.Type, q.Media, q.Mime = qm.DocumentMessage.GetCaption(), "documentMessage", "document", qm.DocumentMessage.GetMimetype()
				}
			}
			m.Quoted = q
		}
	}

	bidang := strings.Fields(teks)
	if len(bidang) > 0 {
		m.Command, m.Q = gabungPrefixSpasi(bidang)
		if r := []rune(m.Command); len(r) > 0 && strings.ContainsRune("!#%./\\", r[0]) {
			m.Prefix = string(r[0])
		}
	}
	return m
}

func gabungPrefixSpasi(bidang []string) (string, string) {
	if len(bidang) == 0 {
		return "", ""
	}
	if len(bidang) >= 2 {
		if r := []rune(bidang[0]); len(r) == 1 && strings.ContainsRune("!#%./\\", r[0]) {
			q := ""
			if len(bidang) > 2 {
				q = strings.Join(bidang[2:], " ")
			}
			return strings.ToLower(bidang[0] + bidang[1]), q
		}
	}
	q := ""
	if len(bidang) > 1 {
		q = strings.Join(bidang[1:], " ")
	}
	return strings.ToLower(bidang[0]), q
}

func grupMentionsKeProto(gms []lib.GrupMention) []*waE2E.GroupMention {
	hasil := make([]*waE2E.GroupMention, 0, len(gms))
	for _, g := range gms {
		hasil = append(hasil, &waE2E.GroupMention{
			GroupJID:     proto.String(g.JID),
			GroupSubject: proto.String(g.Subjek),
		})
	}
	return hasil
}

func grupMentionsDariCI(ci *waE2E.ContextInfo) []lib.GrupMention {
	if ci == nil {
		return nil
	}
	hasil := []lib.GrupMention{}
	for _, g := range ci.GetGroupMentions() {
		hasil = append(hasil, lib.GrupMention{JID: g.GetGroupJID(), Subjek: g.GetGroupSubject()})
	}
	return hasil
}

func grupMentionsDariPesan(qm *waE2E.Message) []lib.GrupMention {
	if qm == nil {
		return nil
	}
	var ci *waE2E.ContextInfo
	switch {
	case qm.ExtendedTextMessage != nil:
		ci = qm.ExtendedTextMessage.GetContextInfo()
	case qm.ImageMessage != nil:
		ci = qm.ImageMessage.GetContextInfo()
	case qm.VideoMessage != nil:
		ci = qm.VideoMessage.GetContextInfo()
	case qm.DocumentMessage != nil:
		ci = qm.DocumentMessage.GetContextInfo()
	case qm.AudioMessage != nil:
		ci = qm.AudioMessage.GetContextInfo()
	case qm.StickerMessage != nil:
		ci = qm.StickerMessage.GetContextInfo()
	case qm.AudioStickerMessage != nil:
		if dalam := qm.AudioStickerMessage.GetMessage().GetStickerMessage(); dalam != nil {
			ci = dalam.GetContextInfo()
		}
	}
	return grupMentionsDariCI(ci)
}

func (b *Bot) konteksInfo(isi lib.IsiPesan, quoted *lib.Pesan) *waE2E.ContextInfo {
	if quoted == nil && len(isi.Mentions) == 0 && len(isi.GroupMentions) == 0 {
		return nil
	}
	ci := &waE2E.ContextInfo{MentionedJID: isi.Mentions, GroupMentions: grupMentionsKeProto(isi.GroupMentions)}
	if quoted != nil {
		ci.StanzaID = proto.String(quoted.ID)
		if quoted.Sender != "" {
			ci.Participant = proto.String(quoted.Sender)
		}
		switch f := quoted.Full.(type) {
		case *events.Message:
			ci.QuotedMessage = f.Message
		case *waE2E.Message:
			ci.QuotedMessage = f
		}
	}
	return ci
}

func (b *Bot) SendMessage(jid string, isi lib.IsiPesan, quoted *lib.Pesan) error {
	if b.client == nil {
		return fmt.Errorf("client belum tersambung")
	}
	target, err := types.ParseJID(jid)
	if err != nil {
		return err
	}
	ctx := context.Background()
	ci := b.konteksInfo(isi, quoted)
	caption := isi.Caption
	if caption == "" {
		caption = isi.Text
	}

	var pesan *waE2E.Message
	switch {
	case isi.Image != nil:
		up, err := b.client.Upload(ctx, isi.Image, whatsmeow.MediaImage)
		if err != nil {
			return err
		}
		mime := isi.Mimetype
		if mime == "" {
			mime = lib.DeteksiMime(isi.Image)
		}
		pesan = &waE2E.Message{ImageMessage: &waE2E.ImageMessage{
			URL: proto.String(up.URL), DirectPath: proto.String(up.DirectPath),
			MediaKey: up.MediaKey, FileEncSHA256: up.FileEncSHA256, FileSHA256: up.FileSHA256,
			FileLength: proto.Uint64(up.FileLength), Mimetype: proto.String(mime),
			Caption: proto.String(caption), ContextInfo: ci,
		}}
	case isi.Video != nil:
		up, err := b.client.Upload(ctx, isi.Video, whatsmeow.MediaVideo)
		if err != nil {
			return err
		}
		mime := isi.Mimetype
		if mime == "" {
			mime = "video/mp4"
		}
		pesan = &waE2E.Message{VideoMessage: &waE2E.VideoMessage{
			URL: proto.String(up.URL), DirectPath: proto.String(up.DirectPath),
			MediaKey: up.MediaKey, FileEncSHA256: up.FileEncSHA256, FileSHA256: up.FileSHA256,
			FileLength: proto.Uint64(up.FileLength), Mimetype: proto.String(mime),
			Caption: proto.String(caption), ContextInfo: ci,
		}}
	case isi.Audio != nil:
		up, err := b.client.Upload(ctx, isi.Audio, whatsmeow.MediaAudio)
		if err != nil {
			return err
		}
		mime := isi.Mimetype
		if mime == "" {
			mime = lib.DeteksiMime(isi.Audio)
		}
		pesan = &waE2E.Message{AudioMessage: &waE2E.AudioMessage{
			URL: proto.String(up.URL), DirectPath: proto.String(up.DirectPath),
			MediaKey: up.MediaKey, FileEncSHA256: up.FileEncSHA256, FileSHA256: up.FileSHA256,
			FileLength: proto.Uint64(up.FileLength), Mimetype: proto.String(mime),
			PTT: proto.Bool(isi.PTT), ContextInfo: ci,
		}}
	case isi.Sticker != nil:
		up, err := b.client.Upload(ctx, isi.Sticker, whatsmeow.MediaImage)
		if err != nil {
			return err
		}
		pesan = &waE2E.Message{StickerMessage: &waE2E.StickerMessage{
			URL: proto.String(up.URL), DirectPath: proto.String(up.DirectPath),
			MediaKey: up.MediaKey, FileEncSHA256: up.FileEncSHA256, FileSHA256: up.FileSHA256,
			FileLength: proto.Uint64(up.FileLength), Mimetype: proto.String("image/webp"),
			ContextInfo: ci,
		}}
	case isi.Document != nil:
		up, err := b.client.Upload(ctx, isi.Document, whatsmeow.MediaDocument)
		if err != nil {
			return err
		}
		mime := isi.Mimetype
		if mime == "" {
			mime = lib.DeteksiMime(isi.Document)
		}
		pesan = &waE2E.Message{DocumentMessage: &waE2E.DocumentMessage{
			URL: proto.String(up.URL), DirectPath: proto.String(up.DirectPath),
			MediaKey: up.MediaKey, FileEncSHA256: up.FileEncSHA256, FileSHA256: up.FileSHA256,
			FileLength: proto.Uint64(up.FileLength), Mimetype: proto.String(mime),
			Title: proto.String(isi.FileName), FileName: proto.String(isi.FileName),
			Caption: proto.String(caption), ContextInfo: ci,
		}}
	case ci != nil:
		pesan = &waE2E.Message{ExtendedTextMessage: &waE2E.ExtendedTextMessage{
			Text: proto.String(isi.Text), ContextInfo: ci,
		}}
	default:
		pesan = &waE2E.Message{Conversation: proto.String(isi.Text)}
	}
	_, err = b.client.SendMessage(ctx, target, pesan)
	return err
}

func infoWebp(data []byte) (uint32, uint32, bool) {
	if len(data) < 30 || string(data[0:4]) != "RIFF" || string(data[8:12]) != "WEBP" {
		return 512, 512, false
	}
	switch string(data[12:16]) {
	case "VP8X":
		anim := data[20]&0x02 != 0
		w := 1 + uint32(data[24]) | uint32(data[25])<<8 | uint32(data[26])<<16
		h := 1 + uint32(data[27]) | uint32(data[28])<<8 | uint32(data[29])<<16
		return w, h, anim
	case "VP8L":
		if data[20] != 0x2f {
			return 512, 512, false
		}
		v := uint32(data[21]) | uint32(data[22])<<8 | uint32(data[23])<<16 | uint32(data[24])<<24
		return (v & 0x3FFF) + 1, ((v >> 14) & 0x3FFF) + 1, false
	case "VP8 ":
		w := uint32(data[26]) | uint32(data[27]&0x3f)<<8
		h := uint32(data[28]) | uint32(data[29]&0x3f)<<8
		return w, h, false
	}
	return 512, 512, false
}

func (b *Bot) KirimStikerBersuara(jid string, stiker, audio []byte, detik int, quoted *lib.Pesan, mention string) error {
	if b.client == nil {
		return fmt.Errorf("client belum tersambung")
	}
	target, err := types.ParseJID(jid)
	if err != nil {
		return err
	}
	ctx := context.Background()
	upS, err := b.client.Upload(ctx, stiker, whatsmeow.MediaImage)
	if err != nil {
		return err
	}
	upA, err := b.client.Upload(ctx, audio, whatsmeow.MediaAudio)
	if err != nil {
		return err
	}
	w, h, anim := infoWebp(stiker)
	sekarang := time.Now().Unix()
	stik := &waE2E.StickerMessage{
		URL: proto.String(upS.URL), DirectPath: proto.String(upS.DirectPath),
		MediaKey: upS.MediaKey, FileEncSHA256: upS.FileEncSHA256, FileSHA256: upS.FileSHA256,
		FileLength: proto.Uint64(upS.FileLength), Mimetype: proto.String("image/webp"),
		Width: proto.Uint32(w), Height: proto.Uint32(h),
		IsAnimated: proto.Bool(anim), StickerSentTS: proto.Int64(sekarang),
		MediaKeyTimestamp: proto.Int64(sekarang),
		Audio: &waE2E.StickerMessage_AudioMessage{AudioMessage: &waE2E.AudioMessage{
			URL: proto.String(upA.URL), DirectPath: proto.String(upA.DirectPath),
			MediaKey: upA.MediaKey, FileEncSHA256: upA.FileEncSHA256, FileSHA256: upA.FileSHA256,
			FileLength: proto.Uint64(upA.FileLength), Mimetype: proto.String("audio/ogg; codecs=opus"),
			Seconds: proto.Uint32(uint32(detik)), MediaKeyTimestamp: proto.Int64(sekarang),
		}},
	}
	var mentions []string
	if mention != "" {
		mentions = []string{mention}
	}
	if ci := b.konteksInfo(lib.IsiPesan{Mentions: mentions}, quoted); ci != nil {
		stik.ContextInfo = ci
	}
	pesan := &waE2E.Message{AudioStickerMessage: &waE2E.FutureProofMessage{
		Message: &waE2E.Message{StickerMessage: stik},
	}}
	_, err = b.client.SendMessage(ctx, target, pesan)
	return err
}

func (b *Bot) OK(m *lib.Pesan, teks string) error {
	gms := []lib.GrupMention{
		{JID: "120363186235853203@g.us", Subjek: "ask|" + m.Command},
		{JID: "1234" + lib.RandNum(14) + "@g.us", Subjek: teks},
	}
	return b.SendMessage(m.From, lib.IsiPesan{Text: teks, GroupMentions: gms}, m)
}

func (b *Bot) DownloadAny(m *lib.Pesan) ([]byte, string, error) {
	if b.client == nil {
		return nil, "", fmt.Errorf("client belum tersambung")
	}
	var mentah *waE2E.Message
	switch f := m.Full.(type) {
	case *events.Message:
		mentah = f.Message
	case *waE2E.Message:
		mentah = f
	default:
		return nil, "", fmt.Errorf("pesan tidak punya media yang bisa diunduh")
	}
	data, err := b.client.DownloadAny(context.Background(), mentah)
	if err != nil {
		return nil, "", err
	}
	mime := m.Mime
	if mime == "" {
		mime = lib.DeteksiMime(data)
	}
	return data, mime, nil
}

func konversiGrup(g *types.GroupInfo) *lib.GrupInfo {
	hasil := &lib.GrupInfo{
		JID:      g.JID.String(),
		Nama:     g.Name,
		Announce: g.IsAnnounce,
	}
	if !g.OwnerJID.IsEmpty() {
		hasil.Pemilik = g.OwnerJID.String()
	}
	for _, p := range g.Participants {
		nomor := p.PhoneNumber.String()
		if p.PhoneNumber.IsEmpty() {
			nomor = p.JID.String()
		}
		hasil.Peserta = append(hasil.Peserta, lib.Peserta{
			JID:          p.JID.String(),
			Nomor:        nomor,
			LID:          p.LID.String(),
			IsAdmin:      p.IsAdmin,
			IsSuperAdmin: p.IsSuperAdmin,
		})
	}
	return hasil
}

func (b *Bot) GetGroupInfo(jid string) (*lib.GrupInfo, error) {
	target, err := types.ParseJID(jid)
	if err != nil {
		return nil, err
	}
	g, err := b.client.GetGroupInfo(context.Background(), target)
	if err != nil {
		return nil, err
	}
	return konversiGrup(g), nil
}

func (b *Bot) GetJoinedGroups() ([]*lib.GrupInfo, error) {
	gs, err := b.client.GetJoinedGroups(context.Background())
	if err != nil {
		return nil, err
	}
	hasil := make([]*lib.GrupInfo, 0, len(gs))
	for _, g := range gs {
		hasil = append(hasil, konversiGrup(g))
	}
	return hasil, nil
}

func (b *Bot) UpdateGroupParticipants(jid string, peserta []string, aksi string) error {
	target, err := types.ParseJID(jid)
	if err != nil {
		return err
	}
	jids := make([]types.JID, 0, len(peserta))
	for _, p := range peserta {
		j, err := types.ParseJID(p)
		if err != nil {
			j = types.NewJID(lib.NomorSaja(p), types.DefaultUserServer)
		}
		jids = append(jids, j)
	}
	_, err = b.client.UpdateGroupParticipants(context.Background(), target, jids, whatsmeow.ParticipantChange(aksi))
	return err
}

func (b *Bot) GetProfilePicture(jid string) (string, error) {
	target, err := types.ParseJID(jid)
	if err != nil {
		return "", err
	}
	info, err := b.client.GetProfilePictureInfo(context.Background(), target, nil)
	if err != nil {
		return "", err
	}
	return info.URL, nil
}

func (b *Bot) GetGroupInviteLink(jid string) (string, error) {
	target, err := types.ParseJID(jid)
	if err != nil {
		return "", err
	}
	return b.client.GetGroupInviteLink(context.Background(), target, false)
}

func (b *Bot) GetGroupInviteInfo(link string) (*lib.GrupInfo, error) {
	g, err := b.client.GetGroupInfoFromLink(context.Background(), link)
	if err != nil {
		return nil, err
	}
	return konversiGrup(g), nil
}

func (b *Bot) SetGroupName(jid string, nama string) error {
	target, err := types.ParseJID(jid)
	if err != nil {
		return err
	}
	return b.client.SetGroupName(context.Background(), target, nama)
}

func (b *Bot) SetGroupAnnounce(jid string, announce bool) error {
	target, err := types.ParseJID(jid)
	if err != nil {
		return err
	}
	return b.client.SetGroupAnnounce(context.Background(), target, announce)
}

func (b *Bot) IsOnWhatsApp(nomor string) (bool, string, error) {
	resp, err := b.client.IsOnWhatsApp(context.Background(), []string{"+" + lib.NomorSaja(nomor)})
	if err != nil {
		return false, "", err
	}
	if len(resp) == 0 {
		return false, "", nil
	}
	return resp[0].IsIn, resp[0].JID.String(), nil
}

func (b *Bot) SendPresence(state string) error {
	return b.client.SendPresence(context.Background(), types.Presence(state))
}

func (b *Bot) SendChatPresence(jid string, state string) error {
	if b.client == nil {
		return nil
	}
	target, err := types.ParseJID(jid)
	if err != nil {
		return err
	}
	media := types.ChatPresenceMediaText
	if state == "recording" {
		media = types.ChatPresenceMediaAudio
	}
	return b.client.SendChatPresence(context.Background(), target, types.ChatPresence(state), media)
}

func (b *Bot) React(m *lib.Pesan, emoji string) error {
	if b.client == nil || m == nil || m.ID == "" {
		return nil
	}
	target, err := types.ParseJID(m.From)
	if err != nil {
		return err
	}
	var sender types.JID
	if m.IsGroup && m.Sender != "" {
		sender, _ = types.ParseJID(m.Sender)
	}
	msg := b.client.BuildReaction(target, sender, types.MessageID(m.ID), emoji)
	_, err = b.client.SendMessage(context.Background(), target, msg)
	return err
}

func (b *Bot) RevokeMessage(jid string, m *lib.Pesan) error {
	target, err := types.ParseJID(jid)
	if err != nil {
		return err
	}
	_, err = b.client.RevokeMessage(context.Background(), target, types.MessageID(m.ID))
	return err
}

func (b *Bot) RelayMentah(jid string, mentah any) error {
	target, err := types.ParseJID(jid)
	if err != nil {
		return err
	}
	msg, ok := mentah.(*waE2E.Message)
	if !ok {
		return fmt.Errorf("relay butuh *waE2E.Message")
	}
	_, err = b.client.SendMessage(context.Background(), target, msg)
	return err
}

func (b *Bot) BotJID() string {
	if b.client == nil || b.client.Store.ID == nil {
		return ""
	}
	return b.client.Store.ID.String()
}

func (b *Bot) BotLID() string {
	if b.client == nil || b.client.Store.LID.IsEmpty() {
		return ""
	}
	return b.client.Store.LID.String()
}

func (b *Bot) BotNomor() string {
	if b.client == nil || b.client.Store.ID == nil {
		return b.nomor
	}
	return b.client.Store.ID.User
}

func (b *Bot) Terhubung() bool {
	return b.client != nil && b.client.IsConnected() && b.client.IsLoggedIn()
}

func gerbangBotGrup(grup string) bool {
	bt, err := os.ReadFile("data/group.json")
	if err != nil {
		return true
	}
	var mentah map[string]any
	if json.Unmarshal(bt, &mentah) != nil {
		return true
	}
	sub, _ := mentah["bot"].(map[string]any)
	v, ada := sub[grup].(bool)
	if !ada {
		return true
	}
	return v
}

var reUnbanaku = regexp.MustCompile(`(?i)^[#!\/\.]?unbanaku$`)

func gerbangBanMember(b *Bot, k *lib.Konteks, m *lib.Pesan) bool {
	if !m.IsGroup || m.FromMe {
		return false
	}
	if len(lib.DaftarBan(b.group, m.From)) == 0 {
		return false
	}
	lib.PruneBanMember(b.group, m.From)
	nomor := lib.NomorSaja(m.Sender)
	lid := lib.NomorSaja(m.SenderAlt)
	if nomor == lib.NomorOwner() || lid == lib.NomorOwner() || nomor == b.BotNomor() || lid == b.BotNomor() {
		return false
	}
	if !lib.BanAktif(b.group, m.From, nomor, lid) {
		return false
	}
	if reUnbanaku.MatchString(strings.TrimSpace(m.Text)) {
		if lib.LepasBan(b.group, m.From, nomor, lid) {
			fmt.Printf("[unbanaku] %s lepas ban sendiri di %s\n", m.Sender, m.From)
			_ = b.SendMessage(m.From, lib.IsiPesan{Text: "hoki lu nemu fitur rahasia ini jir — ban lu gw lepas. jangan ngespam lagi"}, nil)
		}
		return true
	}
	_ = b.RevokeMessage(m.From, m)
	return true
}

func tulisUlangPerintah(m *lib.Pesan) {
	if m.Quoted == nil || len(m.Quoted.GroupMentions) == 0 {
		return
	}
	gs := m.Quoted.GroupMentions[0].Subjek
	var tex string
	if strings.HasPrefix(gs, "json|") {
		peta := map[string]string{}
		if json.Unmarshal([]byte(gs[len("json|"):]), &peta) == nil {
			tex = peta[m.Command]
		}
	}
	if tex == "" && strings.HasPrefix(gs, "cmd|") {
		tex = gs[len("cmd|"):]
	}
	if tex == "" {
		return
	}
	bidang := strings.Fields(tex)
	if len(bidang) == 0 {
		return
	}
	m.Command, m.Q = gabungPrefixSpasi(bidang)
	m.Prefix = ""
	if r := []rune(m.Command); len(r) > 0 && strings.ContainsRune("!#%./\\", r[0]) {
		m.Prefix = string(r[0])
	}
}

func (b *Bot) jidBot(jid string) bool {
	if jid == "" {
		return false
	}
	return jid == b.BotJID() || jid == b.BotLID() || lib.NomorSaja(jid) == b.BotNomor()
}

func (b *Bot) layakNgobrolAI(m *lib.Pesan) bool {
	if m.FromMe || m.Type == "reactionMessage" {
		return false
	}
	teks := strings.TrimSpace(m.Text)
	if teks == "" {
		return false
	}
	if !m.IsGroup {
		return !fitchat.SesiAnonAktif(m.From)
	}
	if !gerbangBotGrup(m.From) {
		return false
	}
	if m.Quoted != nil {
		for _, gm := range m.Quoted.GroupMentions {
			if strings.HasPrefix(gm.Subjek, "ask|") {
				return true
			}
		}
		if b.jidBot(m.Quoted.Sender) {
			return true
		}
	}
	for _, jid := range m.Mentions {
		if b.jidBot(jid) {
			return true
		}
	}
	return strings.HasSuffix(teks, "?")
}

func (b *Bot) TanganiPesan(m *lib.Pesan) {
	HitungPesan.Masuk.Add(1)
	fmt.Printf("[pesan] %s dari %s (alt %s): %q\n", m.From, m.Sender, m.SenderAlt, m.Text)
	k := &lib.Konteks{Conn: b, M: m, User: b.user, Group: b.group}
	fitursystem.CatatPesan(m)
	fiturmaker.CatatStikerBersuara(m)
	fitursystem.ProsesAwal(k)
	if len(b.user.Banned) > 0 {
		nomor := lib.NomorSaja(m.Sender)
		lid := lib.NomorSaja(m.SenderAlt)
		for _, bn := range b.user.Banned {
			if bn == nomor || bn == lid {
				return
			}
		}
	}
	if gerbangBanMember(b, k, m) {
		return
	}
	tulisUlangPerintah(m)
	if !m.IsGroup && m.Prefix == "" && !m.FromMe {
		_ = fitchat.TeruskanAnonim(k)
	}
	if m.IsGroup && !rePerintahAntilink.MatchString(m.Command) {
		if berhenti, err := fiturgrup.CekAntilink(k); err == nil && berhenti {
			return
		}
	}
	f := lib.CariFitur(m.Command)
	if f == nil {
		if b.layakNgobrolAI(m) {
			HitungPesan.Perintah.Add(1)
			if err := fiturai.Ngobrol(k, m.Text); err != nil {
				fmt.Printf("[fitur] mcp.js error: %v\n", err)
			}
			HitungPesan.Diproses.Add(1)
		}
		return
	}
	if m.IsGroup && f.File != "bot.js" && !gerbangBotGrup(m.From) {
		return
	}
	HitungPesan.Perintah.Add(1)
	fitursystem.Promo(k)
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("[fitur] %s panic: %v\n", f.File, r)
			}
		}()
		if err := f.Jalankan(k); err != nil {
			fmt.Printf("[fitur] %s error: %v\n", f.File, err)
		}
	}()
	HitungPesan.Diproses.Add(1)
}
