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

#### 思路一：通過單階段之間的關係確定狀態轉移方程
由於一次只能上1個或2個臺階，所以假如我當前在第n個臺階上，那麼我上一次應該在第n-1個或者第n-2個臺階上。換句話說，當我知道到達第n-1個臺階的方案數f(n-1)，到達第n-2個臺階的方案數是f(n-2)。
那麼，我們就自然知道到達第n個階梯時的方案數，即：f(n)=f(n-1)+f(n-2)，這也被稱為此問題的狀態轉移方程。

接著只需要界定邊界或特殊狀況</br>
* f(0) = 1 => 爬0階等於只有一種(不用爬)
* f(1) = 1 => 爬一階只有一種
* f(2) = f(2-1) + f(2-2) = f(1) + f(0) = 2

[文章參考](https://blog.chairco.me/posts/2017/08/Understand_Dynamic_Programming_Algorithm_usining_Python.html?fbclid=IwAR0zBv1Vqbk9u7yi2q3gFFNQ5Que_wRyeeN29kt3HgBmYmoK5LjO82_BQsw)