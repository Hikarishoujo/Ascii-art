package utils

import (
	"fmt"
	"strings"
)

func ExhibitArt(userText string, asciiTable [][]string) {

	var output string

	userLine := strings.Split(userText, `\n`)
	if userText == "" {
		output = ""
	} else {
		for _, newLine := range userLine {
			if newLine == "" {
				output += "\n"
				continue
			} else {
				for i := 0; i < 8; i++ {
					for _, char := range newLine {
						output += asciiTable[int(char)-32][i]
					}
					output += "\n"
				}
			}
		}
	}
	fmt.Print(output)

}
