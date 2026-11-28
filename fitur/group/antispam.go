package group

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"gow1/lib"

	"go.mau.fi/whatsmeow/types/events"
)

const batasSpam = 5
const jendelaSpam = 10 * time.Second
const banSpamMs = int64(3600 * 1000)

var jejakSpam = struct {
	sync.Mutex
	data map[string]map[string][]time.Time
}{data: map[string]map[string][]time.Time{}}

func TanganiAntispam(k *lib.Konteks) error {
	m := k.M
	if !m.IsGroup {
		return k.Conn.OK(m, "Cuma buat grup")
	}
	arg := strings.ToLower(argPertama(m))
	nyala := toggleGrupAktif("antispam", m.From)
	gc, err := k.Conn.GetGroupInfo(m.From)
	if err != nil {
		return err
	}
	if !pengirimAdmin(k, gc) && !pengirimOwner(k) {
		return k.Conn.OK(m, "Lu bukan admin")
	}
	prefix := m.Prefix
	if prefix == "" {
		prefix = "#"
	}
	switch arg {
	case "on":
		if nyala {
			return k.Conn.OK(m, "Udah nyala kok")
		}
		if !botAdmin(k, gc) {
			return k.Conn.OK(m, "Gw bukan admin, gak bisa kick")
		}
		setToggleGrup("antispam", m.From, true)
		return k.Conn.OK(m, "Antispam nyala. Yang nyepam (>5 pesan/10 detik) langsung gw ban. Udah kena ban masih nyepam juga -> gw kick.")
	case "off":
		if !nyala {
			return k.Conn.OK(m, "Emang udah mati")
		}
		setToggleGrup("antispam", m.From, false)
		return k.Conn.OK(m, "Antispam mati")
	}
	status := "off"
	if nyala {
		status = "on"
	}
	return k.Conn.OK(m, fmt.Sprintf("Antispam: %s\nPakai %santispam on/off", status, prefix))
}

func CekAntispam(k *lib.Konteks) (bool, error) {
	m := k.M
	if m == nil || !m.IsGroup || m.FromMe {
		return false, nil
	}
	if pengirimOwner(k) {
		return false, nil
	}
	if ev, ok := m.Full.(*events.Message); ok && ev != nil && ev.Message != nil {
		if ev.Message.ReactionMessage != nil || ev.Message.ProtocolMessage != nil || ev.Message.EditedMessage != nil {
			return false, nil
		}
	}
	num := lib.NomorSaja(m.Sender)
	if num == "" {
		return false, nil
	}
	lib.PruneBanMember(k.Group, m.From)
	sudahBan := lib.BanAktif(k.Group, m.From, lib.NomorSaja(m.Sender), lib.NomorSaja(m.SenderAlt))
	nyala := toggleGrupAktif("antispam", m.From)
	if !sudahBan && !nyala {
		return false, nil
	}
	gc, err := k.Conn.GetGroupInfo(m.From)
	if err != nil {
		return false, nil
	}
	pelakuAdmin := pengirimAdmin(k, gc)
	if !botAdmin(k, gc) {
		return false, nil
	}
	ts := m.Timestamp
	if ts.IsZero() {
		ts = time.Now()
	}
	if time.Since(ts) > jendelaSpam {
		return false, nil
	}
	jejakSpam.Lock()
	perGrup := jejakSpam.data[m.From]
	if perGrup == nil {
		perGrup = map[string][]time.Time{}
		jejakSpam.data[m.From] = perGrup
	}
	arr := append(perGrup[num], ts)
	for len(arr) > 0 && ts.Sub(arr[0]) > jendelaSpam {
		arr = arr[1:]
	}
	if len(arr) < batasSpam {
		perGrup[num] = arr
		jejakSpam.Unlock()
		return false, nil
	}
	perGrup[num] = nil
	jejakSpam.Unlock()

	jid := m.Sender
	if sudahBan {
		if err := k.Conn.UpdateGroupParticipants(m.From, []string{jid}, "remove"); err != nil {
			fmt.Printf("[antispam] kick gagal: %v\n", err)
			return true, nil
		}
		fmt.Printf("[antispam] kick %s di %s (banmember nyepam)\n", num, m.From)
		_ = k.Conn.SendMessage(m.From, lib.IsiPesan{Text: fmt.Sprintf("@%s dikick karena nyepam padahal udah kena ban.", num), Mentions: []string{jid}}, nil)
		return true, nil
	}
	dipecat := false
	if pelakuAdmin {
		if err := k.Conn.UpdateGroupParticipants(m.From, []string{jid}, "demote"); err != nil {
			fmt.Printf("[antispam] pecat admin gagal: %v\n", err)
		} else {
			dipecat = true
			fmt.Printf("[antispam] pecat admin %s di %s (nyepam)\n", num, m.From)
		}
	}
	entri := map[string]any{"j": jid, "t": float64(time.Now().UnixMilli()), "d": float64(banSpamMs)}
	daftar := append(lib.DaftarBan(k.Group, m.From), entri)
	if k.Group.BanMember == nil {
		k.Group.BanMember = map[string]any{}
	}
	k.Group.BanMember[m.From] = daftar
	_ = k.Group.Simpan()
	fmt.Printf("[antispam] ban %s di %s\n", num, m.From)
	teks := fmt.Sprintf("@%s kena ban 1 jam karena nyepam (>5 pesan/10 detik).", num)
	if dipecat {
		teks = fmt.Sprintf("@%s dipecat dari admin + kena ban 1 jam karena nyepam (>5 pesan/10 detik).", num)
	}
	_ = k.Conn.SendMessage(m.From, lib.IsiPesan{Text: teks, Mentions: []string{jid}}, nil)
	return true, nil
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "antispam.js",
		Jenis:    "group",
		Nama:     []string{"#antispam on/off"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(antispam)$"),
		Jalankan: TanganiAntispam,
	})
}
