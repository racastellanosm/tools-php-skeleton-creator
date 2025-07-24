package utilities

import (
	"testing"
	"time"
)

// mockSpinner implements Spinner for testing
// It records calls to Start and Stop

type mockSpinner struct {
	started bool
	stopped bool
}

func (m *mockSpinner) Start() {
	m.started = true
}

func (m *mockSpinner) Stop() {
	m.stopped = true
}

func TestMockSpinner_StartStop(t *testing.T) {
	ms := &mockSpinner{}
	ms.Start()
	if !ms.started {
		t.Error("Start() did not set started to true")
	}
	ms.Stop()
	if !ms.stopped {
		t.Error("Stop() did not set stopped to true")
	}
}

func TestBriandownsSpinner_StartStop(t *testing.T) {
	sp := NewSpinner()
	// Start and Stop should not panic
	sp.Start()
	time.Sleep(100 * time.Millisecond)
	sp.Stop()
}
