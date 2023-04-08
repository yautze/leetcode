# [206. Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)

## 題目所屬
- Degree of difficulty: Easy
- Category: Linked List

## 題目大意
反轉一個Linked List

## 邏輯思維
每次將head的next節點(nextNode)指向反轉後的前一個節點(preNode)，然後再將反轉後的前一個節點更新為當前節點(head)，最後將head更新為原先的next節點。

## 複雜度
- 時間複雜度: O(n)