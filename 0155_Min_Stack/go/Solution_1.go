package main

// 思路: 通过两个栈实现, 第二个辅助栈存放当前栈内的最小值

// 时间复杂度：O(1)  空间复杂度：O(n)

type MinStack struct {
	stack1 []int // 数据栈
	stack2 []int // 辅助栈
}

func Constructor() MinStack {
	return MinStack{
		stack1: make([]int, 0),
		stack2: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack1 = append(this.stack1, val)
	if len(this.stack2) > 0 && this.stack2[len(this.stack2)-1] < val {
		this.stack2 = append(this.stack2, this.stack2[len(this.stack2)-1])
	} else {
		this.stack2 = append(this.stack2, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack1) > 0 {
		this.stack1 = this.stack1[:len(this.stack1)-1]
		this.stack2 = this.stack2[:len(this.stack2)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack1) > 0 {
		return this.stack1[len(this.stack1)-1]
	}

	return 0
}

func (this *MinStack) GetMin() int {
	if len(this.stack2) > 0 {
		return this.stack2[len(this.stack2)-1]
	}

	return 0
}
