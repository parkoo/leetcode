package main

// 思路:  栈 通过栈记录当前重复的字母的个数

// 迭代字符串：
// 如果当前字符与前一个相同，栈顶元素加 1。
// 否则，往栈中压入 1。
// 如果栈顶元素等于 k，则从字符串中删除这 k 个字符，并将 k 从栈顶移除。

// 时间复杂度: O(n)    空间复杂度: O(n)

func removeDuplicates_1(s string, k int) string {
	res := make([]rune, 0)

	ss := []rune(s)
	stack := make([]int, 0)

	for _, c := range ss {
		if len(res) > 0 && c == res[len(res)-1] {
			stack[len(stack)-1] = stack[len(stack)-1] + 1

		} else {
			stack = append(stack, 1)
		}
		res = append(res, c)

		top := stack[len(stack)-1]
		if top == k {
			stack = stack[:len(stack)-1]
			res = res[:len(res)-k]
		}
	}

	return string(res)
}
