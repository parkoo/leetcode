package main

// 思路: lc0198打家劫舍 变形
// 在选择了元素 x 后，该元素以及所有等于 x−1 或 x+1 的元素会从数组中删去。
// 若还有多个值为 x 的元素，由于所有等于 x−1 或 x+1 的元素已经被删除，
// 我们可以直接删除 x 并获得其点数。因此若选择了 x，所有等于 x 的元素也应一同被选择，以尽可能多地获得点数

// 2,2,3,3,3,4 -> 4为最大值，统计1~4
// 1出现0次，2出现2次，3出现3次，4出现1次 -> sum = [0, 0, 4, 9, 4]
// 构造sumArr, 然后打家劫舍即可

// 时间复杂度: O(m+n)    空间复杂度: O(m)

func deleteAndEarn(nums []int) int {
	// nums中的最大值
	maxVal := 0
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}

	// 构造sumArr
	sumArr := make([]int, maxVal+1)
	for _, num := range nums {
		sumArr[num] += num
	}

	return rob(sumArr)
}

// 数组偷盗
func rob(arr []int) int {
	pre1, pre2 := 0, 0
	res := 0
	for i := 0; i < len(arr); i++ {
		res = max(arr[i]+pre1, pre2)

		pre1 = pre2
		pre2 = res
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
