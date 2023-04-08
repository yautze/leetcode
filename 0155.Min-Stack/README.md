# [155. Min Stack](https://leetcode.com/problems/min-stack/)

## 題目所屬
- Degree of difficulty: Medium
- Category: Stack

## 題目大意
設計一個stack，支持push、pop、top，並在O(1)時間內取得stack內最小的元素。

## 邏輯思維
透過兩個slice來紀錄, 一個用來儲存現在stack內的所有元素, 另一個則是紀錄stack內最小元素得歷程
維持這兩個slice長度會是一致的, 這樣可以簡化處理最小元素的邏輯
  1. 每次push元素時，都會與最小值歷程中的最頂部數字進行比較，如果小於，則將該元素加入最小值歷程中，否則加入最小值歷程中的最頂部數字。可以理解為每次添加時都記錄當時的最小值。
  2. pop時, 則是同時將這兩個slice最頂部的數字一起移
