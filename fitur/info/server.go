package info

import (
	"encoding/json"
	"fmt"
	"math"
	"net"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"

	"gow1/lib"

	"go.mau.fi/whatsmeow/types/events"
)

var waktuMulaiInfo = time.Now()

var cacheIPInfo struct {
	waktu time.Time
	data  *infoIPServer
}

func kyun(s float64) string {
	detik := int64(math.Floor(s)) % 60
	menit := int64(math.Floor(s/60)) % 60
	jam := int64(math.Floor(s/(60*60))) % 24
	hari := int64(math.Floor(s / (60 * 60 * 24)))
	hitungan := ""
	if hari > 0 {
		hitungan += fmt.Sprintf("%d hari ", hari)
	}
	if jam > 0 {
		hitungan += fmt.Sprintf("%d jam ", jam)
	}
	if menit > 0 {
		hitungan += fmt.Sprintf("%d menit ", menit)
	}
	if detik > 0 {
		hitungan += fmt.Sprintf("%d detik", detik)
	}
	return hitungan
}

func kyunAtauNol(s float64) string {
	if h := kyun(s); h != "" {
		return h
	}
	return "0 detik"
}

func emojiFlag(kode string) string {
	bagian := strings.Split(kode, "-")
	negara := strings.ToUpper(bagian[len(bagian)-1])
	var sb strings.Builder
	for _, r := range negara {
		sb.WriteRune(unicode.ToLower(r) + 127365)
	}
	return sb.String()
}

func formatBytesJS(b int64) string {
	if b == 0 {
		return "0 B"
	}
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	i := int(math.Floor(math.Log(float64(b)) / math.Log(1024)))
	if i < 0 {
		i = 0
	}
	if i >= len(sizes) {
		i = len(sizes) - 1
	}
	v := float64(b) / math.Pow(1024, float64(i))
	teks := strconv.FormatFloat(v, 'f', 2, 64)
	teks = strings.TrimRight(teks, "0")
	teks = strings.TrimRight(teks, ".")
	return teks + " " + sizes[i]
}

func arsitekturJS() string {
	switch runtime.GOARCH {
	case "amd64":
		return "x64"
	case "386":
		return "ia32"
	default:
		return runtime.GOARCH
	}
}

func cpuNamaServer() string {
	if b, err := os.ReadFile("/proc/cpuinfo"); err == nil {
		for _, baris := range strings.Split(string(b), "\n") {
			if strings.HasPrefix(baris, "model name") {
				pecah := strings.SplitN(baris, ":", 2)
				if len(pecah) == 2 && strings.TrimSpace(pecah[1]) != "" {
					return strings.TrimSpace(pecah[1])
				}
			}
		}
	}
	out, _ := lib.SpawnSync("getprop ro.product.model")
	if s := strings.TrimSpace(out); s != "" {
		return s
	}
	return arsitekturJS()
}

func lokasiServer() string {
	if run := os.Getenv("GITHUB_RUN_ID"); run != "" {
		return "GitHub Actions (run #" + run + ")"
	}
	if strings.Contains(os.Getenv("PREFIX"), "com.termux") {
		return "Termux (HP)"
	}
	return runtime.GOOS
}

func sistemServer() string {
	tipe, _ := lib.SpawnSync("uname -s")
	rilisan, _ := lib.SpawnSync("uname -r")
	tipe = strings.TrimSpace(tipe)
	rilisan = strings.TrimSpace(rilisan)
	if tipe == "" {
		tipe = runtime.GOOS
	}
	return fmt.Sprintf("%s %s (%s)", tipe, rilisan, arsitekturJS())
}



func ramServer() (total, sisa int64) {
	b, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return 0, 0
	}
	for _, baris := range strings.Split(string(b), "\n") {
		pecah := strings.Fields(baris)
		if len(pecah) < 2 {
			continue
		}
		nilai, _ := strconv.ParseInt(pecah[1], 10, 64)
		switch pecah[0] {
		case "MemTotal:":
			total = nilai * 1024
		case "MemFree:":
			sisa = nilai * 1024
		}
	}
	return total, sisa
}

func bebanServer() string {
	b, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		return "0.00 / 0.00 / 0.00"
	}
	pecah := strings.Fields(string(b))
	if len(pecah) < 3 {
		return "0.00 / 0.00 / 0.00"
	}
	var bagian []string
	for _, p := range pecah[:3] {
		f, _ := strconv.ParseFloat(p, 64)
		bagian = append(bagian, strconv.FormatFloat(f, 'f', 2, 64))
	}
	return strings.Join(bagian, " / ")
}

func nyalaServer() float64 {
	b, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return 0
	}
	pecah := strings.Fields(string(b))
	if len(pecah) == 0 {
		return 0
	}
	f, _ := strconv.ParseFloat(pecah[0], 64)
	return f
}

func pingTCPServer(host string) *int64 {
	mulai := time.Now()
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, "443"), 5*time.Second)
	if err != nil {
		return nil
	}
	conn.Close()
	ms := time.Since(mulai).Milliseconds()
	return &ms
}

type infoIPServer struct {
	IP     string
	Kota   string
	Negara string
	Kode   string
	ISP    string
}

func ambilInfoIPServer() *infoIPServer {
	if cacheIPInfo.data != nil && time.Since(cacheIPInfo.waktu) < 5*time.Minute {
		return cacheIPInfo.data
	}
	var data *infoIPServer
	var who struct {
		Success     bool   `json:"success"`
		IP          string `json:"ip"`
		City        string `json:"city"`
		Region      string `json:"region"`
		Country     string `json:"country"`
		CountryCode string `json:"country_code"`
		Connection  struct {
			ISP string `json:"isp"`
			Org string `json:"org"`
		} `json:"connection"`
	}
	if err := lib.AmbilJSON("https://ipwho.is/", nil, &who); err == nil && who.Success {
		isp := who.Connection.ISP
		if isp == "" {
			isp = who.Connection.Org
		}
		data = &infoIPServer{
			IP:     who.IP,
			Kota:   gabungNonKosong(", ", who.City, who.Region),
			Negara: who.Country,
			Kode:   who.CountryCode,
			ISP:    isp,
		}
	}
	if data == nil {
		var api struct {
			Status      string `json:"status"`
			Query       string `json:"query"`
			City        string `json:"city"`
			RegionName  string `json:"regionName"`
			Country     string `json:"country"`
			CountryCode string `json:"countryCode"`
			ISP         string `json:"isp"`
		}
		if err := lib.AmbilJSON("http://ip-api.com/json", nil, &api); err == nil && api.Status == "success" {
			data = &infoIPServer{
				IP:     api.Query,
				Kota:   gabungNonKosong(", ", api.City, api.RegionName),
				Negara: api.Country,
				Kode:   api.CountryCode,
				ISP:    api.ISP,
			}
		}
	}
	if data != nil {
		cacheIPInfo.waktu = time.Now()
		cacheIPInfo.data = data
	}
	return data
}

func gabungNonKosong(sep string, isi ...string) string {
	var bagian []string
	for _, s := range isi {
		if s != "" {
			bagian = append(bagian, s)
		}
	}
	return strings.Join(bagian, sep)
}

func barisIPServer(ip *infoIPServer) string {
	if ip == nil {
		return "gagal ambil"
	}
	bendera := ""
	if ip.Kode != "" {
		bendera = emojiFlag(ip.Kode) + " "
	}
	tempat := gabungNonKosong(" - ", ip.Kota, ip.Negara)
	isp := ip.ISP
	if isp == "" {
		isp = "-"
	}
	return fmt.Sprintf("%s\nLokasi IP: %s%s\nISP      : %s", ip.IP, bendera, tempat, isp)
}

func sehatProxyServer() string {
	base := os.Getenv("AI_BASE_URL")
	if base == "" {
		base = "http://127.0.0.1:20128/v1"
	}
	mulai := time.Now()
	klien := &http.Client{Timeout: 3 * time.Second}
	resp, err := klien.Get(base + "/models")
	if err != nil {
		sebab := "gagal konek"
		if ne, ok := err.(net.Error); ok && ne.Timeout() {
			sebab = "timeout"
		} else if strings.Contains(err.Error(), "connection refused") {
			sebab = "ECONNREFUSED"
		}
		return "mati (" + sebab + ")"
	}
	defer resp.Body.Close()
	var body struct {
		Data []json.RawMessage `json:"data"`
	}
	_ = json.NewDecoder(resp.Body).Decode(&body)
	hasil := fmt.Sprintf("hidup %d ms", time.Since(mulai).Milliseconds())
	if len(body.Data) > 0 {
		hasil += fmt.Sprintf(" (%d model)", len(body.Data))
	}
	return hasil
}

func umurBerkasServer(path string) (int64, float64, bool) {
	st, err := os.Stat(path)
	if err != nil {
		return 0, 0, false
	}
	return st.Size(), time.Since(st.ModTime()).Seconds(), true
}

func barisSesiServer() string {
	potong := func(path string) string {
		ukuran, lalu, ok := umurBerkasServer(path)
		if !ok {
			return "gak ada"
		}
		return fmt.Sprintf("%s (%s lalu)", formatBytesJS(ukuran), kyunAtauNol(lalu))
	}
	return fmt.Sprintf("%s\nSnapshot : %s", potong("data/.auth.db"), potong("data/auth.db"))
}

func versiWhatsmeowServer() string {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return "?"
	}
	for _, dep := range bi.Deps {
		if dep.Path == "go.mau.fi/whatsmeow" {
			return dep.Version
		}
	}
	return "?"
}

func infoGrupServer(k *lib.Konteks) string {
	grup, err := k.Conn.GetJoinedGroups()
	if err != nil {
		return "gagal (" + err.Error() + ")"
	}
	member := 0
	for _, g := range grup {
		member += len(g.Peserta)
	}
	return fmt.Sprintf("%d grup, %d member", len(grup), member)
}

func TanganiServer(k *lib.Konteks) error {
	m := k.M
	ramTotal, ramSisa := ramServer()
	diskTotal, diskKepake, diskSisa := diskServer()
	var rss uint64
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	rss = ms.Sys

	kirim := time.Now()
	if ev, ok := m.Full.(*events.Message); ok && ev != nil && !ev.Info.Timestamp.IsZero() {
		kirim = ev.Info.Timestamp
	}
	jeda := time.Since(kirim).Milliseconds()
	if jeda < 0 {
		jeda = 0
	}

	var pingNet, pingWa *int64
	var ip *infoIPServer
	var proxy, grup string
	var wg sync.WaitGroup
	wg.Add(5)
	go func() { defer wg.Done(); pingNet = pingTCPServer("1.1.1.1") }()
	go func() { defer wg.Done(); pingWa = pingTCPServer("g.whatsapp.net") }()
	go func() { defer wg.Done(); ip = ambilInfoIPServer() }()
	go func() { defer wg.Done(); proxy = sehatProxyServer() }()
	go func() { defer wg.Done(); grup = infoGrupServer(k) }()
	wg.Wait()

	jaringan := "nyambung (mulai kapan: -)"
	if !k.Conn.Terhubung() {
		jaringan = "putus"
	}
	akun := k.Conn.BotNomor()
	if akun == "" {
		akun = "-"
	}
	fiturTeks := fmt.Sprintf("%d kepasang", lib.Jumlah())
	if len(lib.ErorMuat) > 0 {
		fiturTeks += fmt.Sprintf(", %d error", len(lib.ErorMuat))
	}
	pingNetTeks := "gagal"
	if pingNet != nil {
		pingNetTeks = fmt.Sprintf("%d ms", *pingNet)
	}
	pingWaTeks := "gagal"
	if pingWa != nil {
		pingWaTeks = fmt.Sprintf("%d ms", *pingWa)
	}
	jam := time.Now()
	if lok, err := time.LoadLocation("Asia/Jakarta"); err == nil {
		jam = jam.In(lok)
	}
	host, _ := os.Hostname()
	teks := fmt.Sprintf(`INFORMASI SERVER

Lokasi   : %s
Host     : %s
Sistem   : %s
CPU      : %s x%d
Beban    : %s
RAM      : %s / %s kepake
Disk     : %s / %s kepake (sisa %s)
Nyala    : %s

IP publik: %s
Jaringan : %s
Ping     : internet %s | WA %s
Proxy AI : %s

Go       : %s
Whatsmeow: %s
Bot hidup: %s
Memori   : %s
Fitur    : %s
Akun     : %s
Grup     : %s
Pesan    : %s
Sesi     : %s
Respon   : %d ms
Jam      : %s WIB`,
		lokasiServer(),
		host,
		sistemServer(),
		cpuNamaServer(), runtime.NumCPU(),
		bebanServer(),
		formatBytesJS(ramTotal-ramSisa), formatBytesJS(ramTotal),
		formatBytesJS(diskKepake), formatBytesJS(diskTotal), formatBytesJS(diskSisa),
		kyunAtauNol(nyalaServer()),
		barisIPServer(ip),
		jaringan,
		pingNetTeks, pingWaTeks,
		proxy,
		runtime.Version(),
		versiWhatsmeowServer(),
		kyunAtauNol(time.Since(waktuMulaiInfo).Seconds()),
		formatBytesJS(int64(rss)),
		fiturTeks,
		akun,
		grup,
		"-",
		barisSesiServer(),
		jeda,
		jam.Format("2/1/2006, 15.04.05"),
	)
	return k.Conn.OK(m, teks)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "server.js",
		Jenis:    "info",
		Nama:     []string{"#server"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(server|sinfo|status)$"),
		Jalankan: TanganiServer,
	})
}
