package server

import (
	"time"

	"github.com/google/uuid"

	gocache "github.com/patrickmn/go-cache"
)

const (
	defaultObjectTimeOut   = 30 * time.Second
	defaultCleanupInterval = 10 * time.Minute
)

type ObjectCache struct {
	cache *gocache.Cache
}

func NewObjectCache(objectTimeout time.Duration) *ObjectCache {
	if objectTimeout <= 0 {
		objectTimeout = defaultObjectTimeOut
	}

	return &ObjectCache{
		cache: gocache.New(objectTimeout, defaultCleanupInterval),
	}
}

func (c *ObjectCache) Set(obj *EncodedObject) {
	key := obj.UUID()
	c.cache.SetDefault(key, obj)
}

func (c *ObjectCache) Get(uuid string) (*EncodedObject, bool) {
	ee, ok := c.cache.Get(uuid)
	if !ok {
		return nil, false
	}
	obj, ok := ee.(*EncodedObject)
	if !ok {
		return nil, false
	}
	return obj, true
}

func buildUUID() string {
	t, n, _ := uuid.GetTime()
	s := t.UnixTime()
	b := s & n
	uint16
	return
}
