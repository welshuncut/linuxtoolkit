package main

import (
	"github.com/TrolleyTrampInc/LinuxToolKit/davincisetup"
)

func main() {
	var choice int
	for {
		clearTerminal()
		printIntroMessage()
		choice = getUserInput()
		switch choice {
		case 1:
			davincisetup.Main()
		default:
			return
		}
	}
}
