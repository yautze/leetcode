# [34. Find First and Last Position of Element in Sorted Array](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/)

## 題目
Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with `O(log n)` runtime complexity.

Example1:
```
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
```

Example2:
```
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
```

Example3:
```
Input: nums = [], target = 0
Output: [-1,-1]
```

Constraints:

* 0 <= nums.length <= 10^5
* -10^9 <= nums[i] <= 10^9
* nums is a non-decreasing array.
* -10^9 <= target <= 10^9

## 題目大意
给一個升升冪排序的nums arry，和一個目標值target,找出target在arr中開始位置和結束位置.

要求: 時間複雜度O(log n)

## 邏輯思維
* 經典二分搜尋法的題目

用兩個二分搜尋法,一個找開始位置,一個找結束位置,當找到有值=target時,需特別判斷是否是起始或結束位置.

#### 找到起始位置的條件
> ex: [`1`, 1, 2, 3, 4] => 往前找到第一個數,且這個數等於target<br/>
  `nums[mid] == target && mid == 0`
>
>  ex: [0, `1`, 1, 2, 3] => 找到一個數,這個數等於target,且這個數的`前`一個數字不等於target <br/>
> `nums[mid] == target && nums[mid - 1] != target`


#### 找到結束位置的條件
> ex: [0, 0, 1, 2, `2`] => 往後找到最後一個數,且這個數等於target <br/>
  nums[mid] == target && mid == len(nums) - 1
> 
> ex: [1, `1`, 2, 3, 4] => 找到一個數,這個數等於target,且這個數的`下`一個數字不等於target <br/>
> nums[mid] == target && nums[mid + 1] != target

