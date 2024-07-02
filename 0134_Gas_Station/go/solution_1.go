package main

// 思路： 因为每个位置的油量是可以加上由上一个位置剩余的油量, 假设 x<y, 若从位置x可以到达的最远位置为y, 则从x~y之间的任一位置[x,y]出发最远可到达的位置不超过y
// 所以，若从位置x可以到达的最远位置为y，则下一个位置可以从y+1开始尝试, 直接过滤掉[x+1,y]的位置，这样数组中的每个元素只访问一次，时间复杂度O为(n)

// 时间复杂度: O(n)    空间复杂度: O(1)

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)

	// 计算从pos位置出发可以移动的最远距离， 当距离超过n时，说明可以完成移动一圈，可以直接返回
	maxDistance := func(pos int) int {
		distance := 0  // 当前移动距离
		curAmount := 0 // 当前还未在加油站加油的油量

		// remainder表示从当前位置移动到下一个位置后所剩余的油量, 若大于等于0说明可以移动到下一个位置
		remainder := curAmount + gas[pos] - cost[pos]
		for remainder >= 0 {
			pos = (pos + 1) % n
			curAmount = remainder
			remainder = curAmount + gas[pos] - cost[pos]

			distance++
			if distance >= n {
				return distance
			}
		}

		return distance
	}

	for start := 0; start < n; {
		curMaxDistance := maxDistance(start) // 从start开始可以移动的最远距离
		if curMaxDistance >= n {
			return start
		}

		start = start + curMaxDistance + 1
	}

	return -1
}
