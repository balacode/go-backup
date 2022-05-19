// -----------------------------------------------------------------------------
// go-backup/archive/scan_archive.go

package archive

import (
	"bytes"
	"errors"
	"io"
	"os"

	"github.com/balacode/go-backup/logging"
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
	if err := enc.Validate(); err != nil {
		return logging.Error(0xE8A8B6, err)
	}
	if _, err := os.Stat(archiveFile); errors.Is(err, os.ErrNotExist) {
		msg := "file does not exist: " + archiveFile
		return logging.Error(0xE2F2C5, msg)
	}
	archive, err := os.Open(archiveFile)
	if err != nil {
		return logging.Error(0xE4E32E, err)
	}
	defer archive.Close()
	var (
		buf    = make([]byte, 16*1024*1024) // 16 MiB
		pos    = int64(0)
		lenBOF = len(storage.BOFMark)
	)
	for {
		n, err := archive.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return logging.Error(0xE7BA59, err)
		}
		ar := buf[:n]
		findBOF := func() int { return bytes.Index(ar, storage.BOFMark) }
		for i := findBOF(); i != -1 && len(ar) >= 8; i = findBOF() {
			fl, err := storage.ReadArchivedFile(
				archive, pos+int64(i), loadContent, enc,
			)
			if err != nil {
				return logging.Error(0xE8A6E8, err)
			}
			pos += int64(i)
			actor(pos, fl)
			pos += int64(lenBOF)
			ar = ar[i+lenBOF:]
		}
		i := len(ar)
		if i > (lenBOF + 8) {
			i -= (lenBOF + 8)
		}
		pos += int64(i)
		if _, err := archive.Seek(pos, 0); err != nil {
			return logging.Error(0xE3E7A5, err)
		}
	}
	return nil
}

// end
