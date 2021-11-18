package main

type LRUCache struct {
	store    map[int]*Node
	capacity int
	head     *Node
	tail     *Node
}

type Node struct {
	key   int
	value int
	pre   *Node
	next  *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.pre = head
	return LRUCache{
		store:    make(map[int]*Node),
		capacity: capacity,
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.store[key]
	if !ok {
		return -1
	}

	this.unlink(node)
	this.pushHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.store[key]
	if ok { // 若key存在 直接修改value 并更新节点位置
		node.value = value
		this.unlink(node)
		this.pushHead(node)
	} else { // 若key不存在 则在头部新增节点
		newNode := &Node{key: key, value: value}
		this.store[key] = newNode
		this.pushHead(newNode)

		// 新增后 检查超容删除
		if len(this.store) > this.capacity {
			invalidNode := this.tail.pre
			this.unlink(invalidNode)
			delete(this.store, invalidNode.key)
		}
	}
}

// 删除节点
func (this *LRUCache) unlink(node *Node) {
	pre := node.pre
	next := node.next
	pre.next = next
	next.pre = pre
	node.pre = nil
	node.next = nil
}

// 将节点置于头部
func (this *LRUCache) pushHead(node *Node) {
	next := this.head.next
	node.next = next
	next.pre = node
	this.head.next = node
	node.pre = this.head
}
