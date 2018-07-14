package designpattern

import (
	"sync"
)

// singleton struct
type singleton struct{}

func (s singleton) String() string {
	return "singleton"
}

var instance *singleton
var once sync.Once

// Singleton instance
func Singleton() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}
