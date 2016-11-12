package main

import (
	"os"

	"github.com/256dpi/gcode"
)

func loadFile(path string) *gcode.File {
	// open g-code file
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	// make sure file gets closed
	defer file.Close()

	// parse file
	f, err := gcode.ParseFile(file)
	if err != nil {
		panic(err)
	}

	return f
}

func writeFile(path string, f *gcode.File) {
	// create g-code file
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	// make sure file gets closed
	defer file.Close()

	// parse file
	err = gcode.WriteFile(file, f)
	if err != nil {
		panic(err)
	}
}
