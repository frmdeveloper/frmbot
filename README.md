# FRM Bot — Bot WhatsApp (Go)

**Repo:** https://github.com/frmdeveloper/frmbot

Bot WhatsApp multi-fitur yang ditulis dengan **Go** memakai transport
**whatsmeow upstream** (`go.mau.fi/whatsmeow`). 59 fitur di 10 folder:
audio, chat, converter, downloader, game, group, info, maker, system,
utility.

Di grup, perintah mati secara default — admin/owner harus ketik `#bot on`
dulu di grup itu.

## Tools yang dibutuhkan

### Wajib

| Tool | Dipakai oleh | Cek |
|---|---|---|
| **Go ≥ 1.26** | build bot | `go version` |
| **GCC / build-essential** | CGO untuk driver SQLite (`mattn/go-sqlite3`) | `gcc --version` |
| **ffmpeg** | `#ytmp3`, `#tomp3`, `#toimg`, `#tovid`, stiker video, efek audio | `ffmpeg -version` |
| **ImageMagick** (`convert`) | `#brat`, `#smeme`, `#nulis`, `#quote` | `convert --version` |
| **yt-dlp** (di PATH sistem) | `#ytdl`, `#ytmp3`, `#ytplay`, fallback video `#pin` | `yt-dlp --version` |
| **Node.js** | JS runtime untuk yt-dlp (`--js-runtimes node`) | `node --version` |
| **Python 3** | runtime yt-dlp (rilis resmi = zipapp Python) | `python3 --version` |
| **zbar-tools** (`zbarimg`) | `#readqr` (baca QR dari gambar) | `zbarimg --version` |

### Opsional

| Tool / layanan | Dipakai oleh |
|---|---|
| **pm2** | menjalankan bot sebagai daemon (cara produksi) |
| **WARP mode proxy** (SOCKS5 `127.0.0.1:40000`) | egress YouTube & Pinterest — IP server/VPS biasanya kena bot-wall YouTube; unduhan `#ytdl`/`#ytplay` otomatis lewat WARP dulu, direct jadi fallback terakhir |
| **rmbg-ai lokal** (HTTP `127.0.0.1:8773`) | layanan hapus background untuk `#rmbg` (rembg/BiRefNet jalan terpisah) |
| **cookie YouTube** (`bantuan/cookie-yt.txt`) | sesi login YouTube untuk `#ytdl`/`#ytplay`; tanpa ini banyak video kena bot-wall |

## Instalasi (Ubuntu/Debian)

### 1. Install tools sistem

```sh
sudo apt update
sudo apt install -y build-essential git ffmpeg imagemagick zbar-tools \
  python3 nodejs curl
```

Go (pakai versi terbaru dari situs resmi bila distro ketinggalan):

```sh
curl -fsSL https://go.dev/dl/go1.26.0.linux-amd64.tar.gz -o /tmp/go.tgz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf /tmp/go.tgz
export PATH=$PATH:/usr/local/go/bin
```

yt-dlp (rilis resmi, ditaruh di PATH sistem — bot memanggil `yt-dlp`,
**bukan** binary bundel):

```sh
curl -fsSL https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp -o /tmp/yt-dlp
sudo install -m 755 /tmp/yt-dlp /usr/local/bin/yt-dlp
yt-dlp --version
```

pm2 (opsional, untuk produksi):

```sh
npm install -g pm2
```

### 2. Ambil source & siapkan konfigurasi

```sh
git clone https://github.com/frmdeveloper/frmbot.git gow1
cd gow1
cp .env.example .env
```

Edit `.env`, isi nomor WhatsApp yang akan jadi bot (format kode negara,
tanpa `+`):

```sh
NOMOR=628xxxxxxxxxx
```

Opsional: taruh cookie YouTube hasil export browser di
`bantuan/cookie-yt.txt` (format Netscape, sama seperti untuk yt-dlp biasa).

### 3. Build

```sh
go build -o gow1bin .
```

### 4. Jalankan & pairing

```sh
./gow1bin
```

Kalau belum ada sesi, bot mencetak **kode pairing**. Di HP nomor tersebut:
WhatsApp → Perangkat tertaut → Tautkan perangkat → *Tautkan dengan nomor
telepon* → masukkan kode. Sesi tersimpan di `data/auth.db`; jalan berikutnya
langsung menyambung tanpa pairing.

Jalankan sebagai daemon dengan pm2:

```sh
TTWS_PORT=4001 pm2 start ./gow1bin --name gow1
pm2 save
```

`TTWS_PORT` = port server WebSocket TikTok internal (default 4000; ganti
kalau bentrok dengan proses lain).

Subcommand lain:

```sh
./gow1bin logout   # putuskan sesi / ganti nomor
```

## Konfigurasi (environment / .env)

| Variabel | Wajib | Keterangan |
|---|---|---|
| `NOMOR` | ya | nomor bot untuk pairing |
| `TTWS_PORT` | tidak | port server TikTok WS (default 4000) |
| `TTWS_HOST` | tidak | host bind server TikTok WS |
| `IG_SESSION_ID`, `IG_DS_USER_ID` | tidak | sesi Instagram untuk jalur unduhan IG tertentu |

Fitur eval (`.e`, `.ee`, `.ex`) hanya untuk owner — nomor owner diatur di
kode (`fitur/utility/eval.go` dan pengecekan owner lain di `conn.go`).

## Struktur

```text
main.go            muat .env + data, muat fitur, start
conn.go            koneksi whatsmeow, gate (ban, banmember, antilink), routing
lib/               pustaka inti: tipe Pesan/Conn, registry fitur, data JSON
fitur/<folder>/    59 fitur, satu berkas per fitur (init() mendaftarkan diri)
data/              auth.db (sesi), group.json, user.json, data runtime lain
bantuan/           font, cookie-yt.txt, aset stiker/audio, sebelumkiri.jpg
sampah/            folder unduhan sementara
```

## Keamanan

Jangan pernah commit / bagikan:

- `data/auth.db` — kunci sesi WhatsApp (setara login penuh nomor bot)
- `bantuan/cookie-yt.txt` — cookie sesi Google/YouTube
- `.env` — konfigurasi & nomor

`.gitignore` bawaan sudah menutup ketiganya plus data runtime
(`data/*.json`), binary, dan log. Cek lagi dengan `git status` sebelum push.

## Troubleshooting

- **Pairing diminta terus / sesi rusak** — hentikan bot, jalankan
  `./gow1bin logout` (atau hapus `data/auth.db*`), lalu jalankan lagi.
- **`#ytdl` error "Sign in to confirm you're not a bot"** — IP kena
  bot-wall YouTube: pastikan WARP proxy hidup di `127.0.0.1:40000` dan
  `bantuan/cookie-yt.txt` masih valid (cookie browser rutin terotasi —
  export ulang).
- **`#readqr` gagal** — pastikan `zbarimg` terinstall (`zbar-tools`).
- **Perintah di grup tidak merespons** — grup itu belum diaktifkan:
  ketik `#bot on` di grup.
