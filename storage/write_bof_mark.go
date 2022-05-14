// -----------------------------------------------------------------------------
// go-backup/storage/write_bof_mark.go

package storage

import (
	"io"

	"github.com/balacode/go-backup/logging"
)

// WriteBOFMark writes the Beginning-Of-File Marker to writer 'wr'.
func WriteBOFMark(wr io.Writer) error {
	if err := WriteBytes(wr, BOFMark); err != nil {
		return logging.Error(0xE4E1F7, err)
	}
	return nil
}

// end
