package app

import "remx/internal/terminal"

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
