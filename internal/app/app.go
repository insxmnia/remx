package app

import (
	"remx/internal/screens"
	"remx/internal/ui"

	title "github.com/lxi1400/GoTitle"
)

func Entry() {
	versions := GetDependenciesVersions()
	title.SetTitle("REMx - Resource and Environment Manager")
	ui.InLoop(func() {

		if screen, ok := screens.Screens["main"]; ok {
			screen.Show(versions)
		}
	}, 100)

}
