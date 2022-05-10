// -----------------------------------------------------------------------------
// go-backup/archive/create_archive.go

package archive

import (
	"github.com/balacode/go-backup/security"
)

// CreateArchive creates an archive in 'archiveFile', by compressing and
// encrypting all files in 'backupPath' using 'enc' for encryption.
func CreateArchive(
	backupPath string,
	archiveFile string,
	enc *security.Encryption,
) error {
	panic("NOT IMPLEMENTED: CreateArchive") // TODO
}

// end
