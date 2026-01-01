package main

import (
	"os"
	"sync"

	"some_code/tests"
)

var testNames map[string]bool
var allTests []tests.Test

func init() {
	testNames = make(map[string]bool)
	for _, arg := range os.Args[1:] {
		testNames[arg] = true
	}

	allTests = []tests.Test{
		tests.T1{},
		&tests.T2{},
		tests.T3{},
		tests.T4{},
	}
}

func main() {
	var wg sync.WaitGroup

	for _, test := range allTests {
		if len(testNames) > 0 && !testNames[test.Name()] {
			continue
		}
		wg.Add(1)
		go test.Execute(&wg)
	}

	wg.Wait()
}
