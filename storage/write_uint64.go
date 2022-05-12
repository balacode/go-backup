// -----------------------------------------------------------------------------
// go-backup/storage/write_uint64.go

package storage

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/balacode/go-backup/logging"
)

// WriteUint64 writes a uint64 value to writer 'wr'
// using the minimum number of bytes.
//
// The written bytes consist of two parts:
//
// * The first byte specifies the number of additional bytes needed to store
// the unsigned. It can range from 0 (for 0) to 8 bytes for largest numbers.
//
// * The second part contains from 0 to 8 bytes that represent the number.
//
// The function saves space provided most numbers are not large:
// 0 only needs 1 byte, 0-255 only needs two bytes, etc.
// But math.MaxUint64 would need 9 bytes instead of 8.
//
func WriteUint64(wr io.Writer, num uint64) error {
	var (
		size byte
		ar   []byte
	)
	if num > 0 {
		buf := bytes.NewBuffer(make([]byte, 0, 8))
		err := binary.Write(buf, binary.BigEndian, num)
		if err != nil {
			return logging.Error(0xE9DB47, err)
		}
		ar = buf.Bytes()
		size = byte(8)
		for i := 0; i < 8; i++ {
			if ar[i] != 0 {
				break
			}
			size--
		}
		if size < 8 {
			ar = ar[8-size:]
		}
	}
	n, err := wr.Write([]byte{size})
	if err != nil {
		return logging.Error(0xE93E89, err)
	}
	if n != 1 {
		return logging.Error(0xE57A07, err)
	}
	if num > 0 {
		n, err = wr.Write(ar)
		if err != nil {
			return logging.Error(0xE89C03, err)
		}
		if n != len(ar) {
			return logging.Error(0xE92C1D, err)
		}
	}
	return nil
}

// end
