package pokeapi

import (
	"encoding/json"
	"log"
	"net/http"
)

type LocationAreaList struct {
	Count    int `json:"count"`
	Next     any `json:"next"`
	Previous any `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreaList(url string) (*LocationAreaList, error) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching location area list: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	var locationAreaList LocationAreaList
	if err := json.NewDecoder(resp.Body).Decode(&locationAreaList); err != nil {
		log.Printf("Error decoding location area list JSON: %v", err)
		return nil, err
	}

	return &locationAreaList, nil
}
