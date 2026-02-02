package pokecache

import (
	"log"
	"sync"
	"time"
)

func NewCache(interval time.Duration) *PokeCache {
	cache := PokeCache{
		mux:      &sync.Mutex{},
		interval: interval,
		entries:  make(map[string]cacheEntry),
	}

	// Start a goroutine to clean up old cache entries periodically
	go cache.reapLoop()

	return &cache
}

func (c *PokeCache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	if entry, exists := c.entries[key]; exists {
		log.Printf("Cache hit for key: %s", key)
		return entry.val, true
	}

	log.Printf("Cache miss for key: %s", key)
	return nil, false
}

func (c *PokeCache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	log.Printf("Cache added for key: %s", key)
}

func (c *PokeCache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		now := time.Now()
		c.mux.Lock()
		for key, entry := range c.entries {
			if now.Sub(entry.createdAt) >= c.interval {
				delete(c.entries, key)
				log.Printf("Cache entry expired for key: %s", key)
			}
		}
		c.mux.Unlock()
	}
}
