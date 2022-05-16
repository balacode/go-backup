// -----------------------------------------------------------------------------
// go-backup/storage/file.go

package storage

import (
	"bytes"
	"io"
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

// WriteEncryptedMetadata encrypts and writes file metadata to writer 'wr'.
// The metadata consists of file size, modified time, hash and path.
func (fl *File) WriteEncryptedMetadata(
	wr io.Writer,
	enc *security.Encryption,
) error {
	var plaintext []byte
	{
		buf := bytes.NewBuffer(make([]byte, 0, 8+8+64+len(fl.Path)))
		if err := WriteUint64(buf, uint64(fl.Size)); err != nil {
			return logging.Error(0xE0B5D1, err)
		}
		modTimeUnix := fl.ModTime.Unix()
		if err := WriteUint64(buf, uint64(modTimeUnix)); err != nil {
			return logging.Error(0xE56E03, err)
		}
		if err := WriteBytes(buf, fl.Hash[:]); err != nil {
			return logging.Error(0xE51F11, err)
		}
		if err := WriteString(buf, fl.Path); err != nil {
			return logging.Error(0xE1C6E6, err)
		}
		plaintext = buf.Bytes()
	}
	ciphertext, err := enc.EncryptBytes(plaintext)
	if err != nil {
		return logging.Error(0xE8A2DA, err)
	}
	n := uint64(len(ciphertext))
	if err := WriteUint64(wr, n); err != nil {
		return logging.Error(0xE6AF25, err)
	}
	if err := WriteBytes(wr, ciphertext); err != nil {
		return logging.Error(0xE8F66C, err)
	}
	return nil
}

// end
