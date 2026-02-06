package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *cliConfig, params ...string) error {
	pokemonName := params[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	baseExperience, err := cfg.pokeapiClient.GetPokemonBaseExperience(pokemonName)
	if err != nil {
		return errors.New("failed to fetch pokemon: " + pokemonName + ": " + err.Error())
	}
	// Simulate catching the Pokemon with a random chance of success
	if rand.Intn(400) < baseExperience {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}
	fmt.Printf(" %s was caught!\n", pokemonName)
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return errors.New("failed to fetch pokemon: " + pokemonName + ": " + err.Error())
	}
	cfg.pokeapiClient.Pokedex.Add(pokemonName, pokemon)
	return nil
}
