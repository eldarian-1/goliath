package main

import (
	"context"
	"sync"

	"some_code/tests"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go tests.T1(&wg)
	go tests.T2(&wg)

	wg.Wait()

	context.Background().Done()
}
