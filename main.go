// -----------------------------------------------------------------------------
// go-backup/main.go

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/balacode/go-backup/consts"
)

func main() {
	runCommand(os.Args)
}

// runCommand runs the command, passing it arguments from the comand line.
// (You could pass it a different set of arguments for testing.)
func runCommand(osArgs []string) {
	started := time.Now()
	fmt.Println(started.Format(consts.TimestampFormat) + ": started")
	{
		// TODO: add code to do actual work here
	}
	ended := time.Now()
	dur := ended.Sub(started)
	fmt.Println(started.Format(consts.TimestampFormat) + ": started")
	fmt.Println(ended.Format(consts.TimestampFormat) + ": ended")
	fmt.Println("time spent:", dur)
}

// end
