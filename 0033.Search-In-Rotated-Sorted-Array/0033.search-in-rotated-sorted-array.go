package leetcode

// O(n) - map對應表
func searchByMap(nums []int, target int) int {
	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}

	index, ok := m[target]
	if ok {
		return index
	}

	return -1
}

// O(log n) - 二分搜尋法
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > nums[left] {
			// [4, 5, 6, 7, 0, 1, 2]
			// 中間值落在大的那邊(大的最右邊) [4, 5, 6, 7]

			// target >= 4 && target < 7 => right移動到mid左邊一位
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < nums[right] {
			// [5, 6, 7, 0, 1, 2, 3]
			// 中間值在小的一邊(小的最左邊) [0, 1, 2, 3]

			// target > 0 && target <= 3 => left移動到mid右邊一位
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			// 等於時的情況
			if nums[mid] == nums[left] {
				left++
			} else {
				right--
			}
		}
	}

	return -1
}
