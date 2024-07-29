package main

// 思路:  分解质因数 贪心

// 数学本质：n为质数只能复制1次然后不停粘贴，dp[i] = i
// 只需要对 n 进行质因数分解，统计所有质因数的和，即为最终的答案。

// 时间复杂度: O(n^(1/2))    空间复杂度: O(1)

func minSteps_3(n int) int {
	res := 0

	for i := 2; i*i <= n; i++ {
		for n%i == 0 { // i是n的一个质因数
			res += i
			n /= i // 更新n
		}
	}

	if n > 1 { // 此时的n也是一个质因数
		res += n
	}

	return res
}
