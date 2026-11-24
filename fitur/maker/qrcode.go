package maker

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"regexp"

	"gow1/lib"
)

var gfExpQR [512]byte
var gfLogQR [256]byte

func init() {
	x := 1
	for i := 0; i < 255; i++ {
		gfExpQR[i] = byte(x)
		gfLogQR[byte(x)] = byte(i)
		x <<= 1
		if x >= 256 {
			x ^= 0x11D
		}
	}
	for i := 255; i < 512; i++ {
		gfExpQR[i] = gfExpQR[i-255]
	}
}

func gfMulQR(a byte, b byte) byte {
	if a == 0 || b == 0 {
		return 0
	}
	return gfExpQR[int(gfLogQR[a])+int(gfLogQR[b])]
}

func rsGenQR(deg int) []byte {
	poly := []byte{1}
	for i := 0; i < deg; i++ {
		baru := make([]byte, len(poly)+1)
		for a := range poly {
			for b, q := range []byte{1, gfExpQR[i]} {
				baru[a+b] ^= gfMulQR(poly[a], q)
			}
		}
		poly = baru
	}
	return poly
}

func rsEncodeQR(data []byte, deg int) []byte {
	gen := rsGenQR(deg)
	msg := make([]byte, len(data)+deg)
	copy(msg, data)
	for i := 0; i < len(data); i++ {
		coef := msg[i]
		if coef == 0 {
			continue
		}
		for j := range gen {
			msg[i+j] ^= gfMulQR(gen[j], coef)
		}
	}
	return msg[len(data):]
}

type infoBlokQR struct {
	ecPerBlok int
	grup      [][2]int
}

var tabelBlokH = map[int]infoBlokQR{
	1:  {17, [][2]int{{1, 9}}},
	2:  {28, [][2]int{{1, 16}}},
	3:  {22, [][2]int{{2, 13}}},
	4:  {16, [][2]int{{4, 9}}},
	5:  {22, [][2]int{{2, 11}, {2, 12}}},
	6:  {28, [][2]int{{4, 15}}},
	7:  {26, [][2]int{{4, 13}, {1, 14}}},
	8:  {26, [][2]int{{4, 14}, {2, 15}}},
	9:  {24, [][2]int{{4, 12}, {4, 13}}},
	10: {28, [][2]int{{6, 15}, {2, 16}}},
}

var sisaBitQR = map[int]int{1: 0, 2: 7, 3: 7, 4: 7, 5: 7, 6: 7, 7: 0, 8: 0, 9: 0, 10: 0}

var tengahAlignmentQR = map[int][]int{
	2:  {6, 18},
	3:  {6, 22},
	4:  {6, 26},
	5:  {6, 30},
	6:  {6, 34},
	7:  {6, 22, 38},
	8:  {6, 24, 42},
	9:  {6, 26, 46},
	10: {6, 28, 50},
}

func totalDataQR(versi int) int {
	total := 0
	for _, g := range tabelBlokH[versi].grup {
		total += g[0] * g[1]
	}
	return total
}

type aliranBit struct {
	data []byte
	n    int
}

func (a *aliranBit) tambah(val int, panjang int) {
	for i := panjang - 1; i >= 0; i-- {
		bit := (val >> i) & 1
		if a.n%8 == 0 {
			a.data = append(a.data, 0)
		}
		if bit == 1 {
			a.data[a.n/8] |= 1 << (7 - a.n%8)
		}
		a.n++
	}
}

func (a *aliranBit) bitKe(i int) bool {
	if i/8 >= len(a.data) {
		return false
	}
	return (a.data[i/8]>>(7-i%8))&1 == 1
}

func bangunDataQR(payload []byte, versi int) []byte {
	total := totalDataQR(versi)
	aliran := &aliranBit{}
	aliran.tambah(0b0100, 4)
	if versi >= 10 {
		aliran.tambah(len(payload), 16)
	} else {
		aliran.tambah(len(payload), 8)
	}
	for _, b := range payload {
		aliran.tambah(int(b), 8)
	}
	kapasitas := total * 8
	sisa := kapasitas - aliran.n
	if sisa > 4 {
		sisa = 4
	}
	aliran.tambah(0, sisa)
	for aliran.n%8 != 0 {
		aliran.tambah(0, 1)
	}
	isiPad := []int{0xEC, 0x11}
	for i := 0; aliran.n < kapasitas; i++ {
		aliran.tambah(isiPad[i%2], 8)
	}
	info := tabelBlokH[versi]
	type blok struct {
		data []byte
		ec   []byte
	}
	blokBlok := []blok{}
	pos := 0
	for _, g := range info.grup {
		for i := 0; i < g[0]; i++ {
			d := aliran.data[pos : pos+g[1]]
			pos += g[1]
			blokBlok = append(blokBlok, blok{d, rsEncodeQR(d, info.ecPerBlok)})
		}
	}
	hasil := []byte{}
	maks := 0
	for _, b := range blokBlok {
		if len(b.data) > maks {
			maks = len(b.data)
		}
	}
	for i := 0; i < maks; i++ {
		for _, b := range blokBlok {
			if i < len(b.data) {
				hasil = append(hasil, b.data[i])
			}
		}
	}
	for i := 0; i < info.ecPerBlok; i++ {
		for _, b := range blokBlok {
			hasil = append(hasil, b.ec[i])
		}
	}
	return hasil
}

type matriksQR struct {
	ukuran int
	gelap  [][]bool
	fungsi [][]bool
}

func matriksBaruQR(versi int) *matriksQR {
	ukuran := versi*4 + 17
	m := &matriksQR{ukuran: ukuran}
	m.gelap = make([][]bool, ukuran)
	m.fungsi = make([][]bool, ukuran)
	for i := range m.gelap {
		m.gelap[i] = make([]bool, ukuran)
		m.fungsi[i] = make([]bool, ukuran)
	}
	return m
}

func (m *matriksQR) setFungsi(x int, y int, gelap bool) {
	if 0 <= x && x < m.ukuran && 0 <= y && y < m.ukuran {
		m.gelap[y][x] = gelap
		m.fungsi[y][x] = true
	}
}

func (m *matriksQR) gambarFungsi(versi int) {
	for i := 0; i < m.ukuran; i++ {
		m.setFungsi(6, i, i%2 == 0)
		m.setFungsi(i, 6, i%2 == 0)
	}
	m.gambarFinder(3, 3)
	m.gambarFinder(m.ukuran-4, 3)
	m.gambarFinder(3, m.ukuran-4)
	if versi >= 2 {
		als := tengahAlignmentQR[versi]
		for _, x := range als {
			for _, y := range als {
				pojok := (x == als[0] && y == als[0]) || (x == als[0] && y == als[len(als)-1]) || (x == als[len(als)-1] && y == als[0])
				if pojok {
					continue
				}
				for dy := -2; dy <= 2; dy++ {
					for dx := -2; dx <= 2; dx++ {
						jarak := maxQR(absQR(dx), absQR(dy))
						m.setFungsi(x+dx, y+dy, jarak != 1)
					}
				}
			}
		}
	}
	for i := 0; i <= 5; i++ {
		m.setFungsi(8, i, true)
	}
	m.setFungsi(8, 7, true)
	m.setFungsi(8, 8, true)
	m.setFungsi(7, 8, true)
	for i := 9; i < 15; i++ {
		m.setFungsi(14-i, 8, true)
	}
	for i := 0; i < 8; i++ {
		m.setFungsi(m.ukuran-1-i, 8, true)
	}
	for i := 8; i < 15; i++ {
		m.setFungsi(8, m.ukuran-15+i, true)
	}
	m.setFungsi(8, m.ukuran-8, true)
	if versi >= 7 {
		for i := 0; i < 18; i++ {
			bit := bitInt(infoVersiQR(versi), i)
			m.setFungsi(m.ukuran-11+i%3, i/3, bit)
			m.setFungsi(i/3, m.ukuran-11+i%3, bit)
		}
	}
}

func (m *matriksQR) gambarFinder(cx int, cy int) {
	for dy := -4; dy <= 4; dy++ {
		for dx := -4; dx <= 4; dx++ {
			jarak := maxQR(absQR(dx), absQR(dy))
			m.setFungsi(cx+dx, cy+dy, jarak != 2 && jarak != 4)
		}
	}
}

func absQR(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func maxQR(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func bitInt(v int, i int) bool {
	return (v>>i)&1 == 1
}

func infoVersiQR(versi int) int {
	bits := versi << 12
	for i := 17; i >= 12; i-- {
		if (bits>>i)&1 == 1 {
			bits ^= 0x1F25 << (i - 12)
		}
	}
	return (versi << 12) | bits
}

func infoFormatQR(mask int) int {
	data := (2 << 3) | mask
	rem := data
	for i := 0; i < 10; i++ {
		rem = (rem << 1) ^ ((rem >> 9) * 0x537)
	}
	return ((data << 10) | rem) ^ 0x5412
}

func (m *matriksQR) gambarFormat(mask int) {
	bits := infoFormatQR(mask)
	for i := 0; i <= 5; i++ {
		m.setFungsi(8, i, bitInt(bits, i))
	}
	m.setFungsi(8, 7, bitInt(bits, 7))
	m.setFungsi(8, 8, bitInt(bits, 8))
	m.setFungsi(7, 8, bitInt(bits, 9))
	for i := 9; i < 15; i++ {
		m.setFungsi(14-i, 8, bitInt(bits, i))
	}
	for i := 0; i < 8; i++ {
		m.setFungsi(m.ukuran-1-i, 8, bitInt(bits, i))
	}
	for i := 8; i < 15; i++ {
		m.setFungsi(8, m.ukuran-15+i, bitInt(bits, i))
	}
	m.setFungsi(8, m.ukuran-8, true)
}

func (m *matriksQR) tempatkanData(data []byte, sisaBit int) {
	totalBit := len(data)*8 + sisaBit
	aliran := &aliranBit{data: data, n: totalBit}
	i := 0
	for kanan := m.ukuran - 1; kanan >= 1; kanan -= 2 {
		if kanan == 6 {
			kanan = 5
		}
		for vert := 0; vert < m.ukuran; vert++ {
			for j := 0; j < 2; j++ {
				x := kanan - j
				naik := ((kanan + 1) & 2) == 0
				y := vert
				if naik {
					y = m.ukuran - 1 - vert
				}
				if !m.fungsi[y][x] {
					if i < totalBit {
						m.gelap[y][x] = aliran.bitKe(i)
					}
					i++
				}
			}
		}
	}
}

func maskCocokQR(mask int, x int, y int) bool {
	switch mask {
	case 0:
		return (y+x)%2 == 0
	case 1:
		return y%2 == 0
	case 2:
		return x%3 == 0
	case 3:
		return (y+x)%3 == 0
	case 4:
		return (y/2+x/3)%2 == 0
	case 5:
		return (y*x)%2+(y*x)%3 == 0
	case 6:
		return ((y*x)%2+(y*x)%3)%2 == 0
	case 7:
		return ((y+x)%2+(y*x)%3)%2 == 0
	}
	return false
}

func (m *matriksQR) terapkanMask(mask int) {
	for y := 0; y < m.ukuran; y++ {
		for x := 0; x < m.ukuran; x++ {
			if !m.fungsi[y][x] && maskCocokQR(mask, x, y) {
				m.gelap[y][x] = !m.gelap[y][x]
			}
		}
	}
}

func (m *matriksQR) penalti() int {
	skor := 0
	barisSkor := func(get func(i int) bool, n int) int {
		s := 0
		lari := 1
		for i := 1; i < n; i++ {
			if get(i) == get(i-1) {
				lari++
			} else {
				if lari >= 5 {
					s += 3 + (lari - 5)
				}
				lari = 1
			}
		}
		if lari >= 5 {
			s += 3 + (lari - 5)
		}
		return s
	}
	for y := 0; y < m.ukuran; y++ {
		yy := y
		skor += barisSkor(func(i int) bool { return m.gelap[yy][i] }, m.ukuran)
	}
	for x := 0; x < m.ukuran; x++ {
		xx := x
		skor += barisSkor(func(i int) bool { return m.gelap[i][xx] }, m.ukuran)
	}
	for y := 0; y+1 < m.ukuran; y++ {
		for x := 0; x+1 < m.ukuran; x++ {
			if m.gelap[y][x] == m.gelap[y][x+1] && m.gelap[y][x] == m.gelap[y+1][x] && m.gelap[y][x] == m.gelap[y+1][x+1] {
				skor += 3
			}
		}
	}
	polaSkor := func(get func(i int) bool, n int) int {
		s := 0
		for i := 0; i+7 <= n; i++ {
			if get(i) && !get(i+1) && get(i+2) && get(i+3) && get(i+4) && !get(i+5) && get(i+6) {
				sebelum := i >= 4
				if sebelum {
					for j := i - 4; j < i; j++ {
						if get(j) {
							sebelum = false
						}
					}
				}
				sesudah := i+11 <= n
				if sesudah {
					for j := i + 7; j < i+11; j++ {
						if get(j) {
							sesudah = false
						}
					}
				}
				if sebelum || sesudah {
					s += 40
				}
			}
		}
		return s
	}
	for y := 0; y < m.ukuran; y++ {
		yy := y
		skor += polaSkor(func(i int) bool { return m.gelap[yy][i] }, m.ukuran)
	}
	for x := 0; x < m.ukuran; x++ {
		xx := x
		skor += polaSkor(func(i int) bool { return m.gelap[i][xx] }, m.ukuran)
	}
	gelap := 0
	for y := 0; y < m.ukuran; y++ {
		for x := 0; x < m.ukuran; x++ {
			if m.gelap[y][x] {
				gelap++
			}
		}
	}
	total := m.ukuran * m.ukuran
	k := (absQR(gelap*20-total*10) + total - 1) / total
	k--
	if k < 0 {
		k = 0
	}
	skor += k * 10
	return skor
}

func salinMatriksQR(m *matriksQR) *matriksQR {
	baru := &matriksQR{ukuran: m.ukuran}
	baru.gelap = make([][]bool, m.ukuran)
	baru.fungsi = m.fungsi
	for i := range m.gelap {
		baru.gelap[i] = append([]bool{}, m.gelap[i]...)
	}
	return baru
}

func buatQR(teks string) ([]byte, error) {
	payload := []byte(teks)
	versi := 0
	for v := 1; v <= 10; v++ {
		bitPanjang := 8
		if v >= 10 {
			bitPanjang = 16
		}
		if 4+bitPanjang+8*len(payload) <= totalDataQR(v)*8 {
			versi = v
			break
		}
	}
	if versi == 0 {
		return nil, fmt.Errorf("teks kepanjangan buat QR (maks 119 byte)")
	}
	data := bangunDataQR(payload, versi)
	dasar := matriksBaruQR(versi)
	dasar.gambarFungsi(versi)
	dasar.tempatkanData(data, sisaBitQR[versi])
	var terbaik *matriksQR
	skorTerbaik := -1
	for mask := 0; mask < 8; mask++ {
		kandidat := salinMatriksQR(dasar)
		kandidat.terapkanMask(mask)
		kandidat.gambarFormat(mask)
		skor := kandidat.penalti()
		if skorTerbaik < 0 || skor < skorTerbaik {
			skorTerbaik = skor
			terbaik = kandidat
		}
	}
	return renderPNGQR(terbaik), nil
}

func renderPNGQR(m *matriksQR) []byte {
	const sisi = 512
	total := m.ukuran + 2
	img := image.NewRGBA(image.Rect(0, 0, sisi, sisi))
	putih := color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
	hitam := color.RGBA{0x00, 0x00, 0x00, 0xFF}
	for y := 0; y < sisi; y++ {
		my := y*total/sisi - 1
		for x := 0; x < sisi; x++ {
			mx := x*total/sisi - 1
			c := putih
			if 0 <= mx && mx < m.ukuran && 0 <= my && my < m.ukuran && m.gelap[my][mx] {
				c = hitam
			}
			img.SetRGBA(x, y, c)
		}
	}
	var buf bytes.Buffer
	_ = png.Encode(&buf, img)
	return buf.Bytes()
}

func TanganiQrcode(k *lib.Konteks) error {
	m := k.M
	if m.Q == "" {
		return k.Conn.OK(m, "Gk gitu")
	}
	qr, err := buatQR(m.Q)
	if err != nil {
		return k.Conn.OK(m, fmt.Sprintf("Error: %s", err.Error()))
	}
	return k.Conn.SendMessage(m.From, lib.IsiPesan{
		Image:   qr,
		Caption: fmt.Sprintf("QR Code untuk:\n%s", m.Q),
	}, m)
}

func init() {
	lib.Daftarkan(&lib.Fitur{
		File:     "qrcode.js",
		Jenis:    "maker",
		Nama:     []string{"#qrcode"},
		Itunya:   regexp.MustCompile("(?i)^(#|\\!|\\/|\\.)(qr(code)?)$"),
		Jalankan: TanganiQrcode,
	})
}
