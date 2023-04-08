# [238. Product of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/)

## 題目所屬
- Degree of difficulty: Medium
- Category: Arrays & Hash

## 題目大意
給定一個整數數組 nums，返回一個數組 res，其中 res[i] 等於 nums 中除 nums[i] 之外所有數的乘積。

## 邏輯思維
先計算 nums[i] 左邊的所有數字的乘積，將結果保存在 res 中。再計算 nums[i] 右邊的所有數字的乘積，並將其與 res 相乘，從而得到最終的結果。

## 複雜度
- 時間複雜度: O(n)
- 空間複雜度: O(1)