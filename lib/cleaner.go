package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// BersihkanDirektori menghapus berkas yang lebih tua dari batasUmur di dalam direktori
func BersihkanDirektori(dir string, batasUmur time.Duration) int {
	entri, err := os.ReadDir(dir)
	if err != nil {
		return 0
	}
	sekarang := time.Now()
	terhapus := 0
	for _, e := range entri {
		if e.IsDir() {
			continue
		}
		if e.Name() == ".gitkeep" || e.Name() == ".gitignore" {
			continue
		}
		path := filepath.Join(dir, e.Name())
		info, err := e.Info()
		if err != nil {
			continue
		}
		if sekarang.Sub(info.ModTime()) > batasUmur {
			if os.Remove(path) == nil {
				terhapus++
			}
		}
	}
	return terhapus
}

// BersihkanSemuaSampah membersihkan folder sampah dan data/tmp
func BersihkanSemuaSampah(batasUmur time.Duration) int {
	total := 0
	total += BersihkanDirektori("sampah", batasUmur)
	total += BersihkanDirektori("data/tmp", batasUmur)
	return total
}

// MulaiPembersihSampah menjalankan pembersihan berkas sementara secara berkala di latar belakang
func MulaiPembersihSampah(interval, batasUmur time.Duration) {
	go func() {
		BersihkanSemuaSampah(batasUmur)

		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for range ticker.C {
			jml := BersihkanSemuaSampah(batasUmur)
			if jml > 0 {
				fmt.Printf("[cleaner] %d berkas sementara terhapus\n", jml)
			}
		}
	}()
}
