package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseUrl = "https://pokeapi.co/api/v2/"

type config struct{
	Next int
	Previous int
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

type namedApiResourceList struct{
	List []namedApiResource `json:"results"`
}

type namedApiResource struct{
	Name string `json:"name"`
	Url string `json:"url"`
}

func parseNamedList(data []byte) ([]namedApiResource, error){
	var resourceList namedApiResourceList
	err := json.Unmarshal(data, &resourceList)
	if err != nil{
		return nil, fmt.Errorf("Error parsing data: %w", err)
	}
	return resourceList.List, nil
}