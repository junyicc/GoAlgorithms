package algo

import (
	"fmt"
	"testing"

	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		str    string
		expect bool
	}{
		{"heooeh", true},
		{"hello", false},
		{"heoeh", true},
		{"a", true},
		{"", false},
	}

	for _, test := range tests {
		l := datastructure.NewLinkedList()
		for _, c := range test.str {
			_ = l.Append(string(c))
		}
		if res := IsPalindrome(l); res != test.expect {
			t.Errorf("test %s: expected %t, got %t", test.str, test.expect, res)
		}
	}
}

var l *datastructure.LinkedList

func init() {
	l = datastructure.NewLinkedList()
	_ = l.Append(5)
	_ = l.Append(4)
	_ = l.Append(3)
	_ = l.Append(2)
	_ = l.Append(1)
}
func TestReverseLinkedList(t *testing.T) {
	fmt.Println(l)
	ReverseLinkedList(l)
	fmt.Println(l)
}

func TestHasCycle(t *testing.T) {
	if res := HasCycle(l); res {
		t.Errorf("expected false, got %t", res)
	}

	l.Head().Next.Next.Next.Next.Next.Next = l.Head().Next.Next
	if res := HasCycle(l); !res {
		t.Errorf("expected true, got %t", res)
	}
}

func TestMergeSortedList(t *testing.T) {
	l1 := datastructure.NewLinkedList()
	_ = l1.Append(1)
	_ = l1.Append(3)
	_ = l1.Append(5)
	_ = l1.Append(7)
	_ = l1.Append(9)

	l2 := datastructure.NewLinkedList()
	_ = l2.Append(2)
	_ = l2.Append(4)
	_ = l2.Append(6)
	_ = l2.Append(8)
	_ = l2.Append(10)

	l := MergeSortedList(l1, l2)
	fmt.Println(l)
}

func TestDeleteBottomN(t *testing.T) {
	fmt.Println(l)
	DeleteBottomN(l, 3)
	fmt.Println(l)
}

func TestFindMiddleNode(t *testing.T) {
	fmt.Println(l)
	fmt.Println(FindMiddleNode(l))

	DeleteBottomN(l, 1)
	fmt.Println(l)
	fmt.Println(FindMiddleNode(l))
}
