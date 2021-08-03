# [55. Jump Game](https://leetcode.com/problems/jump-game/)

## 題目
You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.

Example1:
```
Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
```

Example2:
```
Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
```

Example3:
```
Input: nums = [1]
Output: true
```

## 題目大意
给一個非負數的int arr，從第一個位置,輪詢arr中的每個數字代表可以從該位置跳躍的最大距離。判斷是否能夠到達arr的最後位置

## 邏輯思維
依照題意的邏輯寫for迴圈輪詢