package main

// 思路：递归法

// 时间复杂度：O(n)  空间复杂度：O(n)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree_2(preorder []int, inorder []int) *TreeNode {
	return helper(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func helper(preorder []int, inorder []int, preStart, preEnd, inStart, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	rootVal := preorder[preStart]
	rootIndex := inStart
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == rootVal {
			rootIndex = i
		}
	}
	leftNum := rootIndex - inStart
	rightNum := inEnd - rootIndex

	left := helper(preorder, inorder, preStart+1, preStart+leftNum, inStart, rootIndex-1)
	right := helper(preorder, inorder, preStart+leftNum+1, preStart+leftNum+rightNum, rootIndex+1, inEnd)

	node := new(TreeNode)
	node.Val = rootVal
	node.Left = left
	node.Right = right

	return node
}
