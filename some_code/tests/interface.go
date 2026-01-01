package tests

import (
	"sync"
)

type Test interface {
	Name() string
	Execute(wg *sync.WaitGroup)
}
