package main

// 暴力法
// 时间复杂度：O(n^2)  空间复杂度O(1)

func twoSum_1(nums []int, target int) []int {
	
	for i:=0; i<len(nums); i++ {
		for j:=i+1; j<len(nums); j++ {
			if(nums[i] + nums[j] == target){
				return []int{i, j}
			}
		}
	}

	return nil
}
	
