// -----------------------------------------------------------------------------
// go-backup/archive/extract_archive.go

package archive

import (
	"github.com/balacode/go-backup/security"
)

// ExtractArchive extracts all files in the specified 'archiveFile' to
// the folder specified by 'extractToPath', using 'enc' for decryption.
func ExtractArchive(
	archiveFile string,
	extractToPath string,
	enc *security.Encryption,
) error {
	panic("NOT IMPLEMENTED: ExtractArchive") // TODO
}

// end
