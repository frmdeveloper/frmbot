package group

import (
	"fmt"
	"regexp"

	"gow1/lib"
)

func TanganiResetwallpaper(k *lib.Konteks) error {
	m := k.M
	if !m.IsGroup {
		return k.Conn.OK(m, "ini cuma buat grup jir")
	}
	gc, err := k.Conn.GetGroupInfo(m.From)
	if err != nil {
		return err
	}
	if !pengirimOwner(k) && !pengirimAdmin(k, gc) {
		return k.Conn.OK(m, "lu bukan admin jir")
	}
	if !botAdmin(k, gc) {
		return k.Conn.OK(m, "gw bukan admin di grup ini — jadiin admin dulu biar bisa reset wallpaper buat semua orang")
	}
	if err := k.Conn.ResetGroupWallpaper(m.From); err != nil {
		return k.Conn.OK(m, fmt.Sprintf("gagal reset wallpapernya jir: %v", err))
	}
	hapusWallpaperGrup(m.From)
	return k.Conn.OK(m, "wallpaper grup udah direset ke bawaan")
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "resetwallpaper.js",
		Jenis:    "group",
		Nama:     []string{"#resetwallpaper"},
		Itunya:   regexp.MustCompile(`(?i)^(#|!|/|\.)resetwallpaper$`),
		Jalankan: TanganiResetwallpaper,
	})
}
