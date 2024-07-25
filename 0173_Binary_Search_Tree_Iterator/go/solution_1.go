package main

// 思路:  利用栈实现迭代式的二叉树中序遍历

// 时间复杂度：显然，初始化和调用 hasNext() 都只需要 O(1) 的时间。
// 每次调用 next() 函数最坏情况下需要 O(n) 的时间；但考虑到 n 次调用 next() 函数总共会遍历全部的 n 个节点，
// 因此总的时间复杂度为 O(n)，因此单次调用平均下来的均摊复杂度为 O(1)

// 空间复杂度：O(n)，其中 n 是二叉树的节点数量
// 空间复杂度取决于栈深度，而栈深度在二叉树为一条链的情况下会达到 O(n) 的级别

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
	cur   *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		stack: make([]*TreeNode, 0),
		cur:   root,
	}
}

func (this *BSTIterator) Next() int {
	for this.cur != nil {
		this.stack = append(this.stack, this.cur)
		this.cur = this.cur.Left
	}

	curNode := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	this.cur = curNode.Right
	return curNode.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0 || this.cur != nil
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
