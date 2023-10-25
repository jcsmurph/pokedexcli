package pokecache

import (
	"sync"
	"time"
)

type PkCache struct {
	cacheMap map[string]cacheEntry
	mu       *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) PkCache {

	cache := PkCache{
		cacheMap: make(map[string]cacheEntry),
		mu:       &sync.Mutex{},
	}

    go cache.reapLoop(interval)

	return cache

}

func (cache *PkCache) AddToCache(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.cacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

}

func (cache *PkCache) GetFromCache(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cacheValue, ok := cache.cacheMap[key]

	if !ok {
		return []byte{}, false
	}

	return cacheValue.val, true
}

func (cache *PkCache) reapLoop(interval time.Duration) {
    ticker := time.NewTicker(interval)

    for range ticker.C {
        cache.reap(time.Now().UTC(), interval)
    }
}


func (cache *PkCache) reap(now time.Time, last time.Duration) {
    cache.mu.Lock()
    defer cache.mu.Unlock()
    for key, val := range cache.cacheMap {
        if val.createdAt.Before(now.Add(-last)) {
            delete(cache.cacheMap, key)
        }
    }
}
