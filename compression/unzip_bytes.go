// -----------------------------------------------------------------------------
// go-backup/compression/unzip_bytes.go

package compression

import (
	"bytes"
	"compress/zlib"
	"io"

	"github.com/balacode/go-backup/logging"
)

// UnzipBytes uncompresses a ZLIB-compressed array of bytes.
func UnzipBytes(data []byte) ([]byte, error) {

	rd, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, logging.Error(0xE8F1D0, err)
	}
	buf := bytes.NewBuffer(make([]byte, 0, 8192)) // grows as needed
	_, err = io.Copy(buf, rd)
	if err != nil {
		return nil, logging.Error(0xE5BE26, err)
	}
	err = rd.Close()
	if err != nil {
		return nil, logging.Error(0xE0F0D0, err)
	}
	ret := buf.Bytes()
	return ret, nil
}

// end
