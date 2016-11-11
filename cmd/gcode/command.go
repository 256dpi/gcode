package main

import "github.com/docopt/docopt-go"

type command struct {
	// commands
	cInfo    bool
	cProcess bool

	// arguments
	aInput  string
	aOutput string
}

func parseCommand() *command {
	usage := `gcode.

Usage:
  gcode info <input>
  gcode process <input> <output>

Options:
  -h --help         Show this screen.
`

	a, _ := docopt.Parse(usage, nil, true, "", false)

	return &command{
		// commands
		cInfo:    getBool(a["info"]),
		cProcess: getBool(a["process"]),

		// arguments
		aInput:  getString(a["<input>"]),
		aOutput: getString(a["<output>"]),
	}
}

func getBool(field interface{}) bool {
	if val, ok := field.(bool); ok {
		return val
	}

	return false
}

func getString(field interface{}) string {
	if str, ok := field.(string); ok {
		return str
	}

	return ""
}
