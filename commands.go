package main

import (
	"fmt"
	"os"
	"github.com/Blackthifer/bootdev-pokedex/internal/pokecache"
)

type cliCommand struct{
	name string
	description string
	callback func(*config, *pokecache.Cache) error
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
	}
}

func commandExit(conf *config, cache *pokecache.Cache) error{
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *config, cache *pokecache.Cache) error{
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, v := range allCommands{
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil;
}

func commandMap(conf *config, cache *pokecache.Cache) error{
	fullUrl := baseUrl + "location-area/?offset=" + fmt.Sprint(conf.Next)
	data, ok := cache.Get(fullUrl)
	if !ok{
		newData, err := getData(fullUrl)
		if err != nil{
			return err
		}
		data = newData
		cache.Add(fullUrl, data)
	}
	locationAreas, err := parseNamedList(data)
	if err != nil{
		return err
	}
	for _, area := range locationAreas{
		fmt.Println(area.Name)
	}
	conf.Next += 20
	conf.Previous += 20
	return nil
}

func commandMapb(conf *config, cache *pokecache.Cache) error{
	if conf.Previous < 0{
		fmt.Println("You are on the first page")
		return nil
	}
	fullUrl := baseUrl + "location-area/?offset=" + fmt.Sprint(conf.Previous)
	data, ok := cache.Get(fullUrl)
	if !ok{
		newData, err := getData(fullUrl)
		if err != nil{
			return err
		}
		data = newData
		cache.Add(fullUrl, data)
	}
	locationAreas, err := parseNamedList(data)
	if err != nil{
		return err
	}
	for _, area := range locationAreas{
		fmt.Println(area.Name)
	}
	conf.Next -= 20
	conf.Previous -= 20
	return nil
}