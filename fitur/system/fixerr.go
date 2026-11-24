package system

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"gow1/lib"
)

const nomorOwnerFixerr = "6285188261898"

type EntriError struct {
	Waktu time.Time `json:"waktu"`
	File  string    `json:"file"`
	Error string    `json:"error"`
	Dari  string    `json:"dari"`
}

var (
	muFixerr   sync.Mutex
	terakhirDM = map[string]time.Time{}
)

func potongFixerr(s string, n int) string {
	if len(s) > n {
		return s[:n] + "…"
	}
	return s
}

func Lapor(k *lib.Konteks, err error, file string) {
	if err == nil || k == nil || k.M == nil {
		return
	}
	pesan := err.Error()
	baris := strings.SplitN(pesan, "\n", 2)[0]
	fmt.Printf("[fixerr] %s: %s\n", file, baris)

	muFixerr.Lock()
	entri := EntriError{
		Waktu: time.Now(),
		File:  file,
		Error: potongFixerr(pesan, 1500),
		Dari:  k.M.From,
	}
	path := filepath.Join("data", "logerror.json")
	var daftar []EntriError
	if b, e := os.ReadFile(path); e == nil {
		_ = json.Unmarshal(b, &daftar)
	}
	daftar = append(daftar, entri)
	if len(daftar) > 20 {
		daftar = daftar[len(daftar)-20:]
	}
	if b, e := json.MarshalIndent(daftar, "", "  "); e == nil {
		_ = os.WriteFile(path, b, 0o644)
	}
	kunci := file + "|" + baris
	bolehDM := time.Since(terakhirDM[kunci]) >= 10*time.Minute
	if bolehDM {
		terakhirDM[kunci] = time.Now()
	}
	muFixerr.Unlock()

	if bolehDM && k.Conn != nil {
		teks := fmt.Sprintf("fitur error\n\nfile: %s\nerror: %s", file, potongFixerr(pesan, 1500))
		_ = k.Conn.SendMessage(nomorOwnerFixerr+"@s.whatsapp.net", lib.IsiPesan{Text: teks}, nil)
	}
}

func TanganiFixerr(k *lib.Konteks) error {
	m := k.M
	nomor := lib.NomorSaja(m.Sender)
	if nomor != nomorOwnerFixerr && nomor != k.Conn.BotNomor() {
		return nil
	}
	return k.Conn.OK(m, "konteks error-nya udah ilang, panggil ulang fiturnya biar error-nya kebaca lagi.")
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "fixerr.js",
		Jenis:    "system",
		Nama:     nil,
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(fixerr)$"),
		Jalankan: TanganiFixerr,
	})
}
