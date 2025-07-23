package screens

import (
	"fmt"
)

var screenCount int
var Screens map[string]Screen

type Screen struct {
	id       string
	name     string
	selected map[string]string
	show     func(s *Screen, args any)
}

func (s *Screen) Create(name string, fn func(s *Screen, args any)) {
	screenCount++
	s.id = fmt.Sprintf("%s.%d", name, screenCount)
	s.name = name
	s.selected = make(map[string]string)
	s.show = fn
}
func (s *Screen) Show(args any) {
	s.show(s, args)
}
func (s *Screen) GetSelected(key string) string {
	return s.selected[key]
}

func init() {
	Screens = make(map[string]Screen)

	main := Screen{}
	main.Create("main", MainViewFN)
	Screens["main"] = main

	projects := Screen{}
	projects.Create("projects", ProjectViewFN)
	Screens["projects"] = projects

	webserver := Screen{}
	webserver.Create("webserver", ServerLoggerFN)
	Screens["webserver"] = webserver
}
