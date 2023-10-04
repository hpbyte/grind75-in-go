/**
 * find the height of a node
 * the base would be 0 if the node is nil
 * if the difference between the left node's height and the right
 * is less than or equal to 1
 * the binary tree is balanced
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getHeight(node *TreeNode) int {
    if node == nil {
        return 0
    }

    leftHeight := getHeight(node.Left)
    rightHeight := getHeight(node.Right)

    if leftHeight == -1 || rightHeight == -1 || abs(leftHeight-rightHeight) > 1 {
        return -1
    }

    return max(leftHeight, rightHeight) + 1
}

func isBalanced(root *TreeNode) bool {
    return getHeight(root) != -1
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
