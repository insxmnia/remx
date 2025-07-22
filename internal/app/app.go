package app

import (
	"remx/internal/inmemory"
	"remx/internal/selector"
	"remx/internal/ui"
	"remx/pkg/slogger"
	"remx/pkg/termc"
)

func Entry() {
	termc.Clear()
	versions := GetDependenciesVersions()
	termc.GradientText(ui.Banner, secondaryRGB, primaryRGB, true)
	PrintCombined("Version: ", "1.0.0a")
	PrintCombined("Latest Update: ", "23/07/2025 @ 00:07")
	println()
	PrintCombined("GO: ", versions["go"])
	PrintCombined("PNPM: ", versions["pnpm"])
	PrintCombined("NPM: ", versions["npm"]+"\n")
	println("\n")
	termc.GradientText("Resource & Environment Manager CLI (Developed by Insxmnia)", secondaryRGB, primaryRGB, true)

	secondary.Println(ui.Seperator)

	_, err := selector.Select(termc.UnderlineSprint(primary.Sprint("Select a project:")), inmemory.CF.Project.Types, selector.Options{
		ItemSelectedColour: selector.RGB(selector.RGB{R: 80, G: 80, B: 80}),
		ItemFocusedColour:  selector.RGB(primaryRGB),
		SelectorIcon:       "> ",
	})
	if err != nil {
		slogger.Error("app", "selector failed", "error", err)
	}

}
