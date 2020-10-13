package cache

import (
	"context"
	"errors"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// CacheEngine interface
type CacheEngine interface {
	// Saves the item to cache
	Save(id string, started *timestamppb.Timestamp, lastseen *timestamppb.Timestamp)
	// Retrieves item from cache
	Get(id string) Result
	// Retrieve all items from cache
	GetAll() []Result
}

type Time struct {
	Seconds string
	Nanos   string
	Time    string
}

// Result - struct to hold items in cache
type Result struct {
	Id       string
	Started  Time
	Lastseen Time
	Alive string
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
