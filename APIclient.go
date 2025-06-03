package main

import (
	"fmt"
	"io"
	"net/http"
)

const baseUrl = "https://pokeapi.co/api/v2/"

type config struct{
	Next string
	Previous string
}

func getData(url string) ([]byte, error){
	res, err := http.Get(url)
	if err != nil{
		return nil, fmt.Errorf("Error retrieving data: %w", err)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil{
		return nil, fmt.Errorf("Error reading response body: %w", err)
	}
	return data, nil
}