package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	interval time.Duration
	mux      sync.Mutex
	cache    map[string]cacheEntry
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	entry, ok := c.cache[key]
	c.mux.Unlock()
	return entry.val, ok
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mux.Lock()
		for key, entry := range c.cache {

			if time.Since(entry.createdAt) > c.interval {
				delete(c.cache, key)
			}
		}
		c.mux.Unlock()
	}

}

func NewCache(interval time.Duration) *Cache {
	if interval <= 0 {
		panic("Interval cannot be equal or less than 0")
	}

	c := &Cache{
		interval: interval,
		cache:    make(map[string]cacheEntry),
	}

	go c.reapLoop()

	return c
}
