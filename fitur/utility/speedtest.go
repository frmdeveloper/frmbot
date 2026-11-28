package utility

import (
	"fmt"
	"regexp"

	"gow1/lib"

	"github.com/showwin/speedtest-go/speedtest"
)

func TanganiSpeedtest(k *lib.Konteks) error {
	m := k.M
	_ = k.Conn.SendMessage(m.From, lib.IsiPesan{Text: "⏳ *Speedtest* dimulai, tunggu sebentar..."}, m)
	klien := speedtest.New()
	info, _ := klien.FetchUserInfo()
	serverList, err := klien.FetchServers()
	if err != nil {
		return k.Conn.OK(m, fmt.Sprintf("❌ Speedtest gagal:\n%v", err))
	}
	targets, err := serverList.FindServer([]int{})
	if err != nil || len(targets) == 0 {
		if err == nil {
			err = fmt.Errorf("server speedtest gak ketemu")
		}
		return k.Conn.OK(m, fmt.Sprintf("❌ Speedtest gagal:\n%v", err))
	}
	server := targets[0]
	if err := server.PingTest(nil); err != nil {
		return k.Conn.OK(m, fmt.Sprintf("❌ Speedtest gagal:\n%v", err))
	}
	if err := server.DownloadTest(); err != nil {
		return k.Conn.OK(m, fmt.Sprintf("❌ Speedtest gagal:\n%v", err))
	}
	if err := server.UploadTest(); err != nil {
		return k.Conn.OK(m, fmt.Sprintf("❌ Speedtest gagal:\n%v", err))
	}
	isp := "?"
	namaServer := "?"
	if info != nil && info.Isp != "" {
		isp = info.Isp
	}
	if server.Name != "" {
		namaServer = fmt.Sprintf("%s (%s)", server.Name, server.Country)
	}
	teks := fmt.Sprintf("📶 *SPEEDTEST RESULT*\n\n⬇️ Download : *%.2f Mbps*\n⬆️ Upload   : *%.2f Mbps*\n📡 Ping     : *%d ms* (jitter %d ms)\n🏢 Server   : %s\n🛰 ISP      : %s",
		server.DLSpeed.Mbps(), server.ULSpeed.Mbps(), server.Latency.Milliseconds(), server.Jitter.Milliseconds(), namaServer, isp)
	return k.Conn.OK(m, teks)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "speedtest.js",
		Jenis:    "utility",
		Nama:     nil,
		Itunya:   regexp.MustCompile("(?i)^(#|!|/|\\.)(speedtest|teskecepatan|speed)$"),
		Jalankan: TanganiSpeedtest,
	})
}
