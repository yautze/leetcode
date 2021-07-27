package leetcode

func moveZeroes(nums []int) {
	// count 紀錄第幾個index不是為0的數
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			// 當找到數值為0,且不是該數值的位置時,交換兩數
			if i != count {
				nums[i], nums[count] = nums[count], nums[i]
			}
			// 紀錄數值不是0的位置
			count++
		}
	}
}
