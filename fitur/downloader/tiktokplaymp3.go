package downloader

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	waAICommon "go.mau.fi/whatsmeow/proto/waAICommon"
	waAICommonDeprecated "go.mau.fi/whatsmeow/proto/waAICommonDeprecated"
	waE2E "go.mau.fi/whatsmeow/proto/waE2E"
	"google.golang.org/protobuf/proto"

	"gow1/lib"
)

const berkasRiwayatPlay = "data/riwayatttplay.json"
const maksRiwayatPlay = 30

type entriPlay struct {
	Nama     string `json:"nama"`
	Tii      dataTT `json:"tii"`
	Link     string `json:"link"`
	SisiForm string `json:"sisiForm,omitempty"`
}

func bacaRiwayatPlay() map[string][]entriPlay {
	hasil := map[string][]entriPlay{}
	b, err := os.ReadFile(berkasRiwayatPlay)
	if err != nil {
		return hasil
	}
	_ = json.Unmarshal(b, &hasil)
	return hasil
}

func tulisRiwayatPlay(d map[string][]entriPlay) {
	b, err := json.Marshal(d)
	if err != nil {
		return
	}
	if err := os.WriteFile(berkasRiwayatPlay, b, 0o644); err != nil {
		fmt.Println("[ttplay] simpen riwayat gagal:", err)
	}
}

func simpanRiwayatPlay(chat string, entri entriPlay) {
	d := bacaRiwayatPlay()
	d[chat] = append([]entriPlay{entri}, d[chat]...)
	if len(d[chat]) > maksRiwayatPlay {
		d[chat] = d[chat][:maksRiwayatPlay]
	}
	tulisRiwayatPlay(d)
}

func daftarRiwayatPlay(chat string) []entriPlay {
	return bacaRiwayatPlay()[chat]
}

func contextInfoDariPesan(m *waE2E.Message) *waE2E.ContextInfo {
	if m == nil {
		return nil
	}
	switch {
	case m.GetExtendedTextMessage() != nil:
		return m.GetExtendedTextMessage().GetContextInfo()
	case m.GetImageMessage() != nil:
		return m.GetImageMessage().GetContextInfo()
	case m.GetVideoMessage() != nil:
		return m.GetVideoMessage().GetContextInfo()
	case m.GetDocumentMessage() != nil:
		return m.GetDocumentMessage().GetContextInfo()
	case m.GetAudioMessage() != nil:
		return m.GetAudioMessage().GetContextInfo()
	case m.GetStickerMessage() != nil:
		return m.GetStickerMessage().GetContextInfo()
	case m.GetBotForwardedMessage() != nil:
		return contextInfoDariPesan(m.GetBotForwardedMessage().GetMessage())
	case m.GetRichResponseMessage() != nil:
		return m.GetRichResponseMessage().GetContextInfo()
	}
	return nil
}

func subjekGrupDikutip(k *lib.Konteks) string {
	if k.M.Quoted == nil {
		return ""
	}
	qm, ok := k.M.Quoted.Full.(*waE2E.Message)
	if !ok {
		return ""
	}
	ci := contextInfoDariPesan(qm)
	if ci == nil {
		return ""
	}
	gm := ci.GetGroupMentions()
	if len(gm) == 0 {
		return ""
	}
	return gm[0].GetGroupSubject()
}

func escPlay(s string) string {
	r := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		"\"", "&quot;",
	)
	return r.Replace(s)
}

func jsonStringPlay(s string) string {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(s)
	return strings.TrimRight(buf.String(), "\n")
}

func halamanPlay(entri entriPlay, port int, host string, form bool, cari bool, sisiForm string) string {
	sisi := sisiForm
	if sisi == "" {
		sisi = entri.Nama
	}
	cariQS := ""
	if cari {
		cariQS = "&cari=1"
	}
	formHTML := `<div style="padding:18px;flex:1;display:flex;flex-direction:column;justify-content:center"><input id="url" type="url" placeholder="Tempel link tiktok di sini" style="width:100%;max-width:280px;display:block;margin:0 auto 12px;padding:12px;border-radius:10px;border:1.5px solid rgba(255,255,255,.3);background:rgba(255,255,255,.08);color:#fff;font-size:14px;box-sizing:border-box;outline:none" /><button onclick="kirim()" style="display:block;width:100%;max-width:280px;margin:0 auto;background:rgba(0,210,211,.3);border:1.5px solid #00d2d3;color:#fff;padding:10px;border-radius:10px;font-weight:bold;cursor:pointer">Download 🎵</button><div style="text-align:center;margin-bottom:12px;font-size:11px;color:rgba(255,255,255,.6);margin-top:12px"><span id="dot" style="display:inline-block;width:8px;height:8px;border-radius:50%;background:#ff6b6b;margin-right:6px"></span>WS: <span id="stws">nyambung...</span></div><div id="status" style="text-align:center;font-size:12px;color:rgba(255,255,255,.6)"></div></div><script>var fws=null;function st(e){const d=document.getElementById("stws"),o=document.getElementById("dot");if(!d)return;d.textContent=e;if(o)o.style.background=e.indexOf("siap")>-1?"#00d2d3":e.indexOf("nyambung")>-1?"#ffd32a":"#ff6b6b"}function ganti(h){document.open();document.write(h);document.close()}function hubung(){st("nyambung...");fws=new WebSocket("ws://` + host + `:` + strconv.Itoa(port) + `/?sisi=` + sisi + cariQS + `"),fws.onopen=()=>st("siap ✓"),fws.onmessage=e=>{const j=JSON.parse(e.data);if(j.form)st("siap ✓");if(j.proses)document.getElementById("status").textContent="🎵 "+j.judul+" — diproses...";if(j.player)ganti(j.html)},fws.onclose=()=>st("putus ✗"),fws.onerror=()=>st("error ✗")}function kirim(){const u=document.getElementById("url").value.trim();if(!u)return document.getElementById("status").textContent="isinya kosong";if(!fws||fws.readyState!==1)return document.getElementById("status").textContent="WS belum siap, tunggu bentar";fws.send(JSON.stringify({url:u})),document.getElementById("status").textContent="ngirim ke bot..."}hubung()</script>`
	var isi string
	if form {
		isi = formHTML
	} else {
		formHtmlJSON := strings.ReplaceAll(jsonStringPlay(halamanPlay(entriPlay{Nama: entri.Nama}, port, host, true, true, sisi)), "</script>", "<\\/script>")
		isi = `<div style="padding:18px;flex:1;display:flex;flex-direction:column;justify-content:center"><div style="max-width:280px;margin:0 auto 14px;background:rgba(0,0,0,.35);padding:10px;border-radius:16px;border:1px solid rgba(255,255,255,.15)"><input id="bar" type="range" min="0" max="1000" value="0" style="width:100%;accent-color:#00d2d3;cursor:pointer"></div><div id="status" style="text-align:center;margin-bottom:6px;font-size:13px;color:rgba(255,255,255,.85);font-weight:bold">0:00 / 0:00</div><div style="text-align:center;margin-bottom:12px;font-size:11px;color:rgba(255,255,255,.6)"><span id="dot" style="display:inline-block;width:8px;height:8px;border-radius:50%;background:#ff6b6b;margin-right:6px"></span>WS: <span id="stws">belum nyoba</span></div><button onclick="toggle()" style="display:block;width:100%;max-width:280px;margin:0 auto;background:rgba(255,255,255,.15);border:1.5px solid rgba(255,255,255,.3);color:#fff;padding:10px;border-radius:10px;font-weight:bold;cursor:pointer;margin-bottom:10px">Play / Pause 🎵</button><button onclick="ganti(formHtml)" style="display:block;width:100%;max-width:280px;margin:0 auto;background:rgba(0,210,211,.3);border:1.5px solid #00d2d3;color:#fff;padding:10px;border-radius:10px;font-weight:bold;cursor:pointer">Cari Lain 🔄</button></div><script>var formHtml=` + formHtmlJSON + `;var au=null,wss=null,siap=!1;function fmt(e){return Math.floor(e/60)+":"+String(Math.floor(e%60)).padStart(2,"0")}function st(e){const d=document.getElementById("stws"),o=document.getElementById("dot");if(!d)return;d.textContent=e;if(o)o.style.background=e.indexOf("tersambung")>-1?"#00d2d3":e.indexOf("nyambung")>-1?"#ffd32a":"#ff6b6b"}function hubung(){st("nyambung...");wss=new WebSocket("ws://` + host + `:` + strconv.Itoa(port) + `/?id=` + entri.Nama + `"),wss.binaryType="arraybuffer",wss.onopen=()=>{st("tersambung ✓")},wss.onmessage=e=>{if("done"===e.data)return siap=!0,void(upd());au||(au=new Audio,au.onended=()=>{document.getElementById("status").textContent=fmt(au.duration)+" / "+fmt(au.duration)}),au.src=URL.createObjectURL(new Blob([au.buf?new Blob([au.buf,e.data],{type:"audio/mpeg"}):e.data],{type:"audio/mpeg"})),au.buf=au.buf?new Blob([au.buf,e.data]):e.data,siap||upd()},wss.onclose=()=>{st("putus ✗"),siap||(document.getElementById("status").textContent="gagal nyambung ke server 😔")},wss.onerror=()=>st("error ✗")}function upd(){const b=document.getElementById("bar"),s=document.getElementById("status");if(!b)return;const e=au?au.duration:0;b.value=e?au.currentTime/e*1e3:0,s.textContent=(au?fmt(au.currentTime):"0:00")+" / "+fmt(e||0)}function toggle(){if(au)au.paused?au.play():au.pause();else if(!wss)return hubung();else if(!au&&!siap)return}document.getElementById("bar").oninput=function(e){au&&au.duration&&(au.currentTime=e.target.value/1e3*au.duration)},setInterval(upd,250),hubung()</script>`
	}
	judul := "TikTok Play MP3"
	label := "FORM"
	nilai := "INPUT URL"
	if !form {
		judul = entri.Tii.Music
		if judul == "" {
			judul = entri.Tii.Title
		}
		if judul == "" {
			judul = "TikTok MP3"
		}
		judul = escPlay(judul)
		label = "UPLOADER"
		nilai = entri.Tii.Nickname
		if nilai == "" {
			nilai = "-"
		}
		nilai = escPlay(nilai)
	}
	return `<style>*{-webkit-tap-highlight-color:transparent;-webkit-user-select:none;user-select:none;-webkit-touch-callout:none;box-sizing:border-box}</style><body style="margin:0;background:transparent;font-family:Arial,sans-serif;color:#eee;touch-action:manipulation;cursor:pointer"><div style="width:100%;max-width:620px;aspect-ratio:1;margin:auto;padding:16px;box-sizing:border-box"><div style="background:rgba(20,20,30,.85);backdrop-filter:blur(14px);-webkit-backdrop-filter:blur(14px);border:1.5px solid rgba(255,255,255,.25);border-radius:16px;overflow:hidden;height:100%;display:flex;flex-direction:column"><div style="padding:16px 20px;border-bottom:2px solid rgba(255,255,255,.2);display:flex;justify-content:space-between;align-items:center"><div><div style="font-size:11px;letter-spacing:1.5px;color:rgba(255,255,255,.55)">NIXEL PLAYER</div><div style="font-size:20px;font-weight:bold;color:#fff">` + judul + `</div></div><div style="text-align:right"><div style="font-size:10px;color:rgba(255,255,255,.5)">` + label + `</div><div style="font-size:14px;font-weight:bold;color:#00d2d3">` + nilai + `</div></div></div>` + isi + `</div></div></body>`
}

func kirimHTMLPlay(k *lib.Konteks, entri entriPlay, form bool) error {
	html := halamanPlay(entri, PortTT, HostTT, form, false, "")
	teksSub := "TikTok Play MP3 — tempel link di form"
	if !form {
		teksSub = fmt.Sprintf("%s • %s", entri.Tii.Nickname, entri.Tii.Title)
		if r := []rune(teksSub); len(r) > 100 {
			teksSub = string(r[:100])
		}
	}
	payload, err := json.Marshal(map[string]any{
		"response_id": entri.Nama,
		"sections": []any{
			map[string]any{
				"view_model": map[string]any{
					"primitive": map[string]any{
						"__typename":      "GenAIaeacdsnwHtmlPrimitive",
						"payload":         html,
						"trusted_sources": []any{},
					},
					"__typename": "GenAISingleLayoutViewModel",
				},
			},
		},
	})
	if err != nil {
		return err
	}
	pesan := &waE2E.Message{
		BotForwardedMessage: &waE2E.FutureProofMessage{
			Message: &waE2E.Message{
				RichResponseMessage: &waE2E.AIRichResponseMessage{
					MessageType: waAICommonDeprecated.AIRichResponseMessageType_AI_RICH_RESPONSE_TYPE_STANDARD.Enum(),
					Submessages: []*waAICommonDeprecated.AIRichResponseSubMessage{
						{
							MessageType: waAICommonDeprecated.AIRichResponseSubMessageType_AI_RICH_RESPONSE_TEXT.Enum(),
							MessageText: proto.String(teksSub),
						},
					},
					UnifiedResponse: &waAICommon.AIRichResponseUnifiedResponse{
						Data: []byte(base64.StdEncoding.EncodeToString(payload)),
					},
					ContextInfo: &waE2E.ContextInfo{
						ForwardingScore: proto.Uint32(1),
						IsForwarded:     proto.Bool(true),
						ForwardedAiBotMessageInfo: &waAICommon.ForwardedAIBotMessageInfo{
							BotJID: proto.String("867051314767696@bot"),
						},
						ForwardOrigin: waE2E.ContextInfo_META_AI.Enum(),
					},
				},
			},
		},
	}
	return k.Conn.RelayMentah(k.M.From, pesan)
}

func prosesPlay(k *lib.Konteks, link, sisi string, dorong func(pesanWS) bool, senyap bool) error {
	muTT.Lock()
	ttCariKonteks = k
	muTT.Unlock()
	if !strings.HasPrefix(link, "http") {
		link = "https://" + link
	}
	tii := AmbilDataTT(link)
	if tii.MusicURL == "" && tii.Play == "" {
		return k.Conn.OK(k.M, "Gagal ambil audionya, coba lagi")
	}
	kirim := dorong
	if kirim != nil {
		judul := tii.Music
		if judul == "" {
			judul = tii.Title
		}
		if judul == "" {
			judul = "?"
		}
		kirim(pesanWS{Proses: true, Judul: judul})
	} else if !senyap {
		judul := tii.Music
		if judul == "" {
			judul = tii.Title
		}
		if err := k.Conn.OK(k.M, fmt.Sprintf("%s\nSedang disiapkan", judul)); err != nil {
			return err
		}
	}
	sumber := tii.MusicURL
	if sumber == "" {
		sumber = tii.Play
	}
	buf, err := lib.AmbilURL(sumber, nil)
	if err != nil {
		return k.Conn.OK(k.M, "Gagal download audionya 😔")
	}
	mp3, err := mp42mp3(buf)
	if err != nil {
		return err
	}
	nama := simpanLagu(mp3, 2)
	entri := entriPlay{Nama: nama, Tii: tii, Link: link, SisiForm: sisi}
	simpanRiwayatPlay(k.M.From, entri)
	muTT.Lock()
	ttCariLain = jalankanCariLain
	muTT.Unlock()
	if kirim != nil && kirim(pesanWS{Player: true, HTML: halamanPlay(entri, PortTT, HostTT, false, false, sisi), Nama: nama}) {
		return nil
	}
	if senyap {
		return nil
	}
	return kirimHTMLPlay(k, entri, false)
}

func jalankanCariLain(j dataCariLain) bool {
	muTT.Lock()
	k := ttCariKonteks
	muTT.Unlock()
	if k == nil {
		return false
	}
	baru := polaTiktok.FindString(j.URL)
	if baru == "" {
		return true
	}
	_ = prosesPlay(k, baru, j.Nama, j.Dorong, true)
	return true
}

func TanganiTiktokplaymp3(k *lib.Konteks) error {
	if subjek := subjekGrupDikutip(k); strings.HasPrefix(subjek, "ttplay|") {
		nomor := strings.SplitN(subjek, "|", 2)[1]
		r := daftarRiwayatPlay(k.M.From)
		var entri *entriPlay
		if nomor == "main" {
			if len(r) > 0 {
				entri = &r[0]
			}
		} else if n, err := strconv.Atoi(nomor); err == nil && n >= 1 && n <= len(r) {
			entri = &r[n-1]
		}
		if entri == nil {
			return k.Conn.OK(k.M, "Riwayatnya udah basi")
		}
		return kirimHTMLPlay(k, *entri, false)
	}
	if link := polaTiktok.FindString(k.M.Q); link != "" {
		return prosesPlay(k, link, "", nil, false)
	}
	sisi := acakHex(6)
	type hasilTunggu struct {
		d  dataKirim
		ok bool
	}
	ch := make(chan hasilTunggu, 1)
	go func() {
		d, ok := tungguKirim(sisi, 300)
		ch <- hasilTunggu{d, ok}
	}()
	if err := kirimHTMLPlay(k, entriPlay{Nama: sisi}, true); err != nil {
		return err
	}
	jwb := <-ch
	if !jwb.ok {
		return k.Conn.OK(k.M, "Formnya gak diisi dalam 5 menit, kirim #ttplaymp3 lagi yaa")
	}
	link := polaTiktok.FindString(jwb.d.URL)
	if link == "" {
		return k.Conn.OK(k.M, "Format link TikTok tidak sesuai")
	}
	return prosesPlay(k, link, sisi, func(data pesanWS) bool {
		return dorongForm(sisi, data)
	}, false)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "tiktokplaymp3.js",
		Jenis:    "downloader",
		Nama:     nil,
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(tiktokplaymp3|ttplaymp3)$"),
		Jalankan: TanganiTiktokplaymp3,
	})
}
