package datastructure

import "testing"

func TestLRUCache(t *testing.T) {
	lru := NewLRUCache(2)
	lru.Put("1", 1)
	lru.Put("2", 2)
	if v := lru.Get("2"); v.(int) != 2 {
		t.Errorf("expected 2, got %v", v)
	}
	if v := lru.Get("3"); v != nil {
		t.Errorf("expected nil, got %v", v)
	}
	// deletie 1, and add 3
	lru.Put("3", 3)
	if v := lru.Get("3"); v.(int) != 3 {
		t.Errorf("expected 3, got %v", v)
	}
}
