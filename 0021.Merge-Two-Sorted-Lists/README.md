# [21. Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/)

## 題目所屬
- Degree of difficulty: Easy
- Category: Linked List

## 題目大意
將兩個已經排序好的連結串列,合併成一個排序好的鏈結串列

## 邏輯思維
循環兩個串列直到某一個串列沒有值,將較小的那個數鏈結到res中,最後再將還有剩下的串列也鏈結到res的最後面即時答案

## 複雜度
- 時間複雜度: O(m + n)
- 空間複雜度: O(1)