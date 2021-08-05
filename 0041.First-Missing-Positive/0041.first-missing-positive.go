package leetcode

func firstMissingPositive(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v] = v
	}

	// 從正整數1開始比對是否有在map中
	for i := 1; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}

	return len(nums) + 1
}
