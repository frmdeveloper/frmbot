package info

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"gow1/lib"
)

const miripMinim = 0.87

const anilistQ = "query($id:Int){Media(id:$id,type:ANIME){description(asHtml:false) averageScore\n title{romaji english native} format status episodes duration source genres\n startDate{year month day} endDate{year month day} isAdult\n studios(isMain:true){nodes{name}} coverImage{extraLarge large} siteUrl}}"

var bulanAnime = []string{"Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Agu", "Sep", "Okt", "Nov", "Des"}

var statusIDAnime = map[string]string{
	"FINISHED":         "Selesai tayang",
	"RELEASING":        "Masih tayang",
	"NOT_YET_RELEASED": "Belum tayang",
	"CANCELLED":        "Dibatalin",
	"HIATUS":           "Hiatus",
}

var sumberIDAnime = map[string]string{
	"ORIGINAL":           "Cerita asli",
	"MANGA":              "Manga",
	"LIGHT_NOVEL":        "Light novel",
	"VISUAL_NOVEL":       "Visual novel",
	"VIDEO_GAME":         "Game",
	"GAME":               "Game",
	"NOVEL":              "Novel",
	"DOUJINSHI":          "Doujinshi",
	"ANIME":              "Anime",
	"WEB_NOVEL":          "Web novel",
	"COMIC":              "Komik",
	"LIVE_ACTION":        "Live action",
	"MULTIMEDIA_PROJECT": "Proyek multimedia",
	"PICTURE_BOOK":       "Buku bergambar",
	"OTHER":              "Lain-lain",
}

type tanggalAnilist struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type studioAnilist struct {
	Nodes []struct {
		Name string `json:"name"`
	} `json:"nodes"`
	Edges []struct {
		IsMain bool `json:"isMain"`
		Node   struct {
			Name string `json:"name"`
		} `json:"node"`
	} `json:"edges"`
}

type mediaAnilist struct {
	ID           int    `json:"id"`
	Description  string `json:"description"`
	AverageScore int    `json:"averageScore"`
	Title        struct {
		Romaji  string `json:"romaji"`
		English string `json:"english"`
		Native  string `json:"native"`
	} `json:"title"`
	Format    string         `json:"format"`
	Status    string         `json:"status"`
	Episodes  int            `json:"episodes"`
	Duration  int            `json:"duration"`
	Source    string         `json:"source"`
	Genres    []string       `json:"genres"`
	StartDate tanggalAnilist `json:"startDate"`
	EndDate   tanggalAnilist `json:"endDate"`
	IsAdult   bool           `json:"isAdult"`
	Studios   studioAnilist  `json:"studios"`
	SiteURL   string         `json:"siteUrl"`
}

func tanggalAnilistStr(d tanggalAnilist) string {
	if d.Year == 0 {
		return ""
	}
	bulan := "Jan"
	if d.Month >= 1 && d.Month <= 12 {
		bulan = bulanAnime[d.Month-1]
	}
	var bagian []string
	if d.Day != 0 {
		bagian = append(bagian, strconv.Itoa(d.Day))
	}
	bagian = append(bagian, bulan, strconv.Itoa(d.Year))
	return strings.Join(bagian, " ")
}

var (
	reBrHTML    = regexp.MustCompile(`(?i)<br\s*/?>`)
	reTagHTML2  = regexp.MustCompile(`<[^>]+>`)
	reNbspHTML  = regexp.MustCompile(`&nbsp;`)
	reAmpHTML   = regexp.MustCompile(`&amp;`)
	reQuotHTML  = regexp.MustCompile(`&quot;`)
	reAposHTML  = regexp.MustCompile(`&#0?39;|&apos;`)
	reBarisHTML = regexp.MustCompile(`\n{3,}`)
)

func bersihHTMLAnime(s string) string {
	s = reBrHTML.ReplaceAllString(s, "\n")
	s = reTagHTML2.ReplaceAllString(s, "")
	s = reNbspHTML.ReplaceAllString(s, " ")
	s = reAmpHTML.ReplaceAllString(s, "&")
	s = reQuotHTML.ReplaceAllString(s, "\"")
	s = reAposHTML.ReplaceAllString(s, "'")
	s = reBarisHTML.ReplaceAllString(s, "\n\n")
	return strings.TrimSpace(s)
}

func kirimPOSTMentah(alamat string, body []byte, tipeKonten string) ([]byte, error) {
	req, err := http.NewRequest("POST", alamat, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	if tipeKonten != "" {
		req.Header.Set("Content-Type", tipeKonten)
	}
	resp, err := lib.KlienHTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func anilistMedia(id int) (*mediaAnilist, error) {
	body, _ := json.Marshal(map[string]any{
		"query":     anilistQ,
		"variables": map[string]int{"id": id},
	})
	b, err := kirimPOSTMentah("https://graphql.anilist.co", body, "application/json")
	if err != nil {
		return nil, err
	}
	var resp struct {
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
		Data struct {
			Media *mediaAnilist `json:"Media"`
		} `json:"data"`
	}
	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}
	if len(resp.Errors) > 0 {
		return nil, fmt.Errorf("%s", resp.Errors[0].Message)
	}
	return resp.Data.Media, nil
}

type hasilAnimeApa struct {
	Judul    string
	JudulEn  string
	Genre    string
	Sumber   string
	Tayang   string
	Status   string
	Durasi   string
	Tipe     string
	Studio   string
	Skor     string
	Sinopsis string
	TotalEps int
	Dewasa   bool
	Tautan   string
	Episode  json.RawMessage
	Dari     float64
	Sampai   float64
	Mirip    float64
	Video    string
	Gambar   string
}

func gabungMediaAnilist(a, m mediaAnilist) mediaAnilist {
	if m.Description != "" {
		a.Description = m.Description
	}
	if m.AverageScore != 0 {
		a.AverageScore = m.AverageScore
	}
	if m.Title.Romaji != "" || m.Title.English != "" || m.Title.Native != "" {
		a.Title = m.Title
	}
	if m.Format != "" {
		a.Format = m.Format
	}
	if m.Status != "" {
		a.Status = m.Status
	}
	if m.Episodes != 0 {
		a.Episodes = m.Episodes
	}
	if m.Duration != 0 {
		a.Duration = m.Duration
	}
	if m.Source != "" {
		a.Source = m.Source
	}
	if len(m.Genres) > 0 {
		a.Genres = m.Genres
	}
	if m.StartDate.Year != 0 {
		a.StartDate = m.StartDate
	}
	if m.EndDate.Year != 0 {
		a.EndDate = m.EndDate
	}
	a.IsAdult = m.IsAdult
	if len(m.Studios.Nodes) > 0 || len(m.Studios.Edges) > 0 {
		a.Studios = m.Studios
	}
	if m.SiteURL != "" {
		a.SiteURL = m.SiteURL
	}
	return a
}

func tracemoeCari(buper []byte) (*hasilAnimeApa, error) {
	b, err := kirimPOSTMentah("https://api.trace.moe/search?anilistInfo", buper, "application/octet-stream")
	if err != nil {
		return nil, err
	}
	var resp struct {
		Error  string `json:"error"`
		Result []struct {
			Anilist    mediaAnilist    `json:"anilist"`
			Episode    json.RawMessage `json:"episode"`
			From       float64         `json:"from"`
			To         float64         `json:"to"`
			Similarity float64         `json:"similarity"`
			Video      string          `json:"video"`
			Image      string          `json:"image"`
		} `json:"result"`
	}
	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}
	if resp.Error != "" {
		return nil, fmt.Errorf("%s", resp.Error)
	}
	if len(resp.Result) == 0 {
		return nil, fmt.Errorf("nggak nemu anime yang cocok sama gambar itu")
	}
	hit := resp.Result[0]
	if hit.Anilist.ID == 0 {
		return nil, fmt.Errorf("nggak nemu anime yang cocok sama gambar itu")
	}
	a := hit.Anilist
	if media, err := anilistMedia(hit.Anilist.ID); err == nil && media != nil {
		a = gabungMediaAnilist(a, *media)
	}
	var namaStudio []string
	if len(a.Studios.Nodes) > 0 {
		for _, n := range a.Studios.Nodes {
			if n.Name != "" {
				namaStudio = append(namaStudio, n.Name)
			}
		}
	} else {
		for _, e := range a.Studios.Edges {
			if e.IsMain && e.Node.Name != "" {
				namaStudio = append(namaStudio, e.Node.Name)
			}
		}
	}
	judul := a.Title.Romaji
	if judul == "" {
		judul = a.Title.English
	}
	if judul == "" {
		judul = a.Title.Native
	}
	judulEn := ""
	if a.Title.English != "" && a.Title.English != a.Title.Romaji {
		judulEn = a.Title.English
	}
	sumber := sumberIDAnime[a.Source]
	if sumber == "" {
		sumber = a.Source
	}
	status := statusIDAnime[a.Status]
	if status == "" {
		status = a.Status
	}
	var rentang []string
	if t := tanggalAnilistStr(a.StartDate); t != "" {
		rentang = append(rentang, t)
	}
	if t := tanggalAnilistStr(a.EndDate); t != "" {
		rentang = append(rentang, t)
	}
	durasi := ""
	if a.Duration != 0 {
		durasi = fmt.Sprintf("%d menit/eps", a.Duration)
	}
	skor := ""
	if a.AverageScore != 0 {
		skor = fmt.Sprintf("%.1f/10", float64(a.AverageScore)/10)
	}
	return &hasilAnimeApa{
		Judul:    judul,
		JudulEn:  judulEn,
		Genre:    strings.Join(a.Genres, ", "),
		Sumber:   sumber,
		Tayang:   strings.Join(rentang, " – "),
		Status:   status,
		Durasi:   durasi,
		Tipe:     a.Format,
		Studio:   strings.Join(namaStudio, ", "),
		Skor:     skor,
		Sinopsis: bersihHTMLAnime(a.Description),
		TotalEps: a.Episodes,
		Dewasa:   a.IsAdult,
		Tautan:   a.SiteURL,
		Episode:  hit.Episode,
		Dari:     hit.From,
		Sampai:   hit.To,
		Mirip:    hit.Similarity,
		Video:    hit.Video,
		Gambar:   hit.Image,
	}, nil
}

func jamDariDetik(s float64) string {
	t := int64(s + 0.5)
	if t < 0 {
		t = 0
	}
	return fmt.Sprintf("%02d:%02d", t/60, t%60)
}

func keIDAnime(teks string) (string, bool) {
	hasil, err := lib.Gtr("id", teks)
	if err != nil {
		return teks, false
	}
	return hasil, true
}

func formatSatuEpisode(raw json.RawMessage) string {
	var angka float64
	if err := json.Unmarshal(raw, &angka); err == nil {
		return strconv.FormatFloat(angka, 'f', -1, 64)
	}
	var teks string
	if err := json.Unmarshal(raw, &teks); err == nil {
		return teks
	}
	return ""
}

func nomorEpsAnime(an *hasilAnimeApa) string {
	e := ""
	if len(an.Episode) > 0 && string(an.Episode) != "null" {
		var daftar []json.RawMessage
		if err := json.Unmarshal(an.Episode, &daftar); err == nil {
			var bagian []string
			for _, d := range daftar {
				bagian = append(bagian, formatSatuEpisode(d))
			}
			e = strings.Join(bagian, "-")
		} else {
			e = formatSatuEpisode(an.Episode)
		}
	}
	if e == "" {
		if an.TotalEps != 0 {
			return fmt.Sprintf("? dari %d", an.TotalEps)
		}
		return "?"
	}
	if an.TotalEps != 0 {
		return fmt.Sprintf("%s dari %d", e, an.TotalEps)
	}
	return e
}

func any2mp4Lokal(sumber string) ([]byte, error) {
	data, err := lib.AmbilURL(sumber, nil)
	if err != nil {
		return nil, err
	}
	awal, err := lib.TulisTemp(".bin", data)
	if err != nil {
		return nil, err
	}
	defer os.Remove(awal)
	akhir := awal + ".mp4"
	defer os.Remove(akhir)
	cmd := exec.Command("ffmpeg", "-i", awal, akhir)
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	return os.ReadFile(akhir)
}

func TanganiAnimeapa(k *lib.Konteks) error {
	m := k.M
	down, _, err := k.Conn.DownloadAny(m)
	if err != nil || len(down) == 0 {
		down = nil
		if m.Quoted != nil {
			if d2, _, err2 := k.Conn.DownloadAny(m.Quoted); err2 == nil {
				down = d2
			}
		}
	}
	if len(down) == 0 {
		return k.Conn.OK(m, "Gak gitu caranya")
	}
	an, err := tracemoeCari(down)
	if err != nil {
		return k.Conn.OK(m, "Gagal nyari anime-nya: "+err.Error())
	}
	baris := []string{
		"*CUPLIKAN*",
		"Episode: " + nomorEpsAnime(an),
		"Waktu: " + jamDariDetik(an.Dari) + " - " + jamDariDetik(an.Sampai),
		fmt.Sprintf("Kemiripan: %.1f%%", an.Mirip*100),
	}
	if an.Mirip < miripMinim {
		baris = append(baris, "⚠️ Kemiripannya rendah, bisa jadi salah anime")
	}
	baris = append(baris, "", "*TENTANG ANIME*", "✨️ *Judul:* "+an.Judul)
	tambah := [][2]string{
		{"🏷️ *Judul EN:*", an.JudulEn},
		{"🎗️ *Genre:*", an.Genre},
		{"🧬 *Diadaptasi:*", an.Sumber},
		{"⏰ *Tayang:*", an.Tayang},
		{"📶 *Status:*", an.Status},
		{"⏳ *Durasi:*", an.Durasi},
		{"🎞 *Tipe:*", an.Tipe},
		{"📽 *Studio:*", an.Studio},
		{"💌️ *Skor:*", an.Skor},
		{"🔗 *AniList:*", an.Tautan},
	}
	for _, t := range tambah {
		if t[1] != "" {
			baris = append(baris, t[0]+" "+t[1])
		}
	}
	if an.Dewasa {
		baris = append(baris, "🔞 *Konten dewasa*")
	}
	if an.Sinopsis != "" {
		sin, jadi := keIDAnime(an.Sinopsis)
		label := "_(bhs. Inggris, terjemahan lagi error)_"
		if jadi {
			label = "_(diterjemahkan dari bhs. Inggris)_"
		}
		baris = append(baris, "", "*• Sinopsis:*", label, sin)
	}
	caption := strings.Join(baris, "\n")
	if v, err := any2mp4Lokal(an.Video); err == nil {
		if err := k.Conn.SendMessage(m.From, lib.IsiPesan{Video: v, Caption: caption}, m); err == nil {
			return nil
		}
	}
	if an.Gambar != "" {
		if gbr, err := lib.AmbilURL(an.Gambar, nil); err == nil {
			if err := k.Conn.SendMessage(m.From, lib.IsiPesan{Image: gbr, Caption: caption}, m); err == nil {
				return nil
			}
		}
	}
	return k.Conn.SendMessage(m.From, lib.IsiPesan{Text: caption}, m)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "animeapa.js",
		Jenis:    "info",
		Nama:     []string{"#animeapa"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(wait|animeapa)$"),
		Jalankan: TanganiAnimeapa,
	})
}
