package main

// 思路: 通过map记录当前节点的拷贝节点 迭代遍历

// 时间复杂度: O(n)    空间复杂度: O(n)

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList_1(head *Node) *Node {
	cache := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		newNode := new(Node)
		newNode.Val = cur.Val
		cache[cur] = newNode
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		newNode := cache[cur]
		newNode.Next = cache[cur.Next]
		newNode.Random = cache[cur.Random]
		cur = cur.Next
	}

	return cache[head]
}
