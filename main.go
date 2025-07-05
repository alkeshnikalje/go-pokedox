package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/alkeshnikalje/go-pokedox/internal/pokeapi"
	"github.com/alkeshnikalje/go-pokedox/internal/pokecache"
)

func cleanInput(input string) []string {
	trimmedInput := strings.TrimSpace(input)
	lowerCaseInput := strings.ToLower(trimmedInput)
	return strings.Split(lowerCaseInput, " ")
}

type Config struct {
	Next	 string
	Previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config,c *pokecache.Cache, areaName string) error
}

func exitCommand(config *Config,c *pokecache.Cache, areaName string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	return errors.New("exit")
}


func mapCommand(config *Config,c *pokecache.Cache, areaName string) error {
	locationResponse,err := pokeapi.GetLocationAreas(config.Next,c)	
	if err != nil {
		return err
	}

	for i:=0; i<len(locationResponse.Results); i++ {
		fmt.Println(locationResponse.Results[i].Name)
	}

	config.Next = locationResponse.Next

	if locationResponse.Previous != nil {
		config.Previous = *locationResponse.Previous
	}

	return nil
}

func mapbCommand(config *Config,c *pokecache.Cache, areaName string) error {
	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locationResponse, err := pokeapi.GetLocationAreas(config.Previous,c)

	if err != nil {
		return err
	}
	for i:=0; i<len(locationResponse.Results); i++ {
		fmt.Println(locationResponse.Results[i].Name)
	}
	config.Next = locationResponse.Next
	if locationResponse.Previous != nil {
		config.Previous = *locationResponse.Previous
	}else{
		config.Previous = ""
	}	
	return nil
}

func exploreCommand (config *Config,c *pokecache.Cache, areaName string) error {

	areResponse,statusCode, err := pokeapi.GetArea(areaName,c)

	if err != nil {
		if statusCode == 404 {
			fmt.Println("Area not found. Please enter the valid area name")
			return err
		}
		return err
	}
	fmt.Println("Exploring" + " " + areaName + "...")
	fmt.Println("Found Pokemon:")
	for _,encounter := range areResponse.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
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
	"mapb": {
		name: 		 "mapb",
		description: "It's similar to the map command, however, instead of displaying the next 20 locations, it displays the previous 20 locations. It's a way to go back",
		callback: 	 mapbCommand,
	},
	"explore": {
		name:        "explore",
		description: "this command will help you explore the pokemons in a area you specify, you need to type an area's name after the explore command. For example- explore somearea.",
		callback: 	 exploreCommand,
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
	config := Config{Next: "https://pokeapi.co/api/v2/location-area"}
	cache := pokecache.NewCache(10*time.Second)
	go cache.ReadLoop() 
	for  {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanedInput := cleanInput(userInput)
		firstElement := cleanedInput[0]
		var secondElement = ""
		if len(cleanedInput) > 1 {
			secondElement = cleanedInput[1]
		}
		command, ok := cliCommands[firstElement]
		if !ok && firstElement != "help"{
			fmt.Println("unknown command")
			continue
		}

		if firstElement == "help" {
			helpCommand()
			continue
		}

		if command.name == "exit" {
			err := command.callback(&config,cache,secondElement)
			fmt.Println(err)
			os.Exit(0)
		}

		command.callback(&config,cache,secondElement)

	}	
}























