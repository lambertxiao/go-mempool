package mempool

import (
	"sync/atomic"
	"time"
)

type CasLocker struct {
	flag int32
}

func newCasLocker() Locker {
	return &CasLocker{
		flag: 0,
	}
}

func (g *CasLocker) Lock() {
	for {
		if g.flag == 0 && atomic.CompareAndSwapInt32(&g.flag, 0, 1) {
			break
		}

		time.Sleep(1 * time.Nanosecond)
	}
}

func (g *CasLocker) Unlock() {
	g.flag = 0
}
