package main

import (
	"container/list"
)

// 单调栈　通过维护一个递减栈来求距离,索引入栈方便直接求距离
// 时间复杂度：O(n)  空间复杂度：O(n)

func dailyTemperatures_2(T []int) []int {

    res := make([]int, len(T))

	// 通过list维护一个单调递减栈,栈内存储索引
	stack := list.New()
	
	for i:=0; i<len(T); i++ {
		
		// 栈顶元素满足出栈条件（遇到后边第一个比它大的元素）时，进行出栈处理（计算距离）
		for stack.Len()>0 && T[i]>T[stack.Back().Value.(int)] { 
			res[stack.Back().Value.(int)] = i-stack.Back().Value.(int)
			stack.Remove(stack.Back())
		}

		stack.PushBack(i)
	}
	
	return res
}