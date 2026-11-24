package lib

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"sync"
)

type Handler func(k *Konteks) error

type Fitur struct {
	File     string
	Jenis    string
	Nama     []string
	Itunya   *regexp.Regexp
	Jalankan Handler
}

var (
	mu       sync.RWMutex
	registry = map[string]*Fitur{}

	ErorMuat []error
)

func Daftarkan(f *Fitur) {
	mu.Lock()
	defer mu.Unlock()
	if f == nil || f.File == "" {
		ErorMuat = append(ErorMuat, fmt.Errorf("fitur tanpa nama berkas dilewati"))
		return
	}
	registry[f.File] = f
}

func Semua() map[string]*Fitur {
	mu.RLock()
	defer mu.RUnlock()
	hasil := make(map[string]*Fitur, len(registry))
	for k, v := range registry {
		hasil[k] = v
	}
	return hasil
}

func CariFitur(teks string) *Fitur {
	token := strings.Fields(teks)
	if len(token) == 0 {
		return nil
	}
	perintah := token[0]
	mu.RLock()
	defer mu.RUnlock()
	kunci := make([]string, 0, len(registry))
	for k := range registry {
		kunci = append(kunci, k)
	}
	sort.Strings(kunci)
	for _, k := range kunci {
		f := registry[k]
		if f.Itunya != nil && f.Itunya.MatchString(perintah) {
			return f
		}
	}
	return nil
}

func Jumlah() int {
	mu.RLock()
	defer mu.RUnlock()
	return len(registry)
}
