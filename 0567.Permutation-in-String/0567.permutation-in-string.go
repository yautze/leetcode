package leetcode

func checkInclusion(s1 string, s2 string) bool {
	l1, l2 := len(s1), len(s2)

	var m1, m2 [26]int

	// first slide window
	// v - 'a' => ('b' - 'a' = 98 - 97 = 1) 為了放進長度26的array中
	for i, v := range s1 {
		m1[v-'a']++
		m2[s2[i]-'a']++
	}

	if m1 == m2 {
		return true
	}

	// next slide window
	for i := l1; i < l2; i++ {
		// 加入後面新的一個char ex: ab -> abc
		m2[s2[i]-'a']++

		// 移除最前面一個char ex: abc -> bc
		m2[s2[i-l1]-'a']--
		if m1 == m2 {
			return true
		}
	}

	return false
}
