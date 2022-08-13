package main

// 单调栈 + 贪心

func maxWidthRamp(nums []int) int {
	res := 0
	stack := make([]int, 0)
	// 以 nums[0] 开始生成单调递减栈，坡底必在其中。
	// 反证：若存在 (k, j) 可形成最大坡度，则在 k 之前不存在比 k 位置上得元素更小的元素，否则会出现更大的坡度,
	//      而这样的 k 必定会被压入栈
	// 如：[6, 9, 1, 2, 3, 4] 栈中元素坐标为 [0, 2]
	for i, num := range nums {
		if len(stack) == 0 || num < nums[stack[len(stack)-1]] {
			stack = append(stack, i)
		}
	}
	// 贪心思路，为了使得坡度最大，从末端遍历求坡度，此时以栈顶元素为坡底的坡度是栈顶元素所能形成的最大坡度，可以将其出栈
	for j := len(nums) - 1; j >= 0; j-- {
		for len(stack) > 0 && nums[j] >= nums[stack[len(stack)-1]] {
			ramp := j - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if ramp > res {
				res = ramp
			}
		}
	}
	return res
}
