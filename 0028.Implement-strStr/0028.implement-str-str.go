package leetcode

import "strings"

func strStr1(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func strStr2(haystack string, needle string) int {
	hl, nl := len(haystack), len(needle)

	// 排除特殊狀況
	switch {
	case nl > hl:
		return -1
	case nl == 0:
		return 0
	}

	left, right := 0, nl
	for right <= hl {
		if haystack[left:right] == needle {
			return left
		}

		// 沒找到將slide windows往右移一格,因為固定長度所以左右一起移動
		left++
		right++
	}

	return -1
}
