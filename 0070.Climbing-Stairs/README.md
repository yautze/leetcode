# [70. Climbing Stairs](https://leetcode.com/problems/climbing-stairs/)

## 題目
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example1:
```
Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
```

Example2:
```
Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
```

Constraints:
* 1 <= n <= 45

## 題目大意
假設是在爬樓梯,你必續爬到n階才算爬到頂層,一次只能爬一階或兩階,最多有幾種走法

## 邏輯思維
* 典型DP的樓梯問題
* 可參考費氏數列

爬三階的走法 = 爬兩階加爬一階的相加

[文章參考](https://blog.chairco.me/posts/2017/08/Understand_Dynamic_Programming_Algorithm_usining_Python.html?fbclid=IwAR0zBv1Vqbk9u7yi2q3gFFNQ5Que_wRyeeN29kt3HgBmYmoK5LjO82_BQsw)