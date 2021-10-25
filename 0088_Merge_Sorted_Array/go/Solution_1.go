package main

// 归并思想
// 时间复杂度；O(m+n)  空间复杂度：O(m)

func merge(nums1 []int, m int, nums2 []int, n int) {
	temp := make([]int, m)
	for i := 0; i < m; i++ {
		temp[i] = nums1[i]
	}

	i, j := 0, 0
	for k := 0; k < m+n; k++ {
		if i >= m {
			nums1[k] = nums2[j]
			j++
		} else if j >= n {
			nums1[k] = temp[i]
			i++
		} else {
			if nums2[j] < temp[i] {
				nums1[k] = nums2[j]
				j++
			} else {
				nums1[k] = temp[i]
				i++
			}
		}
	}
}
