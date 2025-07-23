package ui

var Colours map[string]Colour

func init() {
	Colours = make(map[string]Colour)
	RegisterColour("primary", ToRGB(217, 120, 63))
	RegisterColour("secondary", ToRGB(100, 100, 100))
	RegisterColour("gradient:primary", ToRGB(224, 36, 29))
	RegisterColour("gradient:secondary", ToRGB(217, 120, 6))
	RegisterColour("item:focused", ToRGB(204, 58, 0))
	RegisterColour("item:selected", ToRGB(100, 100, 100))
	RegisterColour("item:normal", ToRGB(100, 100, 100))

}
