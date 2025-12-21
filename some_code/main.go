package main

import (
	"context"
	"sync"

	"some_code/tests"
)

func main() {
	tests := []tests.Test{
		tests.T1{},
		&tests.T2{},
	}
	var wg sync.WaitGroup
	wg.Add(len(tests))

	for _, test := range tests {
		go test.Execute(&wg)
		test.Log()
	}

	wg.Wait()

	context.Background().Done()
}
