package leetcode

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	// 次方是負的時候調整
	if n < 0 {
		n = -n
		x = 1 / x
	}

	// 找到次方一半的數 => 2^10 = 2^5 * 2^5
	halfNum := myPow(x, n/2)

	if n%2 == 0 {
		return halfNum * halfNum
	}

	// 奇數多乘一次
	return halfNum * halfNum * x
}
