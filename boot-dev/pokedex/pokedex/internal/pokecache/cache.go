package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	store map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(duration time.Duration) *Cache {
	cache := &Cache{
		store: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}
	go cache.reapLoop(duration)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.store[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, exists := c.store[key]
	return entry.val, exists
}

func (c *Cache) reap(now time.Time, duration time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for key, entry := range c.store {
		if now.Sub(entry.createdAt) > duration {
			delete(c.store, key)
		}
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}
