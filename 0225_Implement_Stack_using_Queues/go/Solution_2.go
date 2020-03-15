package main

import "container/list"

// 使用GO中的list作为FIFO队列, 利用一个队列实现栈,每次push时都将新入队的元素移到队首
// "push()"，"pop()"，"top()"，"empty()"的时间复杂度分别为：O(n),O(1),O(1),O(1)

type MyStack struct {
	q *list.List
}

func Constructor() MyStack {

	return MyStack{q:list.New()}
}

func (this *MyStack) Push(x int)  {

	this.q.PushBack(x)
	for i:=0; i<this.q.Len()-1; i++ {
		e := this.q.Front()
		this.q.Remove(e)
		this.q.PushBack(e.Value.(int))
	}
}

func (this *MyStack) Pop() int {

	if this.q.Len()==0 {
		return -1
	}

	e := this.q.Front()
	this.q.Remove(e)

	return e.Value.(int)	
}

func (this *MyStack) Top() int {

	if this.q.Len()==0 {
		return -1
	}

	e := this.q.Front()
	return e.Value.(int)
}

func (this *MyStack) Empty() bool {

	return this.q.Len()==0
}