package utils

// import (
// 	"flight-service/models"
// 	"sync"
// 	"time"
// )

// type Cache struct {
// 	data map[string]models.FlightData
// 	mu   sync.Mutex
// }

// func (c *Cache) Get(key string) (models.FlightData, bool) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	val, found := c.data[key]
// 	return val, found
// }

// func (c *Cache) Set(key string, value models.FlightData, ttl time.Duration) {
// 	c.mu.Lock()
// 	c.data[key] = value
// 	c.mu.Unlock()
// 	go func() {
// 		time.Sleep(ttl)
// 		c.mu.Lock()
// 		delete(c.data, key)
// 		c.mu.Unlock()
// 	}()
// }

// var CacheInstance = Cache{data: make(map[string]models.FlightData)}
