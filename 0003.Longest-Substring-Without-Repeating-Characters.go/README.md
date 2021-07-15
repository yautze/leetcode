# [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)


## 題目
Given a string s, find the length of the longest substring without repeating characters.

Example1:
```
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
```

Example2:
```
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
```

Example3:
```
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
```

Example4:
```
Input: s = ""
Output: 0
```

## 題目大意
在一個字串中,找尋連續且不重複的最長字串長度

## 邏輯思維
#### Tip : slide window 控制 left, right
* right - left + 1 = slide window length(取最大則是答案)
right不斷向右邊移, 直到遇到重複的字元後, 將left移動到該字元的位置右邊(該字元的index + 1)來移除重複的字元

