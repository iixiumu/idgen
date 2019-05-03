package main

import "sync/atomic"

type MemAtomic struct {
}

func (MemAtomic) GenID() uint64 {
	return atomic.AddUint64(&id, 1)
}
