package solution

// https://leetcode.com/problems/median-of-two-sorted-arrays/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}
	arr := mergeArray(nums1, nums2)
	return getVal(arr)
}

func mergeArray(nums1 []int, nums2 []int) []int {
	le := len(nums1) + len(nums2)
	res := make([]int, le)
	i := 0
	j := 0
	for k := 0; k < le; k++ {
		if i >= len(nums1) {
			res[k] = nums2[j]
			j++
			continue
		} else if j >= len(nums2) {
			res[k] = nums1[i]
			i++
			continue
		}
		if nums1[i] <= nums2[j] {
			res[k] = nums1[i]
			i++
		} else {
			res[k] = nums2[j]
			j++
		}
	}

	return res
}

func getVal(arr []int) float64 {
	le := len(arr)
	mid := le / 2
	if le%2 == 0 {
		return (float64(arr[mid-1]) + float64(arr[mid])) / 2
	}
	return float64(arr[mid])
}
