package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicCache struct {
	mu   sync.Mutex
	data atomic.Value
}

// Set replaces the entire map atomically
func (c *AtomicCache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Load existing data (if any)
	oldData, _ := c.data.Load().(map[string]string)

	// Copy old map and update it
	newData := make(map[string]string, len(oldData)+1)
	for k, v := range oldData {
		newData[k] = v
	}
	newData[key] = value

	// Atomically store the new map
	c.data.Store(newData)
}

// Get retrieves a value atomically
func (c *AtomicCache) Get(key string) (string, bool) {
	data, _ := c.data.Load().(map[string]string)
	value, ok := data[key]
	return value, ok
}

func main() {
	cache := &AtomicCache{}

	cache.Set("foo", "bar")
	cache.Set("hello", "world")

	fmt.Println(cache.Get("foo"))   // Output: bar, true
	fmt.Println(cache.Get("hello")) // Output: world, true

	cache2 := &SyncMapCache{}

	cache2.Set("foo", "bar")
	cache2.Set("hello", "world")

	fmt.Println(cache2.Get("foo"))   // Output: bar, true
	fmt.Println(cache2.Get("hello")) // Output: world, true

}

type SyncMapCache struct {
	data sync.Map
}

// Set stores a key-value pair
func (c *SyncMapCache) Set(key, value string) {
	c.data.Store(key, value)
}

// Get retrieves a value
func (c *SyncMapCache) Get(key string) (string, bool) {
	val, ok := c.data.Load(key)
	if !ok {
		return "", false
	}
	return val.(string), true
}

func (c *SyncMapCache) showAll() {
	// its not atomic, we most use mutex, if we have range
	c.data.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}

func (c *SyncMapCache) IncrementCounter(key string) {
	// its not atomic, we most use mutex, if we have Load() → Modify → Store()
	//c.mu.Lock()
	//defer c.mu.Unlock()
	val, _ := c.data.Load(key)
	count := val.(int) + 1
	c.data.Store(key, count)
}
