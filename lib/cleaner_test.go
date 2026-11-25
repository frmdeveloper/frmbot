package lib

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCleaner(t *testing.T) {
	testDir := filepath.Join("sampah", "test_cleaner")
	_ = os.MkdirAll(testDir, 0o755)
	defer os.RemoveAll(testDir)

	oldFile := filepath.Join(testDir, "old.txt")
	_ = os.WriteFile(oldFile, []byte("old content"), 0o644)

	// Ubah modtime file menjadi 1 jam lalu
	satuJamLalu := time.Now().Add(-1 * time.Hour)
	_ = os.Chtimes(oldFile, satuJamLalu, satuJamLalu)

	newFile := filepath.Join(testDir, "new.txt")
	_ = os.WriteFile(newFile, []byte("new content"), 0o644)

	terhapus := BersihkanDirektori(testDir, 10*time.Minute)
	if terhapus != 1 {
		t.Errorf("Harusnya 1 file terhapus, didapat: %d", terhapus)
	}

	if _, err := os.Stat(oldFile); !os.IsNotExist(err) {
		t.Errorf("File lama harusnya sudah terhapus")
	}

	if _, err := os.Stat(newFile); os.IsNotExist(err) {
		t.Errorf("File baru tidak boleh terhapus")
	}
}
