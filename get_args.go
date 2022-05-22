// -----------------------------------------------------------------------------
// go-backup/get_args.go

package main

import (
	"strings"

	"github.com/balacode/go-backup/consts"
	"github.com/balacode/go-backup/logging"
)

// Args contains parsed command-line arguments passed to the command.
type Args struct {
	Command  consts.Action
	Source   string
	Target   string
	Password string
}

// GetArgs parses command-line arguments and returns an Args struct.
//
// You should pass os.Args in osArgs parameter, except when testing.
//
// Returns an error if any of the required parameters is missing or incorrect.
//
// Also, if a password hasn't been specified, displays a password prompt.
//
func GetArgs(osArgs []string) (*Args, error) {
	var (
		ret             = parseOSArgs(osArgs)
		sourceIsArchive = strings.Contains(ret.Source, consts.ArchiveExt)
		targetIsArchive = strings.Contains(ret.Target, consts.ArchiveExt)
	)
	switch {
	case ret.Source == "":
		const msg = "source not specified"
		return nil, logging.Error(0xE9A3F4, msg)

	case ret.Target == "":
		const msg = "target not specified"
		return nil, logging.Error(0xE3A5C5, msg)

	case !sourceIsArchive && !targetIsArchive:
		const msg = "archive not specified"
		return nil, logging.Error(0xE26A8E, msg)

	case sourceIsArchive && targetIsArchive:
		const msg = "archive can be both source and target"
		return nil, logging.Error(0xE40F46, msg)
	}
	if targetIsArchive {
		ret.Command = consts.CreateArchive
	}
	if sourceIsArchive {
		ret.Command = consts.ExtractArchive
	}
	return ret, nil
}

// parseOSArgs is called by GetArgs() to parse command-line arguments
// into Args without checking their values or doing additional i/o.
func parseOSArgs(osArgs []string) *Args {

	ret := &Args{}
	args := make([]string, len(osArgs)-1)
	copy(args, osArgs[1:])

	for i := 0; i < len(args); i++ {
		arg := args[i]
		if !strings.HasPrefix(arg, "-") {
			continue
		}
		s := strings.TrimLeft(arg, "-")
		if (s == "p" || s == "password") && i < (len(args)-1) {
			ret.Password = args[i+1]
			args[i], args[i+1] = "", ""
			i++
		}
	}
	pull := func() string {
		for i, arg := range args {
			if arg != "" {
				args[i] = ""
				return arg
			}
		}
		return ""
	}
	ret.Source = pull()
	ret.Target = pull()
	return ret
}

// end
