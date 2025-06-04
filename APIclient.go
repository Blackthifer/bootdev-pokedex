package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseUrl = "https://pokeapi.co/api/v2/"

func getData(url string) ([]byte, error){
	res, err := http.Get(url)
	if err != nil{
		return nil, fmt.Errorf("Error retrieving data: %w", err)
	}
	defer res.Body.Close()
	if int(res.StatusCode / 100) != 2{
		return nil, fmt.Errorf("Non-ok response: %s", res.Status)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil{
		return nil, fmt.Errorf("Error reading response body: %w", err)
	}
	return data, nil
}

func parseData[T any](data []byte) (T, error){
	var parsedData T
	err := json.Unmarshal(data, &parsedData)
	if err != nil{
		var zero T
		return zero, fmt.Errorf("Error parsing data: %w", err)
	}
	return parsedData, nil
}

type namedApiResourceList struct{
	List []namedApiResource `json:"results"`
}

type namedApiResource struct{
	Name string `json:"name"`
	Url string `json:"url"`
}

type locationArea struct{
	PokemonEncounters []pokemonEncounter `json:"pokemon_encounters"`
}

type pokemonEncounter struct{
	Pokemon namedApiResource `json:"pokemon"`
}

type pokemon struct{
	BaseExp int `json:"base_experience"`
}