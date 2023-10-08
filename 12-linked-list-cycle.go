/**
 * Use two pointer approach
 * - slow pointer and a fast pointer
 * - loop thru the linked list
 * - when the two pointers matched, means there's a cycle
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
  slow := head
  fast := head

  for fast != nil && slow.Next != nil {
    fast = fast.Next.Next
    slow = slow.Next

    if slow == fast {
      return true
    }
  }

  return false
}
