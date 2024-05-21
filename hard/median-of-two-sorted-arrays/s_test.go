package solution

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	cases := []struct {
		name   string
		nums1  []int
		nums2  []int
		expect float64
	}{
		{
			name:   "case 1",
			nums1:  []int{1, 3},
			nums2:  []int{2},
			expect: 2.00000,
		},
	}

	for _, ca := range cases {
		res := findMedianSortedArrays(ca.nums1, ca.nums2)
		if res != ca.expect {
			t.Fatalf("val: %f is not expected: %f in case: %s", res, ca.expect, ca.name)
		}
	}
}
