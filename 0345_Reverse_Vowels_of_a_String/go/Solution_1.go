package main

// 双指针 对撞指针
// 时间复杂度：O(n)  空间复杂度：O(1)

func reverseVowels(s string) string {
	vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}

	ss := []rune(s)
	i, j := 0, len(ss)-1
	for {
		for i < len(ss) && !vowels[ss[i]] {
			i++
		}
		for j >= 0 && !vowels[ss[j]] {
			j--
		}

		if i > j {
			break
		}

		ss[i], ss[j] = ss[j], ss[i]
		i++
		j--
	}

	return string(ss)
}
