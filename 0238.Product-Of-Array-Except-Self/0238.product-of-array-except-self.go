package leetcode

func productExceptSelf(nums []int) []int {
    res := make([]int, len(nums))
    res[0] = 1

    // nums[i]左邊的所有數字的乘積
    for i := 1; i < len(nums); i++ {
        res[i] = res[i-1] * nums[i-1]
    }
    
    // nums[i]右邊的所有數字的乘積
    right := 1
    for i := len(nums) - 1; i >= 0; i-- {
        res[i] *= right
        right *= nums[i]
    }

    return res
}
