package main

// 思路: 前缀树（Trie） + DFS

// 时间复杂度: O(?)    空间复杂度: O(?)

type Node struct {
	isEnd bool
	next  map[rune]*Node
}

type WordDictionary struct {
	root *Node
}

func Constructor() WordDictionary {
	root := &Node{
		next: make(map[rune]*Node),
	}
	return WordDictionary{
		root: root,
	}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this.root
	for _, c := range []rune(word) {
		_, ok := cur.next[c]
		if !ok {
			newNode := new(Node)
			newNode.next = make(map[rune]*Node)
			cur.next[c] = newNode
		}
		cur = cur.next[c]
	}

	cur.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {

	// dfs 处理"."匹配规则
	var dfs func(node *Node, word []rune, n int) bool
	dfs = func(node *Node, word []rune, n int) bool {
		if n == len(word) {
			return node.isEnd
		}

		cur := word[n]
		if cur != '.' {
			next, ok := node.next[cur]
			if ok && dfs(next, word, n+1) {
				return true
			}

		} else {
			for _, next := range node.next {
				if dfs(next, word, n+1) {
					return true
				}
			}
		}

		return false
	}

	root := this.root
	return dfs(root, []rune(word), 0)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
