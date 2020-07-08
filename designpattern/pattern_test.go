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
	var consumerWg sync.WaitGroup

	producerWg.Add(1)
	go Producer("P1", data, &producerWg)
	consumerWg.Add(1)
	go Consumer("U1", data, &consumerWg)
	producerWg.Add(1)
	go Producer("P2", data, &producerWg)
	consumerWg.Add(1)
	go Consumer("U2", data, &consumerWg)

	producerWg.Wait()
	close(data)
	consumerWg.Wait()
}
