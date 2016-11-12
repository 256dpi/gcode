package main

import (
	"fmt"

	"github.com/256dpi/gcode"
)

func main() {
	command := parseCommand()

	// triage command
	if command.cInfo {
		info(command)
	} else if command.cStrip {
		strip(command)
	} else if command.cOffset {
		offset(command)
	}
}

func info(command *command) {
	// load g-code file
	file := loadFile(command.aInput)

	fmt.Printf("Lines: %d", len(file.Lines))
}

func strip(command *command) {
	// load g-code file
	file := loadFile(command.aInput)

	// strip comments
	gcode.StripComments(file)

	// write g-code file
	writeFile(command.aOutput, file)
}

func offset(command *command) {
	// load g-code file
	file := loadFile(command.aInput)

	// offset all coordinates
	gcode.OffsetXYZ(file, command.aX, command.aY, command.aZ)

	// write g-code file
	writeFile(command.aOutput, file)
}
