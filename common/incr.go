package common

import (
	"sync/atomic"
	"time"
)

var incr *Incr

func init() {
	incr = NewIncr()
}

type Incr struct {
	base uint64
}

func NewIncr() *Incr {
	return &Incr{
		base: uint64(time.Now().Unix()),
	}
}

func (i *Incr) GetID() uint64 {
	atomic.AddUint64(&i.base, 1)
}
