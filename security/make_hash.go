// -----------------------------------------------------------------------------
// go-backup/security/make_hash.go

package security

import (
	"crypto/sha512"
)

// MakeHash returns the SHA-512 hash of byte slice 'data'.
func MakeHash(data []byte) [64]byte {
	if len(data) == 0 {
		return [64]byte{}
	}
	ret := sha512.Sum512(data)
	return ret
}

// end
