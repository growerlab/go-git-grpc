package server

import (
	"container/list"
	"sync"
	"time"

	"github.com/google/uuid"
)

const defaultObjectTimeOut = 30 * time.Second

type Object struct {
	key     string
	obj     *EncodedObject // 对象
	timeout int64          // 对象的超时时间点
}

type ObjectStash struct {
	stash *list.List
	cache *sync.Map

	timeout time.Duration

	// 未来可能要改成多级锁提高性能
	mut sync.Mutex
}

func NewObjectStash(objectTimeout time.Duration) *ObjectStash {
	if objectTimeout <= 0 {
		objectTimeout = defaultObjectTimeOut
	}

	return &ObjectStash{
		stash:   list.New(),
		cache:   &sync.Map{},
		timeout: objectTimeout,
		mut:     sync.Mutex{},
	}
}

func (c *ObjectStash) Put(obj *EncodedObject) {
	c.mut.Lock()
	defer c.mut.Unlock()

	key := obj.UUID()
	stashObj := &Object{
		key:     key,
		obj:     obj,
		timeout: time.Now().Add(c.timeout).Unix(),
	}
	ele := c.stash.PushBack(stashObj)
	c.cache.Store(key, ele)
	c.release()
}

func (c *ObjectStash) Get(key string) (*EncodedObject, bool) {
	c.mut.Lock()
	defer c.mut.Unlock()

	ee, ok := c.cache.Load(key)
	if !ok {
		return nil, false
	}
	ele := ee.(*list.Element)
	obj, ok := ele.Value.(*Object)
	if !ok {
		return nil, false
	}
	if obj.timeout < time.Now().Unix() {
		c.del(ele)
		return nil, false
	}
	return obj.obj, true
}

func (c *ObjectStash) release() {
	for e := c.stash.Front(); e != nil; e = e.Next() {
		o, ok := e.Value.(*Object)
		if !ok {
			continue
		}
		if o.timeout < time.Now().Unix() {
			c.del(e)
			continue
		}
		break
	}
}

func (c *ObjectStash) del(e *list.Element) {
	c.stash.Remove(e)
	obj := e.Value.(*Object)
	c.cache.Delete(obj.key)
}

func buildUUID() string {
	return uuid.NewString()
}
