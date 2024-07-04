package main

// 思路:  从左向右、从右向左两次遍历， 两次结果中取较大的值

// 时间复杂度: O(n)    空间复杂度: O(n)

func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}

	leftRes := make([]int, n)
	for i := 0; i < n; i++ {
		if i > 0 && ratings[i] > ratings[i-1] {
			leftRes[i] = leftRes[i-1] + 1

		} else {
			leftRes[i] = 1
		}
	}

	total := 0
	rightCandy := 0 // 记录从右向左遍历，相邻右侧儿童分得的糖果数，即前驱位置的糖果数
	for i := n - 1; i >= 0; i-- {
		curCandy := 1
		if i < n-1 && ratings[i] > ratings[i+1] {
			curCandy = rightCandy + 1
		}
		rightCandy = curCandy

		if curCandy > leftRes[i] {
			leftRes[i] = curCandy
		}
		total += leftRes[i]
	}

	return total
}
