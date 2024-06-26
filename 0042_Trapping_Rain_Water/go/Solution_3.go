package main

// 双指针 对撞指针，每次根据情况,移动左指针或者右指针
// 时间复杂度：O(n)  空间复杂度：O(１)

func trap_3(height []int) int {
	res := 0

	leftMax, rightMax := height[0], height[len(height)-1]
	left, right := 0, len(height)-1

	// 始终保证 leftMax >= height[left]、 rightMax >= height[right]
	// 因为每次只移动短板一侧的指针，若 height[left] < height[right]，则一定有 leftMax < rightMax
	// 若 height[left] < height[right]，则 leftMax >= height[left] < height[right] <= rightMax, 且 leftMax <= rightMax, 可计算left位置的水量
	// 反之，计算right位置的水量
	for left < right {
		if height[left] > leftMax {
			leftMax = height[left]
		}
		if height[right] > rightMax {
			rightMax = height[right]
		}

		if height[left] < height[right] {
			res += (leftMax - height[left])
			left++

		} else {
			res += (rightMax - height[right])
			right--
		}
	}

	return res
}
