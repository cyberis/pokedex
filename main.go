package main

import (
	"time"

	"github.com/cyberis/pokedex/internal/pokeapi"
)

func main() {
	blank := ""
	pokeapiClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &cliConfig{
		pokeapiClient: pokeapiClient,
		nextURL:       &blank,
		prevURL:       &blank,
	}

	repl(cfg)
}
