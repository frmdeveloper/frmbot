package utility

import (
	"gow1/lib"
)

func TanganiBantuan(k *lib.Konteks) error {

	return nil
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "bantuan.js",
		Jenis:    "utility",
		Nama:     nil,
		Itunya:   nil,
		Jalankan: TanganiBantuan,
	})
}
