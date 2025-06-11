package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"errors"
)

func cleanInput(input string) []string {
	trimmedInput := strings.TrimSpace(input)
	lowerCaseInput := strings.ToLower(trimmedInput)
	return strings.Split(lowerCaseInput, " ")
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func exitCommand() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	return errors.New("exit")
}


var cliCommands = map[string]cliCommand{
    "exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    exitCommand,  
    },
	"map": {
		name: 		 "map",
		description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call to map should display the next 20 locations, and so on.",
		callback: 	 mapCommand,
	},
}


func helpCommand() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	fmt.Println("help: Displays a help message")

	for _,command := range cliCommands {
		desc := fmt.Sprintf("%s: %s",command.name,command.description)
		fmt.Println(desc)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for  {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanedInput := cleanInput(userInput)
		firstElement := cleanedInput[0]
		if firstElement == "help" {
			helpCommand()
		} 
		command, ok := cliCommands[firstElement]
		if !ok && firstElement != "help"{
			fmt.Println("unknown command")
		}
		if command.name == "exit" {
			err := command.callback()
			fmt.Println(err)
			os.Exit(0)
		}
	}
}
