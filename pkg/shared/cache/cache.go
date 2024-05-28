package cache

import (
	"github.com/zeebo/xxh3"
)

type Cache struct {
	storage   Storage
	displacer Displacer
}

func NewCache(storage Storage, displacer Displacer) *Cache {
	c := &Cache{
		storage:   storage,
		displacer: displacer,
	}
	c.displacer.Run(storage)
	return c
}

func (c *Cache) Get(key string, fn func(CacheItem) (data interface{}, err error)) (data interface{}, err error) {
	return c.storage.Get(xxh3.HashString(key), fn)
}

func (c *Cache) Delete(key string) {
	c.storage.Delete(xxh3.HashString(key))
}
