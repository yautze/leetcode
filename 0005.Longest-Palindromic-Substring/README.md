# [5. Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)

## 題目
Given a string s, return the longest palindromic substring in s

Example1:
```
Input: s = "babad"
Output: "bab"
Note: "aba" is also a valid answer.
```

Example2:
```
Input: s = "cbbd"
Output: "bb"
```

Example3:
```
Input: s = "a"
Output: "a"
```

Example4:
```
Input: s = "ac"
Output: "a"
```

## 題目大意
给你一個字串s,找到s中最長的回文子串。

## 邏輯思維
#### 解法1 - 窮舉法
用兩個迴圈列舉出每種字串的各種組合
```
Ex: aba

ab
aba
ba
```

#### 解法2 - 中心點擴散法
回文特性
* 奇數個字串(Ex:aba) => 中心為`b` </br>
往左擴散則是中心點(1)往左移動一格`-1` </br>
往右擴散則是中心點(1)往右移動一格`+1`

* 偶數個字串(Ex:bccb) => 中心點為`cc` </br>
往左擴散則是左中心點(1)往左移動一格`-1` </br>
往右擴散則是右中心點(1)往右移動一格`+1` </br>

由上述可得得知 </br>
奇數個中心點 => left = right </br>
偶數個中心點 => right = left + 1 </br>

所以建立一個func(maxPalindrome) 帶入字串s,並由其中心點向左右擴散
![](https://scontent.frmq2-2.fna.fbcdn.net/v/t1.15752-9/p1080x2048/219229115_253136196322435_1720194273016691161_n.jpg?_nc_cat=104&ccb=1-3&_nc_sid=ae9488&_nc_ohc=S5UsiSXz2TQAX_ioD0i&_nc_ht=scontent.frmq2-2.fna&oh=22f4fdd17406d81902fb74a1652aeb52&oe=61223FAD)