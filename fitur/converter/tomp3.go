package converter

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"gow1/lib"
)

func TanganiTomp3(k *lib.Konteks) error {
	var down []byte
	if b, _, err := k.Conn.DownloadAny(k.M); err == nil {
		down = b
	} else if k.M.Quoted != nil {
		if b2, _, err2 := k.Conn.DownloadAny(k.M.Quoted); err2 == nil {
			down = b2
		}
	}
	ada := len(down) > 0 && strings.HasPrefix(lib.DeteksiMime(down), "video/")
	if !ada {
		return k.Conn.OK(k.M, "Videonya mana 🤪")
	}
	mp3, err := Mp42mp3(down)
	if err != nil {
		return err
	}
	return k.Conn.SendMessage(k.M.From, lib.IsiPesan{Audio: mp3}, k.M)
}

func Mp42mp3(buffer []byte) ([]byte, error) {
	awal, err := lib.TulisTemp(".mp4", buffer)
	if err != nil {
		return nil, err
	}
	defer os.Remove(awal)
	akhir := strings.TrimSuffix(awal, ".mp4") + ".mp3"
	defer os.Remove(akhir)
	_, stderr := lib.SpawnSync(fmt.Sprintf(`ffmpeg -i "%s" "%s"`, awal, akhir))
	if !jadiBerkas(akhir) {
		return nil, fmt.Errorf("%s", ekorBaris(stderr, 2))
	}
	return os.ReadFile(akhir)
}

func Mp42ogg(buffer []byte) ([]byte, error) {
	awal, err := lib.TulisTemp(".mp4", buffer)
	if err != nil {
		return nil, err
	}
	defer os.Remove(awal)
	akhir := strings.TrimSuffix(awal, ".mp4") + ".ogg"
	defer os.Remove(akhir)
	_, stderr := lib.SpawnSync(fmt.Sprintf(`ffmpeg -i "%s" -map_metadata -1 -c:a libopus -b:a 128k -vbr on -compression_level 10 -application voip -compression undictated -frame_duration 60 -f opus "%s"`, awal, akhir))
	if !jadiBerkas(akhir) {
		return nil, fmt.Errorf("%s", ekorBaris(stderr, 2))
	}
	return os.ReadFile(akhir)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "tomp3.js",
		Jenis:    "converter",
		Nama:     []string{"#tomp3"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)((to)?(mp3|audio|vn))$"),
		Jalankan: TanganiTomp3,
	})
}
