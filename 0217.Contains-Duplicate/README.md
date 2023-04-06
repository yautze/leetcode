# [217. Contains Duplicate](https://leetcode.com/problems/contains-duplicate/)

## 題目
- Degree of difficulty: Easy
- Category: Arrays & Hash

## 題目大意
給定一個integer array, 判斷是否有重複的值

## 邏輯思維
使用map來紀錄已出現過的數字, 下次再出現該數字時, 直接return true

## 複雜度
- 時間複雜度: O(n)
- 空間複雜度: O(n)