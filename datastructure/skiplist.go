package datastructure

import (
	"fmt"
	"math"
	"math/rand"
)

const MAX_LEVEL = 16

type SkipListNode struct {
	value interface{}
	level int
	// score for sort
	score int
	// forward pointer on each level
	forwards []*SkipListNode
}

func NewSkiplilstNode(v interface{}, score, level int) *SkipListNode {
	return &SkipListNode{
		value:    v,
		level:    level,
		score:    score,
		forwards: make([]*SkipListNode, level, level),
	}
}

// --- SkipList ---

type SkipList struct {
	head *SkipListNode
	// current level
	// start from 1
	level  int
	length int
}

func NewSkipList() *SkipList {
	head := NewSkiplilstNode(0, math.MinInt32, MAX_LEVEL)
	return &SkipList{
		head:   head,
		level:  1,
		length: 0,
	}
}

// Find skip list node by value
func (sl *SkipList) Find(v interface{}, score int) *SkipListNode {
	if v == nil || sl.length == 0 {
		return nil
	}
	cur := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for cur.forwards[i] != nil {
			if cur.forwards[i].value == v && cur.forwards[i].score == score {
				return cur.forwards[i]
			} else if cur.forwards[i].score > score {
				// jump down
				break
			}
			cur = cur.forwards[i]
		}
	}
	return nil
}

// Insert skip list node, and return level
func (sl *SkipList) Insert(v interface{}, score int) error {
	if v == nil {
		return fmt.Errorf("invalid value")
	}
	cur := sl.head

	// record update node
	// TODO: why max level
	update := [MAX_LEVEL]*SkipListNode{}
	for i := MAX_LEVEL - 1; i >= 0; i-- {
		for cur.forwards[i] != nil {
			if cur.forwards[i].value == v {
				return fmt.Errorf("value %v already exists", v)
			} else if cur.forwards[i].score > score {
				update[i] = cur
				// jump down
				break
			}
			cur = cur.forwards[i]
		}
		// last node on this level
		if cur.forwards[i] == nil {
			update[i] = cur
		}
	}

	kLevel := randomLevel()
	node := NewSkiplilstNode(v, score, kLevel)
	// insert node into all level
	for i := 0; i < kLevel; i++ {
		next := update[i].forwards[i]
		node.forwards[i] = next
		update[i].forwards[i] = node
	}
	// update skip list level
	if kLevel > sl.level {
		sl.level = kLevel
	}
	// update skip list length
	sl.length++

	return nil
}

// Delete skip list node by value
func (sl *SkipList) Delete(v interface{}, score int) {
	if v == nil {
		return
	}
	cur := sl.head
	// record deleting node
	update := make([]*SkipListNode, sl.level)
	for i := sl.level - 1; i >= 0; i-- {
		update[i] = sl.head
		for cur.forwards[i] != nil {
			if cur.forwards[i].value == v && cur.forwards[i].score == score {
				update[i] = cur
				// jump down
				break
			}
			cur = cur.forwards[i]
		}
	}
	// deleting node
	cur = update[0].forwards[0]
	// has not found deleting node
	if cur == nil || cur.value != v {
		return
	}
	// delete node on its each level
	for i := cur.level - 1; i >= 0; i-- {
		// delete the whole level
		if update[i] == sl.head && cur.forwards[i] == nil {
			sl.level = i
		}
		// delete node
		if update[i].forwards[i] != nil && update[i].forwards[i].value == v {
			update[i].forwards[i] = update[i].forwards[i].forwards[i]
		}
	}
	sl.length--
}

func (sl *SkipList) Level() int {
	return sl.level
}

func (sl *SkipList) Length() int {
	return sl.length
}

func (sl *SkipList) String() string {
	return fmt.Sprintf("level: %+v, length: %+v", sl.level, sl.length)
}

func randomLevel() int {
	level := 1
	for i := 1; i < MAX_LEVEL; i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}
	return level
}
