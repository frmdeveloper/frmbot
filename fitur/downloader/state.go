package downloader

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"gow1/lib"
)

var polaTiktok = regexp.MustCompile(`(?i)(?:https?://)?(?:[\w-]+\.)?tiktok\.com/\S+`)

func emojiFlag(kode string) string {
	bagian := strings.Split(kode, "-")
	negara := strings.ToUpper(bagian[len(bagian)-1])
	hasil := ""
	for _, c := range negara {
		lc := []rune(strings.ToLower(string(c)))
		if len(lc) > 0 {
			hasil += string(rune(lc[0]) + 127365)
		}
	}
	return hasil
}

func acakHex(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func mp42mp3(buffer []byte) ([]byte, error) {
	ran := acakHex(14)
	awal := filepath.Join("sampah", ran+".mp4")
	akhir := filepath.Join("sampah", ran+".mp3")
	if err := os.MkdirAll("sampah", 0o755); err != nil {
		return nil, err
	}
	if err := os.WriteFile(awal, buffer, 0o644); err != nil {
		return nil, err
	}
	defer os.RemoveAll(awal)
	defer os.RemoveAll(akhir)
	_, stderr := lib.SpawnSync(fmt.Sprintf("ffmpeg -y -i %s %s", awal, akhir))
	baca, err := os.ReadFile(akhir)
	if err != nil {
		return nil, fmt.Errorf("%s", stderr)
	}
	return baca, nil
}

func kirimVideoURL(k *lib.Konteks, url, caption string) error {
	b, err := lib.AmbilURL(url, nil)
	if err != nil {
		return err
	}
	return k.Conn.SendMessage(k.M.From, lib.IsiPesan{Video: b, Caption: caption}, k.M)
}

func kirimGambarURL(k *lib.Konteks, url, caption string) error {
	b, err := lib.AmbilURL(url, nil)
	if err != nil {
		return err
	}
	return k.Conn.SendMessage(k.M.From, lib.IsiPesan{Image: b, Caption: caption}, k.M)
}

func kirimAudioURL(k *lib.Konteks, url string) error {
	b, err := lib.AmbilURL(url, nil)
	if err != nil {
		return err
	}
	return k.Conn.SendMessage(k.M.From, lib.IsiPesan{Audio: b}, k.M)
}

type dataKirim struct {
	URL  string
	Nama string
}

type pesanWS struct {
	Player bool   `json:"player,omitempty"`
	HTML   string `json:"html,omitempty"`
	Nama   string `json:"nama,omitempty"`
	Proses bool   `json:"proses,omitempty"`
	Judul  string `json:"judul,omitempty"`
	Form   bool   `json:"form,omitempty"`
	Terima bool   `json:"terima,omitempty"`
}

type hasilForm struct {
	HTML  string
	Nama  string
	Waktu time.Time
}

type dataCariLain struct {
	URL    string
	Nama   string
	Dorong func(pesanWS) bool
}

var (
	muTT          sync.Mutex
	ttKirimDi     = map[string]func(dataKirim){}
	ttFormWs      = map[string]*klienWS{}
	ttFormHasil   = map[string]*hasilForm{}
	ttCariLain    func(dataCariLain) bool
	ttCariKonteks *lib.Konteks
)

func tungguKirim(nama string, detik int) (dataKirim, bool) {
	ch := make(chan dataKirim, 1)
	muTT.Lock()
	ttKirimDi[nama] = func(d dataKirim) {
		select {
		case ch <- d:
		default:
		}
	}
	muTT.Unlock()
	select {
	case d := <-ch:
		return d, true
	case <-time.After(time.Duration(detik) * time.Second):
		muTT.Lock()
		delete(ttKirimDi, nama)
		muTT.Unlock()
		return dataKirim{}, false
	}
}

func panggilKirimDi(nama string, d dataKirim) bool {
	muTT.Lock()
	fn, ada := ttKirimDi[nama]
	if ada {
		delete(ttKirimDi, nama)
	}
	muTT.Unlock()
	if !ada {
		return false
	}
	fn(d)
	return true
}

func ambilKirimDi(nama string) (func(dataKirim), bool) {
	muTT.Lock()
	defer muTT.Unlock()
	fn, ada := ttKirimDi[nama]
	return fn, ada
}

func dorongForm(sisi string, data pesanWS) bool {
	muTT.Lock()
	ws := ttFormWs[sisi]
	muTT.Unlock()
	if ws == nil || !ws.siap() {
		return false
	}
	if data.Player {
		muTT.Lock()
		ttFormHasil[sisi] = &hasilForm{HTML: data.HTML, Nama: data.Nama, Waktu: time.Now()}
		delete(ttKirimDi, sisi)
		muTT.Unlock()
		go func() {
			time.Sleep(500 * time.Millisecond)
			muTT.Lock()
			delete(ttFormWs, sisi)
			muTT.Unlock()
		}()
	}
	return ws.kirimJSON(data)
}

func dorongKe(ws *klienWS, sisi string) func(pesanWS) bool {
	return func(data pesanWS) bool {
		if data.Player {
			muTT.Lock()
			ttFormHasil[sisi] = &hasilForm{HTML: data.HTML, Nama: data.Nama, Waktu: time.Now()}
			delete(ttKirimDi, sisi)
			muTT.Unlock()
		}
		if ws == nil || !ws.siap() {
			return false
		}
		return ws.kirimJSON(data)
	}
}

func hosntt() string {
	if v := os.Getenv("TTWS_HOST"); v != "" {
		return v
	}
	ifaces, _ := net.Interfaces()
	cadangan := ""
	for _, iface := range ifaces {
		nama := strings.ToLower(iface.Name)
		addrs, _ := iface.Addrs()
		for _, a := range addrs {
			ipnet, ok := a.(*net.IPNet)
			if !ok || ipnet.IP.IsLoopback() {
				continue
			}
			ip4 := ipnet.IP.To4()
			if ip4 == nil {
				continue
			}
			if strings.Contains(nama, "wlan") || strings.Contains(nama, "rmnet") || strings.Contains(nama, "swlan") {
				return ip4.String()
			}
			if cadangan == "" {
				cadangan = ip4.String()
			}
		}
	}
	if cadangan != "" {
		return cadangan
	}
	return "127.0.0.1"
}

func porttt() int {
	hasil, err := strconv.Atoi(os.Getenv("TTWS_PORT"))
	if err == nil && hasil != 0 {
		return hasil
	}
	return 4000
}

var HostTT = hosntt()
var PortTT = porttt()
