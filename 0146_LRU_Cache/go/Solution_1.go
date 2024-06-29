package main

// 思路：双端链表 + map
// 注意点：
//  1. Node中需记录key,方便从链表定位到map, 在删除过期的尾部节点时，需要定位map中的数据并删除
//  2. Push()时，先新增节点后检查是否超容，否则会出现容量满了后，无法新增节点
//  3. Push()中新增节点、删除节点时需要同步处理对应的map中的数据

type LRUCache struct {
	store    map[int]*Node
	capacity int
	head     *Node
	tail     *Node
}

type Node struct {
	key   int // 记录key,方便从链表定位到map, 在删除过期的尾部节点时，需要定位map中的数据并删除
	value int
	pre   *Node
	next  *Node
}

func Constructor(capacity int) LRUCache {
	head, tail := new(Node), new(Node)
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

// 先新增后检查是否超容，否则会出现容量满了后，无法新增节点
func (this *LRUCache) Put(key int, value int) {
	node, ok := this.store[key]
	if ok { // 若key存在 直接修改value 并更新节点位置
		node.value = value
		this.unlink(node)
		this.pushHead(node)
		return

	}

	// 若key不存在 则在头部新增节点
	newNode := &Node{key: key, value: value}
	this.store[key] = newNode // map中新增
	this.pushHead(newNode)

	// 新增后 检查超容删除
	if len(this.store) > this.capacity {
		invalidNode := this.tail.pre
		this.unlink(invalidNode)
		delete(this.store, invalidNode.key) // map中同步删除，需要用到Node中的key值
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
