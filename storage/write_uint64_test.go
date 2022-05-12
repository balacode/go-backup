// -----------------------------------------------------------------------------
// go-backup/storage/write_uint64_test.go

package storage

import (
	"bytes"
	"math"
	"testing"
)

func Test_WriteUint64_(t *testing.T) {
	for _, tc := range []struct {
		in   uint64
		want []byte
	}{
		{0, []byte{0}},
		{1, []byte{1, 1}},
		{100, []byte{1, 100}},
		{255, []byte{1, 255}},
		{256, []byte{2, 1, 0}},
		{65535, []byte{2, 255, 255}},
		{65536, []byte{3, 1, 0, 0}},
		{math.MaxUint8, []byte{1, 255}},
		{math.MaxUint16, []byte{2, 255, 255}},
		{math.MaxUint32, []byte{4, 255, 255, 255, 255}},
		{math.MaxUint64, []byte{8, 255, 255, 255, 255, 255, 255, 255, 255}},
	} {
		buf := bytes.NewBuffer(make([]byte, 0, 10))
		WriteUint64(buf, tc.in)
		have := buf.Bytes()
		if !bytes.Equal(have, tc.want) {
			t.Fail()
		}
	}
}

// end
