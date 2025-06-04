package pokecache

import (
	"sync"
	"time"
)

type Cache struct{
	cache map[string]cacheEntry
	mu sync.Mutex
}

type cacheEntry struct{
	createdAt time.Time
	data []byte
}

func NewCache(interval time.Duration) *Cache{
	newCache := Cache{
		cache: map[string]cacheEntry{},
		mu: sync.Mutex{},
	}
	go newCache.reapLoop(interval)
	return &newCache
}

func (c *Cache) Add(key string, val []byte){
	newEntry := cacheEntry{
		createdAt: time.Now(),
		data: val,
	}
	c.mu.Lock()
	c.cache[key] = newEntry
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool){
	c.mu.Lock()
	entry, ok := c.cache[key]
	c.mu.Unlock()
	if !ok{
		return nil, false
	}
	return entry.data, true
}

func (c *Cache) reapLoop(interval time.Duration){
	ticker := time.NewTicker(interval)
	for currentTime := range ticker.C{
		c.mu.Lock()
		for key, entry := range c.cache{
			if entry.createdAt.Add(interval).After(currentTime){
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}