package interviews

import (
	"sync"
)

// Singleton struct
type Singleton struct{}

var instance *Singleton
var once sync.Once

// GetSingleton instance
func GetSingleton() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})

	return instance
}
