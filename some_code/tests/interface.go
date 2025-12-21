package tests

import (
	"sync"
)

type Test interface {
	Execute(wg *sync.WaitGroup)
	Log()
}
