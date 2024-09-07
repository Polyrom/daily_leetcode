/*
83. Remove Duplicates from Sorted List

Given the head of a sorted linked list, delete all duplicates such that each element appears only once.
Return the linked list sorted as well.
*/
package linkedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

// Memory: O(n). Space: O(1)
func deleteDuplicates(head *ListNode) *ListNode {
	pntr := head
	for pntr != nil && pntr.Next != nil {
		if pntr.Val == pntr.Next.Val {
			pntr.Next = pntr.Next.Next
		} else {
			pntr = pntr.Next
		}
	}
	return head
}
