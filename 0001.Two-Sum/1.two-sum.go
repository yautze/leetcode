package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if index, ok := m[target-v]; ok {
			return []int{index, i}
		}
		// 未找到的數字將其位置存入map中
		// m[數字] = array index
		m[v] = i
	}
	return nil
}
