// -----------------------------------------------------------------------------
// go-backup/archive/extract_archive.go

package archive

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/balacode/go-backup/consts"
	"github.com/balacode/go-backup/logging"
	"github.com/balacode/go-backup/security"
	"github.com/balacode/go-backup/storage"
)

// ExtractArchive extracts all files in the specified 'archiveFile' to
// the folder specified by 'extractToPath', using 'enc' for decryption.
func ExtractArchive(
	archiveFile string,
	extractToPath string,
	enc *security.Encryption,
) error {
	if !storage.ExistsPath(archiveFile) {
		msg := "archive not found: " + archiveFile
		return logging.Error(0xE99CE8, msg)
	}
	if !strings.HasSuffix(archiveFile, consts.ArchiveExt) {
		msg := fmt.Sprintf("archive name must end with %q: %v",
			consts.ArchiveExt, archiveFile)
		return logging.Error(0xE6DF85, msg)
	}
	destPath := func() string {
		s := filepath.Base(archiveFile)
		s = s[:len(s)-len(consts.ArchiveExt)]
		return filepath.Join(extractToPath, s)
	}()
	if storage.ExistsPath(destPath) {
		msg := "destination folder already exists: " + destPath
		return logging.Error(0xE51D60, msg)
	}
	if err := os.MkdirAll(destPath, os.ModePerm); err != nil {
		return logging.Error(0xE4CD1A, err)
	}
	sn := 0
	actor := func(pos int64, fl *storage.File) error {
		path := filepath.Join(destPath, fl.Path)
		dir, _ := filepath.Split(path)
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return logging.Error(0xE29A98, err)
		}
		err = os.WriteFile(path, fl.Content, os.ModePerm)
		if err != nil {
			return logging.Error(0xE4A32D, err)
		}
		fmt.Printf("extract #%v @%v %vb %v\n", sn, pos, fl.Size, path)
		sn++
		return nil
	}
	const loadContent = true
	err := scanArchive(archiveFile, enc, loadContent, actor)
	if err != nil {
		return logging.Error(0xE05C93, err)
	}
	return nil
}

// end
