package main

// 思路: 二分搜索 + 二进制查找完全二叉树的节点

// [0, h-1]层总节点数为 2^0 + 2^1 + ... + 2^(h-1) = 2^h - 1
// 第h层的节点数最少为1，最大为2^h
// 即对于总层数为h的完全二叉树，总结点个数在[2^h, 2^(h+1) - 1]之间
// 将节点依次从左向右、从上到下编号，起始根节点为1，则最后一层的节点编号范围为
// 由此，可以将节点编号，用二分法查找完全二叉树的最后一个节点 ( [2^h, 2^(h+1)-1]范围内 )

// 对各个节点编号，起始根节点为1，观察各个节点编号的二进制
//            1
//          /    \
//        10      11
//      /   \    /   \
//    100  101  110  111

// 有如下规律：
// 1. 对于第n层的节点，其编号的二进制共有n+1位, 且最高位为1
// 2. 某一个节点的左子节点的二进制编号是该节点的编号后边加一位'0', 而其右子节点的二进制编号是该节点的编号后边加一位'1'
// 根据如上规律可以通过节点的二进制编号判断该编号的节点是否存在

// O((logn)^2)，其中 n 是完全二叉树的节点数。
// 首先需要 O(h) 的时间得到完全二叉树的最大层数，其中 h 是完全二叉树的最大层数。
// 使用二分查找确定节点个数时，需要查找的次数为 O(log(2^h))=O(h)，每次查找需要遍历从根节点开始的一条长度为 h 的路径，需要 O(h) 的时间，因此二分查找的总时间复杂度是 O(h^2)。
// 因此总时间复杂度是 O(h^2)。由于完全二叉树满足 2^h <= n < 2^(h+1)，因此有 O(h)=O(logn)，O(h^2)=O((logn)^2)

// 空间复杂度: O(1)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes_2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 计算树的总层数（起始层数为0）
	h := 0
	for cur := root; cur.Left != nil; cur = cur.Left {
		h++
	}

	// 判断最后一层的节点编号为n的节点是否存在
	exist := func(n int, h int) bool {
		// 获取编号n的二进制的每一位的掩码（n为最后一层节点， 共有h+1位, 从次高位开始取）
		mask := 1 << (h - 1)
		cur := root
		for mask != 0 {
			if n&mask == 0 {
				cur = cur.Left
			} else {
				cur = cur.Right
			}

			mask = mask >> 1
		}

		return cur != nil
	}

	// h == 0 时， left == right, 无法进入循环，符合正确结果
	left, right := 1<<h, (1<<(h+1))-1 // 2^h, 2^(h+1)-1
	for left < right {
		// 这里采用的是左等于中间点模式，在个数为偶数时，中间点取靠右侧的数
		mid := left + (right-left+1)/2

		if exist(mid, h) {
			left = mid // 左等于中间点模式

		} else {
			right = mid - 1
		}
	}

	return left
}
