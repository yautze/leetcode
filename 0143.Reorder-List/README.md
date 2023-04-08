# [143. Reorder List](https://leetcode.com/problems/reorder-list/)

## 題目所屬
- Degree of difficulty: Medium
- Category: Linked List

## 題目大意
重新排序整個Linked List

## 邏輯思維
觀察整個測試資料的排序規律，可以發現到是將Linked List切分成前半段與後半段後，再將反轉的後半段與前半部分交叉合併，所以可以先透過Fast & Slow Pointer找到Linked List的中間位置, 再將後半段的部分反轉，最後交叉合併即可。

反轉Linked List, 可參考leetcode 206題
[Fast & Slow Pointer](https://www.notion.so/Fast-Slow-Pointer-0b024d837c534f0b8f161b118c0fdd48?pvs=4)

## 複雜度
- 時間複雜度: O(n)
- 空間複雜度: O(1)