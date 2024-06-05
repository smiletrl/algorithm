package solution

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	d := make([]int, k)
	nHead := head
	it := head

	j := 0
	for head != nil {
		d[j%k] = head.Val
		j++
		head = head.Next

		if j%k == 0 {
			for i := k - 1; i > -1; i-- {
				it.Val = d[i]
				it = it.Next
			}
		}
	}
	return nHead
}

type ListNode struct {
	Val  int
	Next *ListNode
}
