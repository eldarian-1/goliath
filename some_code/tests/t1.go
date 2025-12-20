package tests

import (
	"sync"
	"fmt"
	"math/rand"
	"time"
)

func setInt(intChan chan int) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	intChan <- 13
}

func setStr(strChan chan string) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	strChan <- "june"
}

func T1(wg *sync.WaitGroup) {
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
