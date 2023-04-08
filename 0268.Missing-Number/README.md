# [268. Missing Number](https://leetcode.com/problems/missing-number/description/)

## 題目所屬
- Degree of difficulty: Easy
- Category: Bit Manipulation

## 題目大意
在一個包含`[0, n]`範圍內，不同數字的int array，返回該範圍內中唯一缺缺少的數字。

## 邏輯思維
### 方法1 透過xor的特性(a xor a = 0)
先xor一次0 ~ n的數字，再依序xor nums裏面的每個數字，最後結果則是缺少的數字
備註: 用xor是為了降低空間複雜度

[XOR](https://www.notion.so/Bitwise-XOR-bab0ba170006420e96e064f813ff8910?pvs=4)

### 方法2 透過cyclic sort
因為題目已有表示是0 ~ n的範圍, 所以可以透過cyclic sort來排序這些數字, 如果數字不在對應的index上, 代表該index是缺少的數字

[cyclic sort](https://www.notion.so/Cyclic-Sort-7cdc2b3906774ab690fbfd71d447f2e8?pvs=4)

## 複雜度
- 方法1
     - 時間複雜度: O(n)
     - 空間複雜度: O(1)
- 方法2
     - 時間複雜度: O(n)
     - 空間複雜度: O(n)