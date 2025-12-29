package tests

import (
	"fmt"
	"sync"
)

type T2 struct{}

func (_ *T2) Execute(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Hello world")
}

func (_ T2) Log() {
	fmt.Println("T2")
}
