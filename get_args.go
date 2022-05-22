// -----------------------------------------------------------------------------
// go-backup/get_args.go

package main

import (
	"strings"

	"github.com/balacode/go-backup/consts"
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

	// copy osArgs ignoring the first arg (program's name)
	args := make([]string, len(osArgs)-1)
	copy(args, osArgs[1:])

	ret := &Args{}
	ret.Password, args = extractNamedArg(args, "p", "pwd", "password")
	ret.Source, args = extractNextArg(args)
	ret.Target, args = extractNextArg(args)

	sourceIsArchive := strings.Contains(ret.Source, consts.ArchiveExt)
	targetIsArchive := strings.Contains(ret.Target, consts.ArchiveExt)

	switch {
	case targetIsArchive:
		ret.Command = consts.CreateArchive

	case sourceIsArchive:
		ret.Command = consts.ExtractArchive
	}
	return ret, nil
}

// extractNamedArg returns the value of the specified named argument,
// as well as 'argsIn' with the argument and value removed from args.
func extractNamedArg(
	args []string,
	argName ...string,
) (
	argValue string,
	argsOut []string,
) {
	argsOut = make([]string, 0, len(args))
	foundName := false
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "-") {
			a := strings.TrimLeft(arg, "-")
			for _, b := range argName {
				b = strings.TrimLeft(arg, "-")
				if a == b {
					foundName = true
					break
				}
			}
			if foundName && i < (len(args)-1) {
				argValue = args[i+1]
				i++
				continue
			}
		}
		argsOut = append(argsOut, args[i])
	}
	return argValue, argsOut
}

// extractNextArg returns the value of the next (i.e. first) argument in args,
// as well as 'argsIn' with the argument value removed from args.
func extractNextArg(argsIn []string) (argValue string, argsOut []string) {
	if len(argsIn) < 1 {
		return "", argsIn
	}
	return argsIn[0], argsIn[1:]
}

// end
