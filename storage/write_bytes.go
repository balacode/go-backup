// -----------------------------------------------------------------------------
// go-backup/storage/write_bytes.go

package storage

import (
	"fmt"
	"io"

	"github.com/balacode/go-backup/logging"
)

// WriteBytes writes the bytes in 'data' to writer 'wr'.
func WriteBytes(wr io.Writer, data []byte) error {
	n, err := wr.Write(data)
	if err != nil {
		return logging.Error(0xE37D1A, err)
	}
	if n != len(data) {
		msg := fmt.Sprintf("n:%v != len(ar):%v", n, len(data))
		return logging.Error(0xE53F0B, msg)
	}
	return nil
}

// end
