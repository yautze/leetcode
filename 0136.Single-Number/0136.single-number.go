package leetcode

func singleNumberByMap(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		val, ok := m[v]
		if ok {
			val++
			m[v] = val
			continue
		}

		m[v] = 1
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}

	return 0
}

func singleNumber(nums []int) int {
	res := 0

	for _, v := range nums {
		res ^= v
	}

	return res
}
