package ui

import (
	"github.com/fatih/color"
)

var (
	BoldText     = color.New(color.Bold).SprintlnFunc()
	BoldTextNoLn = color.New(color.Bold).SprintFunc()
	BlueText     = color.New(color.FgCyan).SprintlnFunc()
	RedText      = color.New(color.FgRed).SprintlnFunc()
	YellowText   = color.New(color.FgYellow).SprintlnFunc()
	PurpleText   = color.New(color.FgMagenta).SprintlnFunc()
	GreenText    = color.New(color.FgGreen).SprintlnFunc()
	NormalText   = color.New(color.FgWhite).SprintlnFunc()
	OrangeText   = color.New(color.FgHiRed).SprintlnFunc()
	BoldBlueText = color.New(color.Bold, color.FgCyan).SprintlnFunc()

	// Same line purple text
	PurpleTextSL = color.New(color.FgMagenta).SprintFunc()
)

func AddUnderDecor(s string) string {
	es := ""

	for i := 0; i < len(s); i++ {
		es += "="
	}

	return s + "\n" + es + "\n"
}
