package tests

import (
	"fmt"
	"sync"
)

func T2(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Hello world")
}
