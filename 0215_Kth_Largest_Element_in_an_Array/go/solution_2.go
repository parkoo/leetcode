package main

// 思路: 对数组依次截取k+1个元素作heapify, 大数据会超时，借此用来练习heapify，并不是此题的可取解法

// 时间复杂度：O(nklogk) 其中heapify的时间复杂度为O(klogk),因为会对截取的片段做复制,可做优化    空间复杂度：O(k)

func findKthLargest_2(nums []int, k int) int {
	for i := k; i < len(nums); i++ {
		heapify(nums, i-k, i)
	}
	heapify(nums, len(nums)-k, len(nums)-1)

	return nums[len(nums)-k]
}

// 对数组arr的[m,n]区间作heapify
func heapify(data []int, m, n int) {
	// 为了实现对数组的heapfy，必须将数组从位置0或1开始计入
	arr := make([]int, n-m+2)
	for i := m; i <= n; i++ {
		arr[i-m+1] = data[i]
	}

	last := len(arr) - 1/2
	for i := last; i >= 1; i-- {
		cur := i
		leftChild := 2 * cur
		for leftChild < len(arr) {
			flag, rightChild := leftChild, leftChild+1

			if rightChild < len(arr) && arr[rightChild] < arr[leftChild] {
				flag = rightChild
			}

			if arr[cur] < arr[flag] {
				break
			}

			arr[cur], arr[flag] = arr[flag], arr[cur]
			cur = flag
			leftChild = 2 * cur
		}
	}

	for i := 1; i < len(arr); i++ {
		data[i+m-1] = arr[i]
	}
}
