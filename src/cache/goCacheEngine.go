package cache

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
	"google.golang.org/protobuf/types/known/timestamppb"
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
func (gc *GoCacheEngine) Save(id string, started *timestamppb.Timestamp, lastseen *timestamppb.Timestamp) {
	startedT := Time{Seconds: fmt.Sprint(started.Seconds), Nanos: fmt.Sprint(started.Nanos), Time: fmt.Sprint(time.Unix(started.Seconds, int64(started.Nanos)))}
	lastseenT := Time{Seconds: fmt.Sprint(lastseen.Seconds), Nanos: fmt.Sprint(lastseen.Nanos), Time: fmt.Sprint(time.Unix(lastseen.Seconds, int64(lastseen.Nanos)))}
	alive := fmt.Sprint(lastseen.Seconds - started.Seconds)
	gc.Cache.Set(id, Result{id, startedT, lastseenT, alive}, cache.DefaultExpiration)
}

// Get - Get from Cache
func (gc *GoCacheEngine) Get(id string) Result {
	if x, found := gc.Cache.Get(id); found {
		item := x.(Result)
		return item
	}
	return Result{}
}

// GetAll - Get all items from cache
func (gc *GoCacheEngine) GetAll() []Result {
	items := gc.Cache.Items()
	res := []Result{}

	for _, item := range items {
		res = append(res, item.Object.(Result))
	}

	fmt.Println(res)

	return res
}
