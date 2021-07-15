# [567. Permutation in String](https://leetcode.com/problems/permutation-in-string/)

## 題目
Given two strings s1 and s2, return true if s2 contains the permutation of s1.

In other words, one of s1's permutations is the substring of s2.

Example1:
```
Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 contains one permutation of s1 ("ba").
```

Example2:
```
Input: s1 = "ab", s2 = "eidboaoo"
Output: false
```

Constraints:
* 1 <= s1.length, s2.length <= 10的4次
* s1 and s2 consist of <b>lowercase</b> English letters.

## 題目大意
在s2字串組合中是否擁有s1字串任何連續的排列組合 </br>
> Ex: s1 = "ab"</br>
排列組合 => "ab", "ba"</br>
> 如果s2字串內有上述的排列組合,即可回true

## 邏輯思維
#### Tip : 固定長度(s1的長度)的Slide Window的題目
Step.1 透過兩個長度為26的int array來儲存s1,s2字串的排列組合 </br></br>
> Ex: s1 = "ab"</br>
=> arr1[0] = 1, arr[1] = 1 </br>
arr1[0] 代表 'a' 出現的次數 </br>
arr1[1] 代表 'b' 出現的次數
>

Step.2 透過<b>slide window</b>來更新s2字串的排列組合 </br></br>
> 原先arr2儲存的排列組合為 'ei', 沒有對到需要往右再移一格 </br>
先往右加一個字元(eid), 再減到最左邊那一個字元(id)<br> 
得到往左移動一格後的arr2結果(id)<br>
以次類推來找是否有符合的排列組合
>

## 參考資訊
* [Slide Window](https://www.zhihu.com/question/314669016/answer/620247024)
* [邏輯思維講解](https://www.youtube.com/watch?v=wpq03MmEHIM)