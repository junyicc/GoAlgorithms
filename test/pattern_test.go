package test

import (
	"testing"

	"github.com/CAIJUNYI/GoAlgorithms/designpattern"
)

func TestSingleton(t *testing.T) {
	s1 := designpattern.Singleton()
	s2 := designpattern.Singleton()
	if s1 != s2 {
		t.Error("failed singleton")
	}
	if result := s1.String(); result != "singleton" {
		t.Errorf("expected singleton and got %s", result)
	}
}

func TestMediator(t *testing.T) {
	wtw := designpattern.WaterTreatmentWork{
		PotableWater: 550,
	}
	wsw := designpattern.WaterStorageWork{
		MaxStorage: 200,
		Storage:    50,
	}
	mediator := designpattern.NewWTWWSWMediator(&wtw, &wsw)
	wsw.OnDemand(mediator, 100)

	if wtw.PotableWater != 450 || wsw.Storage != 150 {
		t.Errorf("failed to transfer water")
	}
}