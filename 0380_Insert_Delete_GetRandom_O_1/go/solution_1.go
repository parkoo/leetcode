package main

import "math/rand"

// 思路: 通过变长数组 + map 实现

// 时间复杂度: O(?)    空间复杂度: O(?)

type RandomizedSet struct {
	nums []int
	m    map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums: make([]int, 0),
		m:    make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.nums = append(this.nums, val)
	this.m[val] = len(this.nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.m[val]
	if !ok {
		return false
	}

	// 注意细节 当index == len(this.nums)-1 == 0 时
	// 需要先修改完m, 再删除nums中最后提个元素，否则会出现索引越界
	this.nums[index] = this.nums[len(this.nums)-1]
	this.m[this.nums[index]] = index
	this.nums = this.nums[:len(this.nums)-1]
	delete(this.m, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	// O(1)复杂度的随机获取 必须通过数组实现
	r := rand.Intn(len(this.nums))
	return this.nums[r]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
