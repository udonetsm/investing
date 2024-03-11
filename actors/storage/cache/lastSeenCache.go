package cache

import (
	"sync"

	"github.com/udonetsm/investing/models"
)

type lastSeenCache models.Cache

var LastSeenInfoCache lastSeenCache

func CreateLastSeenCache() {
	LastSeenInfoCache = newLastSeenCache()
}

func newLastSeenCache() lastSeenCache {
	return lastSeenCache{
		Mux:     new(sync.Mutex),
		Storage: NewStorage(),
	}
}

func insertLastSeen(lsc *lastSeenCache, key, value string) {
	lsc.Storage[key] = value
}

func (lsc *lastSeenCache) Save(key, value string) error {
	insertLastSeen(lsc, key, value)
	return nil
}
