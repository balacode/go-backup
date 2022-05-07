// -----------------------------------------------------------------------------
// go-backup/main.go

package main

import (
	"fmt"
	"os"
	"time"
)

// TimestampFormat specifies the timestamp format for Time.Format().
const TimestampFormat = "2006-01-02 15:04:05"

func main() {
	runCommand(os.Args)
}

// runCommand runs the command, passing it arguments from the comand line.
// (You could pass it a different set of arguments for testing.)
func runCommand(osArgs []string) {
	started := time.Now()
	fmt.Println(started.Format(TimestampFormat) + ": started")
	{
		// TODO: add code to do actual work here
	}
	ended := time.Now()
	dur := ended.Sub(started)
	fmt.Println(started.Format(TimestampFormat) + ": started")
	fmt.Println(ended.Format(TimestampFormat) + ": ended")
	fmt.Println("time spent:", dur)
}

// end
