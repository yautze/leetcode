# [28. Implement strStr()](https://leetcode.com/problems/implement-strstr/)


## 題目
implement `strStr()`.

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().

Example1:
```
Input: haystack = "hello", needle = "ll"
Output: 2
```

Example2:
```
Input: haystack = "aaaaa", needle = "bba"
Output: -1
```

Example3:
```
Input: haystack = "", needle = ""
Output: 0
```

Constraints:

* 0 <= haystack.length, needle.length <= 5 * 10^4
* haystack and needle consist of only `lower-case` English characters.


## 題目大意
給一個字串和一個目標,在這個字串內找尋目標,如果找到目標則回起始的index,否則會傳-1

## 邏輯思維
#### 方法一 - 偷懶使用strings.Index()

#### 方法二 - 使用迴圈來做固定長度的slide windows移找找尋

