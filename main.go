package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(cleanInput("Hello, World!"))
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	trimmedText := strings.TrimSpace(lowerText)
	cleanedInput := strings.Split(trimmedText, " ")
	if len(cleanedInput) == 0 {
		return []string{}
	}
	return cleanedInput
}
