# [22. Generate Parentheses](https://leetcode.com/problems/generate-parentheses/)

## 題目
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

Constraints:
* 1 <= n <= 8

Example1:
```
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
```

Example2:
```
Input: n = 1
Output: ["()"]
```

## 題目大意
给一個數字,寫出一個函式,能夠產生有可能並且左右括號能對稱的組合。

## 邏輯思維
Tip
* 不要做左右括號匹配
* [DFS演算法](https://ithelp.ithome.com.tw/articles/10237752)

透過遞迴來讓左括號與又括號匹配

