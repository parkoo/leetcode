package main

// 思路: 通过掩码mask，将num的每一个二进制位翻转，然后再通过或运算将其结果拼接起来

// 时间复杂度: O(C)    空间复杂度: O(1)

func reverseBits(num uint32) uint32 {
	res := uint32(0)

	mask := 1 << 31 // 通过掩码mask获取num的每一个二进制位（从高位到低位）
	for i := 0; i < 32; i++ {
		if num&uint32(mask) != 0 { // 该二进制位为'1'
			res |= 1 << i // 将结果的对应位设置为'1'
		}
		mask >>= 1
	}

	return res
}
