package leetcode

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	operators := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}

	for _, v := range tokens {
		if handler, ok := operators[v]; ok {
			n := len(stack)
			a, b := stack[n-2], stack[n-1]
			stack = append(stack[:n-2], handler(a, b))
		} else {
			num, _ := strconv.Atoi(v)
			stack = append(stack, num)
		}
	}

	return stack[len(stack)-1]
}
