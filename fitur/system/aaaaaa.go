package system

import (
	"sync"
	"time"

	"gow1/lib"
)

const maksSimpan = 50

type CatatanPesan struct {
	ID      string `json:"id"`
	From    string `json:"from"`
	Sender  string `json:"sender"`
	LID     string `json:"lid,omitempty"`
	Type    string `json:"type"`
	Text    string `json:"text"`
	IsGroup bool   `json:"isGroup"`
	FromMe  bool   `json:"fromMe"`
	Waktu   int64  `json:"waktu"`
}

var (
	muCache    sync.Mutex
	cachePesan []CatatanPesan
	petaLID    = map[string]string{}
)

func CatatPesan(m *lib.Pesan) {
	if m == nil {
		return
	}
	muCache.Lock()
	defer muCache.Unlock()
	sender := lib.NomorSaja(m.Sender)
	lid := lib.NomorSaja(m.SenderAlt)
	if sender != "" && lid != "" && sender != lid {
		petaLID[sender] = lid
		petaLID[lid] = sender
	}
	cachePesan = append(cachePesan, CatatanPesan{
		ID:      m.ID,
		From:    m.From,
		Sender:  m.Sender,
		LID:     m.SenderAlt,
		Type:    m.Type,
		Text:    m.Text,
		IsGroup: m.IsGroup,
		FromMe:  m.FromMe,
		Waktu:   time.Now().UnixMilli(),
	})
	if len(cachePesan) > maksSimpan {
		cachePesan = cachePesan[len(cachePesan)-maksSimpan:]
	}
}

func SemuaCatatan() []CatatanPesan {
	muCache.Lock()
	defer muCache.Unlock()
	hasil := make([]CatatanPesan, len(cachePesan))
	copy(hasil, cachePesan)
	return hasil
}

func PasanganLID(nomor string) string {
	muCache.Lock()
	defer muCache.Unlock()
	return petaLID[nomor]
}

func ProsesAwal(k *lib.Konteks) {
	CatatPesan(k.M)
}

func TanganiAaaaaa(k *lib.Konteks) error {
	ProsesAwal(k)
	return nil
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "aaaaaa.js",
		Jenis:    "system",
		Nama:     nil,
		Itunya:   nil,
		Jalankan: TanganiAaaaaa,
	})
}
