package downloader

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sync"
	"time"

	"github.com/coder/websocket"

	"gow1/lib"
)

type klienWS struct {
	c  *websocket.Conn
	mu sync.Mutex
}

func (k *klienWS) siap() bool {
	return k != nil && k.c != nil
}

func (k *klienWS) kirimJSON(v any) bool {
	b, err := json.Marshal(v)
	if err != nil {
		return false
	}
	return k.kirim(websocket.MessageText, b)
}

func (k *klienWS) kirimTeks(s string) bool {
	return k.kirim(websocket.MessageText, []byte(s))
}

func (k *klienWS) kirim(tipe websocket.MessageType, data []byte) bool {
	if !k.siap() {
		return false
	}
	k.mu.Lock()
	defer k.mu.Unlock()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	return k.c.Write(ctx, tipe, data) == nil
}

var polaNamaLagu = regexp.MustCompile(`^ttws-[a-f0-9]{16}\.mp3$`)

func ambilLagu(nama string) ([]byte, bool) {
	if !polaNamaLagu.MatchString(nama) {
		return nil, false
	}
	b, err := os.ReadFile(filepath.Join("sampah", nama))
	if err != nil {
		return nil, false
	}
	return b, true
}

func simpanLagu(buf []byte, jam int) string {
	if jam <= 0 {
		jam = 2
	}
	nama := "ttws-" + acakHex(8) + ".mp3"
	_ = os.MkdirAll("sampah", 0o755)
	_ = os.WriteFile(filepath.Join("sampah", nama), buf, 0o644)
	go func() {
		time.Sleep(time.Duration(jam) * time.Hour)
		_ = os.Remove(filepath.Join("sampah", nama))
	}()
	return nama
}

var serverTTSiap sync.Once

func mulaiServerTT() {
	serverTTSiap.Do(func() {
		go func() {
			mux := http.NewServeMux()
			mux.HandleFunc("/", tanganiTTWS)
			alamat := fmt.Sprintf("0.0.0.0:%d", PortTT)
			srv := &http.Server{Addr: alamat, Handler: mux}
			fmt.Printf("[ttws] siap di %s (URL webview pakai %s)\n", alamat, HostTT)
			if err := srv.ListenAndServe(); err != nil {
				fmt.Println("[ttws] GAGAL jalan:", err)
			}
		}()
	})
}

func tanganiTTWS(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Upgrade") != "websocket" {
		lagu, ada := ambilLagu(r.URL.Query().Get("id"))
		if !ada {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("kedaluwarsa"))
			return
		}
		w.Header().Set("Content-Type", "audio/mpeg")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		_, _ = w.Write(lagu)
		return
	}
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		return
	}
	klien := &klienWS{c: c}
	q := r.URL.Query()
	if lagu, ada := ambilLagu(q.Get("id")); ada {
		fmt.Println("[ttws] klien masuk")
		defer c.Close(websocket.StatusNormalClosure, "")
		kirim := func() {
			const potongan = 65536
			for pos := 0; pos < len(lagu); pos += potongan {
				ujung := pos + potongan
				if ujung > len(lagu) {
					ujung = len(lagu)
				}
				if !klien.kirim(websocket.MessageBinary, lagu[pos:ujung]) {
					return
				}
			}
			fmt.Println("[ttws] stream selesai", len(lagu), "byte")
			klien.kirimTeks("done")
		}
		go func() {
			for {
				if _, _, err := c.Read(context.Background()); err != nil {
					fmt.Println("[ttws] klien keluar")
					return
				}
				kirim()
			}
		}()
		kirim()
		return
	}
	sisi := q.Get("sisi")
	cari := q.Get("cari") == "1"
	if sisi == "" {
		fmt.Println("[ttws] koneksi ditolak")
		c.Close(websocket.StatusPolicyViolation, "")
		return
	}
	defer c.Close(websocket.StatusNormalClosure, "")
	if cari {
		muTT.Lock()
		delete(ttFormHasil, sisi)
		muTT.Unlock()
	}
	muTT.Lock()
	hasil := ttFormHasil[sisi]
	_, adaKirim := ttKirimDi[sisi]
	muTT.Unlock()
	if hasil != nil && !cari {
		fmt.Println("[ttws] form dibuka lagi", sisi, "replay player")
		klien.kirimJSON(pesanWS{Player: true, HTML: hasil.HTML, Nama: hasil.Nama})
		for {
			if _, _, err := c.Read(context.Background()); err != nil {
				return
			}
		}
	}
	if cari || adaKirim || hasil == nil {
		muTT.Lock()
		if _, ada := ttKirimDi[sisi]; !ada {
			ttKirimDi[sisi] = func(d dataKirim) {
				muTT.Lock()
				fn := ttCariLain
				muTT.Unlock()
				if fn != nil {
					fn(dataCariLain{URL: d.URL, Nama: sisi, Dorong: dorongKe(klien, sisi)})
				}
			}
		}
		ttFormWs[sisi] = klien
		muTT.Unlock()
		fmt.Println("[ttws] form masuk", sisi)
		klien.kirimJSON(pesanWS{Form: true})
		defer func() {
			muTT.Lock()
			if ttFormWs[sisi] == klien {
				delete(ttFormWs, sisi)
			}
			muTT.Unlock()
		}()
		for {
			_, data, err := c.Read(context.Background())
			if err != nil {
				return
			}
			var j struct {
				URL string `json:"url"`
			}
			if json.Unmarshal(data, &j) != nil || j.URL == "" {
				continue
			}
			klien.kirimJSON(pesanWS{Terima: true})
			fmt.Println("[ttws] form kirim url")
			if panggilKirimDi(sisi, dataKirim{URL: j.URL, Nama: sisi}) {
				continue
			}
			muTT.Lock()
			fn := ttCariLain
			muTT.Unlock()
			if fn != nil && fn(dataCariLain{URL: j.URL, Nama: sisi, Dorong: dorongKe(klien, sisi)}) {
				muTT.Lock()
				ttCariLain = nil
				muTT.Unlock()
			}
		}
	}
}

func TanganiTiktokws(k *lib.Konteks) error {

	return nil
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "tiktokws.js",
		Jenis:    "downloader",
		Nama:     nil,
		Itunya:   nil,
		Jalankan: TanganiTiktokws,
	})
}

func init() {
	mulaiServerTT()
}
