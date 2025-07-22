package app

import (
	"remx/pkg/termc"

	"github.com/fatih/color"
)

var primaryRGB = termc.RGB{R: 217, G: 120, B: 63}
var secondaryRGB = termc.RGB{R: 224, G: 36, B: 29}
var secondary = color.RGB(80, 80, 80)
var primary = color.RGB(217, 120, 63)

func PrintCombined(key, value string) {
	secondary.Print(key)
	primary.Print(value + "   ")
}
