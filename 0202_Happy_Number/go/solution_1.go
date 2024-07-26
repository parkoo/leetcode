package main

// 思路:  隐式链表， Floyd判环算法
// 如果不是快乐数，说明出现了环

// 时间复杂度: O(logn)    空间复杂度: O(1)

func isHappy(n int) bool {
	if n == 1 || getNext(n) == 1 {
		return true
	}

	slow, fast := n, getNext(n)

	for slow != fast {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}

	return slow == 1
}

func getNext(n int) int {
	next := 0
	for n > 0 {
		a := n % 10
		n /= 10
		next += a * a
	}

	return next
}
