package main

import "strconv"

// 思路: 分组循环法, 这个写法的好处是，无需特判 nums 是否为空，也无需在循环结束后，再补上处理最后一段区间的逻辑
// 适用场景：按照题目要求，数组会被分割成若干段，且每一段的判断/处理逻辑是一样的

// i, n = 0, len(nums)
// while i < n:
//     start = i
//     while i < n and ...:
//         i += 1
//     # 从 start 到 i-1 是一段
//     # 下一段从 i 开始，无需 i+=1

// 时间复杂度: O(n)    空间复杂度: O(1)

func summaryRanges(nums []int) []string {
	res := make([]string, 0)
	n := len(nums)

	i := 0
	for i < n {
		left := i

		i++
		for i < n && nums[i] == nums[i-1]+1 {
			i++
		}

		item := getRange(nums[left], nums[i-1])
		res = append(res, item)
	}

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
