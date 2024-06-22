package main

// 思路: 小顶堆，在小顶堆内维护k个元素，将数组中的元素依次插入小顶堆，最后弹出堆顶元素即为所求

// 时间复杂度：O(nlogk)     空间复杂度：O(k)

type MinHeap struct {
	arr []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		arr: make([]int, 0),
	}
}

func (mh *MinHeap) Insert(v int) {
	mh.arr = append(mh.arr, v)
	shiftUp(mh.arr, len(mh.arr)-1)
}

func (mh *MinHeap) Extract() int {
	top := mh.arr[0]
	mh.arr[0] = mh.arr[len(mh.arr)-1]
	mh.arr = mh.arr[:len(mh.arr)-1]
	shiftDown(mh.arr, 0)

	return top
}

func shiftUp(arr []int, k int) {
	v := arr[k]
	i := k
	for i > 0 && v < arr[(i-1)/2] { // 这里 i > 0表示当前节点不是根节点（根节点为0）
		arr[i] = arr[(i-1)/2]
		i = (i - 1) / 2
	}

	arr[i] = v
}

func shiftDown(arr []int, k int) {
	leftChild := 2*k + 1
	for leftChild < len(arr) {
		flag, rightChild := leftChild, leftChild+1
		if rightChild < len(arr) && arr[rightChild] < arr[leftChild] {
			flag = rightChild
		}

		if arr[k] < arr[flag] {
			break
		}

		arr[k], arr[flag] = arr[flag], arr[k]
		k = flag
		leftChild = 2*k + 1
	}
}

func findKthLargest(nums []int, k int) int {
	res := 0
	mh := NewMinHeap()
	for i := 0; i < len(nums); i++ {
		mh.Insert(nums[i])
		if i >= k {
			res = mh.Extract()
		}
	}
	res = mh.Extract()

	return res
}
