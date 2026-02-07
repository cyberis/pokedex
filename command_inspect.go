package main

import (
	"errors"
	"fmt"
	"log"
)

func commandInspect(cfg *cliConfig, params ...string) error {
	pokemonName := params[0]
	inPokedex := cfg.pokeapiClient.Pokedex.Exists(pokemonName)
	if !inPokedex {
		fmt.Printf("%s is not in your Pokedex. Try catching it first!\n", pokemonName)
		log.Printf("Attempted to inspect a pokemon not in pokedex: %s", pokemonName)
		return errors.New("pokemon not found in pokedex: " + pokemonName)
	}
	fmt.Printf("Name: %s\n", pokemonName)
	height, err := cfg.pokeapiClient.GetPokemonHeight(pokemonName)
	if err != nil {
		return errors.New("failed to fetch pokemon height: " + pokemonName + ": " + err.Error())
	}
	weight, err := cfg.pokeapiClient.GetPokemonWeight(pokemonName)
	if err != nil {
		return errors.New("failed to fetch pokemon weight: " + pokemonName + ": " + err.Error())
	}
	fmt.Printf("Height: %d\n", height)
	fmt.Printf("Weight: %d\n", weight)
	pokemonStats, err := cfg.pokeapiClient.GetPokemonStats(pokemonName)
	if err != nil {
		return errors.New("failed to fetch pokemon stats: " + pokemonName + ": " + err.Error())
	}
	for statName, statValue := range pokemonStats {
		fmt.Printf("  -%s: %d\n", statName, statValue)
	}
	pokemonTypes, err := cfg.pokeapiClient.GetPokemonTypes(pokemonName)
	if err != nil {
		return errors.New("failed to fetch pokemon types: " + pokemonName + ": " + err.Error())
	}
	fmt.Printf("Types:\n")
	for _, pokemonType := range pokemonTypes {
		fmt.Printf("  - %s\n", pokemonType)
	}
	return nil
}
