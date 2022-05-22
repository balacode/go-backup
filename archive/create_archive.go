// -----------------------------------------------------------------------------
// go-backup/archive/create_archive.go

package archive

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/balacode/go-backup/logging"
	"github.com/balacode/go-backup/security"
	"github.com/balacode/go-backup/storage"
)

// CreateArchive creates an archive in 'archiveFile', by compressing and
// encrypting all files in 'sourcePath' using 'enc' for encryption.
func CreateArchive(
	archiveFile string,
	sourcePath string,
	enc *security.Encryption,
) error {

	archiveFile = strings.TrimSpace(archiveFile)
	sourcePath = strings.TrimSpace(sourcePath)

	// validate arguments
	switch {
	case archiveFile == "":
		const msg = "archive file not specified"
		return logging.Error(0xE5C58C, msg)

	case storage.ExistsFile(archiveFile):
		msg := "file already exists: " + archiveFile
		return logging.Error(0xE85E89, msg)

	case sourcePath == "":
		const msg = "source path not specified"
		return logging.Error(0xE22A34, msg)

	case !storage.ExistsDir(sourcePath):
		msg := "source path not found"
		return logging.Error(0xE51EF7, msg)
	}
	if err := enc.Validate(); err != nil {
		return logging.Error(0xE3F9D1, err)
	}

	// create the archive file
	archive, err := os.Create(archiveFile)
	if err != nil {
		return logging.Error(0xE5C48A, err)
	}
	defer archive.Close()

	// add all files in sourcePath to the archive
	sn := -1
	filepath.Walk(
		sourcePath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return logging.Error(0xE46E8F, err)
			}
			if info.IsDir() {
				return nil
			}
			sn++
			fl, err := storage.ReadFile(path)
			if err != nil {
				return logging.Error(0xE0E1D1, err)
			}
			fl.SetRelativePath(sourcePath)
			if fl.Size != uint64(info.Size()) {
				msg := fmt.Sprintf("fl.Size:%v != info.Size():%v",
					fl.Size, info.Size())
				return logging.Error(0xE73B08, msg)
			}
			pos := int64(0)
			{
				info, err := archive.Stat()
				if err != nil {
					return logging.Error(0xE7E2F8, err)
				}
				pos = info.Size()
			}
			if err := storage.WriteBOFMark(archive); err != nil {
				return logging.Error(0xE62CF9, err)
			}
			if err := fl.WriteEncryptedMetadata(archive, enc); err != nil {
				return logging.Error(0xE9D64C, err)
			}
			if fl.Size > 0 {
				if err := fl.WriteEncryptedContent(archive, enc); err != nil {
					return logging.Error(0xE6CE17, err)
				}
			}
			fmt.Printf("archive #%v @%v %vb %v\n", sn, pos, fl.Size, fl.Path)
			return nil
		},
	)
	return nil
}

// end
