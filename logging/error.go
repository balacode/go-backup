// -----------------------------------------------------------------------------
// go-backup/logging/error.go

package logging

import (
	"errors"
	"fmt"
	"log"
)

// Error creates and logs an error message.
//
// 'id' should be a unique error ID specified in hex,
// for example Error(0xE12345, "some error message")
//
// 'message' contains an error message.
// This should be a string or an error value.
//
// Returns an error containing a hexadecimal ID and error message.
//
func Error(id int, message interface{}) error {
	msg := fmt.Sprintf("%06X %v", id, message)
	log.Print("ERROR " + msg)
	return errors.New(msg)
}

// end
