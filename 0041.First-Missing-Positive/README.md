# [41. First Missing Positive](https://leetcode.com/problems/first-missing-positive/)

## 題目
Given an unsorted integer array nums, return the smallest missing positive integer.

You must implement an algorithm that runs in `O(n)` time and uses constant extra space.

Example1:
```
Input: nums = [1,2,0]
Output: 3
```

Example2:
```
Input: nums = [3,4,-1,1]
Output: 2
```

Example3:
```
Input: nums = [7,8,9,11,12]
Output: 1
```

Constraints:

* 1 <= nums.length <= 5 * 10^5
* -2^31 <= nums[i] <= 2^31 - 1

## 題目大意
在一個為排序的arr中,找尋第一個遺失的正整數

## 邏輯思維
先將arr的所有數值放入map中,再用一個從1開始的for迴圈比對是否有遺失的正整數,沒有遺失則回傳陣列長度+1