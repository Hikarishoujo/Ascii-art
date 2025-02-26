package utils

import (
	"fmt"
	"strings"
)

func CharManip(userText string, asciiTemplateByte []byte) ([][]string, error) {

	// Split asciiTemplate by double newline ("\n\n") to get individual ASCII characters from standard.txt.

	asciiCharacters := strings.Split(string(asciiTemplateByte[1:]), "\n\n")

	// Initialize asciiTable (2D table) (using "make" to avoid out of range).
	asciiTable := make([][]string, len(asciiCharacters))

	for _, userTextChar := range userText {
		asciiIndex := int(userTextChar)
		if asciiIndex-32 < 0 || asciiIndex-32 >= len(asciiTable) {
			return nil, fmt.Errorf("found an Invalid Ascii Character: %d", userTextChar)
		}
	}

	// Populate asciiTable [["1 line of A"...."8th line of A"]["1 line of B"...."8th line of B"]["1 line of C"...."8th line of C"]...["1 line of ~"..."8th line of ~"]].
	for i := range asciiCharacters {
		lines := strings.Split(asciiCharacters[i], "\n")
		asciiTable[i] = append(asciiTable[i], lines...)

	}
	return asciiTable, nil

}
