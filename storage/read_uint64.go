// -----------------------------------------------------------------------------
// go-backup/storage/read_uint64.go

package storage

import (
	"fmt"
	"io"

	"github.com/balacode/go-backup/logging"
)

// ReadUint64 reads a uint64 number from reader 'rd'.
//
// The number should have been previously written by WriteUint64().
// See WriteUint64() for more details how these numbers are stored
// to conserve space.
//
func ReadUint64(rd io.Reader) (uint64, error) {
	ar := [8]byte{}
	//
	// read the length of the number in bytes
	n, err := rd.Read(ar[:1])
	if err != nil {
		return 0, logging.Error(0xE3FE25, err)
	}
	if n != 1 {
		return 0, logging.Error(0xE47DA0, err)
	}
	size := int(ar[0])
	if size < 0 || size > 8 {
		msg := fmt.Sprintf("read invalid size: %v", size)
		return 0, logging.Error(0xE9C55A, msg)
	}
	if size == 0 {
		return 0, nil
	}
	// read the big-endian bytes comprising the number
	n, err = rd.Read(ar[:size])
	if err != nil {
		return 0, logging.Error(0xE3E63F, err)
	}
	if n != size {
		return 0, logging.Error(0xE73F90, err)
	}
	// build result by raising each byte from end by power of 256
	pow, ret := uint64(1), uint64(0)
	for i := size - 1; i >= 0; i-- {
		ret += uint64(ar[i]) * pow
		pow *= 256
	}
	return ret, nil
}

// end
