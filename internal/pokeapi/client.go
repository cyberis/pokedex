package pokeapi

import (
	"net/http"
	"time"

	"github.com/cyberis/pokedex/internal/pokecache"
)

// Client -
type Client struct {
	httpClient http.Client
	cache      *pokecache.PokeCache
	Pokedex    *Pokedex
}

// NewClient -
func NewClient(timeout, cacheExpiresIn time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache:   pokecache.NewCache(cacheExpiresIn),
		Pokedex: NewPokedex(),
	}
}
