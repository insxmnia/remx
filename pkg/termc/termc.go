package termc

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
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

// RGB represents a color with red, green, and blue components
type RGB struct {
	R int
	G int
	B int
}

// gradientText prints a string where each character is colorized in a gradient from start to end
func GradientText(text string, start, end RGB, newline bool) {
	length := len(text)
	if length == 0 {
		return
	}

	for i, r := range text {
		// Linear interpolation of RGB values
		interp := func(start, end, i int) int {
			return start + (end-start)*i/(length-1)
		}

		rr := interp(start.R, end.R, i)
		gg := interp(start.G, end.G, i)
		bb := interp(start.B, end.B, i)

		// Use fatih/color's RGB function
		c := color.RGB(rr, gg, bb)
		c.Print(string(r))
	}

	if newline {
		println()
	}
}

func UnderlinedPrint(text string) {
	on := "\033[4m"
	off := "\033[0m"
	println(on + text + off)
}
func UnderlineSprint(text string) string {
	on := "\033[4m"
	off := "\033[0m"
	return on + text + off
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
