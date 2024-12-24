package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mux  *sync.Mutex
	data map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
