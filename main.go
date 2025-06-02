package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	initCommands()
	inputScanner := bufio.NewScanner(os.Stdin)
	for true{
		fmt.Print("Pokedex > ")
		processInput(inputScanner)
	}
}

func cleanInput(text string) []string{
	words := strings.Fields(text)
	for i := range words{
		words[i] = strings.ToLower(words[i])
	}
	return words
}

func processInput(inputScanner *bufio.Scanner){
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
	err := command.callback()
	if err != nil{
		fmt.Println("err")
	}
}