# [33. Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/)

## 題目
There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is `rotated` at an unknown pivot index k (0 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with `O(log n)` runtime complexity.

Example1:
```
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
```

Example2:
```
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
```

Example3:
```
Input: nums = [1], target = 0
Output: -1
```

Constraints:

* 1 <= nums.length <= 5000
* -10^4 <= nums[i] <= 10^4
* All values of nums are unique.
* nums is guaranteed to be rotated at some pivot.
* -10^4 <= target <= 10^4

## 題目大意
假設按照升冪排序的arr在未知的某個點上進行了選轉,(ex,arr[0,1,2,4,5,6,7]可能變成[4,5,6,7,0,1,2]),給一個指定的目標,如果在arr中找到這個數,則返回它的index,否則返回-1. </br>

要求: 時間複雜度O(log n)

## 邏輯思維
* 方法一: 透過map的方式來比對,時間複雜度O(n)
* 方法二: 透過二分搜尋法來找尋
  * 注意旋轉完,是左邊大還是右邊大
