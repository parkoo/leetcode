package main

// 思路: 贪心法

// 时间复杂度: O(n)    空间复杂度: O(n)

func videoStitching(clips [][]int, time int) int {

	// 转换为跳跃游戏的数组形式
	maxn := make([]int, time)
	for _, clip := range clips {
		l, r := clip[0], clip[1]
		if l < time {
			maxn[l] = max(maxn[l], r-l)
		}
	}

	maxPos := 0
	jumpPos := 0
	step := 0
	for i := 0; i < time; i++ {
		num := maxn[i]
		if i > jumpPos {
			return -1
		}

		if i+num > maxPos {
			maxPos = i + num
		}

		if i == jumpPos {
			step++
			jumpPos = maxPos
		}
	}

	if jumpPos >= time {
		return step
	}

	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
