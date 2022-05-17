// -----------------------------------------------------------------------------
// go-backup/storage/file.go

package storage

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/balacode/go-backup/compression"
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

// WriteEncryptedContent writes the file's content to writer 'wr'.
func (fl *File) WriteEncryptedContent(
	wr io.Writer,
	enc *security.Encryption,
) error {
	plaintext, err := compression.ZipBytes(fl.Content)
	if err != nil {
		return logging.Error(0xE7E12A, err)
	}
	ciphertext, err := enc.EncryptBytes(plaintext)
	if err != nil {
		return logging.Error(0xE07DB6, err)
	}
	n := uint64(len(ciphertext))
	if err := WriteUint64(wr, n); err != nil {
		return logging.Error(0xE79C2C, err)
	}
	if err := WriteBytes(wr, ciphertext); err != nil {
		return logging.Error(0xE7DE95, err)
	}
	return nil
}

// -----------------------------------------------------------------------------

// readArchivedFileMetadata reads metadata from an archive.
func readArchivedFileMetadata(
	rd io.Reader,
	enc *security.Encryption,
) (*File, error) {
	metaLen, err := ReadUint64(rd)
	if err != nil {
		return nil, logging.Error(0xE26B82, err)
	}
	ciphertext := make([]byte, metaLen)
	n, err := rd.Read(ciphertext)
	if err != nil {
		return nil, logging.Error(0xE53EB5, err)
	}
	if uint64(n) != metaLen {
		msg := fmt.Sprintf("n:%v != metaLen:%v", n, metaLen)
		return nil, logging.Error(0xE9A8C7, msg)
	}
	plaintext, err := enc.DecryptBytes(ciphertext)
	if err != nil {
		return nil, logging.Error(0xE07A7C, err)
	}
	buf := bytes.NewBuffer(plaintext)

	size, err := ReadUint64(buf)
	if err != nil {
		return nil, logging.Error(0xE6AD91, err)
	}
	modTimeUnix, err := ReadUint64(buf)
	if err != nil {
		return nil, logging.Error(0xE5FA20, err)
	}
	hash := [64]byte{}
	n, err = buf.Read(hash[:])
	if err != nil {
		return nil, logging.Error(0xE88C50, err)
	}
	if n != 64 {
		const msg = "hash size is not 64 bytes"
		return nil, logging.Error(0xE8D04D, msg)
	}
	path, err := buf.ReadString(0)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, logging.Error(0xE2E33E, err)
	}
	return &File{
		Size:    size,
		ModTime: time.Unix(int64(modTimeUnix), 0),
		Hash:    hash,
		Path:    path,
	}, nil
}

// end
