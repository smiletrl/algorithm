package solution

import (
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		key    int
		expect []int
	}{
		{
			name:   "case 1",
			input:  []int{1, 4, 5, 7, 8},
			key:    2,
			expect: []int{4, 1, 7, 5, 8},
		},
	}
	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			head := &ListNode{}
			ohead := head
			for i, val := range ca.input {
				head.Val = val
				if i != len(ca.input)-1 {
					node := &ListNode{}
					head.Next = node
					head = node
				}
			}
			out := reverseKGroup(ohead, ca.key)
			x := 0
			for out != nil {
				if ca.expect[x] != out.Val {
					t.Fatalf("expect not match: expect: %d, real: %d, x: %d", ca.expect[x], out.Val, x)
				}
				x++
				out = out.Next
			}
		})
	}
}
