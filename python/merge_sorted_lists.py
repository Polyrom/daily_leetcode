"""
21. You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists into one sorted list.
The list should be made by splicing together the nodes of the first two lists.
Return the head of the merged linked list.
"""

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def print_list(list_node: ListNode | None) -> None:
    if not list_node:
        print("Empty list")
        return
    if list_node.next:
        print_list(list_node.next)
    print(list_node.val)

def merge_two_lists(list1: ListNode | None, list2: ListNode | None) -> ListNode | None:
    if list1 is None:
        return list2
    if list2 is None:
        return list1
    if list1.val < list2.val:
        list1.next = merge_two_lists(list1.next, list2)
        return list1
    else:
        list2.next = merge_two_lists(list2.next, list1)
        return list2


print('Recursive approach. Per usual, shoutout to Leetcode Solutions experts')
l1 = ListNode(1, ListNode(2, ListNode(4)))
l2 = ListNode(3, ListNode(4, ListNode(6)))
print_list(merge_two_lists(l1, l2))
