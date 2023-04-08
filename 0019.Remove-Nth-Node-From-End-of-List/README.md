# [19. Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)

## 題目所屬
- Degree of difficulty: Medium
- Category: Linked List

## 題目大意
刪除鏈結中倒數第 n 個Node

## 邏輯思維
當移動到該移除節點的前一個位置時(len(ListNode) - i - 1)，將節點的Next更新成下兩個節點的位置，再將下一個節點的Next取消掉(nil)即可完成

## 複雜度
- 時間複雜度: O(n)
- 空間複雜度: O(1)