// -----------------------------------------------------------------------------
// go-backup/archive/list_archive.go

package archive

import (
	"github.com/balacode/go-backup/security"
)

// ListArchive lists all files in 'archiveFile' using
// 'enc' to decrypt file metadata (path, size, etc.)
func ListArchive(archiveFile string, enc *security.Encryption) error {
	panic("NOT IMPLEMENTED: ListArchive") // TODO
}

// end
