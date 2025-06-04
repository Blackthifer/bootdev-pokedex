package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/Blackthifer/bootdev-pokedex/internal/pokecache"
)

func main(){
	mainConfig, mainCache := initPokedex()
	inputScanner := bufio.NewScanner(os.Stdin)
	for true{
		fmt.Print("Pokedex > ")
		processInput(inputScanner, mainConfig, mainCache)
	}
}

func initPokedex() (*config, *pokecache.Cache){
	initCommands()
	conf := config{
		Next: 0,
		Previous: -40,
	}
	return &conf, pokecache.NewCache(time.Second * 5)
}

func cleanInput(text string) []string{
	words := strings.Fields(text)
	for i := range words{
		words[i] = strings.ToLower(words[i])
	}
	return words
}

func processInput(inputScanner *bufio.Scanner, mainConfig *config, mainCache *pokecache.Cache){
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
	err := command.callback(mainConfig, mainCache)
	if err != nil{
		fmt.Println(err.Error())
	}
}