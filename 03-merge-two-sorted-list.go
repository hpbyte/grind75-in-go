/**
 * Recursive approach
 * - Base case: if list1 is nil, return list2 vice versa
 * - if list1 value is less than list2 value
 * - the new list1 next should be the result of merging current list1 next and current list2
 * - else, the new list2 next should be the result of merging current list1 and current list2 next
 *
 * Naive approach in JS
 * - convert the two list into arrays
 * - join the two arrays and sort
 * - and rebuild the list from scratch
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
  if list1 == nil { return list2 }
  if list2 == nil { return list1 }

  if list1.Val < list2.Val {
    list1.Next = mergeTwoLists(list1.Next, list2)
    return list1
  } else {
    list2.Next = mergeTwoLists(list1, list2.Next)
    return list2
  }
}
