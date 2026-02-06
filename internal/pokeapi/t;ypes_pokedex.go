package pokeapi

import "sync"

type Pokedex struct {
	mux     *sync.Mutex
	entries map[string]Pokemon
}
