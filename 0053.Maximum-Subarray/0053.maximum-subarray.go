package leetcode

func maxSubArray(nums []int) int {
	// 排除特殊狀況
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 0 {
		return 0
	}

	// 預設最大為nums第一個數字
	max := nums[0]
	tmp := 0
	for _, v := range nums {
		tmp += v
		if tmp > max {
			max = tmp
		}

		// 小於 0 則重算
		if tmp < 0 {
			tmp = 0
		}
	}

	return max
}
