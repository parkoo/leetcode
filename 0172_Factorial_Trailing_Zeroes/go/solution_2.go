package main

// 思路: 只有2*5才可以得到一个0，那我们只需要看n!可以分解为多少个2*5就可以
// 2出现的频率肯定是高于5的, 因为每隔 2 个数就会包含因子2，比如2,4,6,.., 而每个 5 个数才会出现一个包含因子5的数，比如5,10,15,..
// 因此，可以转换成求n！中’质因子5‘的个数

// n!=1*2*3*4*(1*5)*...*(2*5)*...*(3*5)*...*(1*5*5)*...*(2*5*5)*...*n
// 计算 n!中质因子5的个数，即 n/5+n/(5*5)+n/(5*5*5)+...
// 其中对于一个因数75来说，75 = 1*5*5*5， 出现了3次’质因子5‘，会经过3轮循环

// https://leetcode.cn/problems/factorial-trailing-zeroes/solutions/2588800/da-bai-hua-jiang-jie-zhao-gui-lu-si-lu-b-tp84/

// 时间复杂度: O(log5(n))    空间复杂度: O(1)

func trailingZeroes_2(n int) int {
	res := 0

	interval := 5
	for n >= interval {
		res += n / interval

		interval *= 5
	}

	return res
}
