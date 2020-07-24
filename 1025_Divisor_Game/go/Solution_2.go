package main

// 动态规划
// 如果Alice最终取胜，则必须使得每一步Alice都是必胜状态，相应使得Bob处于必败状态
// dp[i] 表示当前数字为i时，Alice先手是否可以必胜

// 时间复杂度：O(N^2)  空间复杂度：O(N)

func divisorGame_2(N int) bool {
	if N==1 {
		return false
	}
	
	dp := make([]bool, N+1)
	dp[1] = false
	dp[2] = true

	for i:=3; i<=N; i++ {
		for j:=1; j<i; j++ {
			if i%j==0 && !dp[i-j] {  // dp[i-j]是Bob的状态，必须使其处于必败状态
				dp[i] = true
				break
			}
		}
	}

	return dp[N]
}