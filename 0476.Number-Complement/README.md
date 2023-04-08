# [476. Number Complement](https://leetcode.com/problems/number-complement/description/)

## 題目所屬
- Degree of difficulty: Easy
- Category: Bit Manipulation

## 題目大意
將一個整數轉換為二進位表示後，將其中的 0 變為 1，1 變為 0，所得到的新數字即為原數字的complement
ex: 5 => "101", complement => "010" => 2

## 邏輯思維
將該數字的每一個進位都與1做xor即可得到complement

[XOR](https://www.notion.so/Bitwise-XOR-bab0ba170006420e96e064f813ff8910?pvs=4)

## 複雜度
- 時間複雜度: O(log n)
- 空間複雜度: O(1)