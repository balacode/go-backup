// -----------------------------------------------------------------------------
// go-backup/security/key_from_password.go

package security

// KeyFromPassword creates a 32-byte encryption key from the given password.
// This is the symmetric encryption key for use with NewEncryption().
func KeyFromPassword(pwd string) []byte {
	hash := MakeHash([]byte(pwd))
	key := make([]byte, 32)
	for i := 0; i < 32; i++ {
		key[i] = hash[i*2] ^ hash[i*2+1]
	}
	return key
}

// end
