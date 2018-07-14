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
