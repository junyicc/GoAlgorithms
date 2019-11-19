package datastructure

import (
	"fmt"
	"strconv"
	"testing"
)

var avl AVL

func TestAVLInsert(t *testing.T) {
	data := []int{3, 2, 1, 4, 5, 6, 7, 10, 9, 8}
	for _, item := range data {
		avl.Insert(item, strconv.Itoa(item))
	}

	var result []string
	expInOrder := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	visit := func(e Elem) {
		result = append(result, fmt.Sprintf("%v", e))
	}

	avl.InOrderTraverseIter(avl.Root, visit)
	if len(result) == 0 {
		t.Errorf("expected %v\n got %v", expInOrder, result)
	}
	for i, e := range result {
		if expInOrder[i] != e {
			t.Errorf("expected %v and got %v", expInOrder[i], e)
		}
	}

}

func TestAVLSearch(t *testing.T) {
	if e, ok := avl.Search(9); !ok || (*e).(string) != "9" {
		t.Errorf("expected (true, 9) and got (%v, %v)", ok, *e)
	}
	if e, ok := avl.Search(0); ok || e != nil {
		t.Errorf("expected (false, nil) and got (%v, %v)", ok, *e)
	}
}
func TestAVLPreOrderTraverse(t *testing.T) {
	var result []string
	expPreOrder := []string{"4", "2", "1", "3", "7", "6", "5", "9", "8", "10"}

	visit := func(e Elem) {
		result = append(result, fmt.Sprintf("%v", e))
	}

	avl.PreOrderTraverseRecur(avl.Root, visit)
	if len(result) == 0 {
		t.Errorf("expected %v\n got %v", expPreOrder, result)
	}
	for i, e := range result {
		if expPreOrder[i] != e {
			t.Errorf("expected %v and got %v", expPreOrder[i], e)
		}
	}

	result = []string{}
	avl.PreOrderTraverseIter(avl.Root, visit)
	if len(result) == 0 {
		t.Errorf("expected %v\n got %v", expPreOrder, result)
	}
	for i, e := range result {
		if expPreOrder[i] != e {
			t.Errorf("expected %v and got %v", expPreOrder[i], e)
		}
	}
}

func TestAVLInOrderTraverse(t *testing.T) {
	var result []string
	expInOrder := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	visit := func(e Elem) {
		result = append(result, fmt.Sprintf("%v", e))
	}

	avl.InOrderTraverseRecur(avl.Root, visit)
	if len(result) == 0 {
		t.Errorf("expected %v\n got %v", expInOrder, result)
	}
	for i, e := range result {
		if expInOrder[i] != e {
			t.Errorf("expected %v and got %v", expInOrder[i], e)
		}
	}

	result = []string{}
	avl.InOrderTraverseIter(avl.Root, visit)
	if len(result) == 0 {
		t.Errorf("expected %v\n got %v", expInOrder, result)
	}
	for i, e := range result {
		if expInOrder[i] != e {
			t.Errorf("expected %v and got %v", expInOrder[i], e)
		}
	}
}

func TestAVLPostOrderTraverse(t *testing.T) {
	var result []string
	expPostOrder := []string{"1", "3", "2", "5", "6", "8", "10", "9", "7", "4"}

	visit := func(e Elem) {
		result = append(result, fmt.Sprintf("%v", e))
	}

	avl.PostOrderTraverseRecur(avl.Root, visit)
	if len(result) == 0 {
		t.Errorf("expected %v\n got %v", expPostOrder, result)
	}
	for i, e := range result {
		if expPostOrder[i] != e {
			t.Errorf("expected %v and got %v", expPostOrder[i], e)
		}
	}

	result = []string{}
	avl.PostOrderTraverseIter(avl.Root, visit)
	if len(result) == 0 {
		t.Errorf("expected %v\n got %v", expPostOrder, result)
	}
	for i, e := range result {
		if expPostOrder[i] != e {
			t.Errorf("expected %v and got %v", expPostOrder[i], e)
		}
	}
}

func TestAVLLevelTraverse(t *testing.T) {
	var result []string
	expLevelOrder := []string{"4", "2", "7", "1", "3", "6", "9", "5", "8", "10"}

	visit := func(e Elem) {
		result = append(result, fmt.Sprintf("%v", e))
	}

	avl.LevelTraverse(avl.Root, visit)
	if len(result) == 0 {
		t.Errorf("expected %v\n got %v", expLevelOrder, result)
	}
	for i, e := range result {
		if expLevelOrder[i] != e {
			t.Errorf("expected %v and got %v", expLevelOrder[i], e)
		}
	}
}
