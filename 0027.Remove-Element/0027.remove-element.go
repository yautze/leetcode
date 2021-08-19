package leetcode

func removeElement(nums []int, val int) int {
	// 確定不等於val的個數(下一個要補的位置)
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if i != count {
				// 不等於val,且不等於自己的位置時,將值往前移到確定不等於val的位置
				nums[i], nums[count] = nums[count], nums[i]
			}
			count++
		}
	}

	return count
}
