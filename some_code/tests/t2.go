package tests

import (
	"sync"
	"fmt"
)

func T2(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Hello world")
}
