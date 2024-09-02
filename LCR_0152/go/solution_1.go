package main

import "math"

// 思路: 辅助单调栈
// 二分搜索树的后序遍历的逆序：root > rightTree > leftTree

// 时间复杂度: O(n)    空间复杂度: O(n)

func verifyTreeOrder_1(postorder []int) bool {
	n := len(postorder)

	stack := make([]int, 0)
	// 表示上一个根节点的元素，这里可以把postorder的最后一个元素root看成无穷大节点的左孩子
	preRoot := math.MaxInt

	for i := n - 1; i >= 0; i-- {
		cur := postorder[i]
		if cur > preRoot {
			return false
		}

		// 数组元素小于单调栈的元素了，表示往左子树走了，记录下上个根节点
		for len(stack) > 0 && cur < stack[len(stack)-1] {
			// 找到这个左子树对应的根节点，之前右子树全部弹出，不再记录，因为不可能在往根节点的右子树走了
			preRoot = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		// 新元素入栈
		stack = append(stack, cur)
	}

	return true
}
