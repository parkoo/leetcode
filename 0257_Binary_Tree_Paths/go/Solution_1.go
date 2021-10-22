package main

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 回溯法

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)

	var backtrack func(root *TreeNode, subStr string)
	backtrack = func(root *TreeNode, subStr string) {
		if root.Left == nil && root.Right == nil {
			subStr = fmt.Sprintf("%s->%d", subStr, root.Val)
			subStr = subStr[2:] // 去除最前边的"->"
			res = append(res, subStr)
			return
		}

		if root.Left != nil {
			backtrack(root.Left, fmt.Sprintf("%s->%d", subStr, root.Val))
		}

		if root.Right != nil {
			backtrack(root.Right, fmt.Sprintf("%s->%d", subStr, root.Val))
		}
	}

	backtrack(root, "")
	return res
}
