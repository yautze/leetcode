package leetcode

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		// [1, 2, , 4, 5]
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// [1, 2, 4, 5], target = 3 => left = 2, right = 1
	// 當left > right 時, 回傳left即是答案
	return left
}
