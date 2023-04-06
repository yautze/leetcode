# [0739. Daily Temperatures](https://leetcode.com/problems/daily-temperatures/)

## 題目
- Degree of difficulty: Medium
- Category: Stack

## 題目大意
給定一個整數陣列temperatures，表示每天的氣溫，請回傳一個陣列answer，其中answer[i]表示在第i天後需要等待多少天才能得到更高的氣溫。如果未來沒有更高的氣溫，則answer[i]保持為0。

## 邏輯思維
透過一個stack來記錄尚未找到更高溫度的index, 當stack長度不為0, 且出現比stack頂部index對應的temperatures還大時, 斷從stack頂部取出lastIdx, 並計算i和lastIdx之間的距離, 最後將結果存儲在res[lastIdx]中(要存到原本比對的index中, 因為是找到比原本index還大的數字)

## 複雜度
- 時間複雜度: O(n)
- 空間複雜度: O(n)
