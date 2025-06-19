package core

import "sync"

type SafeMap struct {
	mu   sync.Mutex
	data map[string]float64
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]float64),
	}
}

func (safeMap *SafeMap) Set(key string, value float64) {
	safeMap.mu.Lock()
	defer safeMap.mu.Unlock()
	safeMap.data[key] = value
}

func (safeMap *SafeMap) Get(key string) (float64, bool) {
	safeMap.mu.Lock()
	defer safeMap.mu.Unlock()
	value, ok := safeMap.data[key]
	return value, ok
}
