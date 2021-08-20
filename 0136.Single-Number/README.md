# [136. Single Number](https://leetcode.com/problems/single-number/)

## 題目
Given a non-empty array of integers nums, every element appears `twice` except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.

Example1:
```
Input: nums = [2,2,1]
Output: 1
```

Example2:
```
Input: nums = [4,1,2,1,2]
Output: 4
```

Example3:
```
Input: nums = [1]
Output: 1
```

Constraints:

* 1 <= nums.length <= 3 * 10^4
* -3 * 10^4 <= nums[i] <= 3 * 10^4
* Each element in the array appears `twice` except for one element which appears only once.

## 題目大意
找尋陣列中沒有重複的數字

## 邏輯思維
#### 方法一 - 透過map計算每個數字出現的次數
#### 方法二 - 透過XOR的的二元運算子的特性
因為題目有強調只會重複出現兩次,所以可以透過XOR的特性來找是否有重複
```
// x XOR x = 0

ex: x = 6 (二進位等於1100)

// 二元運算
     1100
XOR) 1100
-----------
     0000

=> 0000 = 0 
```

[二元運算子補充](https://opensourcedoc.com/golang-programming/operator/)