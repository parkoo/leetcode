package main

// 双指针 对撞指针
// 时间复杂度：O(n)  空间复杂度：O(1)

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for numbers[i]+numbers[j] != target {
		if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}
	}

	return []int{i + 1, j + 1}
}
