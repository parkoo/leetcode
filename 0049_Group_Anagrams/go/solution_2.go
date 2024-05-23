package main

import "sort"

// 思路: 对字母异位词做归一化处理（按照字母排序），以归一化的结果作为map的key进行分组

// 时间复杂度: O(nklogk)    空间复杂度: O(nk)  其中n是strs中的字符串的数量，k是strs中的字符串的的最大长度

func groupAnagrams_2(strs []string) [][]string {

	cache := make(map[string][]string)

	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortStr := string(s)
		cache[sortStr] = append(cache[sortStr], str)
	}

	res := make([][]string, 0)
	for _, group := range cache {
		res = append(res, group)
	}

	return res
}
