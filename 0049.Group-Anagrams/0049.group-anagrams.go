package leetcode

import "sort"

// groupAnagramsBySort - 透過排序每個字串的方法
func groupAnagramsBySort(strs []string) [][]string {
	m := make(map[string][]string)
	res := [][]string{}

	for _, v := range strs {
		b := []byte(v)

		// sort b
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		s := string(b)
		m[s] = append(m[s], v)
	}

	for _, v := range m {
		res = append(res, v)
	}

	return res
}

// groupAnagramsByCombination - 透過計算字串的每一個英文字母的次數組合
func groupAnagramsByCombination(strs []string) [][]string {
	m := make(map[[26]int][]string)
	res := [][]string{}

	for _, v := range strs {
		mapArr := [26]int{}
		for _, r := range v {
			mapArr[r-'a'] += 1
		}

		m[mapArr] = append(m[mapArr], v)
	}

	for _, v := range m {
		res = append(res, v)
	}

	return res
}
