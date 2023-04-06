package leetcode

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([]int, 0)

	for i, v := range temperatures {
		for len(stack) > 0 && v > temperatures[stack[len(stack)-1]] {
			lastIdx := stack[len(stack)-1]
			distance := i - lastIdx
			res[lastIdx] = distance
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return res
}
