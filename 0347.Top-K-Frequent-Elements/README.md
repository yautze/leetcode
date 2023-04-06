# [347. Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/)

## 題目
- Degree of difficulty: Easy
- Category: Arrays & Hash

## 題目大意
給定一個整數array nums 和一個整數 k，返回前 k 個出現頻率最高的元素。

## 邏輯思維
先使用map紀錄每個數字出現的頻率，然後把map中所有的key取出來，根據其出現頻率由高到低排序，最後返回前k個元素。

## 複雜度
- 時間複雜度: O(nlogn)
- 空間複雜度: O(n)