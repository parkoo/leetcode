package main

// 数学推导
// 连续异或： a^0 = a, a^a=0, a^a^a = a, a^a^a^a = 0, 0^0...^0 = 0
// 由于每轮每人各取一个数字，若Alice面对的是偶数个数数组，则会一直面对偶数个数的数组，奇数同理；
// 情况1: 若Alice面对偶数个数数组，且这偶数个数字的异或结果不为0，Alice失败的情况是：拿走任意一个数字，剩下的奇数个数字的异或结果都为0；
//       令 Alice面对的偶数个数字的异或结果为S， 取走第i个数字的结果为Si, 则 Si = S^nums[i], S != 0, Si == 0;
//        S0^S1...^Sn = S^nums[0]^S^nums[1]...S^nums[n]
//                    = (S^S...^S)^(nums[0]^nums[1]...^nums[n])
//                    = 0^(nums[0]^nums[1]...^nums[n])
//                    = S
//        因为Si == 0，所以S0^S1...^Sn应该为0， 而S!=0, 故，以上结论“若Alice面对偶数个数数组，且这偶数个数字的异或结果不为0”时， Alice不会失败
//        即，Alice面对偶数个数数组，必胜
// 情况2: 若Alice面对奇数个数数组，且这奇数个数字的异或结果不为0，则对手会面对偶数个数数组，根据情况1，Alice必败

// 时间复杂度：O(n)  空间复杂度：O(1)

func xorGame_1(nums []int) bool {
	if len(nums)%2 == 0 {
		return true
	}

	win := 0
	for _, num := range nums {
		win ^= num
	}
	return win == 0
}
