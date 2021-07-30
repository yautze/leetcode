package leetcode

func isValid(s string) bool {
	// 儲存左邊符號的暫存器
	stack := make([]rune, 0)

	for _, v := range s {
		// 若有左邊的符號則加入stack中,當紀錄
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		} else if (v == ')' && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			(v == '}' && len(stack) > 0 && stack[len(stack)-1] == '{') ||
			(v == ']' && len(stack) > 0 && stack[len(stack)-1] == '[') {
			// 讀取到右邊的符號, 且stack長度大於1, 且最後一個存入的符號 == 對稱的符號, 則清掉最後一個符號
			// 找到符合條件則移除該筆stack的內容
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	// 若是對稱,stack的長度應該為0
	return len(stack) == 0
}
