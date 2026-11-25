package main

import (
	"testing"

	"gow1/lib"
)

func TestRegistryDanRouting(t *testing.T) {
	if got := lib.Jumlah(); got != 61 {
		t.Fatalf("jumlah fitur = %d, mau 61", got)
	}
	kasus := map[string]string{
		"#menu":         "menu.js",
		"#stiker halo":  "stiker.js",
		"#hidetag halo": "hidetag.js",
		"#play lagu":    "ytplay.js",
		"#ttt":          "tictactoe.js",
		".e 1+1":        "eval.js",
	}
	for teks, mau := range kasus {
		f := lib.CariFitur(teks)
		if f == nil {
			t.Errorf("CariFitur(%q) = nil, mau %s", teks, mau)
			continue
		}
		if f.File != mau {
			t.Errorf("CariFitur(%q) = %s, mau %s", teks, f.File, mau)
		}
	}
}
