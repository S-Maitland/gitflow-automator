package cli

import "strings"

type Flags struct {
	Verbose bool
	DryRun bool
	Force bool
	Help bool
}

func ParseFlags(args []string) (Flags, []string) {
	flags := Flags{}
	remaining := []string{}

	for _, arg := range args {
		switch {
		case arg == "-v" || arg == "--verbose":
			flags.Verbose = true
		case arg == "-d" || arg == "--dry-run":
			flags.DryRun = true
		case arg == "-f" || arg == "--force":
			flags.Force = true
		case arg == "-h" || arg == "--help":
			flags.Help = true
		case strings.HasPrefix(arg, "-"):
			remaining = append(remaining, arg)
		default:
			remaining = append(remaining, arg)
		}
	}

	return flags, remaining
}