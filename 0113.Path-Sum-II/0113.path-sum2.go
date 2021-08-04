package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// TreeNode -
type TreeNode = structures.TreeNode

func pathSum(root *TreeNode, targetSum int) [][]int {
	// 儲存結果
	res := [][]int{}
	res = findPaths(root, targetSum, res, []int{})
	return res
}

// findPaths -  stack 暫存結果
func findPaths(node *TreeNode, targetSum int, res [][]int, stack []int) [][]int {
	// 該層判斷
	if node == nil {
		return res
	}

	targetSum -= node.Val
	stack = append(stack, node.Val)

	// 最底層(leaf)
	if targetSum == 0 && node.Left == nil && node.Right == nil {
		res = append(res, append([]int{}, stack...))
	}

	// 往下一層的左右找
	res = findPaths(node.Left, targetSum, res, stack)
	res = findPaths(node.Right, targetSum, res, stack)

	return res
}
