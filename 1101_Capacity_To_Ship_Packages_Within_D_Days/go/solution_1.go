package main

// 思路: 二分法

// 时间复杂度: O(nlog(sum))    空间复杂度: O(1)

func shipWithinDays(weights []int, days int) int {
	sum := 0
	max := weights[0]
	for _, weight := range weights {
		sum += weight
		if weight > max {
			max = weight
		}
	}

	l, r := max, sum
	for l < r {
		mid := l + (r-l)/2

		needDays := costDays(mid, weights)
		if needDays > days {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}

// 计算以cap的装载能力需要多少天
func costDays(cap int, weights []int) int {
	days := 1

	total := 0
	for _, weight := range weights {
		if total+weight > cap {
			days++
			total = weight
		} else {
			total += weight
		}
	}

	return days
}
