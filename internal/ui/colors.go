package ui

import "fmt"

// ANSI color codes
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

// Color functions
func Red(text string) string {
	return ColorRed + text + ColorReset
}

func Green(text string) string {
	return ColorGreen + text + ColorReset
}

func Yellow(text string) string {
	return ColorYellow + text + ColorReset
}

func Cyan(text string) string {
	return ColorCyan + text + ColorReset
}

func Blue(text string) string {
	return ColorBlue + text + ColorReset
}

// Print colored text
func PrintGreen(format string, args ...interface{}) {
	fmt.Printf(ColorGreen+format+ColorReset+"\n", args...)
}

func PrintRed(format string, args ...interface{}) {
	fmt.Printf(ColorRed+format+ColorReset+"\n", args...)
}

func PrintCyan(format string, args ...interface{}) {
	fmt.Printf(ColorCyan+format+ColorReset+"\n", args...)
}

func PrintYellow(format string, args ...interface{}) {
	fmt.Printf(ColorYellow+format+ColorReset+"\n", args...)
}