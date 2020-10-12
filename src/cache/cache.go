package cache

import (
	"context"
	"errors"
)

// CacheEngine interface
type CacheEngine interface {
	// Saves the result to cache
	Save(id string, timestamp string)
	// Retrieves result from cache
	Get(id string) (timestamp string)
}

var ctx = context.Background()

// SetupCache - Create Cache Engine
func SetupCache(cacheEngineFlag string, minutesToExpire int, minutesToDelete int) (output CacheEngine, err error) {
	switch cacheEngine := cacheEngineFlag; cacheEngine {

	// case "redis":
	// 	output, err = NewRedisEngine()
	// 	return
	case "gocache":
		output, err = NewGoCacheEngine(minutesToExpire, minutesToDelete)
		return
	default:
		err = errors.New("cache engine not supported" + cacheEngine)
		return
	}
}
