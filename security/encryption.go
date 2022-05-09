// -----------------------------------------------------------------------------
// go-backup/security/encryption.go

package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"

	"github.com/balacode/go-backup/logging"
)

// Encryption provides symmetric encryption and decryption
// using AES-256 as the encryption cipher.
//
// Create a new Encryption with NewEncryption(key).
// You can then use the returned object to encrypt data
// with EncryptBytes() and decrypt with DecryptBytes().
//
type Encryption struct {
	key []byte
	gcm cipher.AEAD
}

// NewEncryption creates a new Encryption object initialized
// with the specified 'key' for encryption or decryption.
// The key must be exactly 32 bytes in length for AES-256
// encryption which is used by Encryption.
func NewEncryption(key []byte) (*Encryption, error) {
	enc := &Encryption{key: key}
	{
		block, err := aes.NewCipher(enc.key[:])
		if err != nil {
			return nil, logging.Error(0xE8B26B, err)
		}
		gcm, err := cipher.NewGCM(block)
		if err != nil {
			return nil, logging.Error(0xE2FD23, err)
		}
		enc.gcm = gcm
	}
	err := enc.Validate()
	if err != nil {
		return nil, logging.Error(0xE10E17, err)
	}
	return enc, nil
}

// DecryptBytes decrypts the specified ciphertext and returns its plaintext.
// If there was an error decrypting, returns nil and an error.
func (enc *Encryption) DecryptBytes(ciphertext []byte) ([]byte, error) {
	if err := enc.Validate(); err != nil {
		return nil, logging.Error(0xE6A97B, err)
	}
	ns := enc.gcm.NonceSize()
	nonce := ciphertext[:ns]
	ciphertext = ciphertext[ns:]
	plaintext, err := enc.gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, logging.Error(0xE34D37, err)
	}
	return plaintext, nil
}

// EncryptBytes encrypts the specified plaintext and returns its ciphertext.
// If there was an error encrypting, returns nil and an error.
func (enc *Encryption) EncryptBytes(plaintext []byte) ([]byte, error) {
	if err := enc.Validate(); err != nil {
		return nil, logging.Error(0xE49D37, err)
	}
	nonce := make([]byte, enc.gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, logging.Error(0xE4B0EF, err)
	}
	ciphertext := enc.gcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

// Validate checks if this Encryption object was initialized
// properly using NewEncryption(). The symmetric encryption
// key must be exactly 32 bytes in length (needed by AES-256).
func (enc *Encryption) Validate() error {
	if len(enc.key) != 32 {
		msg := fmt.Sprintf("len(enc.key):%v != 32", len(enc.key))
		return logging.Error(0xE8B6A0, msg)
	}
	if enc.gcm == nil {
		const msg = "gcm is nil"
		return logging.Error(0xE8BF8B, msg)
	}
	return nil
}

// end
