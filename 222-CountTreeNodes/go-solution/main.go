/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Yeah this was too easy, I just coded it on leetcode so
// no test cases here
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	} else if root.Left == nil {
		return 1 + countNodes(root.Right)
	} else if root.Right == nil {
		return 1 + countNodes(root.Left)
	} else {
		return 1 + countNodes(root.Left) + countNodes(root.Right)
	}
}
