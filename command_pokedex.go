package main

import (
	"fmt"
)

func commandPokedex(cfg *cliConfig, params ...string) error {
	pokemonList := cfg.pokeapiClient.Pokedex.ListPokedex()
	fmt.Println("Your Pokedex:")
	for _, pokemon := range pokemonList {
		fmt.Printf(" - %s\n", pokemon)
	}
	return nil
}
