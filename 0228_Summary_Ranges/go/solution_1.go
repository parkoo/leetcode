package main

import "strconv"

// 思路: 双指针

// 时间复杂度: O(n)    空间复杂度: O(1)

func summaryRanges_1(nums []int) []string {
	res := make([]string, 0)
	if len(nums) == 0 {
		return res
	}
	if len(nums) == 1 {
		res = append(res, getRange(nums[0], nums[0]))
		return res
	}

	left, right := 0, 1
	for ; right < len(nums); right++ {
		if nums[right] != nums[right-1]+1 {
			item := getRange(nums[left], nums[right-1])
			res = append(res, item)
			left = right
		}
	}

	// 特殊处理最后一段
	item := getRange(nums[left], nums[right-1])
	res = append(res, item)

	return res
}

func getRange(i, j int) string {
	if i == j {
		s := strconv.Itoa(i)
		return s
	}

	s1 := strconv.Itoa(i)
	s2 := strconv.Itoa(j)

	return s1 + "->" + s2
}
