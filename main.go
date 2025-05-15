package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)

		if len(cleanedInput) > 0 {
			firstWord := cleanedInput[0]
			fmt.Printf("Your command was: %s\n", firstWord)
		}
	}
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
