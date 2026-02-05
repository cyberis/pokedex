package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *cliConfig, params ...string) error {
	if cfg.nextURL == nil {
		return errors.New("you're on the last page")
	}
	locationList, err := cfg.pokeapiClient.GetLocationAreaList(cfg.nextURL)
	if err != nil {
		return err
	}

	// Update Next and Previous URLs in the config
	cfg.nextURL = locationList.Next
	cfg.prevURL = locationList.Previous

	// Print the location names
	for _, location := range locationList.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *cliConfig, params ...string) error {
	if cfg.prevURL == nil {
		return errors.New("you're on the first page")
	}
	locationList, err := cfg.pokeapiClient.GetLocationAreaList(cfg.prevURL)
	if err != nil {
		return err
	}

	// Update Next and Previous URLs in the config
	cfg.nextURL = locationList.Next
	cfg.prevURL = locationList.Previous

	for _, location := range locationList.Results {
		fmt.Println(location.Name)
	}
	return nil
}
