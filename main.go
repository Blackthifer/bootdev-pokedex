package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/Blackthifer/bootdev-pokedex/internal/pokecache"
)

type config struct{
	Cache *pokecache.Cache
	Next int
	Previous int
	Arguments []string
	CaughtPokemon map[string]pokemon
}

func main(){
	mainConfig := initPokedex()
	inputScanner := bufio.NewScanner(os.Stdin)
	for true{
		fmt.Print("Pokedex > ")
		processInput(inputScanner, mainConfig)
	}
}

func initPokedex() *config{
	initCommands()
	conf := config{
		Cache: pokecache.NewCache(time.Second * 5),
		CaughtPokemon: map[string]pokemon{},
		Arguments: nil,
		Next: 0,
		Previous: -40,
	}
	return &conf
}

func cleanInput(text string) []string{
	words := strings.Fields(text)
	for i := range words{
		words[i] = strings.ToLower(words[i])
	}
	return words
}

func processInput(inputScanner *bufio.Scanner, mainConfig *config){
	if !inputScanner.Scan() && inputScanner.Err() != nil{
		fmt.Println("Error getting user input", inputScanner.Err().Error())
		os.Exit(1)
	}
	input := inputScanner.Text()
	cleaned := cleanInput(input)
	if len(cleaned) == 0{
		return
	}
	command, ok := allCommands[cleaned[0]]
	if !ok{
		fmt.Println("Unkown command")
		return
	}
	mainConfig.Arguments = cleaned[1:]
	err := command.callback(mainConfig)
	if err != nil{
		fmt.Println(err.Error())
	}
}