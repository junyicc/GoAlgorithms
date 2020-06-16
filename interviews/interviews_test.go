package interviews

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

func Test04ReplaceBlank(t *testing.T) {
	s1 := " abc def  g  "
	r1 := ReplaceBlank(s1)
	exp1 := "%20abc%20def%20%20g%20%20"
	if r1 != exp1 {
		t.Errorf("expected %v and got %v", exp1, r1)
	}

	s2 := ""
	r2 := ReplaceBlank(s2)
	exp2 := ""
	if r2 != exp2 {
		t.Errorf("expected %v and got %v", exp2, r2)
	}

	s3 := "abc"
	r3 := ReplaceBlank(s3)
	exp3 := "abc"
	if r3 != exp3 {
		t.Errorf("expected %v and got %v", exp3, r3)
	}
}

func Test05PrintLinkedList(t *testing.T) {
	ll := datastructure.NewLinkedList()

	ll.Append("1")
	ll.Append("2")
	ll.Append("3")
	ll.Append("4")

	PrintLinkedlistReversely(ll)
}

func Test06ReconstructTree(t *testing.T) {
	preorder := []int{8, 4, 2, 1, 3, 6, 5, 7, 10, 9}
	inorder := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	tree := ReConstructBiTree(preorder, inorder)
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
	min1, ok := MinOfRotatedArray(arr1)
	if !ok || min1 != 1 {
		t.Errorf("expected 1 and got %v", min1)
	}

	arr2 := []int{1, 0, 1, 1, 1}
	min2, ok := MinOfRotatedArray(arr2)
	if !ok || min2 != 0 {
		t.Errorf("expected 0 and got %v", min2)
	}

	arr3 := []int{1, 1, 1, 0, 1}
	min3, ok := MinOfRotatedArray(arr3)
	if !ok || min3 != 0 {
		t.Errorf("expected 0 and got %v", min3)
	}

	arr4 := []int{1, 2, 3, 4, 5}
	min4, ok := MinOfRotatedArray(arr4)
	if !ok || min4 != 1 {
		t.Errorf("expected 1 and got %v", min4)
	}
	arr5 := []int{}
	min5, ok := MinOfRotatedArray(arr5)
	if ok || min5 != -1 {
		t.Errorf("expected -1 and got %v", min5)
	}
}

func Test09Fib(t *testing.T) {
	if n, ok := Fib(3); !ok || n != 2 {
		t.Errorf("expected 2 and got %v", n)
	}
}

func Test10NumOf1(t *testing.T) {
	cnt1 := NumOf1(1)
	if cnt1 != 1 {
		t.Errorf("expected 1 and got %v", cnt1)
	}

	cnt2 := NumOf1(0)
	if cnt2 != 0 {
		t.Errorf("expected 0 and got %v", cnt2)
	}

	cnt3 := NumOf1(-1)
	if cnt3 != 64 {
		t.Errorf("expected 64 and got %v", cnt3)
	}

}

func Test10IsPowerOf2(t *testing.T) {
	b1 := IsPowerOf2(4)
	if !b1 {
		t.Errorf("expected true and got %v", b1)
	}

	b2 := IsPowerOf2(0)
	if b2 {
		t.Errorf("expected false and got %v", b2)
	}

	b3 := IsPowerOf2(3)
	if b3 {
		t.Errorf("expected false and got %v", b3)
	}
}

func Test10NumOfChangingBits(t *testing.T) {
	if n1 := NumOfChangingBits(10, 13); n1 != 3 {
		t.Errorf("expected 3 and got %v", n1)
	}

	if n2 := NumOfChangingBits(-1, 1); n2 != 63 {
		t.Errorf("expected 63 and got %v", n2)
	}
}

func Test11Power(t *testing.T) {
	if n1, ok := Power(2.0, 2); !ok || !Equal(n1, 4) {
		t.Errorf("expected 4 and got %v", n1)
	}
	if n2, ok := Power(0, 2); !ok || !Equal(n2, 0) {
		t.Errorf("expected 0 and got %v", n2)
	}
	if n3, ok := Power(-2.0, 3); !ok || !Equal(n3, -8) {
		t.Errorf("expected -8 and got %v", n3)
	}

	if n4, ok := Power(-2.0, -2); !ok || !Equal(n4, 0.25) {
		t.Errorf("expected 0.25 and got %v", n4)
	}
}

func Test12Print1ToMax(t *testing.T) {
	Print1ToMaxOfNDigits(3)
}

func Test13DeleteNode(t *testing.T) {
	l1 := datastructure.ListNode{Data: 1, Next: nil}
	lP := &l1
	d1 := DeleteNode(&lP, lP)
	if *d1 != 1 || lP != nil {
		t.Errorf("expected 1 and got %v", *d1)
	}

	n3 := datastructure.ListNode{Data: 3, Next: nil}
	n2 := datastructure.ListNode{Data: 2, Next: &n3}
	n1 := datastructure.ListNode{Data: 1, Next: &n2}
	n1P := &n1
	if d2 := DeleteNode(&n1P, &n1); *d2 != 1 || n1P.Data != 2 {
		t.Errorf("expected 1 and got %v", *d2)
	}
	if d3 := DeleteNode(&n1P, &n3); *d3 != 3 || n1P.Data != 2 {
		t.Errorf("expected 3 and got %v", *d3)
	}

	if d := DeleteNode(nil, nil); d != nil {
		t.Errorf("expected nil and got %v", d)
	}
}

func Test14ReorderOddEven(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	seq := ReorderOddEven2(arr)
	exp := []int{1, 5, 3, 4, 2}
	for i, e := range seq {
		if e != exp[i] {
			t.Errorf("expected %v and got %v", exp[i], e)
		}
	}
}

func Test15FindKthToTail(t *testing.T) {
	n5 := datastructure.ListNode{Data: 5, Next: nil}
	n4 := datastructure.ListNode{Data: 4, Next: &n5}
	n3 := datastructure.ListNode{Data: 3, Next: &n4}
	n2 := datastructure.ListNode{Data: 2, Next: &n3}
	n1 := datastructure.ListNode{Data: 1, Next: &n2}

	if n := FindKthToTail(&n1, 1); (*n).Data != 5 {
		t.Errorf("expected 5 and got %v", (*n).Data)
	}
	if n := FindKthToTail(&n1, 5); (*n).Data != 1 {
		t.Errorf("expected 1 and got %v", (*n).Data)
	}
	if n := FindKthToTail(&n1, 0); n != nil {
		t.Errorf("expected nil and got %v", n)
	}

	if n := FindKthToTail(&n1, 6); n != nil {
		t.Errorf("expected nil and got %v", n)
	}
	if n := FindKthToTail(nil, 0); n != nil {
		t.Errorf("expected nil and got %v", n)
	}

}
func Test15FindMidNode(t *testing.T) {
	n5 := datastructure.ListNode{Data: 5, Next: nil}
	n4 := datastructure.ListNode{Data: 4, Next: &n5}
	n3 := datastructure.ListNode{Data: 3, Next: &n4}
	n2 := datastructure.ListNode{Data: 2, Next: &n3}
	n1 := datastructure.ListNode{Data: 1, Next: &n2}

	if n := FindMidNode(&n1); (*n).Data != 3 {
		t.Errorf("expected 3 and got %v", (*n).Data)
	}

	n6 := datastructure.ListNode{Data: 6, Next: nil}
	n5.Next = &n6
	if n := FindMidNode(&n1); (*n).Data != 3 {
		t.Errorf("expected 3 and got %v", (*n).Data)
	}
}

func Test15IsLoop(t *testing.T) {
	n5 := datastructure.ListNode{Data: 5, Next: nil}
	n4 := datastructure.ListNode{Data: 4, Next: &n5}
	n3 := datastructure.ListNode{Data: 3, Next: &n4}
	n2 := datastructure.ListNode{Data: 2, Next: &n3}
	n1 := datastructure.ListNode{Data: 1, Next: &n2}
	if b := IsLoop(&n1); b {
		t.Errorf("expected false and got %v", b)
	}

	n5.Next = &n1
	if b := IsLoop(&n1); !b {
		t.Errorf("expected true and got %v", b)
	}
}

func Test16ReverseLinkedlist(t *testing.T) {
	n3 := datastructure.ListNode{Data: 3, Next: nil}
	n2 := datastructure.ListNode{Data: 2, Next: &n3}
	n1 := datastructure.ListNode{Data: 1, Next: &n2}

	if n := ReverseLinkedlist(&n1); n != &n3 || n.Next != &n2 || n.Next.Next != &n1 {
		t.Errorf("failed reversing linkedlist")
	}

	if n := ReverseLinkedlist(&n1); n != &n1 || n.Next != nil {
		t.Errorf("failed reversing linkedlist")
	}

	if n := ReverseLinkedlist(nil); n != nil {
		t.Errorf("failed reversing linkedlist")
	}
}

func Test17MergeLinkedlist(t *testing.T) {
	n4 := datastructure.ListNode{Data: 5, Next: nil}
	n3 := datastructure.ListNode{Data: 4, Next: nil}
	n2 := datastructure.ListNode{Data: 3, Next: &n4}
	n1 := datastructure.ListNode{Data: 2, Next: &n3}

	if n := MergeLinkedlist(&n1, &n2); n != &n1 || n.Next != &n2 || n.Next.Next != &n3 || n.Next.Next.Next != &n4 {
		t.Errorf("failed merging linkedlist")
	}
}

func Test18HasSubtree(t *testing.T) {
	var t1 datastructure.BST
	t1.Insert(8, "8")
	t1.Insert(4, "4")
	t1.Insert(10, "10")
	t1.Insert(2, "2")
	t1.Insert(6, "6")
	t1.Insert(1, "1")

	var t2 datastructure.BST
	t2.Insert(10, "10")
	t2.Insert(2, "2")
	t2.Insert(6, "6")
	// t2 is a subtree of t1
	if has := HasSubtree(t1.Root, t2.Root); !has {
		t.Errorf("expected true and got %v", has)
	}
	// nil pointer
	if has := HasSubtree(nil, nil); has {
		t.Errorf("expected false and got %v", has)
	}

	var t3 datastructure.BST
	t3.Insert(1, "1")
	var t4 datastructure.BST
	t4.Insert(1, "1")
	// only one node
	if has := HasSubtree(t3.Root, t4.Root); !has {
		t.Errorf("expected true and got %v", has)
	}
	// t4 is not a subtree of t3
	t4.Insert(2, "2")
	if has := HasSubtree(t3.Root, t4.Root); has {
		t.Errorf("expected false and got %v", has)
	}
}

func Test19TreeMirror(t *testing.T) {
	var bst datastructure.BST
	bst.Insert(8, "8")
	bst.Insert(4, "4")
	bst.Insert(10, "10")
	bst.Insert(2, "2")

	BiTreeMirror(bst.Root)
	if bst.Root.LChild.Key != 10 || bst.Root.RChild.Key != 4 || bst.Root.RChild.RChild.Key != 2 || bst.Root.RChild.LChild != nil {
		t.Errorf("failed mirror binary tree")
	}
}

func Test20VisitMatrixClockwisely(t *testing.T) {
	m1 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	result := []int{}
	visit := func(e int) {
		result = append(result, e)
	}

	VisitMatrixClockwisely(m1, 4, 4, visit)
	exp1 := []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10}
	for i, e := range result {
		if e != exp1[i] {
			t.Errorf("expected %v and got %v", exp1[i], e)
		}
	}
	// only one row
	m2 := [][]int{
		{1, 2, 3, 4},
	}
	result = []int{}
	exp2 := []int{1, 2, 3, 4}
	VisitMatrixClockwisely(m2, 1, 4, visit)
	for i, e := range result {
		if e != exp2[i] {
			t.Errorf("expected %v and got %v", exp2[i], e)
		}
	}
	// only one col
	m3 := [][]int{
		{1},
		{4},
		{3},
		{2},
	}
	result = []int{}
	exp3 := []int{1, 4, 3, 2}
	VisitMatrixClockwisely(m3, 4, 1, visit)
	for i, e := range result {
		if e != exp3[i] {
			t.Errorf("expected %v and got %v", exp3[i], e)
		}
	}
	// only one entry
	m4 := [][]int{{3}}
	result = []int{}
	exp4 := []int{3}
	VisitMatrixClockwisely(m4, 1, 1, visit)
	for i, e := range result {
		if e != exp4[i] {
			t.Errorf("expected %v and got %v", exp4[i], e)
		}
	}
}

func Test21StackWithMin(t *testing.T) {
	s := StackWithMin{}
	s.Push(3)
	s.Push(4)
	s.Push(2)

	if m := s.Min(); (m).(int) != 2 {
		t.Errorf("expected 2 and got %v", m)
	}
	s.Pop()
	s.Push(1)
	if m := s.Min(); (m).(int) != 1 {
		t.Errorf("expected 1 and got %v", m)
	}
}

func Test22IsPopOrder(t *testing.T) {
	pushOrder := []int{1, 2, 3, 4, 5}
	popOrder := []int{3, 5, 4, 2, 1}

	if order := IsPopOrder(pushOrder, popOrder); !order {
		t.Errorf("expected true and got %v", order)
	}
}

func Test24VerifyPostSeqOfBST(t *testing.T) {
	arr1 := []int{5, 7, 6, 9, 11, 10, 8}
	if b := VerifyPostSeqOfBST(arr1, 7); !b {
		t.Errorf("expected true and got %v", b)
	}

	arr2 := []int{4, 5, 8, 8}
	if b := VerifyPostSeqOfBST(arr2, 4); !b {
		t.Errorf("expected true and got %v", b)
	}

	arr3 := []int{8, 9, 10}
	if b := VerifyPostSeqOfBST(arr3, 3); !b {
		t.Errorf("expected true and got %v", b)
	}
	arr4 := []int{8}
	if b := VerifyPostSeqOfBST(arr4, 1); !b {
		t.Errorf("expected true and got %v", b)
	}

	arr5 := []int{}
	if b := VerifyPostSeqOfBST(arr5, 0); b {
		t.Errorf("expected false and got %v", b)
	}

	if b := VerifyPostSeqOfBST(nil, 0); b {
		t.Errorf("expected false and got %v", b)
	}
}

func Test24VerifyPreSeqOfBST(t *testing.T) {
	arr1 := []int{8, 6, 5, 7, 10, 9, 11}
	if b := VerifyPreSeqOfBST(arr1, 7); !b {
		t.Errorf("expected true and got %v", b)
	}

	arr2 := []int{8, 8, 5, 4}
	if b := VerifyPreSeqOfBST(arr2, 4); !b {
		t.Errorf("expected true and got %v", b)
	}

	arr3 := []int{8, 9, 10}
	if b := VerifyPreSeqOfBST(arr3, 3); !b {
		t.Errorf("expected true and got %v", b)
	}
	arr4 := []int{8}
	if b := VerifyPreSeqOfBST(arr4, 1); !b {
		t.Errorf("expected true and got %v", b)
	}

	arr5 := []int{}
	if b := VerifyPreSeqOfBST(arr5, 0); b {
		t.Errorf("expected false and got %v", b)
	}

	if b := VerifyPreSeqOfBST(nil, 0); b {
		t.Errorf("expected false and got %v", b)
	}
}

func Test25FindPathOfTree(t *testing.T) {
	var bst datastructure.BST
	bst.Insert(10, "10")
	bst.Insert(5, "5")
	bst.Insert(4, "4")
	bst.Insert(7, "7")
	bst.Insert(12, "12")
	FindPathofTree(bst.Root, 22)

	t.Errorf("testing...")
}

func Test26CloneComplexLinkedlist(t *testing.T) {
	var a, b, c, d, e datastructure.ComplexListNode
	a.Data, a.Next, a.Sibling = "A", &b, &c
	b.Data, b.Next, b.Sibling = "B", &c, &e
	c.Data, c.Next, c.Sibling = "C", &d, nil
	d.Data, d.Next, d.Sibling = "D", &e, &b
	e.Data, e.Next, e.Sibling = "E", nil, nil

	clone := CloneComplexLinkedlist(&a)

	for node := &a; clone != nil && node != nil; clone, node = clone.Next, node.Next {
		if clone.Data != node.Data {
			t.Errorf("expected %v and got %v", node.Data, clone.Data)
		}
		if node.Sibling != nil {
			if clone.Sibling.Data != node.Sibling.Data {
				t.Errorf("expected %v and got %v", node.Data, clone.Data)
			}
		}
	}

	a.Data, a.Next, a.Sibling = "A", &b, &c
	b.Data, b.Next, b.Sibling = "B", &c, &e
	c.Data, c.Next, c.Sibling = "C", &d, &c
	d.Data, d.Next, d.Sibling = "D", &e, &b
	e.Data, e.Next, e.Sibling = "E", nil, &b

	clone = CloneComplexLinkedlist(&a)

	for node := &a; clone != nil && node != nil; clone, node = clone.Next, node.Next {
		if clone.Data != node.Data {
			t.Errorf("expected %v and got %v", node.Data, clone.Data)
		}
		if node.Sibling != nil {
			if clone.Sibling.Data != node.Sibling.Data {
				t.Errorf("expected %v and got %v", node.Data, clone.Data)
			}
		}
	}

	o := datastructure.ComplexListNode{Data: "O", Next: nil, Sibling: nil}
	clone = CloneComplexLinkedlist(&o)
	for node := &o; clone != nil && node != nil; clone, node = clone.Next, node.Next {
		if clone.Data != node.Data {
			t.Errorf("expected %v and got %v", node.Data, clone.Data)
		}
		if node.Sibling != nil {
			if clone.Sibling.Data != node.Sibling.Data {
				t.Errorf("expected %v and got %v", node.Data, clone.Data)
			}
		}
	}
}

func Test27ConvertBSTToDoubleLinkedlist(t *testing.T) {
	var bst datastructure.BST
	bst.Insert(10, "10")
	bst.Insert(6, "6")
	bst.Insert(14, "14")
	bst.Insert(4, "4")
	bst.Insert(8, "8")
	bst.Insert(12, "12")
	bst.Insert(16, "16")

	head := BSTToDoubleLinkedlist(bst.Root)
	if head.Key != 4 || head.LChild != nil || head.RChild.Key != 6 {
		t.Errorf("failed converting bst to double linkedlist")
	}

	var bst2 datastructure.BST
	bst2.Insert(10, "10")
	head = BSTToDoubleLinkedlist(bst2.Root)
	if head.Key != 10 || head.LChild != nil || head.RChild != nil {
		t.Errorf("failed converting bst to double linkedlist")
	}

	bst2.Insert(6, "6")
	bst2.Insert(4, "4")
	head = BSTToDoubleLinkedlist(bst2.Root)
	if head.Key != 4 || head.LChild != nil || head.RChild.Key != 6 || head.RChild.RChild.Key != 10 {
		t.Errorf("failed converting bst to double linkedlist")
	}

	var bst3 datastructure.BST
	bst3.Insert(10, "10")
	bst3.Insert(14, "14")
	bst3.Insert(16, "16")
	head = BSTToDoubleLinkedlist(bst3.Root)
	if head.Key != 10 || head.LChild != nil || head.RChild.Key != 14 || head.RChild.RChild.Key != 16 {
		t.Errorf("failed converting bst to double linkedlist")
	}

	head = BSTToDoubleLinkedlist(nil)
	if head != nil {
		t.Errorf("failed converting bst to double linkedlist")
	}
}

func Test40FindTwoNumsAppearOnce(t *testing.T) {
	arr1 := []int{2, 4, 3, 6, 3, 2, 5, 5}
	var n1, n2 int
	ok := FindTwoNumsAppearOnce(arr1, &n1, &n2)
	if !ok || n1 != 6 || n2 != 4 {
		t.Errorf("failed to find two numbers that appear once")
	}
}

func Test28Permutation(t *testing.T) {
	s := "abc"
	visit := func(str string) {
		fmt.Println(str)
	}
	StringPermutation(s, visit)

	n, m := 5, 2
	if r := Permutation(n, m); r != 20 {
		t.Errorf("expected 20 and got %d", r)
	}
}

func Test28Combination(t *testing.T) {
	if c := CombinationIter(10, 2); c != 45 {
		t.Errorf("expected 45 and got %d", c)
	}
}

func Test28NQueens(t *testing.T) {
	if n := SolutionsOfNQueens(4); n != 2 {
		t.Errorf("expected 2 and got %d", n)
	}
}

func Test29FindNumMoreThanHalf(t *testing.T) {
	arr := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	exp := 2
	if n, ok := FindNumMoreThanHalf(arr); !ok || n != exp {
		t.Errorf("expected %d and got %d", exp, n)
	}

	arr = []int{1, 2, 3, 2, 5, 2, 5, 4, 2}
	if n, ok := FindNumMoreThanHalf(arr); ok || n != 0 {
		t.Errorf("expected false and got %t", ok)
	}
}

func Test30FindKthMin(t *testing.T) {
	arr := []int{4, 5, 1, 6, 2, 7, 3, 8}
	k := 4
	if n, ok := FindKthMin(arr, k); !ok || n != 4 {
		t.Errorf("expected 4 and got %d", n)
	}
}

func Test30FindKthMax(t *testing.T) {
	arr := []int{4, 5, 1, 6, 2, 7, 3, 8}
	k := 4
	if n, ok := FindKthMax(arr, k); !ok || n != 5 {
		t.Errorf("expected 5 and got %d", n)
	}
}

func Test30GetKthLeast(t *testing.T) {
	arr := []int{4, 5, 1, 6, 2, 7, 3, 8}
	exp := []int{1, 2, 3, 4}
	k := 4
	result := GetKthLeast(arr, k)
	for i, e := range result {
		if e != exp[i] {
			t.Errorf("expected %d and got %d", exp[i], e)
		}
	}
}

func Test30GetMinKth(t *testing.T) {
	arr := []int{4, 5, 1, 6, 2, 7, 3, 8}
	exp := []int{4, 3, 1, 2}
	k := 4
	result := GetMinKth(arr, k)
	for i, e := range result {
		if e != exp[i] {
			t.Errorf("expected %d and got %d", exp[i], e)
		}
	}
}

func Test30GetMaxKth(t *testing.T) {
	arr := []int{4, 5, 1, 6, 2, 7, 3, 8}
	exp := []int{5, 7, 6, 8}
	k := 4
	result := GetMaxKth(arr, k)
	for i, e := range result {
		if e != exp[i] {
			t.Errorf("expected %d and got %d", exp[i], e)
		}
	}
}

func Test31FindMaxSumOfSubArray(t *testing.T) {
	arr := []int{1, -2, 3, 10, -4, 7, 2, -5}
	if s, ok := FindMaxSumOfSubArray(arr); !ok || s != 18 {
		t.Errorf("expected 18 and got %d", s)
	}

	if s, ok := FindMaxSumOfSubArray([]int{-3}); !ok || s != -3 {
		t.Errorf("expected -3 and got %d", s)
	}
}

func Test31FindSubArrayWithMaxSum(t *testing.T) {
	arr := []int{1, -2, 3, 10, -4, 7, 2, -5}
	exp := []int{3, 10, -4, 7, 2}
	sub := FindSubArrayWithMaxSum(arr)
	for i, e := range sub {
		if e != exp[i] {
			t.Errorf("expected %d and got %d", exp[i], e)
		}
	}

	sub = FindSubArrayWithMaxSum([]int{-3})
	if len(sub) != 1 || sub[0] != -3 {
		t.Errorf("expected {-3} and got %v", sub)
	}
}

func Test32NumOf1(t *testing.T) {
	if n, ok := NumberOf1(5); !ok || n != 1 {
		t.Errorf("expected 1 and got %d", n)
	}
	if n, ok := NumberOf1(12); !ok || n != 5 {
		t.Errorf("expected 5 and got %d", n)
	}
	if n, ok := NumberOf1(0); ok || n != 0 {
		t.Errorf("expected 0 and got %d", n)
	}
	if n, ok := NumberOf1(1); !ok || n != 1 {
		t.Errorf("expected 1 and got %d", n)
	}
	if n, ok := NumberOf1(21345); !ok || n != 18821 {
		t.Errorf("expected 5 and got %d", n)
	}
}

func Test33FindMinPermutation(t *testing.T) {
	arr := []int{3, 32, 321}
	if s := FindMinPermutation(arr); s != "321323" {
		t.Errorf("expected 321323 and got %s", s)
	}

	arr = []int{3}
	if s := FindMinPermutation(arr); s != "3" {
		t.Errorf("expected 321323 and got %s", s)
	}
}

func Test34UglyNum(t *testing.T) {
	if n, ok := GetUglyNum(1); !ok || n != 1 {
		t.Errorf("expected 1 and got %d", n)
	}
	if n, ok := GetUglyNum(0); ok || n != 0 {
		t.Errorf("expected 0 and got %d", n)
	}
	if n, ok := GetUglyNum(11); !ok || n != 15 {
		t.Errorf("expected 15 and got %d", n)
	}
	if n, ok := GetUglyNum(37); !ok || n != 125 {
		t.Errorf("expected 125 and got %d", n)
	}
}

func Test34SuperUglyNum(t *testing.T) {
	n := 12
	primes := []int{2, 7, 13, 19}
	if u, ok := GetSuperUglyNum(n, primes); !ok || u != 32 {
		t.Errorf("expected 32 and got %d", u)
	}
}

func Test35FindFirstNotRepeatingChar(t *testing.T) {
	if c := FindFirstNotRepeatingChar("abaccdeff"); c != 'b' {
		t.Errorf("expected b and got %c", c)
	}
	if c := FindFirstNotRepeatingChar(""); c != 0 {
		t.Errorf("expected 0 and got %c", c)
	}
	if c := FindFirstNotRepeatingChar("s"); c != 's' {
		t.Errorf("expected s and got %c", c)
	}
	if c := FindFirstNotRepeatingChar("a@baccdeff"); c != '@' {
		t.Errorf("expected @ and got %c", c)
	}
}

func Test35DeleteChar(t *testing.T) {
	if s := DeleteChar("We are students", "aeiou"); s != "W r stdnts" {
		t.Errorf("expected 'W r stdnts' and got %s", s)
	}
	if s := DeleteChar("We are students", ""); s != "We are students" {
		t.Errorf("expected 'We are students' and got %s", s)
	}
	if s := DeleteChar("", "asd"); s != "" {
		t.Errorf("expected '' and got %s", s)
	}
}

func Test35DeleteRepeatedChar(t *testing.T) {
	if s := DeleteRepeatedChar("Google"); s != "Gogle" {
		t.Errorf("expected Gogle and got %s", s)
	}
	if s := DeleteRepeatedChar("hooli"); s != "holi" {
		t.Errorf("expected holi and got %s", s)
	}
	if s := DeleteRepeatedChar(""); s != "" {
		t.Errorf("expected '' and got %s", s)
	}
	if s := DeleteRepeatedChar("We are students"); s != "We arstudn" {
		t.Errorf("expected We arstudn and got %s", s)
	}
}

func Test35IsAnagram(t *testing.T) {
	if b := IsAnagram("", ""); !b {
		t.Errorf("expected true and got %t", b)
	}
	if b := IsAnagram("live", ""); b {
		t.Errorf("expected false and got %t", b)
	}
	if b := IsAnagram("livee", "eevil"); !b {
		t.Errorf("expected true and got %t", b)
	}
}

func Test36InversePairs(t *testing.T) {
	if n := InversePairs([]int{}); n != 0 {
		t.Errorf("expected 0 and got %d", n)
	}
	if n := InversePairs(nil); n != 0 {
		t.Errorf("expected 0 and got %d", n)
	}
	if n := InversePairs([]int{7, 5, 6, 4}); n != 5 {
		t.Errorf("expected 5 and got %d", n)
	}
	if n := InversePairs([]int{1, 3, 2, 3, 1}); n != 4 {
		t.Errorf("expected 4 and got %d", n)
	}
}

func Test37FindFirstCommonNode(t *testing.T) {
	n7 := datastructure.ListNode{Data: 7, Next: nil}
	n6 := datastructure.ListNode{Data: 6, Next: &n7}
	n5 := datastructure.ListNode{Data: 5, Next: &n6}
	h2 := datastructure.ListNode{Data: 4, Next: &n5}
	n3 := datastructure.ListNode{Data: 3, Next: &n6}
	n2 := datastructure.ListNode{Data: 2, Next: &n3}
	h1 := datastructure.ListNode{Data: 1, Next: &n2}

	if n := FindFirstCommonNode(&h1, &h2); n.Data != 6 {
		t.Errorf("expected 6 and got %v", n.Data)
	}
}

func Test38CountingK(t *testing.T) {
	if cnt := CountingK([]int{1, 2, 3, 3, 3, 3, 4, 5}, 3); cnt != 4 {
		t.Errorf("expected 4 and got %d", cnt)
	}
	if cnt := CountingK([]int{1, 2, 3, 3, 3, 3, 4, 5}, 6); cnt != 0 {
		t.Errorf("expected 0 and got %d", cnt)
	}
	if cnt := CountingK([]int{2}, 2); cnt != 1 {
		t.Errorf("expected 1 and got %d", cnt)
	}
	if cnt := CountingK([]int{}, 1); cnt != 0 {
		t.Errorf("expected 0 and got %d", cnt)
	}
	if cnt := CountingK(nil, 1); cnt != 0 {
		t.Errorf("expected 0 and got %d", cnt)
	}
}

func Test39TreeDepth(t *testing.T) {
	n7 := datastructure.TreeNode{Key: 7, Data: 7, LChild: nil, RChild: nil}
	n6 := datastructure.TreeNode{Key: 6, Data: 6, LChild: nil, RChild: nil}
	n5 := datastructure.TreeNode{Key: 5, Data: 5, LChild: &n7, RChild: nil}
	n4 := datastructure.TreeNode{Key: 4, Data: 4, LChild: nil, RChild: nil}
	n3 := datastructure.TreeNode{Key: 3, Data: 3, LChild: nil, RChild: &n6}
	n2 := datastructure.TreeNode{Key: 2, Data: 2, LChild: &n4, RChild: &n5}
	n1 := datastructure.TreeNode{Key: 1, Data: 1, LChild: &n2, RChild: &n3}
	root := n1
	if d := TreeDepth(&root); d != 4 {
		t.Errorf("expected 4 and got %d", d)
	}

	n4 = datastructure.TreeNode{Key: 4, Data: 4, LChild: nil, RChild: nil}
	n2 = datastructure.TreeNode{Key: 2, Data: 2, LChild: &n4, RChild: nil}
	n1 = datastructure.TreeNode{Key: 1, Data: 1, LChild: &n2, RChild: nil}
	root = n1
	if d := TreeDepth(&root); d != 3 {
		t.Errorf("expected 3 and got %d", d)
	}

	n1 = datastructure.TreeNode{Key: 1, Data: 1, LChild: nil, RChild: nil}
	root = n1
	if d := TreeDepth(&root); d != 1 {
		t.Errorf("expected 1 and got %d", d)
	}

	if d := TreeDepth(nil); d != 0 {
		t.Errorf("expected 0 and got %d", d)
	}
}

func Test39IsBalanceTree(t *testing.T) {
	n7 := datastructure.TreeNode{Key: 7, Data: 7, LChild: nil, RChild: nil}
	n6 := datastructure.TreeNode{Key: 6, Data: 6, LChild: nil, RChild: nil}
	n5 := datastructure.TreeNode{Key: 5, Data: 5, LChild: &n7, RChild: nil}
	n4 := datastructure.TreeNode{Key: 4, Data: 4, LChild: nil, RChild: nil}
	n3 := datastructure.TreeNode{Key: 3, Data: 3, LChild: nil, RChild: &n6}
	n2 := datastructure.TreeNode{Key: 2, Data: 2, LChild: &n4, RChild: &n5}
	n1 := datastructure.TreeNode{Key: 1, Data: 1, LChild: &n2, RChild: &n3}
	root := n1
	if b := IsBalanceTree(&root); !b {
		t.Errorf("expected true and got %t", b)
	}

	n4 = datastructure.TreeNode{Key: 4, Data: 4, LChild: nil, RChild: nil}
	n2 = datastructure.TreeNode{Key: 2, Data: 2, LChild: &n4, RChild: nil}
	n1 = datastructure.TreeNode{Key: 1, Data: 1, LChild: &n2, RChild: nil}
	root = n1
	if b := IsBalanceTree(&root); b {
		t.Errorf("expected false and got %t", b)
	}

	n1 = datastructure.TreeNode{Key: 1, Data: 1, LChild: nil, RChild: nil}
	root = n1
	if b := IsBalanceTree(&root); !b {
		t.Errorf("expected true and got %t", b)
	}

	if b := IsBalanceTree(nil); b {
		t.Errorf("expected false and got %t", b)
	}
}

func Test41TwoNumWithSum(t *testing.T) {
	arr := []int{1, 2, 4, 7, 11, 15}
	if n1, n2, ok := TwoNumWithSum(arr, 15); !ok {
		t.Errorf("expected true and got %t", ok)
		t.Errorf("two numbers: %d and %d", n1, n2)
	}
	if n1, n2, ok := TwoNumWithSum(arr, 4); ok {
		t.Errorf("expected false and got %t", ok)
		t.Errorf("two numbers: %d and %d", n1, n2)
	}
	if n1, n2, ok := TwoNumWithSum([]int{1}, 1); ok {
		t.Errorf("expected false and got %t", ok)
		t.Errorf("two numbers: %d and %d", n1, n2)
	}
}

func Test41ContinuousSeqWithSum(t *testing.T) {
	visit := func(n1, n2 int) {
		for i := n1; i <= n2; i++ {
			fmt.Printf("%d ", i)
		}
		fmt.Println()
	}
	if found := ContinuousSeqWithSum(0, visit); found {
		t.Errorf("expected false and got %t", found)
	}
	if found := ContinuousSeqWithSum(4, visit); found {
		t.Errorf("expected false and got %t", found)
	}
	if found := ContinuousSeqWithSum(3, visit); !found {
		t.Errorf("expected true and got %t", found)
	}
	if found := ContinuousSeqWithSum(9, visit); !found {
		t.Errorf("expected true and got %t", found)
	}
	if found := ContinuousSeqWithSum(100, visit); !found {
		t.Errorf("expected true and got %t", found)
	}
}

func Test42ReverseString(t *testing.T) {
	s := "I am a student."
	es := "student. a am I"
	if rs := ReverseSentence2(s); rs != es {
		t.Errorf("expected %q and got %q", es, rs)
	}
	if rs := ReverseSentence2(""); rs != "" {
		t.Errorf("expected %q and got %q", "", rs)
	}

	if rs := ReverseSentence(s); rs != es {
		t.Errorf("expected %q and got %q", es, rs)
	}
	if rs := ReverseSentence(""); rs != "" {
		t.Errorf("expected %q and got %q", "", rs)
	}
	if rs := ReverseSentence("  "); rs != "  " {
		t.Errorf("expected %q and got %q", "  ", rs)
	}

	s = "哇哈哈 火箭赢了"
	es = "火箭赢了 哇哈哈"
	if rs := ReverseSentence2(s); rs != es {
		t.Errorf("expected %q and got %q", es, rs)
	}
	if rs := ReverseSentence(s); rs != es {
		t.Errorf("expected %q and got %q", es, rs)
	}
}

func Test42LeftRotateString(t *testing.T) {
	if s := LeftRotateString("我喜欢你", 2); s != "欢你我喜" {
		t.Errorf("expected %q and got %q", "欢你我喜", s)
	}
	if s := LeftRotateString("abcdefg", 2); s != "cdefgab" {
		t.Errorf("expected %q and got %q", "cdefgab", s)
	}
	if s := LeftRotateString("abcdefg", 8); s != "abcdefg" {
		t.Errorf("expected %q and got %q", "abcdefg", s)
	}
	if s := LeftRotateString("abcdefg", 0); s != "abcdefg" {
		t.Errorf("expected %q and got %q", "abcdefg", s)
	}
	if s := LeftRotateString("", 2); s != "" {
		t.Errorf("expected %q and got %q", "", s)
	}
	if s := LeftRotateString("a", 2); s != "a" {
		t.Errorf("expected %q and got %q", "a", s)
	}
}

func Test43PrintDiceProbability(t *testing.T) {
	PrintDiceProbability(6)
}

func Test44IsStraight(t *testing.T) {
	if b := IsStraight([]int{4, 5, 7, 0, 0, 9}); !b {
		t.Errorf("expected true and got %t", b)
	}

	if b := IsStraight([]int{4, 5, 7, 2, 4}); b {
		t.Errorf("expected false and got %t", b)
	}

	if b := IsStraight(nil); b {
		t.Errorf("expected false and got %t", b)
	}
}

func Test45LastReaminingNum(t *testing.T) {
	if n := LastRemaining(5, 3); n != 3 {
		t.Errorf("expected %d and got %d", 3, n)
	}
}

func Test47AddTwoNum(t *testing.T) {
	if n := AddTwoNum(1, 1); n != 2 {
		t.Errorf("expected %d and got %d", 2, n)
	}
	if n := AddTwoNum(1, 0); n != 1 {
		t.Errorf("expected %d and got %d", 0, n)
	}
	if n := AddTwoNum(0, 0); n != 0 {
		t.Errorf("expected %d and got %d", 0, n)
	}
	if n := AddTwoNum(-1, 1); n != 0 {
		t.Errorf("expected %d and got %d", 0, n)
	}
}

func Test49Atoi(t *testing.T) {
	// empty string
	if n, ok := Atoi(""); ok || n != 0 {
		t.Errorf("expected 0, false and got %d, %t", n, ok)
	}
	// not number
	if n, ok := Atoi("!32"); ok || n != 0 {
		t.Errorf("expected 0, false and got %d, %t", n, ok)
	}
	// with symbol +/-
	if n, ok := Atoi("+"); ok || n != 0 {
		t.Errorf("expected 0, false and got %d, %t", n, ok)
	}
	if n, ok := Atoi("+9"); !ok || n != 9 {
		t.Errorf("expected 9, true and got %d, %t", n, ok)
	}
	// exceeding boundary
	if n, ok := Atoi("-9223372036854775808"); !ok || n != -9223372036854775808 {
		t.Errorf("expected -9223372036854775808, true and got %d, %t", n, ok)
	}
	if n, ok := Atoi("9223372036854775807"); !ok || n != 9223372036854775807 {
		t.Errorf("expected 9223372036854775807, true and got %d, %t", n, ok)
	}
	if n, ok := Atoi("9223372036854775808"); ok || n != 0 {
		t.Errorf("expected 0, false and got %d, %t", n, ok)
	}
	if n, ok := Atoi("-9223372036854775809"); ok || n != 0 {
		t.Errorf("expected 0, false and got %d, %t", n, ok)
	}
}

func Test50LastCommonAncestor(t *testing.T) {
	f := datastructure.TreeNode{Key: 1, Data: 1, LChild: nil, RChild: nil}
	g := datastructure.TreeNode{Key: 3, Data: 3, LChild: nil, RChild: nil}
	d := datastructure.TreeNode{Key: 2, Data: 2, LChild: &f, RChild: &g}
	h := datastructure.TreeNode{Key: 5, Data: 6, LChild: nil, RChild: nil}
	i := datastructure.TreeNode{Key: 7, Data: 7, LChild: nil, RChild: nil}
	e := datastructure.TreeNode{Key: 6, Data: 6, LChild: &h, RChild: &i}
	b := datastructure.TreeNode{Key: 4, Data: 4, LChild: &d, RChild: &e}
	c := datastructure.TreeNode{Key: 9, Data: 9, LChild: nil, RChild: nil}
	a := datastructure.TreeNode{Key: 8, Data: 8, LChild: &b, RChild: &c}

	if node := GetLastCommonAncestorForBST(&a, &g, &h); node != &b {
		t.Errorf("expected %v\ngot%v\n", b, *node)
	}

	if node := GetLastCommonAncestorForBST(&a, &g, &c); node != &a {
		t.Errorf("expected %v\ngot%v\n", a, *node)
	}

	if node := GetLastCommonAncestorForBST(&a, &d, &d); node != &d {
		t.Errorf("expected %v\ngot%v\n", d, *node)
	}

	if node := GetLastCommonAncestorForBST(&a, nil, nil); node != nil {
		t.Errorf("expected nil\ngot%v\n", *node)
	}

	if node := GetLastCommonAncestorForTree(&a, &g, &h); node != &b {
		t.Errorf("expected %v\ngot%v\n", b, *node)
	}

	if node := GetLastCommonAncestorForTree(&a, &g, &c); node != &a {
		t.Errorf("expected %v\ngot%v\n", a, *node)
	}

	if node := GetLastCommonAncestorForTree(&a, &d, &d); node != &d {
		t.Errorf("expected %v\ngot%v\n", d, *node)
	}

	if node := GetLastCommonAncestorForTree(&a, nil, nil); node != nil {
		t.Errorf("expected nil\ngot%v\n", *node)
	}
}

func Test51FindRepeatingNum(t *testing.T) {
	if n, ok := FindRepeatingNum([]int{2, 3, 1, 0, 2, 5, 3}); !ok || (n != 2 && n != 3) {
		t.Errorf("expected 2 or 3 and got %d", n)
	}

	if n, ok := FindRepeatingNum(nil); ok || n != -1 {
		t.Errorf("expected nil and got %d", n)
	}

	if n, ok := FindRepeatingNum([]int{0}); ok || n != -1 {
		t.Errorf("expected nil and got %d", n)
	}
}

func Test52Multiply(t *testing.T) {
	arr := []int{1, 3, 2, 1, 5}
	exp := []int{30, 10, 15, 30, 6}
	targetArr := MultiplyArray(arr)
	for i, e := range targetArr {
		if e != exp[i] {
			t.Errorf("expected %d and got %d", exp[i], e)
		}
	}
	arr = []int{1, 3, 2, 0, 5}
	exp = []int{0, 0, 0, 30, 0}
	targetArr = MultiplyArray(arr)
	for i, e := range targetArr {
		if e != exp[i] {
			t.Errorf("expected %d and got %d", exp[i], e)
		}
	}

	arr = []int{1, 0, 2, 0, 5}
	exp = []int{0, 0, 0, 0, 0}
	targetArr = MultiplyArray(arr)
	for i, e := range targetArr {
		if e != exp[i] {
			t.Errorf("expected %d and got %d", exp[i], e)
		}
	}
}

func Test53Regex(t *testing.T) {
	if b := Match("aaa", "a.a"); !b {
		t.Errorf("expected true and got %t", b)
	}
	if b := Match("", "a.a"); b {
		t.Errorf("expected false and got %t", b)
	}
	if b := Match("aaba", "a*.a"); !b {
		t.Errorf("expected true and got %t", b)
	}
	if b := Match("*aa", "**.a"); !b {
		t.Errorf("expected true and got %t", b)
	}
}

func Test54IsNumeric(t *testing.T) {
	if b := IsNumeric("18"); !b {
		t.Errorf("expected true and got %t", b)
	}
	if b := IsNumeric("-18"); !b {
		t.Errorf("expected true and got %t", b)
	}
	if b := IsNumeric("18.58"); !b {
		t.Errorf("expected true and got %t", b)
	}
	if b := IsNumeric("18.58E-2"); !b {
		t.Errorf("expected true and got %t", b)
	}
	if b := IsNumeric("18e2"); !b {
		t.Errorf("expected true and got %t", b)
	}

	if b := IsNumeric(""); b {
		t.Errorf("expected false and got %t", b)
	}
	if b := IsNumeric("2.2a3"); b {
		t.Errorf("expected false and got %t", b)
	}
	if b := IsNumeric("18.3e2.5"); b {
		t.Errorf("expected false and got %t", b)
	}
}

func Test55FirstAppearingOnce(t *testing.T) {
	PrintFirstAppearingOnce()
}

func Test56EntryOfLoop(t *testing.T) {
	n6 := datastructure.ListNode{Data: 6, Next: nil}
	n5 := datastructure.ListNode{Data: 5, Next: &n6}
	n4 := datastructure.ListNode{Data: 4, Next: &n5}
	n3 := datastructure.ListNode{Data: 3, Next: &n4}
	n2 := datastructure.ListNode{Data: 2, Next: &n3}
	n1 := datastructure.ListNode{Data: 1, Next: &n2}
	n6.Next = &n3
	// entry at n3
	if n := FindEntryOfLoop(&n1); n != &n3 {
		t.Errorf("expected %v and got %v", n3, n)
	}
	// break loop
	n6.Next = nil
	if n := FindEntryOfLoop(&n1); n != nil {
		t.Errorf("expected nil and got %v", n)
	}
	// self-loop
	n7 := datastructure.ListNode{Data: 7, Next: nil}
	n7.Next = &n7
	if n := FindEntryOfLoop(&n7); n != &n7 {
		t.Errorf("expected %v and got %v", n7, n)
	}
	// two nodes loop
	n9 := datastructure.ListNode{Data: 9, Next: nil}
	n8 := datastructure.ListNode{Data: 8, Next: &n9}
	n9.Next = &n8
	if n := FindEntryOfLoop(&n8); n != &n8 {
		t.Errorf("expected %v and got %v", n8, n)
	}
}

func Test57DeleteDuplication(t *testing.T) {
	n4 := datastructure.ListNode{Data: 4, Next: nil}
	n33 := datastructure.ListNode{Data: 3, Next: &n4}
	n3 := datastructure.ListNode{Data: 3, Next: &n33}
	n22 := datastructure.ListNode{Data: 2, Next: &n3}
	n2 := datastructure.ListNode{Data: 2, Next: &n22}
	n1 := datastructure.ListNode{Data: 1, Next: &n2}

	n1Ptr := &n1
	DeleteDuplication(&n1Ptr)
	if n1.Next != &n4 {
		t.Errorf("failed deleting duplication")
	}

	n7 := datastructure.ListNode{Data: 7, Next: nil}
	n6 := datastructure.ListNode{Data: 6, Next: &n7}
	n5 := datastructure.ListNode{Data: 5, Next: &n6}

	n5Ptr := &n5
	DeleteDuplication(&n5Ptr)
	if n5.Next != &n6 || n6.Next != &n7 {
		t.Errorf("failed deleting duplication")
	}
}

func Test57UniqueLinkedlist(t *testing.T) {
	n4 := datastructure.ListNode{Data: 4, Next: nil}
	n33 := datastructure.ListNode{Data: 3, Next: &n4}
	n3 := datastructure.ListNode{Data: 3, Next: &n33}
	n22 := datastructure.ListNode{Data: 2, Next: &n3}
	n2 := datastructure.ListNode{Data: 2, Next: &n22}
	n1 := datastructure.ListNode{Data: 1, Next: &n2}

	UniqueLinkedlist(&n1)
	if n1.Next != &n2 || n2.Next != &n3 || n3.Next != &n4 {
		t.Errorf("failed deleting duplication")
	}

	n7 := datastructure.ListNode{Data: 7, Next: nil}
	n6 := datastructure.ListNode{Data: 6, Next: &n7}
	n5 := datastructure.ListNode{Data: 5, Next: &n6}

	UniqueLinkedlist(&n5)
	if n5.Next != &n6 || n6.Next != &n7 {
		t.Errorf("failed deleting duplication")
	}
}

func Test58Successor(t *testing.T) {
	a := TreeNode{Data: "a"}
	b := TreeNode{Data: "b"}
	c := TreeNode{Data: "c"}
	d := TreeNode{Data: "d"}
	e := TreeNode{Data: "e"}
	f := TreeNode{Data: "f"}
	g := TreeNode{Data: "g"}
	h := TreeNode{Data: "h"}
	i := TreeNode{Data: "i"}
	a.LChild = &b
	a.RChild = &c
	b.LChild = &d
	b.RChild = &e
	b.Parent = &a
	c.LChild = &f
	c.RChild = &g
	c.Parent = &a
	d.Parent = &b
	e.Parent = &b
	e.LChild = &h
	e.RChild = &i
	f.Parent = &c
	g.Parent = &c
	h.Parent = &e
	i.Parent = &e

	if n := Succeesor(&a); n != &f {
		t.Errorf("expected %v and got %v", f, *n)
	}
	if n := Succeesor(&i); n != &a {
		t.Errorf("expected %v and got %v", a, *n)
	}
	if n := Succeesor(&d); n != &b {
		t.Errorf("expected %v and got %v", b, *n)
	}
}

func Test59IsSymmertrical(t *testing.T) {
	n5L := datastructure.TreeNode{Data: 5}
	n7L := datastructure.TreeNode{Data: 7}
	n6L := datastructure.TreeNode{Data: 6, LChild: &n5L, RChild: &n7L}
	n5R := datastructure.TreeNode{Data: 5}
	n7R := datastructure.TreeNode{Data: 7}
	n6R := datastructure.TreeNode{Data: 6, LChild: &n7R, RChild: &n5R}
	n8 := datastructure.TreeNode{Data: 8, LChild: &n6L, RChild: &n6R}

	if b := IsSymmertrical(&n8); !b {
		t.Errorf("expected true and got %t", b)
	}
}

func Test60PrintBiTree(t *testing.T) {
	n5L := datastructure.TreeNode{Data: 5}
	n7L := datastructure.TreeNode{Data: 7}
	n6L := datastructure.TreeNode{Data: 6, LChild: &n5L, RChild: &n7L}
	n5R := datastructure.TreeNode{Data: 5}
	n7R := datastructure.TreeNode{Data: 7}
	n6R := datastructure.TreeNode{Data: 6, LChild: &n7R, RChild: &n5R}
	n8 := datastructure.TreeNode{Data: 8, LChild: &n6L, RChild: &n6R}

	PrintBinaryTree(&n8)
}

func Test61PrintBiTreeInZ(t *testing.T) {
	n5L := datastructure.TreeNode{Data: 5}
	n7L := datastructure.TreeNode{Data: 7}
	n6L := datastructure.TreeNode{Data: 6, LChild: &n5L, RChild: &n7L}
	n5R := datastructure.TreeNode{Data: 11}
	n7R := datastructure.TreeNode{Data: 9}
	n6R := datastructure.TreeNode{Data: 10, LChild: &n7R, RChild: &n5R}
	n8 := datastructure.TreeNode{Data: 8, LChild: &n6L, RChild: &n6R}

	PrintBiTreeInZ(&n8)
	PrintBiTreeInZ(&n6L)
	PrintBiTreeInZ(nil)
	PrintBiTreeInZ(&datastructure.TreeNode{})
}

func Test63FindKthNodeInBST(t *testing.T) {
	n5L := datastructure.TreeNode{Data: 5}
	n7L := datastructure.TreeNode{Data: 7}
	n6L := datastructure.TreeNode{Data: 6, LChild: &n5L, RChild: &n7L}
	n5R := datastructure.TreeNode{Data: 11}
	n7R := datastructure.TreeNode{Data: 9}
	n6R := datastructure.TreeNode{Data: 10, LChild: &n7R, RChild: &n5R}
	n8 := datastructure.TreeNode{Data: 8, LChild: &n6L, RChild: &n6R}

	if n := FindKthNodeInBST(&n8, 3); n != &n7L {
		t.Errorf("expected %v\nand got %v\n", n7L, *n)
	}

	if n := FindKthNodeInBST(&n8, 4); n != &n8 {
		t.Errorf("expected %v\nand got %v\n", n8, *n)
	}

	if n := FindKthNodeInBST(&n8, 8); n != nil {
		t.Errorf("expected %v\nand got %v\n", nil, *n)
	}
}

func Test62Serialize(t *testing.T) {
	n4 := datastructure.TreeNode{Data: "4"}
	n5 := datastructure.TreeNode{Data: "5"}
	n6 := datastructure.TreeNode{Data: "6"}
	n2 := datastructure.TreeNode{Data: "2", LChild: &n4}
	n3 := datastructure.TreeNode{Data: "3", LChild: &n5, RChild: &n6}
	n1 := datastructure.TreeNode{Data: "1", LChild: &n2, RChild: &n3}

	var b bytes.Buffer
	Serialize(&n1, &b)

	root := new(datastructure.TreeNode)
	Deserialize(&root, &b)

	if !datastructure.Equal(root.Data, n1.Data) || !datastructure.Equal(root.LChild.Data, n2.Data) || !datastructure.Equal(root.RChild.Data, n3.Data) {
		t.Errorf("failed deserialzing binary tree")
	}
}

func Test64Median(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	da := DynamicArray{}
	da.Init()
	for _, e := range arr {
		da.Insert(e)
	}
	if m, ok := da.Median(); !ok || m != 3 {
		t.Errorf("expected 3 and got %d", m)
	}

	da.Insert(7)
	if m, ok := da.Median(); !ok || m != 3 {
		t.Errorf("expected 3 and got %d", m)
	}
}

func Test66HasPath(t *testing.T) {
	matrix := [][]byte{
		{'a', 'b', 'c', 'e'},
		{'s', 'f', 'c', 's'},
		{'a', 'd', 'e', 'e'},
	}
	if b := HasPath(matrix, 3, 4, "bcced"); !b {
		t.Errorf("expected true and got %t", b)
	}

	if b := HasPath(matrix, 3, 4, "abcb"); b {
		t.Errorf("expected false and got %t", b)
	}
}

func Test68ReverseStr(t *testing.T) {
	s1 := ReverseStr("abcdefg")
	if s1 != "gfedcba" {
		t.Errorf("expected gfedcba and got %s", s1)
	}

	s2 := ReverseStrByStack("abcdefg")
	if s2 != "gfedcba" {
		t.Errorf("expected gfedcba and got %s", s2)
	}

	if s3 := ReverseStrByBit("abcdefg"); s3 != "gfedcba" {
		t.Errorf("expected gfedcba and got %s", s3)
	}
}

func Test69Maze(t *testing.T) {
	nums := [][]int{
		{0, 1, 2},
		{2, 3, -1},
		{1, 3, 1},
	}
	if n := FindShortestPath(nums, 3, 3); n != 7 {
		t.Errorf("expected 7 and got %d\n", n)
	}
}

func Test70City(t *testing.T) {
	n, m := 4, 3
	paths := [][]int{
		{0, 0, 0, 0},
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
	}
	if n := GetImportantCitys(paths, n, m); n != 2 {
		t.Errorf("expected 2 and got %d\n", n)
	}
}

func TestFindNumAppearOnce(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{
			name: "find num appear once",
			args: args{
				arr: []int{2, 3, 3, 3, 5, 5, 5, 7, 7, 7},
			},
			want:  2,
			want1: true,
		},
		{
			name: "find num appear once",
			args: args{
				arr: []int{2, 2, 3, 3, 3, 5, 5, 5, 7, 7, 7},
			},
			want:  0,
			want1: false,
		},
		{
			name: "find num appear once",
			args: args{
				arr: []int{2},
			},
			want:  2,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FindNumAppearOnce(tt.args.arr)
			if got != tt.want {
				t.Errorf("FindNumAppearOnce() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FindNumAppearOnce() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
