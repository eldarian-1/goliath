package tests

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

const (
	unlocked = false
	locked = true

	mutexRetries = 3

	gorutinesNumber = 1000
)

type MyMutex struct {
	state atomic.Bool
}

func (m *MyMutex) Lock() {
	retries := mutexRetries
	for !m.state.CompareAndSwap(unlocked, locked) {
		if retries == 0 {
			runtime.Gosched()
			retries = mutexRetries
			continue
		}
		retries--
	}
}

func (m *MyMutex) Unlock() {
	m.state.Store(unlocked)
}

type T4 struct{}

func (_ T4) Name() string {
	return "T4"
}

func (_ T4) Execute(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("T4")

	n := 0
	m := MyMutex{}
	subWg := sync.WaitGroup{}
	subWg.Add(gorutinesNumber)

	for i := 0; i < gorutinesNumber; i++ {
		go func() {
			defer subWg.Done()
			m.Lock()
			n++
			m.Unlock()
		}()
	}

	subWg.Wait()

	fmt.Println("Result: ", n)
}
