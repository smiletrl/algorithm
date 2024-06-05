# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

from typing import Optional, List

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def reverseKGroup(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        d: List[int] = list(range(k))
        nHead = head
        it = head
        j = 0
        while head != None:
             d[j%k] = head.val
             j += 1
             head = head.next
             if j%k == 0:
                  for i in range(k-1, -1, -1):
                       it.val = d[i]
                       it = it.next
        return nHead

