package leetcode

func searchRange(nums []int, target int) []int {
	return []int{findFirstIndex(nums, target), findLastIndex(nums, target)}
}

// 找最左邊的index
func findFirstIndex(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			// 找到第一個位數, 或前一個位數不等於自己
			// [`1`, 1, 2, 3, 4], [0, `1`, 1, 2, 3]
			if mid == 0 || nums[mid-1] != target {
				return mid
			}

			right = mid - 1
		}
	}

	return -1
}

// 找最右邊的index
func findLastIndex(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			// 找到最後一個位數, 或下一個位數不等於自己
			// [0, 0, 1, 2, `2`], [1, `1`, 2, 3, 4]
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			}

			left = mid + 1
		}
	}

	return -1
}
