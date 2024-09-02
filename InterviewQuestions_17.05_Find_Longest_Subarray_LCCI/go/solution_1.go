package main

// 思路:  前缀和 + 哈希表

// 一个子数组包含的字母和数字的个数相同，等价于该子数组包含的字母和数字的个数之差为 0。
// 因此可以将原数组做转换，每个字母对应 1，每个数字对应 −1，
// 则转换后的数组中，每个子数组的元素和为该子数组对应的原始子数组的字母和数字的个数之差，
// 如果转换后数组中的一个子数组的元素和为 0，则该子数组对应的原始子数组包含的字母和数字的个数相同。
// 问题等价于在转换后的数组中寻找元素和为 0 的最长子数组

// 为了在转换后的数组中寻找元素和为 0 的子数组，
// 可以计算转换后的数组的前缀和，如果两个下标对应的前缀和相等，
// 则这两个下标之间的子数组的元素和为 0。

// 时间复杂度: O(n)    空间复杂度: O(n)

func findLongestSubarray_1(array []string) []string {
	preSum := 0                           // 记录前缀和
	preSumFirstIndex := make(map[int]int) // 记录前缀和第一次出现的位置
	preSumFirstIndex[0] = -1              // array为空数组
	start := -1
	maxLen := 0
	for i, s := range array {

		if s[0] >= '0' && s[0] <= '9' { // 数字为1
			preSum += 1

		} else { // 非数字为-1
			preSum += -1
		}

		if firstIndex, ok := preSumFirstIndex[preSum]; !ok {
			preSumFirstIndex[preSum] = i

		} else {
			if i-firstIndex > maxLen {
				maxLen = i - firstIndex
				start = firstIndex + 1
			}
		}
	}

	if maxLen == 0 {
		return []string{}
	}

	return array[start : start+maxLen]
}
