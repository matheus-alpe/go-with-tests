package sync

import (
	"sync"
	"sync/atomic"
)

type Counter interface {
	Inc()
	Value() int
}

type CounterMutex struct {
	mu    sync.Mutex
	value int
}

func NewCounterMutex() *CounterMutex {
	return &CounterMutex{}
}

func (c *CounterMutex) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *CounterMutex) Value() int {
	return c.value
}

type CounterAtomic struct {
	value atomic.Int32
}

func NewCounterAtomic() *CounterAtomic {
	return &CounterAtomic{}
}

func (ca *CounterAtomic) Inc() {
	ca.value.Add(1)
}

func (ca *CounterAtomic) Value() int {
	return int(ca.value.Load())
}
