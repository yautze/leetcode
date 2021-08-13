# [49. Group Anagrams](https://leetcode.com/problems/group-anagrams/)

## 題目
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An `Anagram` is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example1:
```
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
```

Example2:
```
Input: strs = [""]
Output: [[""]]
```

Example3:
```
Input: strs = ["a"]
Output: [["a"]]
```

Constraints:

* 1 <= strs.length <= 10^4
* 0 <= strs[i].length <= 100
* strs[i] consists of `lower-case` English letters.

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