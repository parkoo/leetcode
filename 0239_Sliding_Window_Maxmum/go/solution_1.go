package main

// 思路: 时代淘汰法，使用单调递减的双端队列维护窗口内的最大值

// 时间复杂度: O(n) 每个元素一次入队一次出队    空间复杂度: O(k) 队列最多存储k个元素

func maxSlidingWindow_1(nums []int, k int) []int {
	res := make([]int, 0)

	// 处理最开始的k个元素
	dq := make([]int, 0) // 双端递减队列，存储窗口内的元素的下标
	for i := 0; i < k; i++ {
		for len(dq) > 0 && nums[i] > nums[dq[len(dq)-1]] { // 保证队列内元素值的单调递减，做队尾淘汰
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)
	}
	res = append(res, nums[dq[0]])

	for i := k; i < len(nums); i++ {
		for len(dq) > 0 && nums[i] > nums[dq[len(dq)-1]] { // 保证队列内元素值的单调递减，做队尾淘汰
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)

		for len(dq) > 0 && dq[0] <= i-k { // 判断队首元素是否还在当前窗口内，不存在则移除，做队首淘汰
			dq = dq[1:]
		}

		res = append(res, nums[dq[0]])
	}

	return res
}
