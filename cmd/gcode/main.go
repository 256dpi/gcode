package main

import (
	"fmt"

	"github.com/256dpi/gcode"
)

func main() {
	c := parseCommand()

	// triage command
	if c.cInfo {
		info(c)
	} else if c.cStrip {
		strip(c)
	} else if c.cOffset {
		offset(c)
	} else if c.cSVG {
		svg(c)
	}
}

func info(c *command) {
	// load g-code file
	file := loadFile(c.aInput)

	fmt.Printf("Lines: %d", len(file.Lines))
}

func strip(c *command) {
	// load g-code file
	file := loadFile(c.aInput)

	// strip comments
	gcode.StripComments(file)

	// write g-code file
	writeFile(c.aOutput, file)
}

func offset(c *command) {
	// load g-code file
	file := loadFile(c.aInput)

	// offset all coordinates
	gcode.OffsetXYZ(file, c.aX, c.aY, c.aZ)

	// write g-code file
	writeFile(c.aOutput, file)
}

func svg(c *command) {
	// load g-code file
	file := loadFile(c.aInput)

	// get svg code
	svg := gcode.ConvertToSVG(file)

	// write svg file
	writeFileString(c.aOutput, svg)
}
