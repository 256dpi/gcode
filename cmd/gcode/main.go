package main

import "fmt"

func main() {
	command := parseCommand()

	// triage command
	if command.cInfo {
		info(command)
	} else if command.cProcess {
		process(command)
	}
}

func info(command *command) {
	// load g-code file
	file := loadFile(command.aInput)

	fmt.Printf("Lines: %d", len(file))
}

func process(command *command) {
	// load g-code file
	file := loadFile(command.aInput)

	// write g-code file
	writeFile(command.aOutput, file)
}
