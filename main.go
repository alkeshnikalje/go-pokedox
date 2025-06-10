package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func cleanInput(input string) []string {
	trimmedInput := strings.TrimSpace(input)
	lowerCaseInput := strings.ToLower(trimmedInput)
	return strings.Split(lowerCaseInput, " ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for  {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanedInput := cleanInput(userInput)
		formattedString := fmt.Sprintf("Your command was: %s", cleanedInput[0])
		fmt.Println(formattedString)
	}
}
