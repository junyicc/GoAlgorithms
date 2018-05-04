package test

import (
	"algorithm/datastructure"
	"testing"
)

var stack = datastructure.Stack{}

func TestNew(t *testing.T) {
	stack.New(5)
	if stack.IsEmpty() != true {
		t.Errorf("expected true, got %v", stack.IsEmpty())
	}
}

func TestPush(t *testing.T) {
	stack.New(5)
	stack.Push(0)
	if stack.IsEmpty() != false {
		t.Errorf("expected true and got %v", stack.IsEmpty())
	}
	top := stack.GetTop()
	if top != 0 {
		t.Errorf("expected 0 and got %v", top)
	}
}

func TestPop(t *testing.T) {
	stack.New(5)
	stack.Push(0)
	e, err := stack.Pop()
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if e != 0 {
		t.Errorf("expected 0 and got %v", e)
	}
}

func TestClear(t *testing.T) {
	stack.New(5)
	stack.Push(0)
	stack.Push(1)
	stack.Push(2)
	stack.Pop()
	stack.Clear()
	if stack.IsEmpty() != true {
		t.Errorf("expected true and got false")
	}
}
