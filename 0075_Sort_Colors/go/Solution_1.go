package main

// 三指针  三路快排思想
// 时间复杂度：O(n)  空间复杂度：O(1)

func sortColors(nums []int) {
	// 保证nums[0,i]为0，nums[j,len(nums)-1]为2
	i, j := -1, len(nums)
	k := 0
	for k < j {
		if nums[k] == 0 {
			i++
			nums[i], nums[k] = nums[k], nums[i]
			k++
		} else if nums[k] == 2 {
			j--
			nums[j], nums[k] = nums[k], nums[j]
		} else {
			k++
		}
	}
}
