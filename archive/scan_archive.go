// -----------------------------------------------------------------------------
// go-backup/archive/scan_archive.go

package archive

import (
	"github.com/balacode/go-backup/security"
	"github.com/balacode/go-backup/storage"
)

// scanArchive used is by ExtractArchive(), ListArchive() and VerifyArchive().
// It iterates through every file in the archive and calls 'actor' with
// the file information (and optionally the file contents).
func scanArchive(
	archiveFile string,
	enc *security.Encryption,
	loadContent bool,
	actor func(pos int64, fl *storage.File) error,
) error {
	panic("NOT IMPLEMENTED: scanArchive") // TODO
}

// end
