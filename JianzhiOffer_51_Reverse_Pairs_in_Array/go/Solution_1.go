package main

// 归并思想 对两个分别有序的部分归并时，统计出两部分中的逆序数对个数
// 时间复杂度：O(nlgn)  空间复杂度：O(n)
func reversePairs_1(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, l, r int) int {
	if l >= r {
		return 0
	}

	mid := l + (r-l)/2
	lcnt := mergeSort(nums, l, mid)
	rcnt := mergeSort(nums, mid+1, r)
	return merge(nums, l, r, mid) + lcnt + rcnt
}

func merge(nums []int, l, r, mid int) int {
	temp := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		temp[i-l] = nums[i]
	}

	cnt := 0 // 记录逆序数对的个数
	m, n := l, mid+1
	for p := l; p <= r; p++ {
		if m > mid {
			nums[p] = temp[n-l]
			n++
		} else if n > r {
			nums[p] = temp[m-l]
			m++
		} else {
			if temp[m-l] > temp[n-l] {
				nums[p] = temp[m-l]
				m++
				cnt += r - n + 1 // 若 temp[m-l] > temp[n-l] 则 temp[m-l] > temp[i] (n-l<=i<=r)
			} else {
				nums[p] = temp[n-l]
				n++
			}
		}
	}

	return cnt
}
