package main

// 思路：贪心法  正向向外扩展边界

// 时间复杂度：O(n)  空间复杂度：O(1)

func jump(nums []int) int {
	// 如：起始从nums[0]起跳，若nums[0]=3, 则在nums[0]时最远可以到达nums[3]，表示下一次最晚的起跳点在nums[3]
	// 在最晚起跳期间，即在nums[1]~nums[3]之间，可以寻找下一次起跳后可以到达的最远位置，在到达nums[3]时，得到选择出下一次可达最远位置,记做nextMaxPos
	// 在到达nums[3]时做出的跳跃次数增加并不是一定在nums[3]起跳，而是在nums[1]~nums[3]之间的某个位置nums[x]起跳, 在该位置起跳可以到达下一次可达最远位置nextMaxPos，nums[3]只表示上一次跳跃到达的最远位置，在nums[3]之前（包括nums[3])必须做出下一次跳跃
	// 如：[3,8,2,1,5,4] 实际上第一次起跳在nums[0], 下一次起跳可以在[nums[1],nums[3]]之间，但由于在nums[1]起跳可以实现第二跳的最远距离（求该最远距离的过程需要遍历nums[1]~nums[3]，在遍历到nums[3]时，作跳跃次数的增加)，故第二次起跳在nums[1]

	maxPosition := 0  // 可以到达的最远距离，实区间
	stepNum := 0      // 跳跃次数
	jumpPosition := 0 // 上次跳跃后确定的跳跃点
	for i, num := range nums {
		// 在小于上次跳跃后确定的跳跃点前，可以尝试每个位置，来更新下次最远可达距离
		// 更新最远到达距离
		if i+num > maxPosition {
			maxPosition = i + num
		}

		// 注意这里当 i == len(nums)-1 时，表示已经到达数组边界，不需要再跳一次
		if i == jumpPosition && i != len(nums)-1 {
			stepNum++
			jumpPosition = maxPosition
		}
	}

	return stepNum
}
