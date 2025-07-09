package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
	"math/rand"

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
	callback    func(config *Config,c *pokecache.Cache, name string, pokemonMap map[string]pokeapi.Pokemon) error
}

func exitCommand(config *Config,c *pokecache.Cache, name string,pokemonMap map[string]pokeapi.Pokemon) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	return errors.New("exit")
}


func mapCommand(config *Config,c *pokecache.Cache, name string,pokemonMap map[string]pokeapi.Pokemon) error {
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

func mapbCommand(config *Config,c *pokecache.Cache, name string,pokemonMap map[string]pokeapi.Pokemon) error {
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

func exploreCommand (config *Config,c *pokecache.Cache, areaName string,pokemonMap map[string]pokeapi.Pokemon) error {

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

func catchCommand (config *Config,c *pokecache.Cache, pokemonName string,pokemonMap map[string]pokeapi.Pokemon) error {
	
	pokemon, ok := pokemonMap[pokemonName]
	if ok {
		fmt.Println(pokemon.Name,"was already caught")
		return nil
	}

	pokemonInfoResponse,statusCode, err := pokeapi.GetPokemon(pokemonName,c)

	if err != nil {
		if statusCode == 404 {
			fmt.Printf("Pokemon: %s not found.",pokemonName)
			return err
		}
		return err
	}
	fmt.Println("Throwing a Pokeball at",pokemonName + "...")
	
	randomRoll := rand.Intn(100)
	catchThreshold := 100 - (pokemonInfoResponse.BaseExp/2)	
	if randomRoll < catchThreshold {
		pokemonMap[pokemonName] = *pokemonInfoResponse
		fmt.Println(pokemonName,"was caught!")
	}else {
		fmt.Println(pokemonName,"was escaped!")
	} 

	return nil
}

func inspectCommand (config *Config,c *pokecache.Cache, pokemonName string,pokemonMap map[string]pokeapi.Pokemon) error {
	pokemon, ok := pokemonMap[pokemonName]	
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Println("Name:",pokemon.Name)
	fmt.Println("Height:",pokemon.Height)
	fmt.Println("Weight:",pokemon.Weight)
	fmt.Println("Stats:")
	for _,stat := range pokemon.Stats {
		fmt.Println("-",stat.Stat.Name,":",stat.BaseStart)
	}

	fmt.Println("Types:")
	for _,typ := range pokemon.Types {
		fmt.Println("-",typ.Type.Name)
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
	"catch": {
		name:		 "catch",
		description: "this command lets you catch the pokemon (you may not be able to catch in the first go though haha)",
		callback: 	 catchCommand,
	},
	"inspect": {
		name: 		 "inspect",
		description: "It takes the name of a Pokemon and prints the name, height, weight, stats and type(s) of the Pokemon.",
		callback:	 inspectCommand,
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
	pokemonMap := map[string]pokeapi.Pokemon{}
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

		if firstElement == "help" {
			helpCommand()
			continue
		}

		if firstElement == "pokedex" {
			fmt.Println("Your Pokedex:")
			for key,_ := range pokemonMap {
				fmt.Println("-",key)
			}
			continue
		}

		command, ok := cliCommands[firstElement]
		if !ok && firstElement != "help"{
			fmt.Println("unknown command")
			continue
		}

		if command.name == "exit" {
			err := command.callback(&config,cache,secondElement,pokemonMap)
			fmt.Println(err)
			os.Exit(0)
		}

		command.callback(&config,cache,secondElement,pokemonMap)

	}	
}























