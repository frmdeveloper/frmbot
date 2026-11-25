package info

import "testing"

func TestTagTerlarangPixiv(t *testing.T) {
	kasus := []struct {
		tags []string
		mau  bool
	}{
		{[]string{"elaina", "魔女の旅々"}, false},
		{[]string{"ロリ"}, true},
		{[]string{"kawaii", "Loli"}, true},
		{[]string{"JS"}, true},
		{[]string{"jk"}, true},
		{[]string{"jsk"}, false},
		{[]string{"underage"}, true},
		{nil, false},
	}
	for _, c := range kasus {
		if got := adaTagTerlarangPixiv(c.tags); got != c.mau {
			t.Errorf("adaTagTerlarangPixiv(%v) = %v, mau %v", c.tags, got, c.mau)
		}
	}
}

func TestRegexPixiv(t *testing.T) {
	if m := reLinkPixiv.FindStringSubmatch("https://www.pixiv.net/en/artworks/151511113"); m == nil || m[1] != "151511113" {
		t.Errorf("link en: %v", m)
	}
	if m := reLinkPixiv.FindStringSubmatch("pixiv.net/artworks/12345678"); m == nil || m[1] != "12345678" {
		t.Errorf("link polos: %v", m)
	}
	if reLinkPixiv.FindStringSubmatch("https://example.com/artworks/123") != nil {
		t.Error("domain lain harusnya gak cocok")
	}
	if !reIDPixiv.MatchString("151511113") || reIDPixiv.MatchString("12345") || reIDPixiv.MatchString("elaina") {
		t.Error("regex ID salah")
	}
}
