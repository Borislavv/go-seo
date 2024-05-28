package cache

import (
	"context"
	"sync"
	"time"
)

type MapCacheStorage struct {
	*mapCacheStorage
}
type mapCacheStorage struct {
	ctx      context.Context
	mu       sync.RWMutex
	storage  map[uint64]*Item
	capacity int64
}

// NewMapCacheStorage is a constructor of MapCacheStorage structure.
func NewMapCacheStorage(ctx context.Context) *MapCacheStorage {
	return &MapCacheStorage{
		mapCacheStorage: &mapCacheStorage{
			ctx:     ctx,
			mu:      sync.RWMutex{},
			storage: map[uint64]*Item{},
		},
	}
}

func (c *MapCacheStorage) Get(key uint64, fn func(CacheItem) (data interface{}, err error)) (data interface{}, err error) {
	item, found := c.get(key)
	if found {
		return item.data, nil
	}

	item, err = c.compute(fn)
	if err != nil {
		return nil, err
	}

	return c.set(key, item), nil
}

func (c *MapCacheStorage) get(key uint64) (item *Item, found bool) {
	defer c.mu.RUnlock()
	c.mu.RLock()
	item, found = c.storage[key]
	return item, found
}

func (c *MapCacheStorage) compute(fn func(CacheItem) (data interface{}, err error)) (item *Item, err error) {
	item = NewCacheItem()
	data, err := fn(item)
	if err != nil {
		return nil, err
	}
	item.data = data
	item.addedAt = time.Now()
	return item, nil
}

func (c *MapCacheStorage) set(key uint64, item *Item) (data interface{}) {
	defer c.mu.Unlock()
	c.mu.Lock()
	cacheItem, found := c.storage[key]
	if found {
		return cacheItem.data
	}
	c.storage[key] = item
	return item.data
}

func (c *MapCacheStorage) Delete(key uint64) {
	defer c.mu.Unlock()
	c.mu.Lock()
	delete(c.storage, key)
}

func (c *MapCacheStorage) Displace() {
	var keys []uint64

	c.mu.RLock()
	for key, item := range c.storage {
		if !item.expiresAt.IsZero() && item.expiresAt.UnixNano() <= time.Now().UnixNano() {
			keys = append(keys, key)
		}
	}
	c.mu.RUnlock()

	c.mu.Lock()
	for _, key := range keys {
		delete(c.storage, key)
	}
	c.mu.Unlock()
}
