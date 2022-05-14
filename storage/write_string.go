// -----------------------------------------------------------------------------
// go-backup/storage/write_string.go

package storage

import (
	"fmt"
	"io"

	"github.com/balacode/go-backup/logging"
)

// WriteString writes a string to writer 'wr'.
func WriteString(wr io.Writer, s string) error {
	n, err := wr.Write([]byte(s))
	if err != nil {
		return logging.Error(0xE8C6E3, err)
	}
	if n != len(s) {
		msg := fmt.Sprintf("n:%v != len(st):%v", n, len(s))
		return logging.Error(0xE90B00, msg)
	}
	return nil
}

// end
