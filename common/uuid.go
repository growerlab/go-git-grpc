package common

import (
	"bytes"
	"encoding/binary"
	"sync"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
)

var id *UUID

func init() {
	id = &UUID{
		base: time.Now().UnixNano(),
		pool: &sync.Pool{New: func() interface{} {
			return make([]byte, 16)
		}},
	}
	id.start()
}

type UUID struct {
	base     int64
	fakeRand chan int64

	pool *sync.Pool
}

func (u *UUID) start() {
	u.fakeRand = make(chan int64, 1024)
	go func() {
		for {
			atomic.AddInt64(&u.base, 1)
			u.fakeRand <- u.base
		}
	}()
}

func (u *UUID) Take() string {
	fake := (<-u.fakeRand) << 32 // 左移8位
	b := u.pool.Get().([]byte)
	defer u.pool.Put(b)

	binary.BigEndian.PutUint64(b, uint64(fake))

	newUUID, err := uuid.NewRandomFromReader(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	return newUUID.String()[:8]
}

func ShortUUID8() string {
	return id.Take()
}
