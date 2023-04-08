package leetcode

func findComplement(num int) int {
    mask := 1
    for mask < num {
        mask = (mask << 1) + 1
    }

    return num ^ mask
}