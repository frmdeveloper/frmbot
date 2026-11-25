package lib

import (
	"sync"
	"time"
)

type PenjagaAntiBan struct {
	mu               sync.Mutex
	terakhirPerintah map[string]time.Time
	jumlahCepat      map[string]int
}

var AntiBan = &PenjagaAntiBan{
	terakhirPerintah: make(map[string]time.Time),
	jumlahCepat:      make(map[string]int),
}

// BolehJalan mengecek batas frekuensi pesan agar tidak terdeteksi spam atau bot oleh WhatsApp
func (a *PenjagaAntiBan) BolehJalan(sender string) bool {
	if sender == "" {
		return true
	}
	nomor := NomorSaja(sender)
	a.mu.Lock()
	defer a.mu.Unlock()

	sekarang := time.Now()
	terakhir, ada := a.terakhirPerintah[nomor]
	if !ada {
		a.terakhirPerintah[nomor] = sekarang
		a.jumlahCepat[nomor] = 1
		return true
	}

	selisih := sekarang.Sub(terakhir)
	a.terakhirPerintah[nomor] = sekarang

	if selisih < 800*time.Millisecond {
		a.jumlahCepat[nomor]++
		if a.jumlahCepat[nomor] > 3 {
			return false
		}
	} else {
		a.jumlahCepat[nomor] = 1
	}

	return true
}

// JedaAman memberikan jeda waktu kecil acak seperti ketikan manusia
func JedaAman(minMs, maxMs int) {
	if maxMs <= minMs {
		time.Sleep(time.Duration(minMs) * time.Millisecond)
		return
	}
	durasi := minMs + int(time.Now().UnixNano()%(int64(maxMs-minMs+1)))
	time.Sleep(time.Duration(durasi) * time.Millisecond)
}
