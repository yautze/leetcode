package leetcode

func missingNumberByXOR(nums []int) int {
    res := 0
    for i := 0; i <= len(nums); i++ {
        res ^= i
    }

    for _, v := range nums {
        res ^= v
    }

    return res
}

func missingNumberByCyclicSort(nums []int) int {
	i := 0
	// cyclic sort
	for i < len(nums) {
		j := nums[i]
		if j < len(nums) && nums[i] != nums[j] {
			// swap the numbers
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
	}

	// find the first missing number
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	
	return len(nums)
}