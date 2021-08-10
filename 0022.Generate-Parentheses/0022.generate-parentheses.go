package leetcode

func generateParenthesis(n int) []string {
	res := []string{}
	if n == 0 {
		return res
	}
	findParenthesis(n, n, "", &res)

	return res
}

func findParenthesis(left, right int, str string, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, str)
		return
	}

	// 左括號
	if left > 0 {
		findParenthesis(left-1, right, str+"(", res)
	}

	// right > left 代表需要補右邊括號
	if right > 0 && right > left {
		findParenthesis(left, right-1, str+")", res)
	}
}
