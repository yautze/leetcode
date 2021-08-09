# [66. Plus One](https://leetcode.com/problems/plus-one/)

## 題目
Given a non-empty array of decimal digits representing a non-negative integer, increment one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contains a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

Example1:
```
Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
```

Example2:
```
Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
```

Example3:
```
Input: digits = [0]
Output: [1]
```

Constraints:

* 1 <= digits.length <= 100
* 0 <= digits[i] <= 9

## 題目大意
給一個沒有複數且沒有null的int arr,在該陣列的基礎上+1.

## 邏輯思維
依照題意, 注意進位的問題