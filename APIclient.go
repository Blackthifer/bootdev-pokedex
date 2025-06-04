package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/Blackthifer/bootdev-pokedex/internal/pokecache"
)

const baseUrl = "https://pokeapi.co/api/v2/"

func getData(url string, cache *pokecache.Cache) ([]byte, error){
	cacheData, ok := cache.Get(url)
	if ok{
		return cacheData, nil
	}
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
	cache.Add(url, data)
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
	Name string `json:"name"`
	BaseExp int `json:"base_experience"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Stats []pokemonStat `json:"stats"`
	Types []pokemonType `json:"types"`
}

type pokemonStat struct{
	Stat namedApiResource `json:"stat"`
	BaseStat int `json:"base_stat"`
}

type pokemonType struct{
	Type namedApiResource `json:"type"`
}