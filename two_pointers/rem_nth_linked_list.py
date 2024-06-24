class ListNode:
    def __init__(self, val, next) -> None:
        self.val = val
        self.next = next


def remove_nth_from_ll(head: ListNode | None, n: int) -> ListNode | None:
    if not head:
        return None
    dummy = ListNode(val=0, next=head)
    fast, slow = dummy, dummy
    for _ in range(n):
        fast = fast.next
    while fast.next:
        fast, slow = fast.next, slow.next
    slow.next = slow.next.next
    return dummy.next
