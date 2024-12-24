package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) Cache {
	c := Cache{
		mux:  &sync.Mutex{},
		data: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	entry := cacheEntry{time.Now(), val}
	c.data[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	entry, exists := c.data[key]
	if !exists {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()

	for key, entry := range c.data {
		if entry.createdAt.Before(now.Add(-last)) {
			delete(c.data, key)
		}
	}
}
