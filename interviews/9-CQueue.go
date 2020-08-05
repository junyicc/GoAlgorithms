package interviews

import "sync"

type CQueue struct {
	first  []int
	second []int
	mu     sync.Mutex
}

func Constructor() CQueue {
	return CQueue{
		first:  make([]int, 0),
		second: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.second = append(this.second, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.first) < 1 && len(this.second) < 1 {
		return -1
	}

	if len(this.first) < 1 {
		this.mu.Lock()
		for len(this.second) > 0 {
			e := this.second[len(this.second)-1]
			this.second = this.second[:len(this.second)-1]
			this.first = append(this.first, e)
		}
		this.mu.Unlock()
	}
	front := this.first[len(this.first)-1]
	this.first = this.first[:len(this.first)-1]
	return front
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
