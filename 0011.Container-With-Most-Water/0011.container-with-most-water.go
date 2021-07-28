package leetcode

func maxArea(height []int) int {
	max := 0
	left, right := 0, len(height)-1

	for left < right {
		width := right - left
		high := 0

		// high是矮的那一邊
		if height[left] < height[right] {
			high = height[left]
			left++
		} else {
			high = height[right]
			right--
		}

		if (high * width) > max {
			max = high * width
		}
	}

	return max
}
