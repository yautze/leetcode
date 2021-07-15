# [53. Maximum Subarray](https://leetcode.com/problems/maximum-subarray/)

## 題目
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example1:
```
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
```

Example2:
```
Input: nums = [1]
Output: 1
```

Example3:
```
Input: nums = [5,4,-1,7,8]
Output: 23
```

## 題目大意
給一個int arr,在裡面找到連續且多個數字的總和相加最大的
* 至少要有一個數字相加

## 邏輯思維
當加到數字後 < 0, 就必須重算 </br>
當加到數字後 > max, 更新max的數字 </br>
當加到數字後 < max 且 > 0, 繼續加下一個數字</br>
(有可能下一個數字加下去就變最大,題目要求需連續加總)
