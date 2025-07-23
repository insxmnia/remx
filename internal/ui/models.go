package ui

import "github.com/fatih/color"

type RGB struct {
	R int
	G int
	B int
}

type Colour struct {
	Name  string
	RGB   RGB
	Faith *color.Color
	Hex   string
	ANSII string
}
