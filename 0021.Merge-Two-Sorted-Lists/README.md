# [21. Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/)

## 題目
Merge two sorted linked lists and return it as a sorted list. The list should be made by splicing together the nodes of the first two lists.

Example1:
![](https://assets.leetcode.com/uploads/2020/10/03/merge_ex1.jpg)
```
Input: l1 = [1,2,4], l2 = [1,3,4]
Output: [1,1,2,3,4,4]
```

Example2:
```
Input: l1 = [], l2 = []
Output: []
```

Example3:
```
Input: l1 = [], l2 = [0]
Output: [0]
```

Constraints:
* The number of nodes in both lists is in the range [0, 50].
* -100 <= Node.val <= 100
* Both l1 and l2 are sorted in non-decreasing order.

## 題目大意
將兩個已經排序好的連結串列,合併成一個排序好的鏈結串列

## 邏輯思維
循環兩個串列直到某一個串列沒有值,將較小的那個數鏈結到res中,最後再將還有剩下的串列也鏈結到res的最後面即時答案