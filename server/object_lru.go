package server

import (
	"container/list"
	"sync"
)

const (
	DefaultMaxSize = 100 * 10000 // 100w
)

type ObjectLRU struct {
	MaxSize int64

	actualSize int64

	ll    *list.List
	cache *sync.Map
	// TODO 未来可能要改成多级锁提高性能
	mut sync.Mutex
}

func NewObjectLRUDefault() *ObjectLRU {
	return &ObjectLRU{MaxSize: DefaultMaxSize}
}

func (c *ObjectLRU) Put(obj *EncodedObject) {
	c.mut.Lock()
	defer c.mut.Unlock()

	if c.cache == nil {
		c.actualSize = 0
		c.cache = &sync.Map{}
		c.ll = list.New()
	}

	key := obj.UUID()

	if int64(c.ll.Len()) > c.MaxSize {
		oldEle := c.ll.Back()
		c.ll.Remove(oldEle)
		c.cache.Delete(oldEle)
		c.actualSize--
	}

	ele := c.ll.PushFront(obj)
	c.cache.Store(key, ele)
	c.actualSize++
}

func (c *ObjectLRU) Get(k string) (*EncodedObject, bool) {
	c.mut.Lock()
	defer c.mut.Unlock()

	ee, ok := c.cache.Load(k)
	if !ok {
		return nil, false
	}
	ele := ee.(*list.Element)

	c.ll.MoveToFront(ele)
	return ele.Value.(*EncodedObject), true
}

func (c *ObjectLRU) Clear() {
	c.mut.Lock()
	defer c.mut.Unlock()

	c.ll = nil
	c.cache = nil
	c.actualSize = 0
}
