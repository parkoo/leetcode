package main

// 寻找顺序对
// 时间复杂度：O(n)  空间复杂度：O(1)

func nextPermutation_1(nums []int) {
	if len(nums) > 1 {
		n := len(nums)
		i := n - 2
		// 重后向前 寻找第一个顺序对 (nums[i]<nums[i+1], i<i+1)
		// 此时, nums[i+1],..,nums[n-1]一定是降序的
		for i >= 0 && nums[i] >= nums[i+1] {
			i--
		}
		if i >= 0 {
			// 由于, nums[i+1],..,nums[n-1]一定是降序的
			// 从后向前，寻找第一个大于nums[i]的元素
			j := n - 1
			for j > i && nums[j] <= nums[i] {
				j--
			}
			nums[i], nums[j] = nums[j], nums[i]
		}

		reverse(nums[i+1:])
	}
}

func reverse(arr []int) {
	for i := 0; i <= (len(arr)-1)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}
