# [19. Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)

## 題目
Given the head of a linked list, remove the nth node from the end of the list and return its head.

Example1:
![](https://assets.leetcode.com/uploads/2020/10/03/remove_ex1.jpg)
```
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
```

Example2:
```
Input: head = [1], n = 1
Output: []
```

Example3:
```
Input: head = [1,2], n = 1
Output: [1]
```

Constraints:

* The number of nodes in the list is sz.
* 1 <= sz <= 30
* 0 <= Node.val <= 100
* 1 <= n <= sz

## 題目大意
刪除鏈結中倒數第 n 個Node

## 邏輯思維
當移動到該移除節點的前一個位置時,將節點的Next = 下兩個節點的位置,再將下一個節點的Next取消掉(nil)即可完成