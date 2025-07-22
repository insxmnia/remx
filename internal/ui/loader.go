package ui

import (
	"time"

	"github.com/briandowns/spinner"
)

var s = spinner.New(spinner.CharSets[14], 100*time.Millisecond)

func WithSpinner(callback func(), message string) {
	s.Prefix = message
	s.Start()
	callback()
	s.Stop()
	s.Prefix = ""
}
