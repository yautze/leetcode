# [283. Move Zeroes](https://leetcode.com/problems/move-zeroes/)

## 題目
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.

Example1:
```
Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
```

Example2:
```
Input: nums = [0]
Output: [0]
```

Constraints:

* 1 <= nums.length <= 10^4
* -2^31 <= nums[i] <= 2&31 - 1


## 題目大意
不能使用额外宣告的array，將array中0的數字都移動到array的尾端，並保持所有非0的數值的相對位置。



## 邏輯思維
透過一個迴圈輪詢,當取得的數值不為0,不等於自己的時候,將此數的位置跟count交換 </br>
> count => 紀錄在count之前的數,都是不為0的數值

![](https://scontent.xx.fbcdn.net/v/t1.15752-9/p206x206/219702454_301328108441108_1964750075020879889_n.jpg?_nc_cat=104&ccb=1-3&_nc_sid=aee45a&_nc_ohc=dGpcYYU_RYUAX_UyO4i&_nc_ad=z-m&_nc_cid=0&_nc_ht=scontent.xx&oh=7dee90b69d8485dcaf366ddcd4992c19&oe=61249445)