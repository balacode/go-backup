// -----------------------------------------------------------------------------
// go-backup/storage/read_bof_mark.go

package storage

import (
	"bytes"
	"fmt"
	"io"

	"github.com/balacode/go-backup/logging"
)

// ReadBOFMark attempts to read the Beginning-Of-File Marker
// from the reader 'rd' at the position specified by 'pos'.
//
// When it succeeds, moves the reader's position to pos+len(BOFMark).
// When it fails, it attempts to move to 'pos'.
//
// Returns nil if successful, or an error if the marker could not be read.
//
func ReadBOFMark(rd io.ReadSeeker, pos int64) error {
	_, err := rd.Seek(pos, io.SeekStart)
	if err != nil {
		msg := fmt.Sprintf("could not seek to position %v", pos)
		return logging.Error(0xE83C91, msg)
	}
	lenBOFMark := len(BOFMark)
	buf := make([]byte, lenBOFMark)
	n, err := rd.Read(buf)
	if err != nil {
		_, _ = rd.Seek(pos, io.SeekStart)
		return logging.Error(0xE78E41, err)
	}
	if n != lenBOFMark {
		_, _ = rd.Seek(pos, io.SeekStart)
		msg := fmt.Sprintf("length read (%v) != BOFMark length (%v)",
			n, lenBOFMark)
		return logging.Error(0xE5C76E, msg)
	}
	if bytes.Compare(buf, BOFMark) != 0 {
		_, _ = rd.Seek(pos, io.SeekStart)
		msg := fmt.Sprintf("BOF mark not found at %v", pos)
		return logging.Error(0xE86FE3, msg)

	}
	return nil
}

// end
