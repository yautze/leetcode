# [1. Two Sum](https://leetcode.com/problems/two-sum/)

## 題目
- Degree of difficulty: Easy
- Category: Arrays & Hash

## 題目大意
在陣列中找到2個數字相加等於給定值得數字,結果返回這兩個數字在陣列中的位置(index)

## 邏輯思維
將給定值減去自己的數字後,再去找Map中是否有等於減去後的值,沒有則放入Map</br>
map key = value -> 比對用</br>
map value = index -> 拿位置

## 複雜度
- 時間複雜度: O(n)
- 空間複雜度: O(n)