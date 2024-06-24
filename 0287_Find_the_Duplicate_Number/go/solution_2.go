package main

// 思路: Floyd判环算法（龟兔赛跑算法）
// 将数组看做链表
// https://leetcode.cn/problems/find-the-duplicate-number/solutions/58841/287xun-zhao-zhong-fu-shu-by-kirsche/?envType=study-plan-v2&envId=top-100-liked

// 时间复杂度: O(n)    空间复杂度: O(1)

func findDuplicate_2(nums []int) int {
	slow, fast := 0, 0
	slow, fast = nums[slow], nums[nums[fast]]

	for slow != fast {
		slow, fast = nums[slow], nums[nums[fast]]
	}

	pre1, pre2 := 0, slow
	for pre1 != pre2 {
		pre1, pre2 = nums[pre1], nums[pre2]
	}

	return pre1 // 返回pre1而不是nums[pre1]
}
