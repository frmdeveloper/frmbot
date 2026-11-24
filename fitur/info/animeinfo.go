package info

import (
	"fmt"
	"html"
	"regexp"
	"strings"

	"gow1/lib"
)

var (
	reLinkMAL     = regexp.MustCompile(`href="(https://myanimelist\.net/anime/\d+/[^"]*)"`)
	reJudulMAL    = regexp.MustCompile(`(?s)h1-title.*?<h1[^>]*>(.*?)</h1>`)
	reGambarMAL   = regexp.MustCompile(`(https://cdn\.myanimelist\.net/images/anime/[^"]+)`)
	reSinopsisMAL = regexp.MustCompile(`(?s)<p itemprop="description">(.*?)</p>`)
)

type infoAnimeMAL struct {
	Title    string
	Image    string
	Synopsis string
	Info     map[string]string
}

func teksBersihHTML(s string) string {
	s = reTagHTML.ReplaceAllString(s, "")
	s = html.UnescapeString(s)
	return strings.TrimSpace(s)
}

func infoLengkapMAL(link string) (*infoAnimeMAL, error) {
	b, err := lib.AmbilURL(link, nil)
	if err != nil {
		return nil, err
	}
	hal := string(b)
	hasil := &infoAnimeMAL{Info: map[string]string{}}
	if mj := reJudulMAL.FindStringSubmatch(hal); mj != nil {
		hasil.Title = teksBersihHTML(mj[1])
	}
	if mg := reGambarMAL.FindStringSubmatch(hal); mg != nil {
		hasil.Image = mg[1]
	}
	if ms := reSinopsisMAL.FindStringSubmatch(hal); ms != nil {
		hasil.Synopsis = teksBersihHTML(ms[1])
	}
	bagian := strings.Split(hal, "class=\"spaceit_pad\"")
	for _, bgn := range bagian[1:] {
		potong := strings.Index(bgn, "</div>")
		if potong < 0 {
			continue
		}
		teks := teksBersihHTML(bgn[:potong])
		pecah := strings.Split(teks, ":")
		if len(pecah) < 2 {
			continue
		}
		nama := strings.TrimSpace(strings.Split(pecah[0], "\n")[0])
		nilai := strings.TrimSpace(strings.Split(pecah[1], "\n")[0])
		if nama != "" {
			hasil.Info[nama] = nilai
		}
	}
	return hasil, nil
}

func cariAnimeMAL(q string) (*infoAnimeMAL, error) {
	b, err := lib.AmbilURL("https://myanimelist.net/anime.php?cat=anime&q="+q, nil)
	if err != nil {
		return nil, err
	}
	ml := reLinkMAL.FindStringSubmatch(string(b))
	if ml == nil {
		return nil, fmt.Errorf("not found")
	}
	return infoLengkapMAL(ml[1])
}

func TanganiAnimeinfo(k *lib.Konteks) error {
	m := k.M
	if m.Q == "" {
		return k.Conn.OK(m, "Judul animenya apa bang🤨")
	}
	an, err := cariAnimeMAL(m.Q)
	if err != nil {
		return err
	}
	statusID, err := lib.Gtr("id", an.Info["Status"])
	if err != nil {
		return err
	}
	sinopsisID, err := lib.Gtr("id", an.Synopsis)
	if err != nil {
		return err
	}
	genre := an.Info["Genre"]
	if genre == "" {
		genre = an.Info["Genres"]
	}
	caption := fmt.Sprintf(`
✨️ *Judul:* %s
🎗️ *Genre:* %s
🧬 *Diadaptasi:* %s
⏰ *Tayang:* %s
⏰ *Status:* %s
⏳ *Durasi:* %s
🎞 *Type:* %s
📽 *Studio:* %s
💌️ *Rating:* %s

*• Sinopsis:*
_(diterjemahkan dari bhs. Inggris)_
%s`, an.Title, genre, an.Info["Source"], an.Info["Aired"], statusID, an.Info["Duration"], an.Info["Type"], an.Info["Studios"], an.Info["Rating"], sinopsisID)
	gambar, err := lib.AmbilURL(an.Image, nil)
	if err != nil {
		return err
	}
	return k.Conn.SendMessage(m.From, lib.IsiPesan{Image: gambar, Caption: caption}, m)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "animeinfo.js",
		Jenis:    "info",
		Nama:     []string{"#infoanime"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)((info)?(anime)(info)?)$"),
		Jalankan: TanganiAnimeinfo,
	})
}
