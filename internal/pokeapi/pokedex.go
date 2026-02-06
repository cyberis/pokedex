package pokeapi

import "sync"

func NewPokedex() *Pokedex {
	return &Pokedex{
		mux:     &sync.Mutex{},
		entries: make(map[string]Pokemon),
	}
}

func (p *Pokedex) Get(name string) (Pokemon, bool) {
	p.mux.Lock()
	defer p.mux.Unlock()

	pokemon, exists := p.entries[name]
	return pokemon, exists
}

func (p *Pokedex) Add(name string, pokemon Pokemon) {
	p.mux.Lock()
	defer p.mux.Unlock()

	p.entries[name] = pokemon
}
