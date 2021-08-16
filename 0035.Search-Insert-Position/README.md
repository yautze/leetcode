# [35. Search Insert Position](https://leetcode.com/problems/search-insert-position/)

## 題目
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.

Example1:
```
Input: nums = [1,3,5,6], target = 5
Output: 2
```

Example2:
```
Input: nums = [1,3,5,6], target = 2
Output: 1
```

Example3:
```
Input: nums = [1,3,5,6], target = 7
Output: 4
```

Example4:
```
Input: nums = [1,3,5,6], target = 0
Output: 0
```

Example5:
```
Input: nums = [1], target = 0
Output: 0
```


Constraints:

* 1 <= nums.length <= 10^4
* -10$4 <= nums[i] <= 10^4
* nums contains distinct values sorted in ascending order.
* -10^4 <= target <= 10^4

## 題目大意
给一個排序後的int arr和一個目標,在arr中找到目標,並返回其索引,如果目標不存在arr中,則返回他應該被插入的位置.</br>

要求: 時間複雜度O(log n)

## 邏輯思維
經典二分搜尋法的題目,注意找不到目標的狀況(left > right時)
```
Ex: [1, 2, 4, 5], target = 3
初始 left = 0, right = 3 , mid = 1
執行第一次後 => left = 2, right = 3, mid = 2
執行第二次後 => left = 2, right = 1 => 找不到目標值此時要回`left`才是應該被插入的值
```