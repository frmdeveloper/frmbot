package lib

import (
	"fmt"
	"math/rand"
)

func RandNum(digit int) string {
	hasil := ""
	for i := 0; i < digit; i++ {
		hasil += fmt.Sprint(rand.Intn(10))
	}
	return hasil
}

func FormatBytes(bytes int64, decimals int) string {
	if bytes == 0 {
		return "0 B"
	}
	const k = 1024.0
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	i := 0
	f := float64(bytes)
	for f >= k && i < len(sizes)-1 {
		f /= k
		i++
	}
	return fmt.Sprintf("%.*f %s", decimals, f, sizes[i])
}

func PickRandom[T any](list []T) T {
	return list[rand.Intn(len(list))]
}
