package utility

import (
	"fmt"
	"strings"
)

var UseColor = true

func ColorPrint(str string, args ...any) {
	fmt.Printf(colorFmt(str), args...)
}

func ColorPrintln(str string, args ...any) {
	fmt.Printf(colorFmt(str)+"\n", args...)
}

func colorFmt(str string) string {
	colors := map[string]string{
		"{RED}": "\033[31m",
		"{YLW}": "\033[33m",
		"{BLU}": "\033[34m",
		"{PUR}": "\033[35m",
		"{GRN}": "\033[32m",
		"{RES}": "\033[0m",
	}

	if !UseColor {
		for k := range colors {
			colors[k] = ""
		}
	}

	return replaceAll(str+"{RES}", colors)
}

func replaceAll(str string, substrings map[string]string) string {
	for k, v := range substrings {
		str = strings.ReplaceAll(str, k, v)
	}

	return str
}
