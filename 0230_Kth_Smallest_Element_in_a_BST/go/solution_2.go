package main

// 思路: 为了优化题目中的频繁查找，需要做个缓存设计
// 主要缓存以某个节点为根节点的节点个数

// 时间复杂度: O(N)    空间复杂度: O(N)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest_2(root *TreeNode, k int) int {
	searchCache := InitSearchCache(root)
	return searchCache.KthSmall(k)
}

// 为了优化题目中的频繁查找，这里需要做个缓存设计
// 主要缓存以某个节点为根节点的节点个数
type SearchCache struct {
	root    *TreeNode
	nodeNum map[*TreeNode]int // 缓存以某个节点为根节点的节点个数
}

func InitSearchCache(root *TreeNode) *SearchCache {
	nodeNum := make(map[*TreeNode]int)

	// 递归构建缓存
	var helper func(node *TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftNum, rightNum := helper(node.Left), helper(node.Right)
		curNum := leftNum + rightNum + 1
		nodeNum[node] = curNum

		return curNum
	}

	helper(root)

	return &SearchCache{
		root:    root,
		nodeNum: nodeNum,
	}
}

// 通过缓存查找BST中第k小的值 BST的中序遍历是从小到大排列的
func (c *SearchCache) KthSmall(k int) int {
	cur := c.root
	for cur != nil {
		// 如果左子树中节点数刚好等于k-1, 则当前节点即为第k个小的节点
		leftNum := c.nodeNum[cur.Left]
		if leftNum == k-1 {
			return cur.Val
		}

		// 如果左子树的节点数大于 k-1 个，说明第k个小的节点在左子树中，令当前节点为左子树
		if leftNum > k-1 {
			cur = cur.Left

			// 如果左子树的节点数小于 k-1 个，说明第k个小的节点在右子树中，令当前节点为右子树，同时减去当前节点和左子树节点的个数（1 + leftNum）
		} else {
			k = k - leftNum - 1
			cur = cur.Right // 在右子树中寻找第 k - leftNum - 1 小的数
		}
	}

	return -1
}
