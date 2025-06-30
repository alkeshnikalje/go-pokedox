package pokecache

import (
	"time"
	"sync"
)

type CacheEntry struct {
    CreatedAt time.Time
    Val       []byte
}

type Cache struct {
    Entries  map[string]CacheEntry
    Mu       sync.Mutex
	Interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		Entries: make(map[string]CacheEntry),
		Interval: interval,
	} 		
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	c.Entries[key] = CacheEntry{
		CreatedAt: time.Now(),
		Val: val,
	}
}

func (c *Cache) Get(key string) ([]byte,bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	value,ok := c.Entries[key]
	if !ok {
		return []byte{},false
	}
	return value.Val,true
}

func (c *Cache) ReadLoop() {
	ticker := time.NewTicker(c.Interval)
	defer ticker.Stop()

	for {
		<-ticker.C
		c.Mu.Lock()
		for k,entry := range c.Entries {
			if time.Since(entry.CreatedAt) > c.Interval {
				delete(c.Entries,k)
			}
		}
		c.Mu.Unlock()
	}
} 























