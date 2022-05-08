// -----------------------------------------------------------------------------
// go-backup/compression/zip_bytes.go

package compression

import (
	"bytes"
	"compress/zlib"
	"fmt"

	"github.com/balacode/go-backup/logging"
)

// ZipBytes compresses an array of bytes and
// returns the ZLIB-compressed array of bytes.
func ZipBytes(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, nil
	}
	var buf bytes.Buffer
	wr := zlib.NewWriter(&buf)
	_, err := wr.Write(data)
	if err != nil {
		return nil, logging.Error(0xE5FC77, err)
	}
	err = wr.Close()
	if err != nil {
		return nil, logging.Error(0xE2D18D, err)
	}
	ret := buf.Bytes()
	if len(ret) < 3 {
		msg := fmt.Sprintf("len(ret):%v < 3", len(ret))
		return nil, logging.Error(0xE0AF0A, msg)
	}
	return ret, nil
}

// end
