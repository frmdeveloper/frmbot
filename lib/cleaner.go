package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

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

func BersihkanSemuaSampah(batasUmur time.Duration) int {
	total := 0
	total += BersihkanDirektori("sampah", batasUmur)
	total += BersihkanDirektori("data/tmp", batasUmur)
	return total
}

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
