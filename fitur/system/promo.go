package system

import (
	"math/rand"
	"sync"
	"time"

	"gow1/lib"
)

const linkPromo = "https://chat.whatsapp.com/FxQ1k5wmQsNLkMcP7z3dX2"

const peluangPromo = 0.15

const jedaPromo = 6 * time.Hour

var teksPromo = []string{
	"Btw mampir grup gua dong, rame kok\n" + linkPromo,
	"Eh iya, gabung sini yuk biar bisa bareng-bareng\n" + linkPromo,
	"Kalo mau nanya-nanya lagi enakan di grup, join aja\n" + linkPromo,
	"Numpang promo bentar ya, grupnya di sini\n" + linkPromo,
	"Sekalian join grup gua lah, banyak yang seru\n" + linkPromo,
}

var (
	muPromo       sync.Mutex
	promoTerakhir = map[string]time.Time{}
)

func Promo(k *lib.Konteks) {
	m := k.M
	if m == nil || m.IsGroup || m.FromMe {
		return
	}
	kunci := lib.NomorSaja(m.Sender)
	if kunci == "" {
		return
	}
	muPromo.Lock()
	defer muPromo.Unlock()
	if time.Since(promoTerakhir[kunci]) < jedaPromo {
		return
	}
	if rand.Float64() > peluangPromo {
		return
	}
	promoTerakhir[kunci] = time.Now()
}

func TanganiPromo(k *lib.Konteks) error {
	Promo(k)
	return nil
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "promo.js",
		Jenis:    "system",
		Nama:     nil,
		Itunya:   nil,
		Jalankan: TanganiPromo,
	})
}
