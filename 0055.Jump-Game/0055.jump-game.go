package leetcode

func canJump(nums []int) bool {
	l := len(nums)

	if l == 0 {
		return false
	}
	if l == 1 {
		return true
	}

	jumps := 0
	for i, v := range nums {
		if i > jumps {
			return false
		}
		if jumps < i+v {
			jumps = i + v
		}
	}

	return true
}
