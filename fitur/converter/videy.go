package converter

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"gow1/lib"
)

type hasilVidey struct {
	res  string
	ress string
}

func unggahVidey(buffer []byte) (hasilVidey, error) {
	ext := strings.TrimPrefix(lib.EkstensiMime(lib.DeteksiMime(buffer)), ".")
	nama := fmt.Sprintf("%d.%s", time.Now().UnixMilli(), ext)
	body, err := kirimMultipart("https://videy.co/api/upload", "file", nama, buffer, nil)
	if err != nil {
		return hasilVidey{}, err
	}
	var r struct {
		ID string `json:"id"`
	}
	if err := json.Unmarshal(body, &r); err != nil {
		return hasilVidey{}, err
	}
	return hasilVidey{
		res:  "https://videy.co/v?id=" + r.ID,
		ress: "https://cdn.videy.co/" + r.ID + ".mp4",
	}, nil
}

func TanganiVidey(k *lib.Konteks) error {
	var down []byte
	if b, _, err := k.Conn.DownloadAny(k.M); err == nil {
		down = b
	} else if k.M.Quoted != nil {
		if b2, _, err2 := k.Conn.DownloadAny(k.M.Quoted); err2 == nil {
			down = b2
		}
	}
	if len(down) == 0 {
		return nil
	}
	up, err := unggahVidey(down)
	if err != nil {
		return err
	}
	return k.Conn.SendMessage(k.M.From, lib.IsiPesan{Text: up.res + "\n" + up.ress}, k.M)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "videy.js",
		Jenis:    "converter",
		Nama:     nil,
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)((upload)?videy(upload)?)$"),
		Jalankan: TanganiVidey,
	})
}
