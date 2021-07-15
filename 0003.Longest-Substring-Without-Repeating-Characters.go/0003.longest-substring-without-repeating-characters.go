package leetcode

func lengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	}

	max := 0
	left := 0 // 起始位置
	r := []rune(s)

	for right, current := range s {
		for i := left; i < right; i++ {
			// 如果找到重複, 將left調整至此index + 1 (往左排除到他)
			if r[i] == current {
				left = i + 1
			}
			// 沒找到right持續加一
		}

		if right-left+1 > max {
			max = right - left + 1
		}
	}

	return max
}
