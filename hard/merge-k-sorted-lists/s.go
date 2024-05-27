package solution

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	re := lists[0]

	for i := 1; i < len(lists); i++ {
		re = sort(re, lists[i])
	}
	return re
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sort(a, b *ListNode) *ListNode {
	start := &ListNode{}

	c := start

	for a != nil || b != nil {
		if a == nil && b != nil {
			node := &ListNode{
				Val: b.Val,
			}
			c.Next = node
			c = c.Next
			b = b.Next
			continue
		}

		if a != nil && b == nil {
			node := &ListNode{
				Val: a.Val,
			}
			c.Next = node
			c = c.Next
			a = a.Next
			continue
		}

		if a.Val < b.Val {
			node := &ListNode{
				Val: a.Val,
			}
			c.Next = node
			c = c.Next
			a = a.Next
		} else {
			node := &ListNode{
				Val: b.Val,
			}
			c.Next = node
			c = c.Next
			b = b.Next
		}
	}
	return start.Next
}
