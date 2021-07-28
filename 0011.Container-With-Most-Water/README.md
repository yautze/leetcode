# [11. Container With Most Water](https://leetcode.com/problems/container-with-most-water/)

## 題目
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of the line i is at (i, ai) and (i, 0). Find two lines, which, together with the x-axis forms a container, such that the container contains the most water.

Notice that you may not slant the container.

Example1:
![](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/17/question_11.jpg)
```
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
```

Example2:
```
Input: height = [1,1]
Output: 1
```

Example3:
```
Input: height = [4,3,2,1,4]
Output: 16
```

Example 4:
```
Input: height = [1,2,1]
Output: 2
```

Constraints:

* n == height.length
* 2 <= n <= 10^5
* 0 <= height[i] <= 10^4


## 題目大意
給一個array,找尋面積最大數字


## 邏輯思維
透過left,right兩指針來計算面積

Tip: 
1. width = 兩指針相減(right - left)
2. hight = 兩指針所對應到數字,且是較小的那一邊
    * Ex: a[0] = 1, a[8] = 7 => hight = 1
    * 當確定較短的那一邊後,將該邊的指針向中新移動``(left++, right--)``