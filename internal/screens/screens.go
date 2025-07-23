package screens

import (
	"os"
	"remx/internal/inmemory"
	"remx/internal/selector"
	"remx/internal/server"
	"remx/internal/terminal"
	"remx/internal/ui"
	"remx/pkg/slogger"
	"remx/pkg/termc"
	"strings"
	"time"
)

func GetDependenciesVersions() map[string]string {
	golang, err := terminal.GetGoVersion()
	if err != nil {
		golang = "none"
	}

	pnpm, err := terminal.GetPNPMVersion()
	if err != nil {
		pnpm = "none"
	}

	npm, err := terminal.GetNPMVersion()
	if err != nil {
		npm = "none"
	}

	return map[string]string{
		"go": golang, "pnpm": pnpm, "npm": npm,
	}
}
func Banner(versions map[string]string) {
	termc.Clear()
	ui.PrintGradient(ui.Banner, ui.RGB(ui.Colours["gradient:primary"].RGB), ui.RGB(ui.Colours["gradient:secondary"].RGB), true)
	ui.PrintStacked("Version: ", "1.0.0a", 4, false)
	ui.PrintStacked("Latest Update: ", "23/07/2025 @ 00:07", 4, true)
	ui.PrintStacked("GO: ", versions["go"], 2, false)
	ui.PrintStacked("PNPM: ", versions["pnpm"], 2, false)
	ui.PrintStacked("NPM: ", versions["npm"], 2, true)
	println()
	ui.PrintGradient("Resource & Environment Manager CLI (Developed by Insxmnia)", ui.RGB(ui.Colours["gradient:primary"].RGB), ui.RGB(ui.Colours["gradient:secondary"].RGB), true)
	ui.Colours["secondary"].Faith.Println(ui.Seperator)
}
func ProjectViewFN(s *Screen, args any) {
	Banner(args.(map[string]string))
	pname := termc.GetInput(ui.Colours["primary"].Faith.Sprint("What's the project called? "))
	s.selected["project:name"] = pname
	ptype, err := selector.Select(ui.UnderlineSprint(ui.Colours["primary"].Faith.Sprint("What project do you want?")), inmemory.CF.Project.Types, selector.Options{
		ItemSelectedColour: selector.RGB(ui.Colours["item:selected"].RGB),
		ItemFocusedColour:  selector.RGB(ui.Colours["item:focused"].RGB),
		SelectorIcon:       "• ",
	})
	if err != nil {
		slogger.Error("app", "project selection failed", "error", err)
	}
	project := strings.ReplaceAll(ptype, "(", "")
	project = strings.ReplaceAll(project, ")", "")
	project = strings.ReplaceAll(project, " ", "-")
	project = strings.ToLower(project)
	s.selected["project:type"] = project

	variant, err := selector.Select(ui.UnderlineSprint(ui.Colours["primary"].Faith.Sprint("Which variant?")), inmemory.CF.Project.Variants[project], selector.Options{
		ItemSelectedColour: selector.RGB(ui.Colours["item:selected"].RGB),
		ItemFocusedColour:  selector.RGB(ui.Colours["item:focused"].RGB),
		SelectorIcon:       "• ",
	})
	if err != nil {
		slogger.Error("app", "variant selection failed", "error", err)
	}
	variant = strings.Split(variant, "(")[0]
	if len(variant) > 0 && variant[len(variant)-1] == ' ' {
		variant = strings.TrimSuffix(variant, " ")
	}
	variant = strings.ReplaceAll(variant, " ", "-")
	variant = strings.ToLower(variant)
	s.selected["project:variant"] = variant
	template := inmemory.CF.Project.Templates[project][variant]
	dependencies, err := selector.Select(ui.UnderlineSprint(ui.Colours["primary"].Faith.Sprint("Would you like us to install the dependencies?")), []string{"Yes", "No"}, selector.Options{
		ItemSelectedColour: selector.RGB(ui.Colours["item:selected"].RGB),
		ItemFocusedColour:  selector.RGB(ui.Colours["item:focused"].RGB),
		SelectorIcon:       "• ",
	})
	if err != nil {
		slogger.Error("app", "variant selection failed", "error", err)
	}
	s.selected["project:installDependancies"] = dependencies
	ready, err := selector.Select(ui.UnderlineSprint(ui.Colours["primary"].Faith.Sprint("Ready to start setup?")), []string{"Yes", "No"}, selector.Options{
		ItemSelectedColour: selector.RGB(ui.Colours["item:selected"].RGB),
		ItemFocusedColour:  selector.RGB(ui.Colours["item:focused"].RGB),
		SelectorIcon:       "• ",
	})
	if err != nil {
		slogger.Error("app", "variant selection failed", "error", err)
	}
	s.selected["project:ready"] = ready
	if ready == "No" {
		ui.Colours["secondary"].Faith.Println("Well tough shit, we're starting 🔥")
	}
	ui.WithSpinner(func() {
		time.Sleep(time.Second * 2)
	}, ui.Colours["secondary"].Faith.Sprint("Cloning repository "))
	ui.WithSpinner(func() {
		time.Sleep(time.Second * 2)
	}, ui.Colours["secondary"].Faith.Sprint("Installing dependancies "))

	println(template, dependencies, pname)
}

func MainViewFN(s *Screen, args any) {
	Banner(args.(map[string]string))
	options := []string{
		"Authenticate",
		"New Project",
		"Web Server (Listening only)",
		"Quit",
	}
	app, err := selector.Select(ui.UnderlineSprint(ui.Colours["primary"].Faith.Sprint("What do you want to do?")), options, selector.Options{
		ItemSelectedColour: selector.RGB(ui.Colours["item:selected"].RGB),
		ItemFocusedColour:  selector.RGB(ui.Colours["item:focused"].RGB),
		SelectorIcon:       "• ",
	})
	if err != nil {
		slogger.Error("main-view", "variant selection failed", "error", err)
	}
	s.selected["application"] = app
	switch strings.ToLower(app) {
	case "authenticate":
		ui.Colours["secondary"].Faith.Println("This option is currently unavailable")
	case "web server (listening only)":
		if screen, ok := Screens["webserver"]; ok {
			screen.Show(args)
		}
	case "new project":
		if screen, ok := Screens["projects"]; ok {
			screen.Show(args)
		}
	case "quit":
		os.Exit(1)
	}
}

func ServerLoggerFN(s *Screen, args any) {
	Banner(args.(map[string]string))
	port := termc.GetInput(ui.Colours["primary"].Faith.Sprint("What port to listen on? "))
	Banner(args.(map[string]string))
	server.Run(port)
}
