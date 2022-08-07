package main

func isMonotonic(nums []int) bool {
	isIncr := true
	isDecr := true

	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			isDecr = false
		}
		if nums[i+1] < nums[i] {
			isIncr = false
		}
	}

	return isIncr || isDecr
}
