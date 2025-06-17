package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct {
    createdAt time.Time
    val       []byte
}

type Cache struct {
    entries  map[string]cacheEntry
    mu       sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		entries: make(map[string]cacheEntry),
		interval: interval,
	} 		
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte,bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value,ok := c.entries[key]
	if !ok {
		return []byte{},false
	}
	return value.val,true
}

func (c *Cache) readLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		<-ticker.C
		c.mu.Lock()
		for k,entry := range c.entries {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.entries,k)
			}
		}
		c.mu.Unlock()
	}
} 























