package test

import (
	"sync"
	"testing"
	"time"
)

var (
	a     *A
	mutex sync.Mutex
	once  sync.Once
)

type A struct {
	N string
}

func TestSingle(t *testing.T) {
	t1 := time.Now()
	for i := 0; i < 100; i++ {
		println(NewAA())
	}
	println(time.Since(t1))
}

func NewA() *A {
	if a == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if a == nil {
			a = &A{}
		}
	}
	return a
}

func NewAA() *A {
	once.Do(func() {
		a = &A{}
	})
	return a
}
