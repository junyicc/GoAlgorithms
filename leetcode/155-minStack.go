package leetcode

type MinStack struct {
	data []int
	min  []int
}

/** initialize your data structure here. */
func StackConstructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	minLen := len(this.min)
	if minLen == 0 || x < this.min[minLen-1] {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, this.min[minLen-1])
	}
}

func (this *MinStack) Pop() {
	if this.IsEmpty() {
		panic("empty stack")
	}

	this.data = this.data[:len(this.data)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	if this.IsEmpty() {
		panic("empty stack")
	}
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	if this.IsEmpty() {
		panic("empty stack")
	}
	return this.min[len(this.min)-1]
}

func (this *MinStack) IsEmpty() bool {
	return len(this.data) == 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
