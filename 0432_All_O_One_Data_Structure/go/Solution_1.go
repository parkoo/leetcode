package main

// map + 双向链表

type AllOne struct {
	store map[string]*Node
	head  *Node
	tail  *Node
}

// 在一个Node里存储了cnt相同的所有key
type Node struct {
	cnt  int
	keys map[string]bool
	pre  *Node
	next *Node
}

func Constructor() AllOne {
	head, tail := &Node{cnt: -1}, &Node{cnt: -1}
	head.next = tail
	tail.pre = head
	return AllOne{
		store: make(map[string]*Node),
		head:  head,
		tail:  tail,
	}
}

func (this *AllOne) Inc(key string) {
	oldNode, ok := this.store[key]
	if ok {
		var newNode *Node
		if oldNode.next.cnt != oldNode.cnt+1 {
			newNode = new(Node)
			newNode.keys = make(map[string]bool)
			newNode.cnt = oldNode.cnt + 1
			this.addNode(oldNode, newNode)
		} else {
			newNode = oldNode.next
		}
		newNode.keys[key] = true
		this.store[key] = newNode

		delete(oldNode.keys, key)
		if len(oldNode.keys) == 0 {
			this.delNode(oldNode)
		}
	} else {
		if this.head.next.cnt == 1 {
			this.head.next.keys[key] = true
			this.store[key] = this.head.next
		} else {
			newNode := new(Node)
			newNode.keys = make(map[string]bool)
			newNode.keys[key] = true
			newNode.cnt = 1
			this.addNode(this.head, newNode)
			this.store[key] = newNode
		}
	}
}

func (this *AllOne) Dec(key string) {
	oldNode, ok := this.store[key]
	if ok {
		var newNode *Node
		if oldNode.pre.cnt != oldNode.cnt-1 {
			if oldNode.cnt-1 != 0 { // 需判断cnt == 0 的情况
				newNode = new(Node)
				newNode.keys = make(map[string]bool)
				newNode.cnt = oldNode.cnt - 1
				newNode.keys[key] = true
				this.store[key] = newNode
				this.addNode(oldNode.pre, newNode)
			} else {
				delete(this.store, key) // 需要delete()
			}

		} else {
			newNode = oldNode.pre
			newNode.keys[key] = true
			this.store[key] = newNode
		}

		delete(oldNode.keys, key)
		if len(oldNode.keys) == 0 {
			this.delNode(oldNode)
		}
	}
}

func (this *AllOne) GetMaxKey() string {
	for k, _ := range this.tail.pre.keys {
		return k
	}
	return ""
}

func (this *AllOne) GetMinKey() string {
	for k, _ := range this.head.next.keys {
		return k

	}
	return ""
}

func (this *AllOne) delNode(node *Node) {
	pre := node.pre
	next := node.next
	pre.next = next
	next.pre = pre
	node.pre = nil
	node.next = nil
}

func (this *AllOne) addNode(cur, node *Node) {
	next := cur.next
	cur.next = node
	node.next = next
	next.pre = node
	node.pre = cur
}
