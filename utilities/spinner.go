package utilities

import (
	"time"

	"github.com/briandowns/spinner"
)

type Spinner interface {
	Start()
	Stop()
}

// BriandownsSpinner is an adapter for github.com/briandowns/spinner
type BriandownsSpinner struct {
	s *spinner.Spinner
}

func NewSpinner(color ...string) Spinner {
	sp := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
	sp.Color(color...)
	return &BriandownsSpinner{s: sp}
}

func (b *BriandownsSpinner) Start() {
	b.s.Start()
}

func (b *BriandownsSpinner) Stop() {
	b.s.Stop()
}
