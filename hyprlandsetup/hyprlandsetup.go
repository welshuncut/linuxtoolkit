package hyprlandsetup

import (
	"fmt"
	"os/exec"
)

func Main() {
	fmt.Println("Now installing Hyprland....")
	editPacmanConf()
	updateSystem()
	err := installPackages()
	if err != nil {
		panic(err)
	}
}

func updateSystem() {
	fmt.Println("")
	fmt.Print("Updating system, Please enter ")

	cmd := exec.Command("bash", "-c", "sudo pacman -Syu")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Unable to update system")
		panic("ERROR")
	}
	fmt.Println("")
	fmt.Println(string(output))
	temp := ""
	for temp == "" {
		fmt.Println("Enter anything to continue....")
		fmt.Scan(&temp)
	}
}

func editPacmanConf() {
	cmd := exec.Command("bash", "-c", "sudo ./hyprlandsetup/editPacmanConf.sh")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		fmt.Println(err)
	} else {
		fmt.Println("Pacman conf successfully edited")
	}
}

func installPackages() error {
	fmt.Println("Installing base packages")
	cmd := exec.Command("bash", "-c", "sudo pacman -S --noconfirm hyprland hyprpaper xdg-desktop-portal-hyprland polkit-gnome swaync wofi nautilus waybar gnome-console firefox gnu-free-fonts qt5-wayland qt6-wayland")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}
