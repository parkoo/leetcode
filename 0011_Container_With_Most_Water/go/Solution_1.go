package main

import "math"

// 双指针之对撞指针
// 时间复杂度：O(n)  空间复杂度：O(1)

func maxArea_1(height []int) int {

	max_area := 0
	i := 0
	j := len(height)-1

	for i<j{
		area := int(math.Min(float64(height[i]), float64(height[j]))) * (j-i)
		if area>max_area {
			max_area = area
		}

		if height[i]<height[j]{  // 选择短板向内扩展
			i++
		}else {
			j--
		}
	}

	return max_area
}