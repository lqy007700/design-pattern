package singleton

import (
	"fmt"
	"sync"
)

type Singleton interface {
	foo()
}

type singleton struct{}

func (s singleton) foo() {}

var (
	instance *singleton
	once     sync.Once
)

func New() Singleton {
	once.Do(func() {
		fmt.Println(123)
		instance = &singleton{}
	})

	return instance
}
