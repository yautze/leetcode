package leetcode

import "sort"

func topKFrequent(nums []int, k int) []int {
	res := []int{}
	m := make(map[int]int)

	for _, v := range nums {
		m[v] += 1
	}

	keys := make([]int, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	for i := 0; i < k; i++ {
		res = append(res, keys[i])
	}

	return res
}
