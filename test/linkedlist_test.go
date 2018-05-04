package test

import (
	"fmt"
	"testing"

	"algorithm/datastructure"
)

var ll datastructure.LinkedList

func TestAppend(t *testing.T) {
	if !ll.IsEmpty() {
		t.Errorf("list should be empty")
	}

	ll.Append("first")
	if ll.IsEmpty() {
		t.Errorf("list should not be empty")
	}

	if size := ll.Len(); size != 1 {
		t.Errorf("wrong count, expected 1 and got %d", size)
	}

	ll.Append("second")
	ll.Append("third")

	if size := ll.Len(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
}

func TestRemove(t *testing.T) {
	_, err := ll.Remove(1)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}

	if size := ll.Len(); size != 2 {
		t.Errorf("wrong count, expected 2 and got %d", size)
	}
}

func TestInsert(t *testing.T) {
	//test inserting in the middle
	err := ll.Insert(1, "second2")
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
	if size := ll.Len(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}

	//test inserting at head position
	err = ll.Insert(0, "zero")
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
}

func TestHead(t *testing.T) {
	h := ll.Head()
	if "zero" != fmt.Sprint(h.Data) {
		t.Errorf("Expected 'zero' but got %s", fmt.Sprint(h.Data))
	}
}

func TestString(t *testing.T) {
	s := ll.String()
	fmt.Println(s)
	if s != "zero->first->second2->third" {
		t.Errorf("wrong data, expected zero->first->second2->third, and got %s", s)
	}
}

func TestGet(t *testing.T) {
	if size := ll.Len(); size != 3 {
		t.Errorf("wrong count, expected 2 and got %d", size)
	}
	if e, _ := ll.Get(3); *e != "third" {
		t.Errorf("wrong data, expected third, and got %s", *e)
	}
}

func TestIndexOf(t *testing.T) {
	if i, _ := ll.IndexOf("zero"); i != 0 {
		t.Errorf("expected position 0 but got %d", i)
	}
	if i, _ := ll.IndexOf("first"); i != 1 {
		t.Errorf("expected position 1 but got %d", i)
	}
	if i, _ := ll.IndexOf("second2"); i != 2 {
		t.Errorf("expected position 2 but got %d", i)
	}
	if i, _ := ll.IndexOf("third"); i != 3 {
		t.Errorf("expected position 3 but got %d", i)
	}
}
