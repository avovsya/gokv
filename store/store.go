package store

import (
	"sync"
)

var storeLock sync.RWMutex
var storeMap map[string]string

func init() {
	storeMap = make(map[string]string)
}

func Put(key string, value string) error {
	storeLock.Lock()
	defer storeLock.Unlock()

	storeMap[key] = value
	return nil
}

func Get(key string) (string, error) {
	storeLock.RLock()
	defer storeLock.RUnlock()

	return storeMap[key], nil
}

func Delete(key string) error {
	storeLock.Lock()
	defer storeLock.Unlock()

	delete(storeMap, key)

	return nil
}
