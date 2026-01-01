package tests

import (
	"fmt"
	"sync"
)

type T2 struct{}

func (_ T2) Name() string {
	return "T2"
}

func (_ *T2) Execute(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("T2")
	fmt.Println("Hello world")
}
