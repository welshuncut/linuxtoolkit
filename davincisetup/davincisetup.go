package davincisetup

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

func Main() {
	var packages = [6]string{"fuse2", "libxcrypt-compat", "rocm-opencl-runtime", "unzip", "qt5-wayland", "qt6-wayland"}

	fmt.Println("Now attempting to install needed packages")

	for _, packageName := range packages {
		checkAndInstallPackage(packageName)
	}

	checkPolkitAgentIsRunning()

	davinciFileName, err := lookForAndPrepareDavinciZip()
	if err != nil {
		panic(err)
	}

	launchInstaller(davinciFileName)
	err = removeProblemLibraries()
	if err != nil {
		fmt.Printf("Couldn't remove libraries. Error message %v", err)
	}

	fmt.Println("*****Setup Complete!!!*****")
	fmt.Println("                                                ")
	temp := ""
	for temp == "" {
		fmt.Println("Enter anything to continue....")
		fmt.Scan(&temp)
	}
}

func checkPolkitAgentIsRunning() {
	var polkits = []string{"polkit-kde-agent", "mate-polkit", "pantheon-polkit-agent", "deepin-polkit-agent", "polkit-gnome"}
	for _, agents := range polkits {
		cmd := exec.Command("bash", "-c", "pacman -Q |grep "+agents)
		output, _ := cmd.CombinedOutput()
		fmt.Println(string(output))

		if string(output) == "" {
			fmt.Println("Polkit not found, continue searching")
			continue
		} else {
			fmt.Printf("%v is installed!\n", agents)
			cmd := exec.Command("bash", "-c", "grep "+agents)
			output, _ := cmd.CombinedOutput()
			if output != nil {
				fmt.Printf("%v Polkit found and running\n", agents)
				fmt.Println("")
				break
			}
		}
	}
}

func removeProblemLibraries() error {

	fmt.Println("")
	var libraryArray = []string{"libglib-2.0*", "libgio-2.0*", "libgmodule-2.0*", "libgobject-2.0*"}

	fmt.Println("Programme must be launched in administrator mode for correct removal of libraries")
	err := os.Chdir("/opt/resolve/libs")
	if err != nil {
		return err
	}
	cmd := exec.Command("bash", "-c", "pwd")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(output))

	for _, entry := range libraryArray {
		cmd := exec.Command("bash", "-c", "sudo rm -f "+entry)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(string(output))
		}
		fmt.Println(string(output))
	}

	fmt.Println("Library tweaks completed! You should now be able to launch Resolve without any problems!")
	return nil
}

func launchInstaller(filename string) {
	lenStr := len(filename)
	newLen := lenStr - 3
	filenameSlice := filename[:newLen]
	fmt.Println("Now installing DaVinci Resolve")
	fmt.Println("Any following messages are from the Davinci installer, grep or qt warnings can be ignored")
	fmt.Print("\n\n")
	cmd := exec.Command("bash", "-c", "./"+filenameSlice+"run -iy")
	output, _ := cmd.CombinedOutput()
	fmt.Println(string(output))
}

func lookForAndPrepareDavinciZip() (fname string, err error) {
	fmt.Println("Changing directory to ~/Downloads")
	homedir, _ := os.UserHomeDir()
	os.Chdir(homedir + "/Downloads")
	output := exec.Command("bash", "-c", "ls")
	data, _ := output.CombinedOutput()
	fname = extractFilename(data)
	fmt.Println(fname)
	if fname != "" {
		fmt.Println("Unzipping DaVinci_Resolve_" + fname)
		cmd := exec.Command("bash", "-c", "unzip -u "+fname)
		output, _ := cmd.CombinedOutput()
		fmt.Println(string(output))
		fmt.Println("Successfully unzipped!")
		return fname, nil
	}
	println(string(data))
	return "", errors.New("unsuccesfull attempt at unzipping davinci resolve folder")
}

func extractFilename(data []byte) string {
	re := regexp.MustCompile(`DaVinci_Resolve_([A-Za-z0-9]+(\.[A-Za-z0-9]+)+)_Linux\.zip`)
	match := re.FindString(string(data))
	if match != "" {
		fmt.Println(match)
		return match
	}
	fmt.Println(match)
	return ""
}

func checkAndInstallPackage(packageName string) {
	cmd := exec.Command("bash", "-c", "pacman -Q |grep "+packageName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(output))
	if string(output) == "" {
		fmt.Println("Package not found")
		fmt.Printf("Now installing %v\n", packageName)
		installPackage(packageName)
		fmt.Printf("Successfully installed %v\n", packageName)
		return
	} else {
		fmt.Printf("%v is already installed! Skipping\n", packageName)
	}
}

func installPackage(packageName string) {
	cmd := exec.Command("bash", "-c", "sudo pacman -S --noconfirm "+packageName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	println(string(output))
}
