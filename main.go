package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"gow1/lib"

	_ "gow1/fitur/ai"
	_ "gow1/fitur/audio"
	_ "gow1/fitur/chat"
	_ "gow1/fitur/converter"
	_ "gow1/fitur/downloader"
	_ "gow1/fitur/game"
	_ "gow1/fitur/group"
	_ "gow1/fitur/info"
	_ "gow1/fitur/maker"
	_ "gow1/fitur/system"
	_ "gow1/fitur/utility"
)

var Owner []string

func muatEnv() {
	f, err := os.Open(".env")
	if err != nil {
		fmt.Println("Tidak ada .env, pakai default (lihat .env.example)")
		return
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		baris := strings.TrimSpace(sc.Text())
		if baris == "" || strings.HasPrefix(baris, "#") {
			continue
		}
		k, v, ok := strings.Cut(baris, "=")
		if ok {
			os.Setenv(strings.TrimSpace(k), strings.TrimSpace(v))
		}
	}
}

var Gtr = func(bahasa, teks string) (string, error) {
	return "", fmt.Errorf("translate belum diporting")
}

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "logout":
			LogoutMain()
			return
		case "gd":
			GdMain(os.Args[2:])
			return
		case "ghact":
			GhactMain(os.Args[2:])
			return
		case "igsesi":
			IgsesiMain(os.Args[2:])
			return
		}
	}

	muatEnv()
	Owner = []string{
		lib.NomorOwner() + "@s.whatsapp.net",
		"6283837747136@s.whatsapp.net",
	}

	user := lib.MuatUser("data/user.json")
	group := lib.MuatGroup("data/group.json")

	fmt.Printf("[fitur] %d fitur termuat\n", lib.Jumlah())
	for _, e := range lib.ErorMuat {
		fmt.Println("[skip]", e)
	}

	nomor := os.Getenv("NOMOR")
	if nomor == "" {
		nomor = "6283169480682"
	}
	bot := BotBaru(user, group, nomor)
	if err := bot.Start(); err != nil {
		fmt.Println("[conn] gagal start:", err)
		os.Exit(1)
	}

	tunggu := make(chan os.Signal, 1)
	signal.Notify(tunggu, syscall.SIGINT, syscall.SIGTERM)
	<-tunggu
	fmt.Println("[conn] berhenti")
	bot.Stop()
}
