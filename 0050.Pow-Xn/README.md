# [50. Pow(x, n)](https://leetcode.com/problems/powx-n/)

## 題目
Implement pow(x, n), which calculates x raised to the power n (i.e., x^n).

Example1:
```
Input: x = 2.00000, n = 10
Output: 1024.00000
```

Example2:
```
Input: x = 2.10000, n = 3
Output: 9.26100
```

Example3:
```
Input: x = 2.00000, n = -2
Output: 0.25000
Explanation: 2^-2 = 1/2^2 = 1/4 = 0.25
```
Constraints:

* -100.0 < x < 100.0
* -2^31 <= n <= 2^31-1
* -10^4 <= xn <= 10^4


## 題目大意
實作pow(x, n),即計算 x 的 n 次方。

## 邏輯思維
```
2^10 => 2^5 * 2^5 
2^5 => 2^2 * 2^2 * 2
```

用遞徊的方式找到x的(n/2)次方後相乘</br>
若 n 是奇數,則多乘一次x