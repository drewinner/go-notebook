package 常用题目

/**
*	https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/
*	定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。
*	MinStack minStack = new MinStack();
*	minStack.push(-2);
*	minStack.push(0);
*	minStack.push(-3);
*	minStack.min();   --> 返回 -3.
*	minStack.pop();
*	minStack.top();   --> 返回  0
*	minStack.min();   --> 返回 -2.
["MinStack","push","push","push","min","pop","top","min"]
[[],[-2],[0],[-3],[],[],[],[]]
[null,null,null,null,-3,null,0,-2]

-3,0,-2


[null,null,null,null,-3,null,0,-3]

*/
type MinStack struct {
	stack []int
	min   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		min:   []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if l := len(this.min); l == 0 || this.min[l-1] >= x {
		this.min = append(this.min, x)
	}

}

//[3,2,1]
func (this *MinStack) Pop() {
	if l := len(this.stack); l > 0 {
		//弹出时、最小栈有影响、如果最小栈里有值、并且栈内弹出的值等于最小栈的最小值、最小栈弹出最小值作为新的最小值
		//this.stack = this.stack[:l-1]
		if m := len(this.min); m > 0 && this.stack[l-1] == this.min[m-1] {
			this.min = this.min[:m-1]
		}
		this.stack = this.stack[:l-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	if len(this.min) == 0 {
		return 0
	}
	return this.min[len(this.min)-1]

}
