package main

import (
	"strconv"
	"strings"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的构建与遍历 前序实现
// 时间复杂度：O(n)  空间复杂度：O(n)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	sb := &strings.Builder{}
	var preTraversal func(root *TreeNode)
	preTraversal = func(root *TreeNode) {
		if root == nil {
			sb.WriteString("nil,")
			return
		}

		sb.WriteString(strconv.Itoa(root.Val))
		sb.WriteString(",")
		preTraversal(root.Left)
		preTraversal(root.Right)
	}

	preTraversal(root)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")
	var treeBuilder func() *TreeNode
	treeBuilder = func() *TreeNode {
		if len(sp) == 0 {
			return nil
		}

		val := sp[0]
		sp = sp[1:]

		if val == "nil" {
			return nil
		}

		v, _ := strconv.Atoi(val)
		root := &TreeNode{Val: v}
		root.Left = treeBuilder()
		root.Right = treeBuilder()

		return root
	}

	return treeBuilder()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
