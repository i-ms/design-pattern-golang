package strategy_cache

type EvictionAlgo interface {
	evict(c *Cache)
}
