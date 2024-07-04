package main

// 思路: 贪心思想，需要同时考虑左右两侧的情况，连续递增、连续递减的次数

// 时间复杂度: O(n)    空间复杂度: O(1)

func candy_2(ratings []int) int {
	total := 1
	preCandy := 1   // 记录上一个儿童分得的糖果数
	descCnt := 0    // 记录连续递减的次数，连续递减i次就需要多分配i个糖果
	lastAscCnt := 1 // 记录连续递增的的次数
	for i := 1; i < len(ratings); i++ {
		curCandy := 1 // 记录当前需要新增的糖果数
		if ratings[i] >= ratings[i-1] {
			curCandy = preCandy + 1
			lastAscCnt = preCandy + 1

			if ratings[i] == ratings[i-1] {
				curCandy = 1
				lastAscCnt = 1
			}

			descCnt = 0
			preCandy = curCandy
		}

		if ratings[i] < ratings[i-1] {
			descCnt++
			curCandy = descCnt
			// 当递减序列长度和递增序列长度相等时，除了要加上应该新增的糖果数，还需将上一个递增序列最后一个位置在多加1个
			if descCnt >= lastAscCnt {
				curCandy++
			}
			preCandy = 1 // 这里新增的糖果数为 curCandy, 但对下一个儿童来说当前儿童是分得的糖果数是1
		}

		total += curCandy
	}

	return total
}
