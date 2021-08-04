package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

// TreeNode -
type TreeNode = structures.TreeNode

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	// 結束點: 在最後一層(leaf) 且符合上層數值相加後 == sum
	if root.Val == sum && root.Left == nil && root.Right == nil {
		return true
	}

	// 不是最後一層往下ㄧ層左右繼續找
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
