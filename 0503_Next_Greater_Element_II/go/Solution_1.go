package main

// 单调栈 + 循环数组
// 时间复杂度：O(n)  空间复杂度：O(n)

func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(res); i++ {
		res[i] = -1
	}

	stack := make([]int, 0)
	for i := 0; i < 2*len(nums); i++ {
		num := nums[i%len(nums)]
		for len(stack) > 0 && num > nums[stack[len(stack)-1]%len(nums)] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[index%len(nums)] = num
		}
		stack = append(stack, i)
	}

	return res
}
