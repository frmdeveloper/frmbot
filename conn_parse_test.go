package main

import (
	"strings"
	"testing"
)

func TestGabungPrefixSpasi(t *testing.T) {
	kasus := []struct {
		teks    string
		command string
		q       string
	}{
		{"# menu", "#menu", ""},
		{"# menu maker", "#menu", "maker"},
		{"#menu", "#menu", ""},
		{"#menu maker", "#menu", "maker"},
		{"! s", "!s", ""},
		{". e 1+1", ".e", "1+1"},
		{"#  s2", "#s2", ""},
		{"#", "#", ""},
		{"halo elfa", "halo", "elfa"},
		{"# 1", "#1", ""},
		{"biasa aja", "biasa", "aja"},
	}
	for _, c := range kasus {
		cmd, q := gabungPrefixSpasi(strings.Fields(c.teks))
		if cmd != c.command || q != c.q {
			t.Errorf("%q: dapet (%q, %q), mau (%q, %q)", c.teks, cmd, q, c.command, c.q)
		}
	}
}
