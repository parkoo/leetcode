package main

// 思路:  递归

// 时间复杂度: O(n)    空间复杂度: O(n)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	return build(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
}

func build(inorder, postorder []int, inStart, inEnd, postStart, postEnd int) *TreeNode {
	if inStart > inEnd {
		return nil
	}

	rootVal := postorder[postEnd]
	rootIndex := 0
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == rootVal {
			rootIndex = i
		}
	}

	leftCnt := rootIndex - inStart
	rightCnt := inEnd - rootIndex

	left := build(inorder, postorder, inStart, rootIndex-1, postStart, postStart+leftCnt-1)
	right := build(inorder, postorder, rootIndex+1, inEnd, postEnd-rightCnt, postEnd-1)

	root := new(TreeNode)
	root.Val = rootVal
	root.Left = left
	root.Right = right

	return root
}
