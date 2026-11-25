package utility

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/dop251/goja"

	"gow1/lib"
)

var reWarna = regexp.MustCompile("\x1b(?:[@-Z\\\\-_]|\\[[0-?]*[ -/]*[@-~])")

func bersihinWarna(s string) string {
	s = reWarna.ReplaceAllString(s, "")
	return strings.ReplaceAll(s, "\r", "")
}

func adalahOwnerEval(k *lib.Konteks) bool {
	for _, jid := range []string{k.M.Sender, k.M.SenderAlt} {
		n := lib.NomorSaja(jid)
		if n == "6289680237010" || n == "6285188261898" || n == k.Conn.BotNomor() {
			return true
		}
	}
	return false
}

const skripFormat = `
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

func objekPesan(m *lib.Pesan) map[string]any {
	obj := map[string]any{
		"id": m.ID, "from": m.From, "sender": m.Sender, "senderAlt": m.SenderAlt,
		"text": m.Text, "q": m.Q, "command": m.Command, "prefix": m.Prefix,
		"isGroup": m.IsGroup, "fromMe": m.FromMe, "type": m.Type,
	}
	if m.Quoted != nil {
		obj["quoted"] = map[string]any{
			"id": m.Quoted.ID, "sender": m.Quoted.Sender, "text": m.Quoted.Text, "type": m.Quoted.Type,
		}
	}
	return obj
}

func pasangConn(vm *goja.Runtime, k *lib.Konteks) {
	m := k.M
	connObj := vm.NewObject()
	_ = connObj.Set("number", k.Conn.BotNomor())
	_ = connObj.Set("ok", func(teks string) (any, error) {
		err := k.Conn.OK(m, teks)
		return nil, err
	})
	_ = connObj.Set("sendMessage", func(jid string, isi any) (any, error) {
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
		err := k.Conn.SendMessage(jid, pesan, m)
		return nil, err
	})
	_ = vm.Set("conn", connObj)
}

func jalankanJS(k *lib.Konteks, kode string, denganConn bool) string {
	vm := goja.New()
	if _, err := vm.RunString(skripFormat); err != nil {
		return err.Error()
	}
	_ = vm.Set("m", objekPesan(k.M))
	if denganConn {
		pasangConn(vm, k)
	}
	program := "(async () => {\n" + kode + "\n})().then(function(v) { globalThis.__selesai = { ok: true, v: v }; }, function(e) { globalThis.__selesai = { ok: false, v: e }; });"
	if _, err := vm.RunString(program); err != nil {
		return formatGalat(vm, err)
	}
	batas := time.Now().Add(10 * time.Second)
	for time.Now().Before(batas) {
		s := vm.Get("__selesai")
		if s != nil && !goja.IsUndefined(s) && !goja.IsNull(s) {
			obj := s.ToObject(vm)
			hasil := obj.Get("v")
			if ok, _ := obj.Get("ok").Export().(bool); ok {
				return panggilFmt(vm, hasil)
			}
			return formatNilai(vm, hasil)
		}
		if _, err := vm.RunString("void 0"); err != nil {
			break
		}
		time.Sleep(2 * time.Millisecond)
	}
	return "Promise { <pending> }"
}

func panggilFmt(vm *goja.Runtime, v goja.Value) string {
	f, ok := goja.AssertFunction(vm.Get("__fmt"))
	if !ok {
		return fmt.Sprint(v.Export())
	}
	hasil, err := f(goja.Undefined(), v)
	if err != nil {
		return err.Error()
	}
	return hasil.String()
}

func formatNilai(vm *goja.Runtime, v goja.Value) string {
	if v == nil || goja.IsUndefined(v) {
		return "undefined"
	}
	return panggilFmt(vm, v)
}

func formatGalat(vm *goja.Runtime, err error) string {
	if je, ok := err.(*goja.Exception); ok {
		return formatNilai(vm, je.Value())
	}
	return err.Error()
}

func jalankanShell(k *lib.Konteks) error {
	m := k.M
	ctx, batal := context.WithTimeout(context.Background(), 120*time.Second)
	defer batal()
	cmd := exec.CommandContext(ctx, "sh", "-c", m.Q)
	cmd.Env = append(os.Environ(), "from="+m.From)
	var outBuf, errBuf strings.Builder
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	_ = cmd.Run()
	if outBuf.Len() > 0 {
		if err := k.Conn.SendMessage(m.From, lib.IsiPesan{Text: bersihinWarna(outBuf.String())}, m); err != nil {
			return err
		}
	}
	if errBuf.Len() > 0 {
		if err := k.Conn.SendMessage(m.From, lib.IsiPesan{Text: bersihinWarna(errBuf.String())}, m); err != nil {
			return err
		}
	}
	return nil
}

func TanganiEval(k *lib.Konteks) error {
	m := k.M
	if !adalahOwnerEval(k) {
		return nil
	}
	switch {
	case m.Command == ".ex":
		if strings.TrimSpace(m.Q) == "" {
			return nil
		}
		return jalankanShell(k)
	case m.Command == ".ee":
		if strings.TrimSpace(m.Q) == "" {
			return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: "Error: astagfirullahaladzim"}, m)
		}
		hasil := jalankanJS(k, m.Q, false)
		return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: hasil}, m)
	default:
		if strings.TrimSpace(m.Q) == "" {
			return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: "Error: astagfirullahaladzim"}, m)
		}
		hasil := jalankanJS(k, m.Q, true)
		return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: hasil}, m)
	}
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "eval.js",
		Jenis:    "utility",
		Nama:     nil,
		Itunya:   regexp.MustCompile("(?i)^(.e|.ee|.ex|\\x{25aa}[\\x{fe0e}\\x{fe0f}]?)$"),
		Jalankan: TanganiEval,
	})
}
