package main

import (
	"math"
)

// 思路: 寻找中位数可以看做是寻找第k小的数，k=lenth/2
// 递归二分法，每次选择从数组nums1中或者是从数组nums2中排除前k/2个元素，注意处理数组长度不足k/2的情况

// 时间复杂度：O(log(m+n))  空间复杂度：O(1)(尾递归)

func findMedianSortedArrays_1(nums1 []int, nums2 []int) float64 {

	length := len(nums1) + len(nums2)

	//分别令k=(length+1)/2, k=(length+2)/2,然后相加除以2,可以保证奇数长度与偶数长度均可得到中位数
	return float64(findKth(nums1, 0, len(nums1)-1, nums2, 0, len(nums2)-1, (length+1)/2)+
		findKth(nums1, 0, len(nums1)-1, nums2, 0, len(nums2)-1, (length+2)/2)) / 2
}

// 从nums1的[start1, end1]和nums2的[start2, end2]中找到第k小的数
func findKth(nums1 []int, start1 int, end1 int, nums2 []int, start2 int, end2 int, k int) int {

	length1 := end1 - start1 + 1
	length2 := end2 - start2 + 1

	// 递归终止条件
	// 注意处理数组为空的情况
	if length1 <= 0 {
		return nums2[start2+k-1]
	}

	if length2 <= 0 {
		return nums1[start1+k-1]
	}

	if k == 1 {
		return int(math.Min(float64(nums1[start1]), float64(nums2[start2])))
	}

	// 若数组的长度小于k/2,则直接考虑数组最后一个元素
	m := k / 2
	temp1 := start1 + int(math.Min(float64(m), float64(length1))) - 1
	temp2 := start2 + int(math.Min(float64(m), float64(length2))) - 1

	if nums1[temp1] < nums2[temp2] {
		return findKth(nums1, temp1+1, end1, nums2, start2, end2, k-(temp1-start1+1))
	} else {
		return findKth(nums1, start1, end1, nums2, temp2+1, end2, k-(temp2-start2+1))
	}
}
