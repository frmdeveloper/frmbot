//go:build windows

package info

import (
	"golang.org/x/sys/windows"
)

func diskServer() (total, kepake, sisa int64) {
	var freeBytesAvailable, totalNumberOfBytes, totalNumberOfFreeBytes uint64
	path, err := windows.UTF16PtrFromString(".")
	if err != nil {
		return 0, 0, 0
	}
	err = windows.GetDiskFreeSpaceEx(path, &freeBytesAvailable, &totalNumberOfBytes, &totalNumberOfFreeBytes)
	if err != nil {
		return 0, 0, 0
	}
	total = int64(totalNumberOfBytes)
	sisa = int64(freeBytesAvailable)
	return total, total - sisa, sisa
}
