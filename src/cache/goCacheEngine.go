package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// GoCacheEngine structure
type GoCacheEngine struct {
	Cache *cache.Cache
}

// Connect - Create Go Cache
func (gc *GoCacheEngine) Connect(minutesToExpire int, minutesToDelete int) (err error) {
	gc.Cache = cache.New(time.Duration(minutesToExpire)*time.Minute, time.Duration(minutesToDelete)*time.Minute)
	return
}

// NewGoCacheEngine - Instantiate GoCache
func NewGoCacheEngine(minutesToExpire int, minutesToDelete int) (output *GoCacheEngine, err error) {
	var engine GoCacheEngine
	err = engine.Connect(minutesToExpire, minutesToDelete)
	if err != nil {
		return
	}
	return &engine, nil
}

// Save - Save to Cache
func (gc *GoCacheEngine) Save(id string, item string) {
	gc.Cache.Set(id, item, cache.DefaultExpiration)
}

// Get - Get from Cache
func (gc *GoCacheEngine) Get(id string) string {
	if x, found := gc.Cache.Get(id); found {
		item := x.(string)
		return item
	}
	return ""
}
