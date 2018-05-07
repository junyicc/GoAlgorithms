package test

import (
	"fmt"
	"testing"

	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

var bst datastructure.BST

func TestBSTInsert(t *testing.T) {
	bst.Insert(8, "8")
	bst.Insert(4, "4")
	bst.Insert(10, "10")
	bst.Insert(2, "2")
	bst.Insert(6, "6")
	bst.Insert(1, "1")
	bst.Insert(3, "3")
	bst.Insert(5, "5")
	bst.Insert(7, "7")
	bst.Insert(9, "9")
}

func TestBSTSearch(t *testing.T) {
	bst.Insert(8, "8")
	bst.Insert(4, "4")
	bst.Insert(10, "10")
	bst.Insert(2, "2")
	bst.Insert(6, "6")
	bst.Insert(1, "1")
	bst.Insert(3, "3")
	bst.Insert(5, "5")
	bst.Insert(7, "7")
	bst.Insert(9, "9")

	if e, ok := bst.Search(9); !ok || (*e).(string) != "9" {
		t.Errorf("expected (true, 9) and got (%v, %v)", ok, *e)
	}
	if e, ok := bst.Search(0); ok || e != nil {
		t.Errorf("expected (false, nil) and got (%v, %v)", ok, *e)
	}
}

func TestBSTRemove(t *testing.T) {
	bst.Insert(8, "8")
	bst.Insert(4, "4")
	bst.Insert(10, "10")
	bst.Insert(2, "2")
	bst.Insert(6, "6")
	bst.Insert(1, "1")
	bst.Insert(3, "3")
	bst.Insert(5, "5")
	bst.Insert(7, "7")
	bst.Insert(9, "9")

	if e, ok := bst.Remove(4); !ok || (*e).(string) != "4" {
		t.Errorf("expected (true, 4) and got (%v, %v)", ok, *e)
	}

	if e, ok := bst.Search(4); ok || e != nil {
		t.Errorf("expected (false, nil) and got (%v, %v)", ok, *e)
	}
}

func TestBSTPreOrderTraverse(t *testing.T) {
	bst.Insert(8, "8")
	bst.Insert(4, "4")
	bst.Insert(10, "10")
	bst.Insert(2, "2")
	bst.Insert(6, "6")
	bst.Insert(1, "1")
	bst.Insert(3, "3")
	bst.Insert(5, "5")
	bst.Insert(7, "7")
	bst.Insert(9, "9")
	var result []string
	expPreOrder := []string{"8", "4", "2", "1", "3", "6", "5", "7", "10", "9"}

	visit := func(e datastructure.Elem) {
		result = append(result, fmt.Sprintf("%v", e))
	}

	bst.PreOrderTraverseRecur(visit)
	if len(result) == 0 {
		t.Errorf("expected %v\n got %v", expPreOrder, result)
	}
	for i, e := range result {
		if expPreOrder[i] != e {
			t.Errorf("expected %v and got %v", expPreOrder[i], e)
		}
	}

	result = []string{}
	bst.PreOrderTraverseIter(visit)
	if len(result) == 0 {
		t.Errorf("expected %v\n got %v", expPreOrder, result)
	}
	for i, e := range result {
		if expPreOrder[i] != e {
			t.Errorf("expected %v and got %v", expPreOrder[i], e)
		}
	}
}

func TestBSTInOrderTraverse(t *testing.T) {
	bst.Insert(8, "8")
	bst.Insert(4, "4")
	bst.Insert(10, "10")
	bst.Insert(2, "2")
	bst.Insert(6, "6")
	bst.Insert(1, "1")
	bst.Insert(3, "3")
	bst.Insert(5, "5")
	bst.Insert(7, "7")
	bst.Insert(9, "9")
	var result []string
	expInOrder := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	visit := func(e datastructure.Elem) {
		result = append(result, fmt.Sprintf("%v", e))
	}

	bst.InOrderTraverseRecur(visit)
	if len(result) == 0 {
		t.Errorf("expected %v\n got %v", expInOrder, result)
	}
	for i, e := range result {
		if expInOrder[i] != e {
			t.Errorf("expected %v and got %v", expInOrder[i], e)
		}
	}

	result = []string{}
	bst.InOrderTraverseIter(visit)
	if len(result) == 0 {
		t.Errorf("expected %v\n got %v", expInOrder, result)
	}
	for i, e := range result {
		if expInOrder[i] != e {
			t.Errorf("expected %v and got %v", expInOrder[i], e)
		}
	}
}

func TestBSTPostOrderTraverse(t *testing.T) {
	bst.Insert(8, "8")
	bst.Insert(4, "4")
	bst.Insert(10, "10")
	bst.Insert(2, "2")
	bst.Insert(6, "6")
	bst.Insert(1, "1")
	bst.Insert(3, "3")
	bst.Insert(5, "5")
	bst.Insert(7, "7")
	bst.Insert(9, "9")
	var result []string
	expPostOrder := []string{"1", "3", "2", "5", "7", "6", "4", "9", "10", "8"}

	visit := func(e datastructure.Elem) {
		result = append(result, fmt.Sprintf("%v", e))
	}

	bst.PostOrderTraverseRecur(visit)
	if len(result) == 0 {
		t.Errorf("expected %v\n got %v", expPostOrder, result)
	}
	for i, e := range result {
		if expPostOrder[i] != e {
			t.Errorf("expected %v and got %v", expPostOrder[i], e)
		}
	}

	result = []string{}
	bst.PostOrderTraverseIter(visit)
	if len(result) == 0 {
		t.Errorf("expected %v\n got %v", expPostOrder, result)
	}
	for i, e := range result {
		if expPostOrder[i] != e {
			t.Errorf("expected %v and got %v", expPostOrder[i], e)
		}
	}
}

func TestBSTLevelTraverse(t *testing.T) {
	bst.Insert(8, "8")
	bst.Insert(4, "4")
	bst.Insert(10, "10")
	bst.Insert(2, "2")
	bst.Insert(6, "6")
	bst.Insert(1, "1")
	bst.Insert(3, "3")
	bst.Insert(5, "5")
	bst.Insert(7, "7")
	bst.Insert(9, "9")
	var result []string
	expLevelOrder := []string{"8", "4", "10", "2", "6", "9", "1", "3", "5", "7"}

	visit := func(e datastructure.Elem) {
		result = append(result, fmt.Sprintf("%v", e))
	}

	bst.LevelTraverse(visit)
	if len(result) == 0 {
		t.Errorf("expected %v\n got %v", expLevelOrder, result)
	}
	for i, e := range result {
		if expLevelOrder[i] != e {
			t.Errorf("expected %v and got %v", expLevelOrder[i], e)
		}
	}

}
