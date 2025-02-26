package utils

import (
	"fmt"
	"os"
	"regexp"
)

func ValidateInput() (string, []byte, error) {

	args := os.Args[1:]

	if len(args) != 1 {
		return "", nil, fmt.Errorf("please enter a valid number of arguments")
	}

	userText := args[0]

	if userText == "\\n" {
		return "\n", nil, nil
	}

	re := regexp.MustCompile(`\A((\\n)*)||n$`)
	userText = re.ReplaceAllString(userText, "$1")

	asciiTemplateByte, err := os.ReadFile("standard.txt")
	if err != nil {
		return "", nil, fmt.Errorf("error reading file: standard.txt, %w", err)
	}
	if len(userText) == 0 {
		return "", asciiTemplateByte, nil
	}

	return userText, asciiTemplateByte, nil
}
