package leetcode

// 法1. 窮舉法
func longestPalindrome1(s string) string {
	l := len(s)
	if l == 1 {
		return s
	}

	left, right := 0, 0
	for i := 0; i < l; i++ {
		for j := i + 1; j <= l; j++ {
			if isPalindrome(s[i:j]) && j-i > right-left {
				left = i
				right = j
			}
		}
	}

	return s[left:right]
}

// isPalindrome -
func isPalindrome(s string) bool {
	right := len(s)
	for i := 0; i < right; i++ {
		if s[i] != s[right-i-1] {
			return false
		}
	}
	return true
}

// 法2. 中心擴散法
func longestPalindrome2(s string) string {
	res := ""

	for i := range s {
		// 奇數個找法 - aba 從`b`往兩旁擴散 中心點 left = right
		res = maxPalindrome(s, res, i, i)

		// 偶數個找法 - bccb 從`cc`往兩旁擴散 中心點 right = left + 1
		res = maxPalindrome(s, res, i, i+1)
	}

	return res
}

func maxPalindrome(s, res string, left, right int) string {
	tmp := ""
	for left >= 0 && right < len(s) && s[left] == s[right] {
		tmp = s[left : right+1]
		
		// 中心點往左擴散
		left--
		// 中心點往右擴散
		right++
	}

	if len(tmp) > len(res) {
		return tmp
	}

	return res
}
