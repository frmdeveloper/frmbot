package group

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"gow1/lib"
)

type entriAbsen struct {
	Jid  string `json:"jid"`
	Time string `json:"time"`
	Name string `json:"name"`
}

type dataAbsenGrup struct {
	Date string       `json:"date"`
	List []entriAbsen `json:"list"`
}

func ambilDataAbsen(v any) dataAbsenGrup {
	hasil := dataAbsenGrup{List: []entriAbsen{}}
	if v == nil {
		return hasil
	}
	b, err := json.Marshal(v)
	if err != nil {
		return hasil
	}
	_ = json.Unmarshal(b, &hasil)
	if hasil.List == nil {
		hasil.List = []entriAbsen{}
	}
	return hasil
}

func simpanDataAbsen(k *lib.Konteks, grup string, d dataAbsenGrup) {
	b, err := json.Marshal(d)
	if err != nil {
		return
	}
	var v any
	if err := json.Unmarshal(b, &v); err != nil {
		return
	}
	if k.Group.Absen == nil {
		k.Group.Absen = map[string]any{}
	}
	k.Group.Absen[grup] = v
	_ = k.Group.Simpan()
}

func TanganiAbsen(k *lib.Konteks) error {
	m := k.M
	if !m.IsGroup {
		return k.Conn.OK(m, "Bukan grup")
	}
	now := time.Now()
	today := now.Format("2/1/2006")

	d := ambilDataAbsen(k.Group.Absen[m.From])
	if d.Date != today {
		d = dataAbsenGrup{Date: today, List: []entriAbsen{}}
		simpanDataAbsen(k, m.From, d)
	}

	if strings.Contains(m.Command, "listabsen") {
		teks := fmt.Sprintf("📋 *Absen %s*\n\n*Total:* %d orang\n\n", today, len(d.List))
		if len(d.List) > 0 {
			baris := make([]string, len(d.List))
			mentions := make([]string, len(d.List))
			for i, x := range d.List {
				baris[i] = fmt.Sprintf("%d. @%s - %s", i+1, lib.NomorSaja(x.Jid), x.Time)
				mentions[i] = x.Jid
			}
			teks += strings.Join(baris, "\n")
			return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: teks, Mentions: mentions}, m)
		}
		return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: teks + "Sepi hehe"}, m)
	}

	phoneNum := lib.NomorSaja(m.Sender)
	for _, x := range d.List {
		if lib.NomorSaja(x.Jid) == phoneNum {
			return k.Conn.OK(m, "Udah njir")
		}
	}
	d.List = append(d.List, entriAbsen{
		Jid:  m.Sender,
		Time: now.Format("15.04"),
		Name: phoneNum,
	})
	simpanDataAbsen(k, m.From, d)
	teks := fmt.Sprintf("✅ @%s hadir! (%d)", phoneNum, len(d.List))
	return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: teks, Mentions: []string{m.Sender}}, m)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "absen.js",
		Jenis:    "group",
		Nama:     []string{"#hadir", "#absen", "#listabsen"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(hadir|absen|listabsen)$"),
		Jalankan: TanganiAbsen,
	})
}
