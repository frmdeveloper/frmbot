package system

import (
	"encoding/json"
	"regexp"

	"gow1/lib"
)

func kirimJSON(k *lib.Konteks, obj any, nama string) error {
	b, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return err
	}
	if len(b) > 60000 {
		return k.Conn.SendMessage(k.M.From, lib.IsiPesan{
			Document: b,
			FileName: nama,
			Mimetype: "application/json",
		}, k.M)
	}
	return k.Conn.SendMessage(k.M.From, lib.IsiPesan{Text: string(b)}, k.M)
}

func TanganiGetmsg(k *lib.Konteks) error {
	m := k.M
	if m.Quoted != nil && m.Quoted.Full != nil {
		return kirimJSON(k, m.Quoted.Full, "pesan.json")
	}
	semua := SemuaCatatan()
	data := []CatatanPesan{}
	if regexp.MustCompile(`(^|\s)all$`).MatchString(m.Q) {
		data = semua
	} else {
		for _, c := range semua {
			if c.From == m.From {
				data = append(data, c)
			}
		}
	}
	return kirimJSON(k, data, "getmsg.json")
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "getmsg.js",
		Jenis:    "system",
		Nama:     nil,
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)getmsg$"),
		Jalankan: TanganiGetmsg,
	})
}
