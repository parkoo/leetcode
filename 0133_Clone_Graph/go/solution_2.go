package main

// 思路: dfs 返回当前原节点的拷贝节点
// 拷贝新节点与更新新节点的相邻节点同步进行

// 时间复杂度: O(n)    空间复杂度: O(n)

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph_2(node *Node) *Node {
	if node == nil {
		return node
	}

	// 记录原节点到新节点的映射
	cache := make(map[*Node]*Node)

	// dfs 返回当前原节点的拷贝节点
	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		if newNode, ok := cache[node]; ok {
			return newNode
		}

		newNode := new(Node)
		newNode.Val = node.Val
		newNode.Neighbors = make([]*Node, 0)
		cache[node] = newNode // 先记录cache再xiang下进行dfs

		for _, neighbor := range node.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, dfs(neighbor))
		}

		return newNode
	}

	dfs(node)

	return cache[node]
}
