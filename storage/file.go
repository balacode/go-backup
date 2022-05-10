// -----------------------------------------------------------------------------
// go-backup/storage/file.go

package storage

import (
	"os"
	"time"

	"github.com/balacode/go-backup/logging"
	"github.com/balacode/go-backup/security"
)

// File holds file information and content.
type File struct {
	Path    string
	Size    uint64
	ModTime time.Time
	Hash    [64]byte
	Content []byte
}

// ReadFile reads a file from the specified path and creates a
// new File object containing the file's properties and data.
func ReadFile(path string) (*File, error) {
	//
	info, err := os.Stat(path)
	if err != nil {
		return nil, logging.Error(0xE40CC3, err)
	}
	if info.IsDir() {
		msg := "expected a file, but given a directory: " + path
		return nil, logging.Error(0xE4FD7B, msg)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, logging.Error(0xE7F7AC, err)
	}
	fl := &File{
		Path:    path,
		Size:    uint64(len(content)),
		ModTime: info.ModTime(),
		Hash:    security.MakeHash(content),
		Content: content,
	}
	return fl, nil
}

// end
