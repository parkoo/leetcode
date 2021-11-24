package main

// 数组遍历 位操作
// 时间复杂度：O(n)  空间复杂度：O(1)

func validUtf8(data []int) bool {
	for i := 0; i < len(data); {
		bits := helper(data[i])
		if bits == -1 || bits == 1 {
			return false
		} else {
			if bits == 0 {
				i++
				continue
			}

			for j := i + 1; j < i+bits; j++ {
				if j >= len(data) || helper(data[j]) != 1 {
					return false
				}
			}
			i += bits
		}
	}

	return true
}

func helper(num int) int {
	a := (num & (1 << 7)) >> 7
	b := (num & (1 << 6)) >> 6
	c := (num & (1 << 5)) >> 5
	d := (num & (1 << 4)) >> 4
	e := (num & (1 << 3)) >> 3

	if a == 1 && b == 1 && c == 1 && d == 1 && e == 0 {
		return 4
	}
	if a == 1 && b == 1 && c == 1 && d == 0 {
		return 3
	}
	if a == 1 && b == 1 && c == 0 {
		return 2
	}
	if a == 1 && b == 0 {
		return 1
	}
	if a == 0 {
		return 0
	}

	return -1
}
