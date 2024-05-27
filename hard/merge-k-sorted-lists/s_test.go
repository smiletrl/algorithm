package solution

import (
	"testing"
)

func TestMergeKLists(t *testing.T) {
	cases := []struct {
		name   string
		input  [][]int
		expect []int
	}{
		{
			name: "case 1",
			input: [][]int{
				{1, 4, 5},
				{1, 3, 4},
				{2, 6},
			},
			expect: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
	}
	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			lists := make([]*ListNode, len(ca.input))
			for i, valArr := range ca.input {
				if len(valArr) == 0 {
					continue
				}
				lists[i] = &ListNode{}
				head := lists[i]
				for j, val := range valArr {
					head.Val = val
					if j != len(valArr)-1 {
						node := &ListNode{}
						head.Next = node
						head = head.Next
					}
				}
			}
			out := mergeKLists(lists)
			x := 0
			for out != nil {
				if ca.expect[x] != out.Val {
					t.Fatalf("expect not match")
				}
				x++
				out = out.Next
			}
		})
	}
}
