package datastructure

import (
	"testing"
)

func TestNew(t *testing.T) {
	stack, err := NewStack(5)
	if err != nil {
		panic(err)
	}
	if stack.IsEmpty() != true {
		t.Errorf("expected true, got %v", stack.IsEmpty())
	}
}

func TestPush(t *testing.T) {
	stack, err := NewStack(5)
	if err != nil {
		panic(err)
	}
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
	stack, err := NewStack(5)
	if err != nil {
		panic(err)
	}
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
	stack, err := NewStack(5)
	if err != nil {
		panic(err)
	}
	stack.Push(0)
	stack.Push(1)
	stack.Push(2)
	stack.Pop()
	stack.Clear()
	if stack.IsEmpty() != true {
		t.Errorf("expected true and got false")
	}
}
