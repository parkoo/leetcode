package main

// 双指针-滑动窗口,  窗口长度固定为k, [l, (l+k-1)%n]
// 时间复杂度：O(n)  空间复杂度：O(1)

func maxScore_1(cardPoints []int, k int) int {
	maxScore := 0

	n := len(cardPoints)
	l := n - k
	r := l + k - 1
	curScore := 0
	for i := l; i <= r; i++ {
		curScore += cardPoints[i]
		maxScore = curScore
	}

	for i := 0; i < k; i++ {
		r = (r + 1) % n
		curScore = curScore - cardPoints[l] + cardPoints[r]
		l = (l + 1) % n
		if curScore > maxScore {
			maxScore = curScore
		}
	}

	return maxScore
}
