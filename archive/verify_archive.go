// -----------------------------------------------------------------------------
// go-backup/archive/verify_archive.go

package archive

import (
	"github.com/balacode/go-backup/security"
)

// VerifyArchive verifies the integrity of every file in the specified
// archive file by reading each file and checking its SHA-512 hash.
// It uses 'enc' to for decryption.
func VerifyArchive(archiveFile string, enc *security.Encryption) error {
	panic("NOT IMPLEMENTED: VerifyArchive") // TODO
}

// end
