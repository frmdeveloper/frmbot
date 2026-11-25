//go:build !windows

package info

import "syscall"

func diskServer() (total, kepake, sisa int64) {
	var st syscall.Statfs_t
	if err := syscall.Statfs(".", &st); err != nil {
		return 0, 0, 0
	}
	total = int64(st.Blocks) * int64(st.Bsize)
	sisa = int64(st.Bavail) * int64(st.Bsize)
	return total, total - sisa, sisa
}
