// -----------------------------------------------------------------------------
// go-backup/storage/exists_path.go

package storage

import (
	"errors"
	"os"
)

// ExistsPath returns true if a path exists, or false otherwise.
func ExistsPath(path string) bool {
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

// end
