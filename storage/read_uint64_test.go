// -----------------------------------------------------------------------------
// go-backup/storage/read_uint64_test.go

package storage

import (
	"bytes"
	"math"
	"testing"
)

func Test_ReadUint64_(t *testing.T) {
	for _, tc := range []struct {
		input []byte
		want  uint64
	}{
		{[]byte{0}, 0},
		{[]byte{1, 1}, 1},
		{[]byte{1, 100}, 100},
		{[]byte{1, 255}, 255},
		{[]byte{2, 1, 0}, 256},
		{[]byte{2, 255, 255}, 65535},
		{[]byte{3, 1, 0, 0}, 65536},
		{[]byte{1, 255}, math.MaxUint8},
		{[]byte{2, 255, 255}, math.MaxUint16},
		{[]byte{4, 255, 255, 255, 255}, math.MaxUint32},
		{[]byte{8, 255, 255, 255, 255, 255, 255, 255, 255}, math.MaxUint64},
	} {
		buf := bytes.NewBuffer(tc.input)
		have, _ := ReadUint64(buf)
		if have != tc.want {
			t.Fail()
		}
	}
}

// end
