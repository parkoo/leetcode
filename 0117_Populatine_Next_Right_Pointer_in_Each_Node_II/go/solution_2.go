package main

// 思路: 利用已处理好的层构建下一未处理的层

// 时间复杂度: O(n)    空间复杂度: O(1)

// Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect_2(root *Node) *Node {

	// prev 为已经处理好的一层，此时prev这一层通过Next指针相连可以看做链表
	// 通过遍历prev的每个节点可以处理prev的下一层，将下一层也通过Next指针相连
	prev := root

	// 循环遍历每一层，从上至下
	for prev != nil {
		// 构建待处理层的虚拟有节点
		dummyHead := new(Node)
		cur := dummyHead

		// 遍历当前层的同时，连接下一待处理层的节点
		for prev != nil {
			if prev.Left != nil {
				cur.Next = prev.Left
				cur = cur.Next
			}
			if prev.Right != nil {
				cur.Next = prev.Right
				cur = cur.Next
			}

			prev = prev.Next
		}

		// 移动到下一层的实际头节点处
		prev = dummyHead.Next
	}

	return root
}
