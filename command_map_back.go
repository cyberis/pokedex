package main

import (
	"fmt"

	"github.com/cyberis/pokedex/internal/pokeapi"
)

func commandMapBack(cfg *cliConfig) error {
	if cfg.prevURL == "nil" {
		fmt.Println("you're on the first page")
		return nil
	}
	url := cfg.prevURL
	locationList, err := pokeapi.GetLocationAreaList(url)
	if err != nil {
		return err
	}
	if locationList.Next != nil {
		cfg.nextURL = locationList.Next.(string)
	} else {
		cfg.nextURL = "nil"
	}
	if locationList.Previous != nil {
		cfg.prevURL = locationList.Previous.(string)
	} else {
		cfg.prevURL = "nil" //Required because .Prev can be null -> nil, nil is not a string but "nil" is
	}
	for _, location := range locationList.Results {
		fmt.Println(location.Name)
	}
	return nil
}
