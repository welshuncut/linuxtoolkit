package main

import (
	"github.com/TrolleyTrampInc/linuxtoolkit/davincisetup"
	"github.com/TrolleyTrampInc/linuxtoolkit/hyprlandsetup"
)

func main() {
	var choice int
	for {
		clearTerminal()
		printIntroMessage()
		choice = getUserInput()
		switch choice {
		case 1:
			hyprlandsetup.Main()
		case 2:
			davincisetup.Main()
		default:
			return
		}
	}
}
