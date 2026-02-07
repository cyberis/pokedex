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

func (p *Pokedex) Exists(name string) bool {
	p.mux.Lock()
	defer p.mux.Unlock()

	_, exists := p.entries[name]
	return exists
}

func (p *Pokedex) ListPokedex() []string {
	p.mux.Lock()
	defer p.mux.Unlock()

	names := make([]string, 0, len(p.entries))
	for name := range p.entries {
		names = append(names, name)
	}
	return names
}
