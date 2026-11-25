package group

import (
	"encoding/json"
	"os"
	"regexp"
	"strings"
	"sync"

	"gow1/lib"
)

const nomorPapa = "6285188261898"
const nomorOwnerUtama = "6289680237010"

var nonDigit = regexp.MustCompile(`[^0-9]`)

func adalahOwner(k *lib.Konteks, jid string) bool {
	n := lib.NomorSaja(jid)
	if n == "" {
		return false
	}
	if n == nomorOwnerUtama || n == nomorPapa {
		return true
	}
	return k != nil && k.Conn != nil && n == k.Conn.BotNomor()
}

func pengirimOwner(k *lib.Konteks) bool {
	return adalahOwner(k, k.M.Sender) || adalahOwner(k, k.M.SenderAlt)
}

func ituPapa(jid string) bool {
	n := lib.NomorSaja(jid)
	return n == nomorOwnerUtama || n == nomorPapa
}

func cariPeserta(gc *lib.GrupInfo, kandidat ...string) *lib.Peserta {
	if gc == nil {
		return nil
	}
	for i := range gc.Peserta {
		p := &gc.Peserta[i]
		for _, c := range kandidat {
			n := lib.NomorSaja(c)
			if n == "" {
				continue
			}
			if n == p.Nomor || n == p.LID || n == lib.NomorSaja(p.JID) {
				return p
			}
		}
	}
	return nil
}

func pesertaAdmin(p *lib.Peserta) bool {
	return p != nil && (p.IsAdmin || p.IsSuperAdmin)
}

func pengirimAdmin(k *lib.Konteks, gc *lib.GrupInfo) bool {
	return pesertaAdmin(cariPeserta(gc, k.M.Sender, k.M.SenderAlt))
}

func botAdmin(k *lib.Konteks, gc *lib.GrupInfo) bool {
	return pesertaAdmin(cariPeserta(gc, k.Conn.BotJID(), k.Conn.BotLID(), k.Conn.BotNomor()))
}

func argPertama(m *lib.Pesan) string {
	f := strings.Fields(m.Q)
	if len(f) == 0 {
		return ""
	}
	return f[0]
}

func targetDariPesan(k *lib.Konteks, pakaiArg bool) string {
	if k.M.Quoted != nil && k.M.Quoted.Sender != "" {
		return k.M.Quoted.Sender
	}
	if len(k.M.Mentions) > 0 {
		return k.M.Mentions[0]
	}
	if pakaiArg {
		if f := strings.Fields(k.M.Q); len(f) > 0 {
			return nonDigit.ReplaceAllString(f[0], "") + "@s.whatsapp.net"
		}
	}
	return ""
}

func daftarString(v any) []string {
	switch t := v.(type) {
	case nil:
		return nil
	case []string:
		return t
	case []any:
		hasil := []string{}
		for _, x := range t {
			if s, ok := x.(string); ok {
				hasil = append(hasil, s)
			}
		}
		return hasil
	}
	return nil
}

var muGrupMentah sync.Mutex

func bacaGrupMentah() map[string]any {
	hasil := map[string]any{}
	b, err := os.ReadFile("data/group.json")
	if err != nil {
		return hasil
	}
	_ = json.Unmarshal(b, &hasil)
	return hasil
}

func toggleGrupAktif(kunci, grup string) bool {
	muGrupMentah.Lock()
	defer muGrupMentah.Unlock()
	raw := bacaGrupMentah()
	sub, _ := raw[kunci].(map[string]any)
	v, _ := sub[grup].(bool)
	return v
}

func setToggleGrup(kunci, grup string, nyala bool) {
	muGrupMentah.Lock()
	defer muGrupMentah.Unlock()
	raw := bacaGrupMentah()
	sub, _ := raw[kunci].(map[string]any)
	if sub == nil {
		sub = map[string]any{}
	}
	if nyala {
		sub[grup] = true
	} else {
		delete(sub, grup)
	}
	raw[kunci] = sub
	b, err := json.Marshal(raw)
	if err == nil {
		_ = os.WriteFile("data/group.json", b, 0o644)
	}
}
