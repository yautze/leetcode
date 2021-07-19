package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	cl := l1 + l2
	arr := []int{}
	isJump := false

	// 如果有某一個arr len = 0, 就跳過兩個陣列相加的還結
	if l1 == 0 {
		arr = append(arr, nums2...)
		isJump = true
	}
	if l2 == 0 {
		arr = append(arr, nums1...)
		isJump = true
	}

	i, j := 0, 0
	if !isJump {
		// 兩個陣列比大小, 小的加入arr中
		for i < l1 && j < l2 {
			if nums1[i] <= nums2[j] {
				arr = append(arr, nums1[i])
				i++
				continue
			}
			arr = append(arr, nums2[j])
			j++
		}

		// 將剩下的那一個陣列加回去
		if i == l1 {
			arr = append(arr, nums2[j:l2]...)
		} else {
			arr = append(arr, nums1[i:l1]...)
		}
	}

	c := cl / 2
	if cl%2 == 0 {
		return float64(arr[c]+arr[c-1]) / 2
	}

	return float64(arr[c])
}
