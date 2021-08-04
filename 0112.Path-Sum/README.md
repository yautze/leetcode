# [112. Path Sum](https://leetcode.com/problems/path-sum/)

## 題目
Given the root of a binary tree and an integer targetSum, return true if the tree has a root-to-leaf path such that adding up all the values along the path equals targetSum.

A leaf is a node with no children.

Example1:
![](https://assets.leetcode.com/uploads/2021/01/18/pathsum1.jpg)
```
Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
Output: true
```

Example2:
![](https://assets.leetcode.com/uploads/2021/01/18/pathsum2.jpg)
```
Input: root = [1,2,3], targetSum = 5
Output: false
```

Example3:
```
Input: root = [1,2], targetSum = 0
Output: false
```

Constraints:

* The number of nodes in the tree is in the range [0, 5000].
* -1000 <= Node.val <= 1000
* -1000 <= targetSum <= 1000


## 題目大意
給一個二元樹和一個目標總和,判斷該二元樹中是否有從根節點root到葉節點leaf的路徑中所有的值相加等於目標總和</br>

葉節點 => 沒有子節點的節點

## 邏輯思維
用遞迴去輪詢不同的路徑,並將目標總和減去該路徑節點的直</br>

備註： 一個遞迴往左找,另一個遞迴往右找,其中一邊找到即可