# [242. Valid Anagram](https://leetcode.com/problems/valid-anagram/description/)

## 題目
- Degree of difficulty: Easy
- Category: Arrays & Hash

## 題目大意
檢查這兩個字串是否是Anagram
Anagram: 通過重新排列不同單詞或短語的字母而形成的單詞或短語，通常只使用所有原始字母一次。

## 邏輯思維
先檢查兩個字串的長度是否相同，若不相同則直接返回 false。接著利用一個長度為 26 的整數陣列 freq，來記錄字母在字串 s 和 t 中出現的頻率。對於字串 s 中的每個字母，將對應的頻率加 1，對於字串 t 中的每個字母，將對應的頻率減 1。最後檢查 freq 中的每個元素是否都為 0，若都為 0 則表示兩個字串為 anagram，否則不是

## 複雜度
- 時間複雜度: O(n)
- 空間複雜度: O(1)