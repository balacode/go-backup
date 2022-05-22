// -----------------------------------------------------------------------------
// go-backup/main.go

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/balacode/go-backup/archive"
	"github.com/balacode/go-backup/consts"
	"github.com/balacode/go-backup/logging"
	"github.com/balacode/go-backup/security"
)

func main() {
	runCommand(os.Args)
}

// runCommand runs the command, passing it arguments from the command line.
// (You could pass it a different set of arguments for testing.)
func runCommand(osArgs []string) {
	args, err := GetArgs(osArgs)
	if err != nil {
		fmt.Println(err)
		return
	}
	started := time.Now()
	key := security.KeyFromPassword(args.Password)
	enc, err := security.NewEncryption(key)
	if err != nil {
		logging.Error(0xE8B7EE, err)
		return
	}
	fmt.Println(started.Format(consts.TimestampFormat) + ": started")
	switch args.Command {
	case consts.CreateArchive:
		err := archive.CreateArchive(args.Target, args.Source, enc)
		if err != nil {
			logging.Error(0xE8D6D4, err)
			return
		}
	case consts.ExtractArchive:
		err := archive.ExtractArchive(args.Source, args.Target, enc)
		if err != nil {
			logging.Error(0xE40E36, err)
			return
		}
	}
	ended := time.Now()
	dur := ended.Sub(started)
	fmt.Println(started.Format(consts.TimestampFormat) + ": started")
	fmt.Println(ended.Format(consts.TimestampFormat) + ": ended")
	fmt.Println("time spent:", dur)
}

// end
