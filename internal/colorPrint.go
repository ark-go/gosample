package internal

import (
	"fmt"
)

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed    Color = "\u001b[31m"
	ColorGreen  Color = "\u001b[32m"
	ColorYellow Color = "\u001b[33m"
	ColorBlue   Color = "\u001b[34m"
	ColorReset  Color = "\u001b[0m"
)

func PrintColor(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}

// func ColorPrint(ече) {
// 	useColor := flag.Bool("color", false, "display colorized output")
// 	flag.Parse()
// 	if *useColor {
// 		colorize(ColorBlue, "Hello, Darling!")
// 		return
// 	}
// 	fmt.Println("Hello, Darling!")
// }
