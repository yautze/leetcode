package leetcode

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] != 10 {
			// 沒有進位
			return digits
		}
		// 有進位
		digits[i] = 0
	}

	// 9999 + 1 = 0000 => 所以要把第一位變成1, 後面補一個0
	digits[0] = 1
	digits = append(digits, 0)
	return digits
}
