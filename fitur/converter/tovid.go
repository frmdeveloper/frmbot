package converter

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"gow1/lib"
)

func TanganiTovid(k *lib.Konteks) error {
	var down []byte
	if k.M.Quoted != nil {
		if b, _, err := k.Conn.DownloadAny(k.M.Quoted); err == nil {
			down = b
		}
	}
	ada := len(down) > 0 && lib.DeteksiMime(down) == "image/webp"
	if !ada {
		return k.Conn.OK(k.M, "Stikernya mana bos 😅")
	}
	if err := k.Conn.OK(k.M, "⏳ Tunggu sebentar, lagi diproses..."); err != nil {
		return err
	}
	mp4, err := Webp2mp4Local(down)
	if err != nil {
		return err
	}
	return k.Conn.SendMessage(k.M.From, lib.IsiPesan{Video: mp4}, k.M)
}

func Webp2mp4Local(buffer []byte) ([]byte, error) {
	awal, err := lib.TulisTemp(".webp", buffer)
	if err != nil {
		return nil, err
	}
	defer os.Remove(awal)
	dasar := strings.TrimSuffix(awal, ".webp")
	gif := dasar + ".gif"
	akhir := dasar + ".mp4"
	defer os.Remove(gif)
	defer os.Remove(akhir)
	filter := "scale=trunc(iw/2)*2:trunc(ih/2)*2"
	lib.SpawnSync(fmt.Sprintf(`ffmpeg -y -i "%s" -vf "%s" -c:v libx264 -pix_fmt yuv420p -movflags +faststart "%s"`, awal, filter, akhir))
	if !jadiBerkas(akhir) {
		lib.SpawnSync(fmt.Sprintf(`convert "%s" -coalesce "%s"`, awal, gif))
		lib.SpawnSync(fmt.Sprintf(`ffmpeg -y -i "%s" -vf "%s" -c:v libx264 -pix_fmt yuv420p -movflags +faststart "%s"`, gif, filter, akhir))
	}
	if !jadiBerkas(akhir) {
		return nil, fmt.Errorf("ffmpeg gagal konversi webp ke mp4")
	}
	return os.ReadFile(akhir)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "tovid.js",
		Jenis:    "converter",
		Nama:     []string{"#tovid"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)to(vid|video|mp4)$"),
		Jalankan: TanganiTovid,
	})
}
