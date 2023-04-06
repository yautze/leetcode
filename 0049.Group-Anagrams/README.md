# [49. Group Anagrams](https://leetcode.com/problems/group-anagrams/)

## 題目
- Degree of difficulty: Medium
- Category: Arrays & Hash

## 題目大意
给一個str arry,要求對str arr中有Anagrams的關係字串進行分類.<br/>

* Anagrams是指兩個字串的字符完全相同,但顺序可以不同,兩個字串的英文字母排列組合是一樣的
    * ex: "str", "rts"

## 邏輯思維
#### 方法一 : 排序字串的英文字母順序
* ex: bac, cba </br>
將兩個字串的字母由小排序到大 => "abc", "abc", 即可判斷是否符合條件

#### 方法二 : 透過計算每個字串出現的英文字母順序
* ex: bac => a:1, b:1, c:1</br>
* ex: cba => a:1, b:1, c:1</br>
將兩個字串的字母出現次數放進一個長度為26的int arr中,判斷兩個字串的arr是否相等

## 複雜度
1. 方法ㄧ
    - 時間複雜度: O(n * k * log(k))
    - 空間複雜度: O(n * k)
2. 方法二
    - 時間複雜度: O(n * m)
    - 空間複雜度: O(n * m)