package main

// 组合问题　回溯 + 剪枝
// 时间复杂度：O(k*Cₙᵏ)  空间复杂度：O(Cₙᵏ)  其中,k表示需要将每个长度为k的组合都添加到输出中
                                    
var res [][]int

func combine_1(n int, k int) [][]int {
	res = [][]int{}
	if n<k || n==0 || k==0 {
		return res
	}

	var t []int
	getCombinations(n, k, 1, 0, t)
	return res
}

func getCombinations(n int, k int, start int, cnt int, t []int) {
	if cnt==k {
		s := make([]int, len(t))
		copy(s, t)
		res = append(res, s)
		return 
	}

	// 剪枝："i<=n"  -->  "i<=n-(k-cnt)+1" 
	for i:=start; i<=n-(k-cnt)+1; i++ {  
		t = append(t, i)
		getCombinations(n, k, i+1, cnt+1, t)
		t = t[:len(t)-1]
	}
}