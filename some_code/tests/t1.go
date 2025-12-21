package tests

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type T1 struct {}

func (_ T1) Execute(wg *sync.WaitGroup) {
	defer wg.Done()

	intChan := make(chan int)
	strChan := make(chan string)

	go setInt(intChan)
	go setStr(strChan)

	for i := 0; i < 2; i++ {
		select {
		case intV := <-intChan:
			fmt.Println("Int value: ", intV)
		case strV := <-strChan:
			fmt.Println("Str value: ", strV)
		}
	}
}

func (_ T1) Log() {
	fmt.Println("T1")
	log(1, "dwa", T1{})
}

func log(args ...interface{}) {
	for i, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Printf("%d) Int value: %d\n", i, arg)
		case string:
			fmt.Printf("%d) Str value: %s\n", i, arg)
		default:
			fmt.Printf("%d) Unknown type: %w\n", i, arg)
		}
	}
}

func setInt(intChan chan int) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	intChan <- 13
}

func setStr(strChan chan string) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	strChan <- "june"
}
