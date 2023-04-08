# [136. Single Number](https://leetcode.com/problems/single-number/)

## 題目所屬
- Degree of difficulty: Easy
- Category: Bit Manipulation

## 題目大意
找尋陣列中沒有重複的數字

## 邏輯思維
#### 方法一 - 透過map計算每個數字出現的次數
#### 方法二 - 透過XOR的的二元運算子的特性
因為題目有強調只會重複出現兩次,所以可以透過XOR的特性(x ^ x = 0 && x ^ 0 = x)來找是否有重複
ex: 2 xor 2 xor 1 => 因為2 xor 2 = 0, 而0 xor 1 =1, 由此推斷出哪一個數字沒有重複
```
// x XOR x = 0

ex: x = 6 (二進位等於0110)

// 二元運算
     0110
XOR) 0110
-----------
     0000

=> 0000 = 0 
```

[XOR](https://www.notion.so/Bitwise-XOR-bab0ba170006420e96e064f813ff8910?pvs=4)

## 複雜度
- 方法1
     - 時間複雜度: O(n)
     - 空間複雜度: O(n)
- 方法2
     - 時間複雜度: O(n)
     - 空間複雜度: O(1)