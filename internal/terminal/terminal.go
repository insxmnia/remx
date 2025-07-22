package terminal

import (
	"os/exec"
	"strings"
)

func GetPNPMVersion() (string, error) {
	cmd := exec.Command("pnpm", "--version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(string(out), "\n", ""), nil
}
func GetNPMVersion() (string, error) {
	cmd := exec.Command("npm", "--version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(string(out), "\n", ""), nil
}
func GetGoVersion() (string, error) {
	cmd := exec.Command("go", "version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	output := strings.ReplaceAll(string(out), "\n", "")
	version := strings.ReplaceAll(strings.Split(output, " ")[2], "go", "")
	return version, nil
}
