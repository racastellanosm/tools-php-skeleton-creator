package utilities

import (
	"bytes"
	"regexp"
	"strings"
	"testing"
)

// stripANSI removes ANSI color codes from a string
func stripANSI(input string) string {
	ansi := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return ansi.ReplaceAllString(input, "")
}

func TestFormatBox_CentersMessage(t *testing.T) {
	message := "Hello"
	width := 20
	padding := 2
	box := FormatBox(message, width, padding)
	if len(box) != 3 {
		t.Errorf("Expected 3 lines in box, got %d", len(box))
	}
	// The message should be centered
	contentLine := box[1]
	if !strings.Contains(contentLine, message) {
		t.Errorf("Box does not contain the message: %s", message)
	}
}

func TestPrintSuccessBox_PrintsGreenBox(t *testing.T) {
	var buf bytes.Buffer
	message := "Success!"
	PrintSuccessBox(&buf, message)
	output := stripANSI(buf.String())
	if !strings.Contains(output, message) {
		t.Errorf("Output does not contain message: %s", message)
	}
}

func TestPrintErrorBox_PrintsRedBox(t *testing.T) {
	var buf bytes.Buffer
	message := "Error!"
	PrintErrorBox(&buf, message)
	t.Log(buf.String())
	output := stripANSI(buf.String())
	if !strings.Contains(output, message) {
		t.Errorf("Output does not contain message: %s", message)
	}
}
