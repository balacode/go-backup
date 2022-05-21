// -----------------------------------------------------------------------------
// go-backup/storage/exists_file.go

package storage

import (
	"errors"
	"os"
)

// ExistsFile returns true if a file exists, or false otherwise.
func ExistsFile(path string) bool {
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) || info.IsDir() {
		return false
	}
	return true
}

// end
