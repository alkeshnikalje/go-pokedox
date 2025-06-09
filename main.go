package main

import (
	"fmt"
	"strings"
)

func cleanInput(input string) []string {
	trimmedInput := strings.TrimSpace(input)
	lowerCaseInput := strings.ToLower(trimmedInput)
	return strings.Split(lowerCaseInput, " ")
}

func main() {
	fmt.Println("Hello, World!")
}
