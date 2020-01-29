package datastructure

import "testing"

func TestSkipList(t *testing.T) {
	sl := NewSkipList()

	t.Log("----insert-------------------------")
	sl.Insert("leo", 95)
	t.Log(sl.head.forwards[0])
	t.Log(sl.head.forwards[0].forwards[0])
	t.Log(sl)

	t.Log("---insert--------------------------")
	sl.Insert("jack", 88)
	t.Log(sl.head.forwards[0])
	t.Log(sl.head.forwards[0].forwards[0])
	t.Log(sl.head.forwards[0].forwards[0].forwards[0])
	t.Log(sl)

	t.Log("---insert--------------------------")
	sl.Insert("lily", 100)
	t.Log(sl.head.forwards[0])
	t.Log(sl.head.forwards[0].forwards[0])
	t.Log(sl.head.forwards[0].forwards[0].forwards[0])
	t.Log(sl.head.forwards[0].forwards[0].forwards[0].forwards[0])
	t.Log(sl)

	t.Log("---find--------------------------")
	t.Log(sl.Find("jack", 88))

	t.Log("---delete--------------------------")
	sl.Delete("leo", 95)
	t.Log(sl.head.forwards[0])
	t.Log(sl.head.forwards[0].forwards[0])
	t.Log(sl.head.forwards[0].forwards[0].forwards[0])
	t.Log(sl)
}
