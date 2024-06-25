package main

// 思路: 对每个节点插入一个后继节点，然后拆分成两个链表

// 时间复杂度: O(n)    空间复杂度: O(1)

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList_3(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 节点复制
	cur := head
	for cur != nil {
		next := cur.Next
		newNode := new(Node)
		newNode.Val = cur.Val
		cur.Next = newNode
		newNode.Next = next
		cur = cur.Next.Next
	}

	// 随机指针复制
	cur = head
	for cur != nil {
		random := cur.Random
		newNode := cur.Next
		if random != nil { // 注意 random为空的情况
			newNode.Random = random.Next
		}
		cur = cur.Next.Next
	}

	// 拆分
	cur = head
	newHead := cur.Next // 注意 head为空的情况
	for cur != nil {
		newNode := cur.Next
		cur.Next = cur.Next.Next
		if cur.Next != nil { // 注意 cur.Next为空的情况
			newNode.Next = cur.Next.Next
		}
		cur = cur.Next
	}

	return newHead
}
