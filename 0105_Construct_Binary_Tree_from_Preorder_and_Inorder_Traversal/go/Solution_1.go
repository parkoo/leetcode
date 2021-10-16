package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归法
// 时间复杂度：O(n)  空间复杂度：O(n)

func buildTree_1(preorder []int, inorder []int) *TreeNode {
	return helper(preorder, inorder)
}

func helper(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 && len(inorder) == 0 {
		return nil
	}

	mark := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			mark = i
			break
		}
	}

	root := &TreeNode{Val: inorder[mark]}
	root.Left = helper(preorder[1:1+mark], inorder[:mark])
	root.Right = helper(preorder[1+mark:], inorder[mark+1:])

	return root
}
