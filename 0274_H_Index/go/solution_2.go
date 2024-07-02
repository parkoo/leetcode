package main

// 思路: 二分法，因为H值最大不能超过文章的篇数也即数组的长度n, 所以H值的范围为[0,n], 可以看做是在[0,n]范围内寻找一个值k, 使得k满足时H值的要求
// 可以使用二分法在[0,n]作搜索

// 时间复杂度: O(nlogn)    空间复杂度: O(1)

func hIndex(citations []int) int {
	n := len(citations)

	// 返回引用因子大于等于target的文章的数量
	getCnt := func(target int) int {
		cnt := 0
		for _, citation := range citations {
			if citation >= target {
				cnt++
			}
		}

		return cnt
	}

	left, right := 0, n
	for left < right {
		// 这里做个+1处理， 使得偶数个数的范围时, mid取中间靠右侧的值, 防止后面因为 left == mid 导致死循环
		mid := left + (right-left)/2 + 1
		cnt := getCnt(mid) // 以mid作为引用因子，大于等于该引用因子的文章数量cnt, 如果该数量过大，cnt>=mid, 说说明引用因子mid 选择小了，可尝试扩大，反之缩小

		// 最终返回的结果是left, left需要是一个有效值
		// 这里是 ‘left = mid、right = mid - 1’ 而不是 ‘left = mid + 1、right = mid’
		if cnt >= mid {
			left = mid

		} else {
			right = mid - 1
		}
	}

	return left
}
