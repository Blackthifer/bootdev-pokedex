package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

type cliCommand struct{
	name string
	description string
	callback func(*config) error
}

var allCommands map[string]cliCommand

func initCommands(){
	allCommands = map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Diplays the next page of maps",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the previous page of maps",
			callback: commandMapb,
		},
		"explore": {
			name: "explore",
			description: "Usage: explore <locationArea>; Displays all pokemon found in locationArea",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Usage: catch <pokemon>; Attempts to catch the pokemon",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Usage: inspect <pokemon>; Displays information about the pokemon if you've caught it",
			callback: commandInspect,
		},
	}
}

func commandExit(conf *config) error{
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *config) error{
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, v := range allCommands{
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil;
}

func commandMap(conf *config) error{
	fullUrl := baseUrl + "location-area/?offset=" + fmt.Sprint(conf.Next)
	data, err := getData(fullUrl, conf.Cache)
	if err != nil{
		return err
	}
	locationAreas, err := parseData[namedApiResourceList](data)
	if err != nil{
		return err
	}
	for _, area := range locationAreas.List{
		fmt.Println(area.Name)
	}
	conf.Next += 20
	conf.Previous += 20
	return nil
}

func commandMapb(conf *config) error{
	if conf.Previous < 0{
		fmt.Println("You are on the first page")
		return nil
	}
	fullUrl := baseUrl + "location-area/?offset=" + fmt.Sprint(conf.Previous)
	data, err := getData(fullUrl, conf.Cache)
	if err != nil{
		return err
	}
	locationAreas, err := parseData[namedApiResourceList](data)
	if err != nil{
		return err
	}
	for _, area := range locationAreas.List{
		fmt.Println(area.Name)
	}
	conf.Next -= 20
	conf.Previous -= 20
	return nil
}

func commandExplore(conf *config) error{
	err := checkArguments(conf, "explore")
	if err != nil{
		return err
	}
	fullUrl := baseUrl + "location-area/" + conf.Arguments[0]
	data, err := getData(fullUrl, conf.Cache)
	if err != nil{
		return err
	}
	exploredArea, err := parseData[locationArea](data)
	if err != nil{
		return err
	}
	for _, encounter := range exploredArea.PokemonEncounters{
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}

func commandCatch(conf *config) error{
	err := checkArguments(conf, "catch")
	if err != nil{
		return err
	}
	fullUrl := baseUrl + "pokemon/" + conf.Arguments[0]
	data, err := getData(fullUrl, conf.Cache)
	if err != nil{
		return err
	}
	pokemon, err := parseData[pokemon](data)
	if err != nil{
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", conf.Arguments[0])
	difficulty := math.Log2(float64(pokemon.BaseExp)) / 1.5
	if rand.Float64() * 10 > difficulty{
		fmt.Printf("%s was caught!\n", conf.Arguments[0])
	} else{
		fmt.Printf("%s escaped!\n", conf.Arguments[0])
	}
	return nil
}

func commandInspect(conf *config) error{
	err := checkArguments(conf, "inspect")
	if err != nil{
		return err
	}
	return nil
}

func checkArguments(conf *config, command string) error{
	if conf.Arguments == nil || len(conf.Arguments) < 1{
		return fmt.Errorf("wrong usage:\n%s", allCommands[command].description)
	}
	return nil
}