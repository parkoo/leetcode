package main

// 思路: 同solution_1, 时代淘汰法，使用单调递减的双端队列维护窗口内的最大值

// 时间复杂度: O(n) 每个元素一次入队一次出队    空间复杂度: O(k) 队列最多存储k个元素

func maxSlidingWindow_2(nums []int, k int) []int {
	res := make([]int, 0)

	dq := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		// 保证队列内元素值的单调递减，做队尾淘汰
		for len(dq) > 0 && nums[dq[len(dq)-1]] < nums[i] {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)

		// 前k-1个元素不做处理
		if i < k-1 {
			continue
		}

		// 判断队首元素是否还在当前窗口内，不存在则移除，做队首淘汰
		for i-dq[0]+1 > k {
			dq = dq[1:]
		}
		res = append(res, nums[dq[0]])
	}

	return res
}
