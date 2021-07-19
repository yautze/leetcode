# [4. Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/)

## 題目
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).

Example1:
```
Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
```

Example2:
```
Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
```

Example3:
```
Input: nums1 = [0,0], nums2 = [0,0]
Output: 0.00000
```

Example4:
```
Input: nums1 = [], nums2 = [1]
Output: 1.00000
```

Example5:
```
Input: nums1 = [2], nums2 = []
Output: 2.00000
```

Constraints:
* nums1.length == m
* nums2.length == n
* 0 <= m <= 1000
* 0 <= n <= 1000
* 1 <= m + n <= 2000
* -10^6 <= nums1[i], nums2[i] <= 10^6

## 題目大意
給兩個以排序過的陣列,從這兩個陣列中找出中位數
時間複雜度 < O(log (m+n))

## 邏輯思維
![](https://scontent.xx.fbcdn.net/v/t1.15752-0/s206x206/210645383_836091150669381_1524726745091486022_n.jpg?_nc_cat=101&ccb=1-3&_nc_sid=aee45a&_nc_ohc=xkAd5xFjrVkAX9efiZc&_nc_ad=z-m&_nc_cid=0&_nc_ht=scontent.xx&oh=38509d45bfd3ba46cf962e0b624f0ace&oe=60FA47A9)