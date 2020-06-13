package main

import (
	"math"
)

// 动态规划  dp_l[i]表示索引i的左边部分的最大高度，dp_r[i]表示索引i的右边部分的最大高度
// 时间复杂度：O(n)  空间复杂度：O(n)

func trap_2(height []int) int {

	res := 0

	dp_l := make([]int, len(height))
	dp_r := make([]int, len(height))

	// 从左往右，生成dp_l[i]
	max_l := 0
	for i:=1; i<len(height); i++ {
		if height[i-1]>max_l {		
			max_l = height[i-1]
		}
		dp_l[i] = max_l
	}

	// 从右往左，生成dp_r[i]
	max_r := 0
	for i:=len(height)-2; i>=0; i-- {
		if height[i+1]>max_r {
			max_r = height[i+1]
		}
		dp_r[i] = max_r
	}

	// 计算每根柱子可以承接的雨水
	for i:=1; i<len(height)-1; i++ {
		// 只有两边的最高高度均大于当前柱子的高度时，当前柱子才可以承接雨水
		if dp_l[i]>height[i] && dp_r[i]>height[i]{ 
			res += int(math.Min(float64(dp_l[i]), float64(dp_r[i]))) - height[i]
		}
	}

	return res
}