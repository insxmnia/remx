package termc

import (
	"fmt"
	"os"
	"os/exec"
	"remx/internal/ui"
	"runtime"
)

const (
	// Reset - will reset terminal formatting back to default
	Reset = "\x1b[0m"
	// Red - will set terminal text background colour to red
	Red = "\033[97m\033[41m"
	// Green - will set terminal text background colour to green
	Green = "\033[42m\033[97m"
	// Yellow - will set terminal text background colour to yellow
	Yellow = "\033[43m\033[97m"
)

var Custom map[string]string = make(map[string]string)

func RegisterCustom(name string, colourCode string) error {
	if Custom[name] != "" {
		return fmt.Errorf("colour with name %s already exists", name)
	}
	Custom[name] = colourCode
	return nil
}

func Clear() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls") // Windows
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear") // Unix-like (Linux, macOS)
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func GetInput(key string) string {
	fmt.Print(key)
	var input string
	print(ui.Colours["secondary"].ANSII)
	fmt.Scanln(&input)
	print(Reset)
	return input
}
