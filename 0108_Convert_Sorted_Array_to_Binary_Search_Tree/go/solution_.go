package main

// 思路: 中序遍历构建法

// 时间复杂度: O(n)    空间复杂度: O(log2(n))

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return toBTS(nums, 0, len(nums)-1)
}

// 递归构建
func toBTS(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	midIndex := (left + right) / 2
	curNode := new(TreeNode)
	curNode.Val = nums[midIndex]
	curNode.Left = toBTS(nums, left, midIndex-1)
	curNode.Right = toBTS(nums, midIndex+1, right)

	return curNode
}
