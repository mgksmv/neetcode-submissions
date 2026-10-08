# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:
        if not head or not head.next:
            return

        l1 = head

        center_node = self.findCenter(head)
        l2 = self.reverseList(center_node)

        while l2.next:
            l1_next = l1.next
            l2_next = l2.next
            l1.next = l2
            l2.next = l1_next

            l1 = l1_next
            l2 = l2_next

    def findCenter(self, head: ListNode) -> ListNode:
        slow = head
        fast = head

        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next

        return slow

    def reverseList(self, head: ListNode) -> ListNode:
        curr_node = head
        prev_node = None

        while curr_node:
            next_node = curr_node.next
            curr_node.next = prev_node
            prev_node = curr_node
            curr_node = next_node

        return prev_node
