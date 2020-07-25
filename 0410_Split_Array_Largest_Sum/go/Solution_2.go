package main

// 二分查找
// 选定一个值 x，可以线性地验证是否存在一种分割方案，满足其最大分割子数组和不超过 x.
// 贪心地模拟分割的过程，从前到后遍历数组，用 sum 表示当前分割子数组的和，cnt 表示已经分割出的子数组的数量（包括当前子数组），
// 那么每当 sum 加上当前值超过了 x，我们就把当前取的值作为新的一段分割子数组的开头，并将 cnt 加 1.

// 时间复杂度：O(n*log(sum-max))  空间复杂度：O(1)  sum表示数组nums中所有元素的和，max表示nums中的最大值

func splitArray_2(nums []int, m int) int {
	n := len(nums)

	// 确定二分查找的区间，l=Max(nums) r=Sum(nums)
	l := nums[0]
	r := 0
	for i:=0; i<n; i++ {
		r += nums[i]
		if nums[i]>l {
			l = nums[i]
		}
	}

	// 二分查找
	for l<r {
		mid := l + (r-l)/2
		num_of_subarrays := getNumOfSubarrays(nums, mid)
		if num_of_subarrays<=m {
			r = mid
		}else {
			l = mid + 1
		}	
	}

	return l
}

// 假设划分出的区间中，区间和的最大值为 x 时，所得出的区间数
func getNumOfSubarrays(arr []int, x int) int {
	cnt := 1
	sum := 0
	for i:=0; i<len(arr); i++ {
		sum = sum + arr[i]
		if sum > x {
			cnt++
			sum = arr[i]
		}
	}

	return cnt
}