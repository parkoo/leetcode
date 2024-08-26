package main

// 思路: 二分法

// 时间复杂度: O(nlogm)    空间复杂度: O(1)

func minEatingSpeed(piles []int, h int) int {
	max := piles[0]
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}

	l, r := 1, max
	for l < r {
		mid := l + (r-l)/2

		time := costTime(piles, mid)
		if time > h {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}

// 计算以speed的速度吃香蕉，吃完所有香蕉需要花费的时间
func costTime(piles []int, speed int) int {
	res := 0
	for _, pile := range piles {
		a, b := pile/speed, pile%speed
		if b == 0 {
			res += a
		} else {
			res = res + a + 1
		}
	}

	return res
}
