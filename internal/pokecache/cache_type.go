package pokecache

import (
	"sync"
	"time"
)

type PokeCache struct {
	mux      *sync.Mutex
	interval time.Duration
	entries  map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
