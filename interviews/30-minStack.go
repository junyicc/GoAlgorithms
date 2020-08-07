package interviews

// 定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。
//
// 示例:
//
// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.min();   --> 返回 -3.
// minStack.pop();
// minStack.top();      --> 返回 0.
// minStack.min();   --> 返回 -2.

type MinStack struct {
	data    []int
	minData []int
}

/** initialize your data structure here. */
func MinStackConstructor() MinStack {
	return MinStack{
		data:    make([]int, 0),
		minData: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	minVal := x
	if len(this.minData) > 0 && this.minData[len(this.minData)-1] < minVal {
		minVal = this.minData[len(this.minData)-1]
	}
	this.minData = append(this.minData, minVal)
	this.data = append(this.data, x)
}

func (this *MinStack) Pop() {
	if len(this.data) > 0 {
		this.data = this.data[:len(this.data)-1]
		this.minData = this.minData[:len(this.minData)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.data) > 0 {
		return this.data[len(this.data)-1]
	}
	return 0
}

func (this *MinStack) Min() int {
	if len(this.data) > 0 {
		return this.minData[len(this.minData)-1]
	}
	return 0
}
