package designpattern

import (
	"sync"
	"testing"
)

func TestSingleton(t *testing.T) {
	s1 := Singleton()
	s2 := Singleton()
	if s1 != s2 {
		t.Error("failed singleton")
	}
	if result := s1.String(); result != "singleton" {
		t.Errorf("expected singleton and got %s", result)
	}
}

func TestMediator(t *testing.T) {
	wtw := WaterTreatmentWork{
		PotableWater: 550,
	}
	wsw := WaterStorageWork{
		MaxStorage: 200,
		Storage:    50,
	}
	mediator := NewWTWWSWMediator(&wtw, &wsw)
	wsw.OnDemand(mediator, 100)

	if wtw.PotableWater != 450 || wsw.Storage != 150 {
		t.Errorf("failed to transfer water")
	}
}

func TestProducerConsumer(t *testing.T) {
	data := make(chan int, 5)
	var producerWg sync.WaitGroup

	// Consumer
	c1 := Consumer{
		Name: "C1",
		Data: data,
	}
	c2 := Consumer{
		Name: "C2",
		Data: data,
	}
	go c1.Consume()
	go c2.Consume()

	// produce
	quit := make(chan struct{})
	p1 := Producer{
		Name: "P1",
		Data: data,
		Wg:   &producerWg,
	}
	p2 := Producer{
		Name: "P2",
		Data: data,
		Wg:   &producerWg,
	}
	producerWg.Add(2)
	go p1.Produce(quit)
	go p2.Produce(quit)

	producerWg.Wait()
	// notify producer to quit
	quit <- struct{}{}
	close(data)
}
