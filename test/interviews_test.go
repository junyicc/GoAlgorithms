package test

import (
	"fmt"
	"testing"

	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
	"github.com/CAIJUNYI/GoAlgorithms/interviews"
)

func Test04ReplaceBlank(t *testing.T) {
	s1 := " abc def  g  "
	r1 := interviews.ReplaceBlank(s1)
	exp1 := "%20abc%20def%20%20g%20%20"
	if r1 != exp1 {
		t.Errorf("expected %v and got %v", exp1, r1)
	}

	s2 := ""
	r2 := interviews.ReplaceBlank(s2)
	exp2 := ""
	if r2 != exp2 {
		t.Errorf("expected %v and got %v", exp2, r2)
	}

	s3 := "abc"
	r3 := interviews.ReplaceBlank(s3)
	exp3 := "abc"
	if r3 != exp3 {
		t.Errorf("expected %v and got %v", exp3, r3)
	}
}

func Test05PrintLinkedList(t *testing.T) {
	var ll datastructure.LinkedList
	ll.Append("1")
	ll.Append("2")
	ll.Append("3")
	ll.Append("4")

	interviews.PrintLinkedlistReversely(&ll)
}

func Test06ReconstructTree(t *testing.T) {
	preorder := []int{8, 4, 2, 1, 3, 6, 5, 7, 10, 9}
	inorder := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	tree := interviews.ReConstructBiTree(preorder, inorder)
	var result []string
	expPostOrder := []string{"1", "3", "2", "5", "7", "6", "4", "9", "10", "8"}

	visit := func(e datastructure.Elem) {
		result = append(result, fmt.Sprintf("%v", e))
	}

	tree.PostOrderTraverseIter(tree.Root, visit)
	if len(result) == 0 {
		t.Errorf("expected %v\n got %v", expPostOrder, result)
	}
	for i, e := range result {
		if expPostOrder[i] != e {
			t.Errorf("expected %v and got %v", expPostOrder[i], e)
		}
	}
}

func Test08MinOfRotatedArray(t *testing.T) {
	arr1 := []int{3, 4, 5, 1, 2}
	min1, ok := interviews.MinOfRotatedArray(arr1)
	if !ok || min1 != 1 {
		t.Errorf("expected 1 and got %v", min1)
	}

	arr2 := []int{1, 0, 1, 1, 1}
	min2, ok := interviews.MinOfRotatedArray(arr2)
	if !ok || min2 != 0 {
		t.Errorf("expected 0 and got %v", min2)
	}

	arr3 := []int{1, 1, 1, 0, 1}
	min3, ok := interviews.MinOfRotatedArray(arr3)
	if !ok || min3 != 0 {
		t.Errorf("expected 0 and got %v", min3)
	}

	arr4 := []int{1, 2, 3, 4, 5}
	min4, ok := interviews.MinOfRotatedArray(arr4)
	if !ok || min4 != 1 {
		t.Errorf("expected 1 and got %v", min4)
	}
	arr5 := []int{}
	min5, ok := interviews.MinOfRotatedArray(arr5)
	if ok || min5 != -1 {
		t.Errorf("expected -1 and got %v", min5)
	}
}

func Test09Fib(t *testing.T) {
	if n, ok := interviews.Fib(3); !ok || n != 2 {
		t.Errorf("expected 2 and got %v", n)
	}
}

func Test10NumOf1(t *testing.T) {
	cnt1 := interviews.NumOf1(1)
	if cnt1 != 1 {
		t.Errorf("expected 1 and got %v", cnt1)
	}

	cnt2 := interviews.NumOf1(0)
	if cnt2 != 0 {
		t.Errorf("expected 0 and got %v", cnt2)
	}

	cnt3 := interviews.NumOf1(-1)
	if cnt3 != 64 {
		t.Errorf("expected 64 and got %v", cnt3)
	}

}

func Test10IsPowerOf2(t *testing.T) {
	b1 := interviews.IsPowerOf2(4)
	if !b1 {
		t.Errorf("expected true and got %v", b1)
	}

	b2 := interviews.IsPowerOf2(0)
	if b2 {
		t.Errorf("expected false and got %v", b2)
	}

	b3 := interviews.IsPowerOf2(3)
	if b3 {
		t.Errorf("expected false and got %v", b3)
	}
}

func Test10NumOfChangingBits(t *testing.T) {
	if n1 := interviews.NumOfChangingBits(10, 13); n1 != 3 {
		t.Errorf("expected 3 and got %v", n1)
	}

	if n2 := interviews.NumOfChangingBits(-1, 1); n2 != 63 {
		t.Errorf("expected 63 and got %v", n2)
	}
}

func Test11Power(t *testing.T) {
	if n1, ok := interviews.Power(2.0, 2); !ok || !interviews.Equal(n1, 4) {
		t.Errorf("expected 4 and got %v", n1)
	}
	if n2, ok := interviews.Power(0, 2); !ok || !interviews.Equal(n2, 0) {
		t.Errorf("expected 0 and got %v", n2)
	}
	if n3, ok := interviews.Power(-2.0, 3); !ok || !interviews.Equal(n3, -8) {
		t.Errorf("expected -8 and got %v", n3)
	}

	if n4, ok := interviews.Power(-2.0, -2); !ok || !interviews.Equal(n4, 0.25) {
		t.Errorf("expected 0.25 and got %v", n4)
	}
}

func Test12Print1ToMax(t *testing.T) {
	interviews.Print1ToMaxOfNDigits(3)
}

func Test13DeleteNode(t *testing.T) {
	l1 := datastructure.Node{Data: 1, Next: nil}
	lP := &l1
	d1 := interviews.DeleteNode(&lP, lP)
	if *d1 != 1 || lP != nil {
		t.Errorf("expected 1 and got %v", *d1)
	}

	n3 := datastructure.Node{Data: 3, Next: nil}
	n2 := datastructure.Node{Data: 2, Next: &n3}
	n1 := datastructure.Node{Data: 1, Next: &n2}
	n1P := &n1
	if d2 := interviews.DeleteNode(&n1P, &n1); *d2 != 1 || n1P.Data != 2 {
		t.Errorf("expected 1 and got %v", *d2)
	}
	if d3 := interviews.DeleteNode(&n1P, &n3); *d3 != 3 || n1P.Data != 2 {
		t.Errorf("expected 3 and got %v", *d3)
	}

	if d := interviews.DeleteNode(nil, nil); d != nil {
		t.Errorf("expected nil and got %v", d)
	}
}

func Test14ReorderOddEven(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	seq := interviews.ReorderOddEven2(arr)
	exp := []int{1, 5, 3, 4, 2}
	for i, e := range seq {
		if e != exp[i] {
			t.Errorf("expected %v and got %v", exp[i], e)
		}
	}
}

func Test15FindKthToTail(t *testing.T) {
	n5 := datastructure.Node{Data: 5, Next: nil}
	n4 := datastructure.Node{Data: 4, Next: &n5}
	n3 := datastructure.Node{Data: 3, Next: &n4}
	n2 := datastructure.Node{Data: 2, Next: &n3}
	n1 := datastructure.Node{Data: 1, Next: &n2}

	if n := interviews.FindKthToTail(&n1, 1); (*n).Data != 5 {
		t.Errorf("expected 5 and got %v", (*n).Data)
	}
	if n := interviews.FindKthToTail(&n1, 5); (*n).Data != 1 {
		t.Errorf("expected 1 and got %v", (*n).Data)
	}
	if n := interviews.FindKthToTail(&n1, 0); n != nil {
		t.Errorf("expected nil and got %v", n)
	}

	if n := interviews.FindKthToTail(&n1, 6); n != nil {
		t.Errorf("expected nil and got %v", n)
	}
	if n := interviews.FindKthToTail(nil, 0); n != nil {
		t.Errorf("expected nil and got %v", n)
	}

}
