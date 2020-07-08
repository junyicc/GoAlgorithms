package designpattern

import (
	"fmt"
	"sync"
)

// ProducerWG wait group
// var ProducerWG sync.WaitGroup

// ConsumerWG wait group
// var ConsumerWG sync.WaitGroup

// Producer side
func Producer(name string, data chan int, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Printf("[Producer %s]: producing %d\n", name, i)
		data <- i
	}
	fmt.Println(name, "closed")
	wg.Done()
}

// Consumer side
func Consumer(name string, data chan int, wg *sync.WaitGroup) {
	for i := range data {
		fmt.Printf("[Consumer %s]: consuming %d\n", name, i)
	}
	fmt.Println(name, "closed")
	wg.Done()
}
