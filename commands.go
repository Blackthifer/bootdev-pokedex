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
		"pokedex": {
			name: "pokedex",
			description: "Displays all unique pokemon you have caught",
			callback: commandPokedex,
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
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	difficulty := math.Log2(float64(pokemon.BaseExp)) / 1.5
	if rand.Float64() * 10 > difficulty{
		fmt.Printf("%s was caught!\n", pokemon.Name)
		conf.CaughtPokemon[pokemon.Name] = pokemon
	} else{
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}

func commandInspect(conf *config) error{
	err := checkArguments(conf, "inspect")
	if err != nil{
		return err
	}
	pokemon, ok := conf.CaughtPokemon[conf.Arguments[0]]
	if !ok{
		return fmt.Errorf("you have not caught that pokemon")
	}
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats{
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types{
		fmt.Println("  -", t.Type.Name)
	}
	return nil
}

func commandPokedex(conf *config) error{
	if conf.CaughtPokemon == nil || len(conf.CaughtPokemon) == 0{
		return fmt.Errorf("You haven't caught any pokemon yet!")
	}
	fmt.Println("Your Pokemon:")
	for name := range conf.CaughtPokemon{
		fmt.Println("  -", name)
	}
	return nil
}

func checkArguments(conf *config, command string) error{
	if conf.Arguments == nil || len(conf.Arguments) < 1{
		return fmt.Errorf("wrong usage:\n%s", allCommands[command].description)
	}
	return nil
}