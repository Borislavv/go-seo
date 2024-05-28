package cache

type Storage interface {
	Get(key uint64, fn func(CacheItem) (data interface{}, err error)) (data interface{}, err error)
	Delete(key uint64)
	Displace()
}
