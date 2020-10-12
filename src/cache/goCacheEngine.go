package cache

import (
	"time"

	"github.com/DanMHammer/statusmonitor/status"
	"github.com/patrickmn/go-cache"
)

// GoCacheEngine structure
type GoCacheEngine struct {
	Cache *cache.Cache
}

// Connect - Create Go Cache
func (gc *GoCacheEngine) Connect(minutesToExpire int, minutesToDelete int) (err error) {
	gc.Cache = cache.New(minutesToExpire*time.Minute, minutesToDelete*time.Minute)
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

// SaveStatus - Save Result to Cache
func (gc *GoCacheEngine) SaveStatus(status status.Status) {
	id := status.Id
	ts := str(status.Timestamp)
	gc.Cache.Set(id, ts, cache.DefaultExpiration)
}

// GetStatus - Get Result from Cache
func (gc *GoCacheEngine) GetStatus(id string) status.Status {
	if x, found := gc.Cache.Get(id); found {
		status := x.(Status)
		return status
	}
	return Status{}
}
