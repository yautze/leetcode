# [20. Valid Parentheses](https://leetcode.com/problems/valid-parentheses/)

## 題目
- Degree of difficulty: Easy
- Category: Stack

## 題目大意
字串內是否符合括號的匹配

## 邏輯思維
建立一個暫存的stack arr來儲存左邊的括號, 且當遇到右邊的括號, 則比對最後一個存入stack中的括號是否是相對應的, 如果是則移除掉最後一個存入stack的括號.  </br>

最後stack的長度應為0才會代表這個字串內的括號是匹配的.

## 複雜度
- 時間複雜度: O(n)
- 空間複雜度: O(n)