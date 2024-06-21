package main

// 思路: 栈内存放元素的差值 通过一个栈实现

// 时间复杂度：O(1)  空间复杂度：O(1)

type MinStack_2 struct {
	innerStack []int64 // 内部普通栈，入栈元素与最小值的差值
	minVal     int64   // 维护当前栈内的元素的最小值
}

func Constructor() MinStack_2 {
	return MinStack_2{
		innerStack: make([]int64, 0),
		minVal:     0,
	}
}

func (this *MinStack_2) Push(val int) {
	// 栈为空时
	if len(this.innerStack) == 0 {
		this.minVal = int64(val)
		this.innerStack = append(this.innerStack, int64(0))
		return
	}

	// 栈不为空时
	// 先将元素(当前元素与当前栈内最小值的差值,记为curVal)入栈，入栈后再更行栈内元素的最小值minVal，此两步不可交换顺序
	// 如此，若栈顶元素curVal小于或等于0，表示此栈顶元素curVal便是当前的站内最小值minVal，即此时栈顶元素的原始值val==minVal
	// 若此时将该元素出栈，则需要将栈内最小值更新为上个栈内最小值lastMinVal，如何得到lastMinVal呢
	// curVal = val - lastMinVal  =>  lastMinVal = val - curVal,
	// 而 val == minVal, 所以, lastMinVal =  minVal - curVal
	// 即，上一个栈内最小值为当前栈内最小值与当前栈顶元素之差

	curVal := int64(val) - this.minVal
	this.innerStack = append(this.innerStack, curVal)

	if int64(val) < this.minVal {
		this.minVal = int64(val)
	}
}

func (this *MinStack_2) Pop() {
	if len(this.innerStack) == 0 {
		return
	}

	cur := this.innerStack[len(this.innerStack)-1]
	this.innerStack = this.innerStack[:len(this.innerStack)-1]

	if cur <= 0 { // 出栈元素为栈内最小值对应元素，出栈后需要更新栈内最小值
		lastMinVal := this.minVal - cur
		this.minVal = lastMinVal
	}
}

func (this *MinStack_2) Top() int {
	cur := this.innerStack[len(this.innerStack)-1]
	if cur <= 0 { // 此时的栈顶元素对应栈内最小值元素， 直接返回栈内最小值即可
		return int(this.minVal)
	}

	return int(this.innerStack[len(this.innerStack)-1] + this.minVal)
}

func (this *MinStack_2) GetMin() int {
	return int(this.minVal)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
