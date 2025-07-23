package app

import (
	"remx/internal/inmemory"
	"remx/internal/selector"
	"remx/internal/ui"
	"remx/pkg/slogger"
	"remx/pkg/termc"
	"strings"
	"time"
)

func Entry() {
	termc.Clear()
	versions := GetDependenciesVersions()
	ui.PrintGradient(ui.Banner, ui.RGB(ui.Colours["gradient:primary"].RGB), ui.RGB(ui.Colours["gradient:secondary"].RGB), true)
	ui.PrintStacked("Version: ", "1.0.0a", 4, false)
	ui.PrintStacked("Latest Update: ", "23/07/2025 @ 00:07", 4, true)
	ui.PrintStacked("GO: ", versions["go"], 2, false)
	ui.PrintStacked("PNPM: ", versions["pnpm"], 2, false)
	ui.PrintStacked("NPM: ", versions["npm"], 2, true)
	println("\n")
	ui.PrintGradient("Resource & Environment Manager CLI (Developed by Insxmnia)", ui.RGB(ui.Colours["gradient:primary"].RGB), ui.RGB(ui.Colours["gradient:secondary"].RGB), true)

	ui.Colours["secondary"].Faith.Println(ui.Seperator)
	pname := termc.GetInput(ui.Colours["primary"].Faith.Sprint("What's the project called? "))
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

	template := inmemory.CF.Project.Templates[project][variant]
	dependencies, err := selector.Select(ui.UnderlineSprint(ui.Colours["primary"].Faith.Sprint("Would you like us to install the dependencies?")), []string{"Yes", "No"}, selector.Options{
		ItemSelectedColour: selector.RGB(ui.Colours["item:selected"].RGB),
		ItemFocusedColour:  selector.RGB(ui.Colours["item:focused"].RGB),
		SelectorIcon:       "• ",
	})
	if err != nil {
		slogger.Error("app", "variant selection failed", "error", err)
	}
	ready, err := selector.Select(ui.UnderlineSprint(ui.Colours["primary"].Faith.Sprint("Ready to start setup?")), []string{"Yes", "No"}, selector.Options{
		ItemSelectedColour: selector.RGB(ui.Colours["item:selected"].RGB),
		ItemFocusedColour:  selector.RGB(ui.Colours["item:focused"].RGB),
		SelectorIcon:       "• ",
	})
	if err != nil {
		slogger.Error("app", "variant selection failed", "error", err)
	}
	if ready == "No" {
		ui.Colours["secondary"].Faith.Println("Well tough shit, we're starting 🔥\n")
	}
	ui.WithSpinner(func() {
		time.Sleep(time.Second * 2)
	}, "Cloning repository ")
	ui.WithSpinner(func() {
		time.Sleep(time.Second * 2)
	}, "Installing dependencies ")

	println(template, dependencies, pname)
}
