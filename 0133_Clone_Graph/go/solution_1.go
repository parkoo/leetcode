package main

// 思路: dfs 记录原节点到新节点的映射， 再通过cache更新新节点的相邻节点
// 拷贝新节点与更新新节点的相邻节点分为两步

// 时间复杂度: O(n)    空间复杂度: O(n)

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph_1(node *Node) *Node {
	if node == nil {
		return node
	}

	// 记录原节点到新节点的映射
	cache := make(map[*Node]*Node)

	// 将原节点拷贝一份到cache中，这里不处理新节点的相邻节点
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if _, ok := cache[node]; ok {
			return
		}

		newNode := new(Node)
		newNode.Val = node.Val
		newNode.Neighbors = make([]*Node, 0)
		cache[node] = newNode

		for _, neighbor := range node.Neighbors {
			dfs(neighbor)
		}
	}

	dfs(node)

	// 通过cache更新新节点的相邻节点
	for node, newNode := range cache {
		for _, neighbor := range node.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, cache[neighbor])
		}
	}

	return cache[node]
}
