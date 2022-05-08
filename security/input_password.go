// -----------------------------------------------------------------------------
// go-backup/security/input_password.go

package security

import (
	"fmt"
	"strings"
	"syscall"

	"golang.org/x/term"

	"github.com/balacode/go-backup/logging"
)

// InputPassword displays a command-line prompt to enter
// a password without displaying the entered characters.
//
// If 'confirm' is true, also prompts for a confirmation. If the password
// and confirmation don't match, returns a blank string and an error.
//
// This function may not work if 'Stdin' is redirected,
// for example when running in debug mode in VS Code.
//
func InputPassword(confirm bool) (string, error) {

	fmt.Print("Enter password: ")
	ar, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		return "", logging.Error(0xE8D7E9, err)
	}
	pwd := strings.TrimSpace(string(ar))
	if confirm {
		fmt.Print("Confirm password: ")
		ar, err := term.ReadPassword(int(syscall.Stdin))
		fmt.Println()
		if err != nil {
			return "", logging.Error(0xE8E62B, err)
		}
		confirm := strings.TrimSpace(string(ar))
		if confirm != pwd {
			msg := "passwords don't match"
			return "", logging.Error(0xE0AD60, msg)
		}
	}
	return pwd, nil
}

// end
