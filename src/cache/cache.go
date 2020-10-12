package cache

import (
	"context"
	"errors"

	"github.com/DanMHammer/statusmonitor/status"
)

// CacheEngine interface
type CacheEngine interface {
	// Saves the result to cache
	SaveStatus(status status.Status)
	// Retrieves result from cache
	GetStatus(id string) (status status.Status)
}

var ctx = context.Background()

// SetupCache - Create Cache Engine
func SetupCache() (output CacheEngine, err error) {
	switch cacheEngine := *cacheEngineFlag; cacheEngine {

	// case "redis":
	// 	output, err = NewRedisEngine()
	// 	return
	case "gocache":
		output, err = NewGoCacheEngine()
		return
	default:
		err = errors.New("cache engine not supported" + cacheEngine)
		return
	}
}
