package main

// 思路: 只有2*5才可以得到一个0，那我们只需要看n!可以分解为多少个2*5就可以
// 2出现的频率肯定是高于5的, 因为每隔 2 个数就会包含因子2，比如2,4,6,.., 而每个 5 个数才会出现一个包含因子5的数，比如5,10,15,..
// 因此，可以转换成求n！中'质因子5’的个数

// 时间复杂度: O(n)    空间复杂度: O(1)

func trailingZeroes_1(n int) int {
	res := 0

	for i := 5; i <= n; i += 5 {
		for j := i; j%5 == 0 && j > 0; j /= 5 {
			res++
		}
	}

	return res
}
