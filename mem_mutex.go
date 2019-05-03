package main

import (
	"sync"
)

var mutex sync.Mutex

type MemMutex struct {
}

func (MemMutex) GenID() uint64 {
	mutex.Lock()
	defer mutex.Unlock()
	id++
	return id
}
