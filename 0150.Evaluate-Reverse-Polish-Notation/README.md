# [150. Evaluate Reverse Polish Notation](https://leetcode.com/problems/evaluate-reverse-polish-notation/)

## 題目所屬
- Degree of difficulty: Medium
- Category: Stack

## 題目大意
實作Reverse Polish Notation表達式的計算邏輯

## 邏輯思維
使用一個stack來處理表達式，當遇到數字時，將其推入堆疊中，當遇到操作符(+、-、*、/)時，從stack中取出最上面的兩個數字進行運算，將結果推回堆疊中。最後，stack中只會剩下一個數字，就是表達式的值。

## 複雜度
- 時間複雜度: O(n)
- 空間複雜度: O(n)
