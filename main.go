package cachecon

import (
	"sync"
	"time"
)

type Cache struct {
	data map[string]cacheValue
	mu *sync.Mutex 
}

type cacheValue struct {
	value interface{}
}

func New() *Cache {
	return &Cache{
		data: make(map[string]cacheValue),
		mu: new(sync.Mutex),
	}
}

func (c *Cache) Delete(key string, ttl time.Duration) {
	time.Sleep(ttl)
	c.mu.Lock()
	delete(c.data, key)
	c.mu.Unlock()
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.data[key] = cacheValue{
		value: value,
	}
	go c.Delete(key, ttl)
}

func (c *Cache) Get(key string) interface{} {
	c.mu.Lock()
	value := c.data[key]
	c.mu.Unlock()
	return value.value
}