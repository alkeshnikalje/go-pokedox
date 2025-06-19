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
    Entries  map[string]cacheEntry
    Mu       sync.Mutex
	Interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		Entries: make(map[string]cacheEntry),
		Interval: interval,
	} 		
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	c.Entries[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte,bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	value,ok := c.Entries[key]
	if !ok {
		return []byte{},false
	}
	return value.val,true
}

func (c *Cache) readLoop() {
	ticker := time.NewTicker(c.Interval)
	defer ticker.Stop()

	for {
		<-ticker.C
		c.Mu.Lock()
		for k,entry := range c.Entries {
			if time.Since(entry.createdAt) > c.Interval {
				delete(c.Entries,k)
			}
		}
		c.Mu.Unlock()
	}
} 























