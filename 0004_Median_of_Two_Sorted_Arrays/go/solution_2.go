package main

// 思路: 二分查找
// https://leetcode.cn/problems/median-of-two-sorted-arrays/solutions/8999/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-w-2/?envType=study-plan-v2&envId=top-100-liked

// 时间复杂度：O(log(min(m, n)))  空间复杂度：O(1)

func findMedianSortedArrays_2(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	// 保证入参的第一个数组长度小于等于第二个数组的长度
	if m > n {
		return findMedianSortedArrays_2(nums2, nums1)
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	iLeft, iRight := 0, len(nums1)
	for iLeft <= iRight {
		i := iLeft + (iRight-iLeft)/2
		j := (m+n+1)/2 - i

		if j != 0 && i != m && nums2[j-1] > nums1[i] {
			iLeft = i + 1

		} else if i != 0 && j != n && nums1[i-1] > nums2[j] {
			iRight = i - 1

		} else {
			leftMax := 0
			if i == 0 {
				leftMax = nums2[j-1]

			} else if j == 0 {
				leftMax = nums1[i-1]

			} else {
				leftMax = max(nums1[i-1], nums2[j-1])
			}

			if (m+n)%2 == 1 {
				return float64(leftMax)
			}

			rightMin := 0
			if i == m {
				rightMin = nums2[j]

			} else if j == n {
				rightMin = nums1[i]

			} else {
				rightMin = min(nums1[i], nums2[j])
			}

			return float64(leftMax+rightMin) / 2
		}
	}

	return 0.0
}
