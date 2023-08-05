/**
 * The input here is a Binary Search Tree (BST)
 * which means that the left nodes are always less than the right nodes
 * min = min(p, q), max = max(p, q)
 * So a lowest common ancestor is the one that is
 * min <= LCA <= max
 * if the current node doesn't resolve the condition
 * move to the left branch if the current node is greater than the max
 * else, move to the right branch
 */

import "math"

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
  Min := int(math.Min(float64(p.Val), float64(q.Val)))
  Max := int(math.Max(float64(p.Val), float64(q.Val)))

  for true {
    if Min <= root.Val && root.Val <= Max {
      break
    }

    if (root.Val > Max) {
      root = root.Left
    } else {
      root = root.Right
    }
  }

  return root
}
