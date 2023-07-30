/**
 * Recursive approach:
 * - Base case:
 *   - if the incoming node is nil, return the node
 * - put either left or right of the node (left in this case) in a temporary var
 * - swap left and right
 * - recursively call the function again with left and right respectively
 * - in the end return the node
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
  if root == nil {
    return root
  }

  left := root.Left
  root.Left = root.Right
  root.Right = left

  invertTree(root.Left)
  invertTree(root.Right)

  return root
}
