// -----------------------------------------------------------------------------
// go-backup/storage/exists_dir.go

package storage

import (
	"errors"
	"os"
)

// ExistsDir returns true if a directory exists, or false otherwise.
func ExistsDir(path string) bool {
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) || !info.IsDir() {
		return false
	}
	return true
}

// end
