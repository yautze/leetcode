# [2. Add Two Numbers](https://leetcode.com/problems/add-two-numbers/)

## 題目
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example1:
![](https://assets.leetcode.com/uploads/2020/10/02/addtwonumber1.jpg)
```
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
```

Example2:
```
Input: l1 = [0], l2 = [0]
Output: [0]
```

Example3:
```
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
```

## 題目大意
2個逆序的鏈結，要求從低位數開始相加，得出结果也是逆序輸出，返回结果是逆序的兩個鏈結相加。

## 邏輯思維
迴圈取出兩個數字的每一個位數(n1,n2)和上一個位數的進位(carry)相加,若有超出總位數的相加必須往後再進一位的鏈結