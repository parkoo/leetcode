package main

// 分治法  线段树思想（区间问题）
// 时间复杂度：O(lgn)  空间复杂度：O(lgn)

//  对于一个区间[l, r], 维护四个变量以便进行两个区间的合并
//  lSum: 区间[l, r]内以l为左端点的最大子段和
//  rSum: 区间[l, r]内以r为右端点的最大子段和
//  iSum: 区间[l, r]内所有数的和
//  mSum: 区间[l, r]内的最大子段和

func maxSubArray_3(nums []int) int {
	_, _, _, msum := sub(nums, 0, len(nums)-1)
	return msum
}

func sub(nums []int, l, r int) (int, int, int, int) {
	if l >= r {
		return nums[l], nums[l], nums[l], nums[l]
	}

	mid := l + (r-l)/2
	lmax1, rmax1, isum1, msum1 := sub(nums, l, mid)
	lmax2, rmax2, isum2, msum2 := sub(nums, mid+1, r)

	// 通过四个变量进行向上的区间合并
	lmax := max(lmax1, isum1+lmax2)
	rmax := max(rmax2, isum2+rmax1)
	isum := isum1 + isum2
	msum := max(max(msum1, msum2), rmax1+lmax2)
	return lmax, rmax, isum, msum
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
