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
type Producer struct {
	Name string
	Data chan int
	Wg   *sync.WaitGroup
}

func (p *Producer) Produce(quit chan struct{}) {
	i := 0
	for {
		select {
		case <-quit:
			return
		default:
			if i < 5 {
				fmt.Printf("[Producer %s]: producing %d\n", p.Name, i)
				p.Data <- i
			} else if i == 5 {
				fmt.Println(p.Name, "finished")
				p.Wg.Done()
			}
			i++
		}
	}
}

// Consumer side
type Consumer struct {
	Name string
	Data chan int
}

func (c *Consumer) Consume() {
	for i := range c.Data {
		fmt.Printf("[Consumer %s]: consuming %d\n", c.Name, i)
	}
	fmt.Println(c.Name, "finished")
}
