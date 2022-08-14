package main

// 单调栈 + 哈希表
// 时间复杂度：O(m+n)  空间复杂度：O(n),  m = len(nums1) n = len(nums2)

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	nums1Map := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		res[i] = -1
		nums1Map[nums1[i]] = i
	}

	stack := make([]int, 0)
	for _, num := range nums2 {
		for len(stack) > 0 && num > stack[len(stack)-1] {
			e := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if index, ok := nums1Map[e]; ok {
				res[index] = num
			}
		}

		stack = append(stack, num)
	}

	return res
}
