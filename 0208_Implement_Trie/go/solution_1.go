package main

// 思路: 实现Trie

// 时间复杂度: 初始化为 O(1)，其余操作为 O(∣S∣)，其中 ∣S∣是每次插入或查询的字符串的长度

type Node struct {
	isWord bool
	next   map[rune]*Node
}

type Trie struct {
	root *Node
	size int // 额外记录一下不重复word的数量，本题不涉及
}

func Constructor() Trie {
	root := &Node{
		isWord: false,
		next:   make(map[rune]*Node),
	}

	return Trie{root: root}
}

func (this *Trie) Insert(word string) {
	cur := this.root
	for _, letter := range []rune(word) {
		if node, ok := cur.next[letter]; ok {
			cur = node
			continue
		}

		newNode := &Node{
			isWord: false,
			next:   make(map[rune]*Node),
		}
		cur.next[letter] = newNode
		cur = newNode
	}

	if !cur.isWord { // 这里为了防止对重复字符串进行计数，需要做一个if判断
		cur.isWord = true
		this.size++
	}
}

func (this *Trie) Search(word string) bool {
	cur := this.root
	for _, letter := range []rune(word) {
		node, ok := cur.next[letter]
		if !ok {
			return false
		}
		cur = node
	}

	return cur.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root
	for _, letter := range []rune(prefix) {
		node, ok := cur.next[letter]
		if !ok {
			return false
		}
		cur = node
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
