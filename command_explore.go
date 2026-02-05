package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *cliConfig, params []string) error {
	areaName := params[0]
	pokemonNames, err := cfg.pokeapiClient.ListPokemonNamesInLocationArea(areaName)
	if err != nil {
		return err
	}
	if len(pokemonNames) == 0 {
		return errors.New("no pokemon found in this location area")
	}
	for _, name := range pokemonNames {
		fmt.Println(name)
	}
	return nil
}
