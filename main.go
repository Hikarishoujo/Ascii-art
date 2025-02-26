package main

import (
	"ascii-art/utils"
	"fmt"
)

func main() {
	userText, asciiTemplateByte, err := utils.ValidateInput()
	if err != nil {
		fmt.Println("Failed to get input because:", err)
		return
	}

	if userText == "\n" {
		fmt.Print("\n")
		return
	}

	asciiTable, err := utils.CharManip(userText, asciiTemplateByte)
	if err != nil {
		fmt.Println("Unable to print Ascii Art because:", err)
		return
	}
	utils.ExhibitArt(userText, asciiTable)

}
