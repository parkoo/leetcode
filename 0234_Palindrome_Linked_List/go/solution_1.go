package main

// 思路: 将链表转换为数组，判断数组是否为回文

// 时间复杂度: O(n)    空间复杂度: O(n)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome_1(head *ListNode) bool {
	arr := make([]int, 0)
	cur := head
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}

	return checkPalindrome(arr)
}

func checkPalindrome(arr []int) bool {
	i, j := 0, len(arr)-1
	for i < j {
		if arr[i] != arr[j] {
			return false
		}
		i++
		j--
	}

	return true
}
