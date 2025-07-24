package utilities

import (
	"fmt"
	"io"
	"strings"

	"github.com/fatih/color"
)

// PrintSuccessBox prints a green success box to the provided writer.
func PrintSuccessBox(w io.Writer, message string) {
	printMessageBox(w, message, len(message), 2, color.BgGreen, color.FgBlack)
}

// PrintErrorBox prints a red error box to the provided writer.
func PrintErrorBox(w io.Writer, message string) {
	printMessageBox(w, message, len(message), 2, color.BgRed, color.FgWhite)
}

// printMessageBox prints the formatted message box to the writer, using color.
func printMessageBox(w io.Writer, message string, width int, padding int, bgColor color.Attribute, fgColor color.Attribute) {
	lines := FormatBox(message, width, padding)
	boxPrinter := color.New(bgColor, fgColor)
	fmt.Fprintln(w)
	boxPrinter.Fprintln(w, lines[0])
	boxPrinter.Fprintf(w, "%s\n", lines[1])
	boxPrinter.Fprintln(w, lines[2])
	fmt.Fprintln(w)
}

// FormatBox returns the lines of a message box as a slice of strings.
func FormatBox(message string, width int, padding int) []string {
	// Dynamically calculate the width based on the message length and padding
	contentWidth := len(message)
	boxWidth := contentWidth + 2*padding + 2
	// calculate left and right padding
	leftPadding := (contentWidth - len(message)) / 2
	rightPadding := contentWidth - len(message) - leftPadding
	// Use spaces for padding to ensure the message is centered
	paddedMessage := fmt.Sprintf("%s%s%s", strings.Repeat(" ", leftPadding), message, strings.Repeat(" ", rightPadding))
	// Create the top and bottom border of the box
	borderLine := strings.Repeat(" ", boxWidth)
	// Return the box as a slice of strings
	return []string{
		borderLine,
		fmt.Sprintf("%s %s %s", strings.Repeat(" ", padding), paddedMessage, strings.Repeat(" ", padding)),
		borderLine,
	}
}
