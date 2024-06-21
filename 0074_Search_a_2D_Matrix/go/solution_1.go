package main

// 思路: 合并为一个数组，然后二分查找
// 注：该题目可以同lc240一样，从右上角开始搜索，每次淘汰一整行或一整列

// 时间复杂度: O(log(mn))    空间复杂度: O(1)

func searchMatrix(matrix [][]int, target int) bool {
	nums := make([]int, 0)
	for i := 0; i < len(matrix); i++ {
		nums = append(nums, matrix[i]...)
	}

	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}

	if l >= len(nums) {
		return false
	}
	return nums[l] == target
}
