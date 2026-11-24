package maker

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	"gow1/lib"

	waE2E "go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

type fldMediaSb struct {
	URL               string `json:"URL,omitempty"`
	DirectPath        string `json:"directPath,omitempty"`
	FileSHA256        []byte `json:"fileSHA256,omitempty"`
	FileEncSHA256     []byte `json:"fileEncSHA256,omitempty"`
	MediaKey          []byte `json:"mediaKey,omitempty"`
	MediaKeyTimestamp int64  `json:"mediaKeyTimestamp,omitempty"`
	Mimetype          string `json:"mimetype,omitempty"`
	FileLength        uint64 `json:"fileLength,omitempty"`
	Width             uint32 `json:"width,omitempty"`
	Height            uint32 `json:"height,omitempty"`
	Seconds           uint32 `json:"seconds,omitempty"`
}

type entriAudioStiker struct {
	S fldMediaSb `json:"s"`
	A fldMediaSb `json:"a"`
	T int64      `json:"t"`
}

var (
	muCacheSb    sync.Mutex
	isiCacheSb   map[string]*entriAudioStiker
	cacheSbMuat  bool
	berkasCacheSb = "data/cacheaudiostiker.json"
)

func muatCacheSb() map[string]*entriAudioStiker {
	if cacheSbMuat {
		return isiCacheSb
	}
	cacheSbMuat = true
	isiCacheSb = map[string]*entriAudioStiker{}
	if b, err := os.ReadFile(berkasCacheSb); err == nil {
		_ = json.Unmarshal(b, &isiCacheSb)
	}
	return isiCacheSb
}

func simpanCacheSb() {
	if b, err := json.Marshal(isiCacheSb); err == nil {
		_ = os.WriteFile(berkasCacheSb, b, 0o644)
	}
}

func fldDariStiker(s *waE2E.StickerMessage) fldMediaSb {
	return fldMediaSb{
		URL: s.GetURL(), DirectPath: s.GetDirectPath(),
		FileSHA256: s.GetFileSHA256(), FileEncSHA256: s.GetFileEncSHA256(),
		MediaKey: s.GetMediaKey(), MediaKeyTimestamp: s.GetMediaKeyTimestamp(),
		Mimetype: s.GetMimetype(), FileLength: s.GetFileLength(),
		Width: s.GetWidth(), Height: s.GetHeight(),
	}
}

func fldDariAudio(a *waE2E.AudioMessage) fldMediaSb {
	return fldMediaSb{
		URL: a.GetURL(), DirectPath: a.GetDirectPath(),
		FileSHA256: a.GetFileSHA256(), FileEncSHA256: a.GetFileEncSHA256(),
		MediaKey: a.GetMediaKey(), MediaKeyTimestamp: a.GetMediaKeyTimestamp(),
		Mimetype: a.GetMimetype(), FileLength: a.GetFileLength(),
		Seconds: a.GetSeconds(),
	}
}

func stikerDariFld(f fldMediaSb) *waE2E.StickerMessage {
	return &waE2E.StickerMessage{
		URL: proto.String(f.URL), DirectPath: proto.String(f.DirectPath),
		FileSHA256: f.FileSHA256, FileEncSHA256: f.FileEncSHA256,
		MediaKey: f.MediaKey, MediaKeyTimestamp: proto.Int64(f.MediaKeyTimestamp),
		Mimetype: proto.String(f.Mimetype), FileLength: proto.Uint64(f.FileLength),
		Width: proto.Uint32(f.Width), Height: proto.Uint32(f.Height),
	}
}

func audioDariFld(f fldMediaSb) *waE2E.AudioMessage {
	return &waE2E.AudioMessage{
		URL: proto.String(f.URL), DirectPath: proto.String(f.DirectPath),
		FileSHA256: f.FileSHA256, FileEncSHA256: f.FileEncSHA256,
		MediaKey: f.MediaKey, MediaKeyTimestamp: proto.Int64(f.MediaKeyTimestamp),
		Mimetype: proto.String(f.Mimetype), FileLength: proto.Uint64(f.FileLength),
		Seconds: proto.Uint32(f.Seconds),
	}
}

func mentahDariFull(f any) *waE2E.Message {
	switch v := f.(type) {
	case *events.Message:
		return v.Message
	case *waE2E.Message:
		return v
	}
	return nil
}

func mentahDariPesan(m *lib.Pesan) *waE2E.Message {
	if m == nil {
		return nil
	}
	return mentahDariFull(m.Full)
}

func AudioDariStiker(s *waE2E.StickerMessage) *waE2E.AudioMessage {
	if s == nil {
		return nil
	}
	if am, ok := s.GetAudio().(*waE2E.StickerMessage_AudioMessage); ok {
		return am.AudioMessage
	}
	return nil
}

func StikerAudioDari(mentah *waE2E.Message) *waE2E.StickerMessage {
	if mentah == nil {
		return nil
	}
	if mentah.GetAudioStickerMessage() != nil {
		return mentah.GetAudioStickerMessage().GetMessage().GetStickerMessage()
	}
	if s := mentah.GetStickerMessage(); s != nil && AudioDariStiker(s) != nil {
		return s
	}
	return nil
}

func CatatStikerBersuara(m *lib.Pesan) {
	if m == nil || m.Type != "audioStickerMessage" || m.ID == "" {
		return
	}
	st := StikerAudioDari(mentahDariPesan(m))
	if st == nil || AudioDariStiker(st) == nil {
		return
	}
	muCacheSb.Lock()
	defer muCacheSb.Unlock()
	c := muatCacheSb()
	c[m.ID] = &entriAudioStiker{S: fldDariStiker(st), A: fldDariAudio(AudioDariStiker(st)), T: time.Now().UnixMilli()}
	batas := time.Now().Add(-7 * 24 * time.Hour).UnixMilli()
	for k, v := range c {
		if v == nil || v.T < batas {
			delete(c, k)
		}
	}
	for len(c) > 300 {
		kunciTua, waktuTua := "", int64(0)
		for k, v := range c {
			if waktuTua == 0 || v.T < waktuTua {
				kunciTua, waktuTua = k, v.T
			}
		}
		delete(c, kunciTua)
	}
	simpanCacheSb()
}

func AmbilCacheStikerBersuara(id string) (*waE2E.StickerMessage, *waE2E.AudioMessage, bool) {
	if id == "" {
		return nil, nil, false
	}
	muCacheSb.Lock()
	defer muCacheSb.Unlock()
	e := muatCacheSb()[id]
	if e == nil {
		return nil, nil, false
	}
	st := stikerDariFld(e.S)
	st.Audio = &waE2E.StickerMessage_AudioMessage{AudioMessage: audioDariFld(e.A)}
	return st, audioDariFld(e.A), true
}
