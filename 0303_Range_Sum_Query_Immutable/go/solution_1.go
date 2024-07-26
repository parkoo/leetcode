package main

// 思路: 前缀和 求区间值

// 时间复杂度: 初始化 O(n)，每次检索 O(1)    空间复杂度: O(n)

type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	preSum := make([]int, len(nums))
	preSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i]
	}

	return NumArray{
		preSum: preSum,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.preSum[right]
	}

	return this.preSum[right] - this.preSum[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
