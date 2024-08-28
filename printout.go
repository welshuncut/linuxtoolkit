package main

import (
	"fmt"
	"os/exec"
)

func printIntroMessage() {
	fmt.Println("Welcome to Trolley's Linux toolkit!")
	fmt.Println("                                    ")
	fmt.Println("1. Hyprland Installer/Config")
	fmt.Println("2. Davinci Resolve Installer")
	fmt.Println("                                    ")
	fmt.Print("Please select an option: ")
}

func clearTerminal() {
	cmd := exec.Command("bash", "-c", "clear")
	output, _ := cmd.CombinedOutput()
	fmt.Println(string(output))
}
