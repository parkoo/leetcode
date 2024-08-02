package main

// 思路: Floyd判环算法（龟兔赛跑算法）
// 将数组看做链表
// 注意Floyd判环算法的起点一定要在环的外边
// 数组长度为 n+1, 数组中的元素范围在[1,n]
// 由于数字从1开始计数，0这个位置不可能被其它位置的数字指向，所以0位置不可能在环内，所以从0开始遍历一定会指向某个包含重复数字的环，
// 这个题还可以再改改，比如数字从0计数到n-1，那就得从n开始遍历了，少了哪个数就从哪开始遍历
// https://leetcode.cn/problems/find-the-duplicate-number/solutions/58841/287xun-zhao-zhong-fu-shu-by-kirsche/?envType=study-plan-v2&envId=top-100-liked

// 时间复杂度: O(n)    空间复杂度: O(1)

func findDuplicate_2(nums []int) int {
	// slow和fast必须从环外的节点开始，因为数组中的值在[1,n]，故0位置一定在环外
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
