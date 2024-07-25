package main

// 思路: 递归

// 时间复杂度: O(n^2*logn) (近O(n^2))    空间复杂度: O(logn)

// Definition for a QuadTree node.
type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct_1(grid [][]int) *Node {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return nil
	}

	n := len(grid)

	// 判断 从[r0行 ~ r1行]、[c0列 ~ c1列] 的矩形区域是否为叶子节点(全为0或者全为1)
	// 返回是否是叶子节点以及当前节点的val
	leafArea := func(r0, r1, c0, c1 int) (bool, bool) {
		isLeaf := true
		flag := grid[r0][c0]
		val := flag == 1

		for i := r0; i <= r1; i++ {
			for j := c0; j <= c1; j++ {
				if grid[i][j] != flag {
					isLeaf, val = false, false
					return isLeaf, val
				}
			}
		}

		return isLeaf, val
	}

	var dfs func(r0, r1, c0, c1 int) *Node
	dfs = func(r0, r1, c0, c1 int) *Node {
		if isLeaf, val := leafArea(r0, r1, c0, c1); isLeaf {
			return &Node{
				Val:    val,
				IsLeaf: isLeaf,
			}
		}

		rMid := r0 + (r1-r0)/2
		cMid := c0 + (c1-c0)/2

		curNode := &Node{
			Val:         false,
			IsLeaf:      false,
			TopLeft:     dfs(r0, rMid, c0, cMid),
			TopRight:    dfs(r0, rMid, cMid+1, c1),
			BottomLeft:  dfs(rMid+1, r1, c0, cMid),
			BottomRight: dfs(rMid+1, r1, cMid+1, c1),
		}
		return curNode
	}

	return dfs(0, n-1, 0, n-1)
}
