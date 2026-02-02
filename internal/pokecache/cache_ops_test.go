package pokecache

import (
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	interval := 10 * time.Second
	cache := NewCache(interval)

	if cache == nil {
		t.Fatalf("Expected non-nil cache, got nil")
	}

	if cache.interval != interval {
		t.Errorf("Expected interval %v, got %v", interval, cache.interval)
	}

	if cache.entries == nil {
		t.Errorf("Expected non-nil entries map, got nil")
	}
}

func TestCacheAddAndGet(t *testing.T) {
	cache := NewCache(5 * time.Second)
	key := "testKey"
	value := []byte("testValue")

	cache.Add(key, value)

	retrievedValue, found := cache.Get(key)
	if !found {
		t.Fatalf("Expected to find key %s in cache", key)
	}

	if string(retrievedValue) != string(value) {
		t.Errorf("Expected value %s, got %s", value, retrievedValue)
	}

	miss := "missingKey"
	_, found = cache.Get(miss)
	if found {
		t.Errorf("Did not expect to find key %s in cache", miss)
	}
}

func TestCacheExpiration(t *testing.T) {
	interval := 1 * time.Second
	cache := NewCache(interval)
	key := "expireKey"
	value := []byte("expireValue")

	cache.Add(key, value)

	time.Sleep(2 * time.Second) // Wait for the entry to expire

	_, found := cache.Get(key)
	if found {
		t.Errorf("Expected key %s to be expired and not found in cache", key)
	}
}

func TestCacheConcurrentAccess(t *testing.T) {
	cache := NewCache(5 * time.Second)
	key := "concurrentKey"
	value := []byte("concurrentValue")

	done := make(chan bool)

	// Writer goroutine
	go func() {
		for i := 0; i < 100; i++ {
			cache.Add(key, value)
		}
		done <- true
	}()

	// Reader goroutine
	go func() {
		for i := 0; i < 100; i++ {
			cache.Get(key)
		}
		done <- true
	}()

	// Wait for both goroutines to finish
	<-done
	<-done
}
