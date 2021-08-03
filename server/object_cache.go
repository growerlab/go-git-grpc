package server

import (
	"time"

	"github.com/growerlab/go-git-grpc/common"

	gocache "github.com/patrickmn/go-cache"
)

const (
	defaultObjectTimeOut   = 30 * time.Second
	defaultCleanupInterval = 10 * time.Minute
)

type Object interface {
	UUID() string
}

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

func (c *ObjectCache) Set(obj Object) {
	key := obj.UUID()
	c.cache.SetDefault(key, obj)
}

func (c *ObjectCache) Get(uuid string) (Object, bool) {
	ee, ok := c.cache.Get(uuid)
	if !ok {
		return nil, false
	}
	obj, ok := ee.(Object)
	if !ok {
		return nil, false
	}
	c.Set(obj) // 延期
	return obj, true
}

func buildUUID() string {
	return common.ShortUUID()
}
