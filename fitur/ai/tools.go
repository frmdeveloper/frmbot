package ai

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/dop251/goja"

	"gow1/fitur/maker"
	"gow1/lib"
)

func toolFungsi(nama, desk string, props map[string]any, wajib ...string) map[string]any {
	params := map[string]any{"type": "object", "properties": props}
	if len(wajib) > 0 {
		params["required"] = wajib
	}
	return map[string]any{
		"type": "function",
		"function": map[string]any{
			"name":        nama,
			"description": desk,
			"parameters":  params,
		},
	}
}

func propTeks(desk string) map[string]any {
	return map[string]any{"type": "string", "description": desk}
}

var skemaTools = []map[string]any{
	toolFungsi("jalankanKode", "Jalankan kode JavaScript buat ngerjain hal yang gak bisa dilakuin tool lain. Ada global `conn` (ok(teks), sendMessage(isi) — cuma bisa ke chat asal, number) dan `m` (pesan pemicu). Return value terakhir jadi hasil. Ini tool paling kuat, pakai buat logika custom, hitungan, olah teks/data.", map[string]any{
		"kode": propTeks("kode JavaScript yang dijalankan (boleh async/await)"),
	}, "kode"),
	toolFungsi("ok", "Kirim pesan teks ke chat ini sekarang (misal laporan progres sebelum lanjut kerja).", map[string]any{
		"text": propTeks("isi pesan"),
	}, "text"),
	toolFungsi("sendMessage", "Kirim pesan teks ke chat ini atau ke DM owner (6285188261898). JID lain ditolak.", map[string]any{
		"jid":  propTeks("JID tujuan, kosongkan buat chat ini"),
		"text": propTeks("isi pesan"),
	}, "text"),
	toolFungsi("listFitur", "Daftar semua fitur/perintah bot yang terdaftar, dikelompokkan per kategori.", map[string]any{}),
	toolFungsi("runFitur", "Jalankan perintah bot kayak user yang ngetik, misal '#stiker' atau '#menu maker'. Hasilnya kekirim langsung ke chat.", map[string]any{
		"perintah": propTeks("perintah lengkap dengan prefix, misal '#ytmp3 https://...'"),
	}, "perintah"),
	toolFungsi("GetGroupInfo", "Info grup: nama, pemilik, jumlah peserta, daftar admin.", map[string]any{
		"jid": propTeks("JID grup, kosongkan buat grup ini"),
	}),
	toolFungsi("GetJoinedGroups", "Daftar grup yang bot ikuti.", map[string]any{}),
	toolFungsi("GetGroupInviteLink", "Link undangan grup.", map[string]any{
		"jid": propTeks("JID grup, kosongkan buat grup ini"),
	}),
	toolFungsi("GetGroupInviteInfo", "Info grup dari link undangan.", map[string]any{
		"link": propTeks("link undangan (chat.whatsapp.com/...)"),
	}, "link"),
	toolFungsi("IsOnWhatsApp", "Cek nomor terdaftar di WhatsApp atau nggak.", map[string]any{
		"nomor": propTeks("nomor telepon format internasional"),
	}, "nomor"),
	toolFungsi("GetProfilePictureInfo", "URL foto profil sebuah JID/nomor.", map[string]any{
		"jid": propTeks("JID atau nomor"),
	}, "jid"),
	toolFungsi("SetGroupAnnounce", "Atur grup jadi mode pengumuman (cuma admin yang bisa kirim) atau buka lagi.", map[string]any{
		"jid":      propTeks("JID grup, kosongkan buat grup ini"),
		"announce": map[string]any{"type": "boolean", "description": "true = cuma admin, false = semua member"},
	}, "announce"),
	toolFungsi("downloadMediaMessage", "Unduh media dari pesan yang lagi dibalas / pesan pemicu, simpan ke folder sampah, balikin info file-nya (mime, ukuran, path).", map[string]any{}),
	toolFungsi("kirimStiker", "Bikin stiker dari gambar di pesan yang dibalas / pesan pemicu dan kirim ke chat.", map[string]any{}),
}

type argTool struct {
	Kode     string `json:"kode"`
	Text     string `json:"text"`
	JID      string `json:"jid"`
	Perintah string `json:"perintah"`
	Link     string `json:"link"`
	Nomor    string `json:"nomor"`
	Announce *bool  `json:"announce"`
}

func jalankanTool(k *lib.Konteks, nama, argJSON string) string {
	m := k.M
	var a argTool
	_ = json.Unmarshal([]byte(argJSON), &a)
	switch nama {
	case "jalankanKode":
		if strings.TrimSpace(a.Kode) == "" {
			return "Error: kode kosong"
		}
		return toolJalankanKode(k, a.Kode)
	case "ok":
		if err := k.Conn.SendMessage(m.From, lib.IsiPesan{Text: a.Text}, m); err != nil {
			return "Error: " + err.Error()
		}
		return "terkirim"
	case "sendMessage":
		tujuan := a.JID
		if tujuan == "" {
			tujuan = m.From
		}
		if tujuan != m.From && lib.NomorSaja(tujuan) != "6285188261898" {
			return "Error: tujuan ditolak, cuma boleh chat ini atau DM owner"
		}
		if err := k.Conn.SendMessage(tujuan, lib.IsiPesan{Text: a.Text}, nil); err != nil {
			return "Error: " + err.Error()
		}
		return "terkirim ke " + tujuan
	case "listFitur":
		return toolListFitur()
	case "runFitur":
		return toolRunFitur(k, a.Perintah)
	case "GetGroupInfo":
		jid := a.JID
		if jid == "" {
			jid = m.From
		}
		info, err := k.Conn.GetGroupInfo(jid)
		if err != nil {
			return "Error: " + err.Error()
		}
		admin := []string{}
		for _, p := range info.Peserta {
			if p.IsAdmin || p.IsSuperAdmin {
				admin = append(admin, lib.NomorSaja(p.Nomor))
			}
		}
		return fmt.Sprintf("nama: %s\njid: %s\npemilik: %s\npeserta: %d\nannounce: %v\nadmin: %s",
			info.Nama, info.JID, lib.NomorSaja(info.Pemilik), len(info.Peserta), info.Announce, strings.Join(admin, ", "))
	case "GetJoinedGroups":
		grup, err := k.Conn.GetJoinedGroups()
		if err != nil {
			return "Error: " + err.Error()
		}
		baris := []string{}
		for _, g := range grup {
			baris = append(baris, g.JID+" — "+g.Nama)
		}
		return strings.Join(baris, "\n")
	case "GetGroupInviteLink":
		jid := a.JID
		if jid == "" {
			jid = m.From
		}
		link, err := k.Conn.GetGroupInviteLink(jid)
		if err != nil {
			return "Error: " + err.Error()
		}
		return link
	case "GetGroupInviteInfo":
		info, err := k.Conn.GetGroupInviteInfo(a.Link)
		if err != nil {
			return "Error: " + err.Error()
		}
		return fmt.Sprintf("nama: %s\njid: %s\npeserta: %d", info.Nama, info.JID, len(info.Peserta))
	case "IsOnWhatsApp":
		ada, jid, err := k.Conn.IsOnWhatsApp(a.Nomor)
		if err != nil {
			return "Error: " + err.Error()
		}
		return fmt.Sprintf("terdaftar: %v, jid: %s", ada, jid)
	case "GetProfilePictureInfo":
		url, err := k.Conn.GetProfilePicture(a.JID)
		if err != nil {
			return "Error: " + err.Error()
		}
		return url
	case "SetGroupAnnounce":
		jid := a.JID
		if jid == "" {
			jid = m.From
		}
		if a.Announce == nil {
			return "Error: announce wajib diisi"
		}
		if err := k.Conn.SetGroupAnnounce(jid, *a.Announce); err != nil {
			return "Error: " + err.Error()
		}
		return "announce diatur ke " + fmt.Sprint(*a.Announce)
	case "downloadMediaMessage":
		target := m
		if m.Quoted != nil && m.Quoted.Media != "" {
			target = m.Quoted
		}
		data, mime, err := k.Conn.DownloadAny(target)
		if err != nil {
			return "Error: " + err.Error()
		}
		_ = os.MkdirAll("sampah", 0o755)
		path := filepath.Join("sampah", "ai-"+m.ID+lib.EkstensiMime(mime))
		if err := os.WriteFile(path, data, 0o644); err != nil {
			return "Error: " + err.Error()
		}
		return fmt.Sprintf("tersimpan di %s (mime %s, %d byte)", path, mime, len(data))
	case "kirimStiker":
		target := m
		if m.Quoted != nil && m.Quoted.Media != "" {
			target = m.Quoted
		}
		data, mime, err := k.Conn.DownloadAny(target)
		if err != nil {
			return "Error: " + err.Error()
		}
		webp, err := maker.BuatStiker(data, mime)
		if err != nil {
			return "Error: " + err.Error()
		}
		if err := k.Conn.SendMessage(m.From, lib.IsiPesan{Sticker: webp}, m); err != nil {
			return "Error: " + err.Error()
		}
		return "stiker terkirim"
	}
	return "Error: tool " + nama + " gak dikenal"
}

func toolListFitur() string {
	perJenis := map[string][]string{}
	for _, f := range lib.Semua() {
		nama := f.File
		if len(f.Nama) > 0 {
			nama += ": " + strings.Join(f.Nama, ", ")
		} else {
			nama += " (tanpa nama menu)"
		}
		perJenis[f.Jenis] = append(perJenis[f.Jenis], nama)
	}
	jenis := []string{}
	for j := range perJenis {
		jenis = append(jenis, j)
	}
	sort.Strings(jenis)
	var sb strings.Builder
	for _, j := range jenis {
		sort.Strings(perJenis[j])
		sb.WriteString(j + ":\n- " + strings.Join(perJenis[j], "\n- ") + "\n")
	}
	return strings.TrimSpace(sb.String())
}

func toolRunFitur(k *lib.Konteks, perintah string) string {
	perintah = strings.TrimSpace(perintah)
	bagian := strings.Fields(perintah)
	if len(bagian) == 0 {
		return "Error: perintah kosong"
	}
	cmd := strings.ToLower(bagian[0])
	f := lib.CariFitur(cmd)
	if f == nil {
		return "Error: fitur buat " + cmd + " gak ketemu"
	}
	m2 := *k.M
	m2.Command = cmd
	m2.Text = perintah
	m2.Q = strings.Join(bagian[1:], " ")
	if len(cmd) > 0 && strings.ContainsRune("#!/.", rune(cmd[0])) {
		m2.Prefix = cmd[:1]
	}
	k2 := &lib.Konteks{Conn: k.Conn, M: &m2, User: k.User, Group: k.Group}
	if err := f.Jalankan(k2); err != nil {
		return "Error: " + err.Error()
	}
	return "fitur " + f.File + " dijalankan"
}

const skripFormatAI = `
globalThis.__fmt = function(v) {
	if (v === undefined) return "undefined";
	if (v === null) return "null";
	if (typeof v === "string") return v;
	if (v instanceof Error) return v.name + ": " + v.message;
	if (typeof v === "function") return "[Function]";
	try {
		var s = JSON.stringify(v, null, 2);
		if (s !== undefined) return s;
	} catch (e) {}
	return String(v);
};
`

func toolJalankanKode(k *lib.Konteks, kode string) string {
	m := k.M
	vm := goja.New()
	if _, err := vm.RunString(skripFormatAI); err != nil {
		return "Error: " + err.Error()
	}
	_ = vm.Set("m", map[string]any{
		"id": m.ID, "from": m.From, "sender": m.Sender, "senderAlt": m.SenderAlt,
		"text": m.Text, "isGroup": m.IsGroup, "fromMe": m.FromMe, "type": m.Type,
	})
	connObj := vm.NewObject()
	_ = connObj.Set("number", k.Conn.BotNomor())
	_ = connObj.Set("ok", func(teks string) (any, error) {
		err := k.Conn.SendMessage(m.From, lib.IsiPesan{Text: teks}, m)
		return nil, err
	})
	_ = connObj.Set("sendMessage", func(isi any) (any, error) {
		pesan := lib.IsiPesan{}
		switch t := isi.(type) {
		case string:
			pesan.Text = t
		case map[string]any:
			if s, ok := t["text"].(string); ok {
				pesan.Text = s
			}
			if s, ok := t["caption"].(string); ok {
				pesan.Caption = s
			}
		default:
			pesan.Text = fmt.Sprint(isi)
		}
		err := k.Conn.SendMessage(m.From, pesan, m)
		return nil, err
	})
	_ = vm.Set("conn", connObj)
	ekor := ")().then(function(v) { globalThis.__selesai = { ok: true, v: v }; }, function(e) { globalThis.__selesai = { ok: false, v: e }; });"
	program := "(async () => { return (\n" + kode + "\n); }" + ekor
	if _, err := goja.Compile("kode", program, false); err != nil {
		program = "(async () => {\n" + kode + "\n}" + ekor
	}
	if _, err := vm.RunString(program); err != nil {
		return "Error: " + err.Error()
	}
	batas := time.Now().Add(10 * time.Second)
	for time.Now().Before(batas) {
		s := vm.Get("__selesai")
		if s != nil && !goja.IsUndefined(s) && !goja.IsNull(s) {
			obj := s.ToObject(vm)
			hasil := obj.Get("v")
			f, ok := goja.AssertFunction(vm.Get("__fmt"))
			if !ok {
				return fmt.Sprint(hasil.Export())
			}
			keluar, err := f(goja.Undefined(), hasil)
			if err != nil {
				return "Error: " + err.Error()
			}
			return keluar.String()
		}
		if _, err := vm.RunString("void 0"); err != nil {
			break
		}
		time.Sleep(2 * time.Millisecond)
	}
	return "Error: kode kelamaan (Promise pending > 10 detik)"
}
