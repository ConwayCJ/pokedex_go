package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	// cache["pokeapi.co/locationsetc"] : {createdAt, val[]byte}
	mu       sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	newEntry := cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}

	c.cache[key] = newEntry

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.cache[key]
	return entry.val, exists
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		<-ticker.C

		c.mu.Lock()
		now := time.Now()

		for key, entry := range c.cache {
			// Check if entry is older than reap timer
			if now.Sub(entry.createdAt) > c.interval {
				delete(c.cache, key)
			}
		}

		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cache:    make(map[string]cacheEntry),
		interval: interval,
	}

	go cache.reapLoop()

	return cache
}
