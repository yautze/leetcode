# [1. Two Sum](https://leetcode.com/problems/two-sum/)

## 題目
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example1:
```
Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1]
```

Example2:
```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

Example3:
```
Input: nums = [3,3], target = 6
Output: [0,1]
```

## 題目大意
在陣列中找到2個數字相加等於給定值得數字,結果返回這兩個數字在陣列中的位置(index)

## 邏輯思維
將給定值減去自己的數字後,再去找Map中是否有等於減去後的值,沒有則放入Map</br>
map key = value -> 比對用</br>
map value = index -> 拿位置