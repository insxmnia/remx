package ui

import (
	"fmt"
	"remx/pkg/termc"
	"remx/pkg/utility"
	"strings"

	"github.com/fatih/color"
)

func PrintStacked(key, value string, gapSize int, newline bool) {
	Colours["primary"].Faith.Print(key)
	Colours["secondary"].Faith.Print(value)
	for i := range gapSize {
		if i == 1000 {
			break
		}
		fmt.Print(" ")
	}
	if newline {
		println()
	}
}

func PrintGradient(text string, start, end RGB, newline bool) {
	length := len(text)
	if length == 0 {
		return
	}

	for i, r := range text {
		interp := func(start, end, i int) int {
			return start + (end-start)*i/(length-1)
		}

		rr := interp(start.R, end.R, i)
		gg := interp(start.G, end.G, i)
		bb := interp(start.B, end.B, i)

		c := color.RGB(rr, gg, bb)
		c.Print(string(r))
	}

	if newline {
		println()
	}
}
func SprintGradient(text string, start, end RGB, newline bool) string {
	chars := []string{}
	length := len(text)
	if length == 0 {
		return text
	}

	for i, r := range text {
		interp := func(start, end, i int) int {
			return start + (end-start)*i/(length-1)
		}

		rr := interp(start.R, end.R, i)
		gg := interp(start.G, end.G, i)
		bb := interp(start.B, end.B, i)

		c := color.RGB(rr, gg, bb)
		chars = append(chars, c.Sprint(string(r)))
	}

	if newline {
		chars = append(chars, "\n")
	}
	return strings.Join(chars, "")
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

func RegisterColour(name string, rgb RGB) {
	faith := color.RGB(rgb.R, rgb.G, rgb.B)
	hex := utility.RGBToHex(rgb.R, rgb.G, rgb.B)
	ansii := fmt.Sprintf("\033[38;2;%d;%d;%dm", rgb.R, rgb.G, rgb.B)
	object := Colour{
		Name:  name,
		RGB:   rgb,
		Hex:   hex,
		Faith: faith,
		ANSII: ansii,
	}
	Colours[name] = object
}

func ToRGB(r, g, b int) RGB {
	return RGB{
		R: r, G: g, B: b,
	}
}

func InLoop(fn func()) {

	for {
		fn()
	}
}

func GetInput(key string) string {
	fmt.Print(key)
	var input string
	print(Colours["secondary"].ANSII)
	fmt.Scanln(&input)
	print(termc.Reset)
	return input
}
