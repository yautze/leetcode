# [20. Valid Parentheses](https://leetcode.com/problems/valid-parentheses/)

## 題目
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.

Example1:
```
Input: s = "()"
Output: true
```

Example2:
```
Input: s = "()[]{}"
Output: true
```

Example3:
```
Input: s = "(]"
Output: false
```

Example4:
```
Input: s = "([)]"
Output: false
```

Example5:
```
Input: s = "{[]}"
Output: true
```

Constraints:
* 1 <= s.length <= 10^4
* s consists of parentheses only '()[]{}'.

## 題目大意
字串內是否符合括號的匹配

## 邏輯思維
建立一個暫存的stack arr來儲存左邊的括號, 且當遇到右邊的括號, 則比對最後一個存入stack中的括號是否是相對應的, 如果是則移除掉最後一個存入stack的括號.  </br>

最後stack的長度應為0才會代表這個字串內的括號是匹配的.