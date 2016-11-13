package main

import (
	"strconv"

	"github.com/docopt/docopt-go"
)

type command struct {
	// commands
	cInfo   bool
	cStrip  bool
	cOffset bool
	cSVG    bool

	// arguments
	aInput  string
	aOutput string
	aX      float64
	aY      float64
	aZ      float64
}

func parseCommand() *command {
	usage := `gcode.

Usage:
  gcode info <input>
  gcode strip <input> <output>
  gcode offset <input> <output> <x> <y> <z>
  gcode svg <input> <output>

Options:
  -h --help  Show this screen.
`

	a, _ := docopt.Parse(usage, nil, true, "", false)

	return &command{
		// commands
		cInfo:   getBool(a["info"]),
		cStrip:  getBool(a["strip"]),
		cOffset: getBool(a["offset"]),
		cSVG:    getBool(a["svg"]),

		// arguments
		aInput:  getString(a["<input>"]),
		aOutput: getString(a["<output>"]),
		aX:      getFloat(a["<x>"]),
		aY:      getFloat(a["<y>"]),
		aZ:      getFloat(a["<z>"]),
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

func getFloat(field interface{}) float64 {
	str := getString(field)

	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}

	return f
}
