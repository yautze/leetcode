package leetcode

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	if s == "" {
		return 0
	}

	// last - 記最後一個數字(for 比對)
	last, total := 0, 0

	// 由陣列最後面倒回來
	for i := len(s) - 1; i >= 0; i-- {
		t := string(s[i])
		num := roman[t]

		// 如果小的數字在大的數字的左邊則減掉
		if num < last {
			total -= num
		} else {
			total += num
		}

		// 記錄最後一個數
		last = num
	}

	return total
}
