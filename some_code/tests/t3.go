package tests

import (
	"fmt"
	"sync"
)

type T3 struct{}

func (_ T3) Name() string {
	return "T3"
}

func (_ T3) Execute(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("T3")
	log(1, "dwa", T1{}, T2{})
}

func log(args ...interface{}) {
	for i, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Printf("%d) Int value: %d\n", i, arg)
		case string:
			fmt.Printf("%d) Str value: %s\n", i, arg)
		case T1:
			fmt.Printf("%d) T1 value: %s\n", i, arg)
		default:
			fmt.Printf("%d) Unknown type: %w\n", i, arg)
		}
	}
}
