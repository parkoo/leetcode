package main

// 思路:  较为通用， 但不是最优
// https://leetcode.cn/problems/implement-rand10-using-rand7/solutions/979495/mo-neng-gou-zao-fa-du-li-sui-ji-shi-jian-9xpz/

// 时间复杂度: O(?)    空间复杂度: O(?)

func rand10_1() int {
	first, second := rand7(), rand7()
	for first > 6 {
		first = rand7()
	}
	for second > 5 {
		second = rand7()
	}

	if first%2 == 0 {
		return second
	} else {
		return 5 + second
	}
}

func rand7() int {
	return 0
}
