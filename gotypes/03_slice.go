package gotypes

import (
	"fmt"
)

func remove(slice []interface{}, s int) []interface{} {
	return append(slice[:s], slice[s+1:]...)
}

/*
DemoSlice show how slice work in go
*/
func DemoSlice() {
	/* empty slice */
	var nums1 []int
	fmt.Printf("Round1 nums1: %v, pntr: %p, len: %d, cap: %d \n", nums1, nums1, len(nums1), cap(nums1))

	nums1 = append(nums1, 0)
	fmt.Printf("Round2 nums1: %v, pntr: %p, len: %d, cap: %d \n", nums1, nums1, len(nums1), cap(nums1))

	nums1 = append(nums1, 1)
	fmt.Printf("Round3 nums1: %v, pntr: %p, len: %d, cap: %d \n", nums1, nums1, len(nums1), cap(nums1))

	/* empty slice with `make` and declare len cap */
	nums2 := make([]int, 5, 6)
	fmt.Printf("Round1 nums2: %v, pntr: %p, len: %d, cap: %d \n", nums2, nums2, len(nums2), cap(nums2))
	nums2[0] = 0
	nums2[1] = 1
	nums2[2] = 2
	nums2[3] = 3
	nums2[4] = 4
	fmt.Printf("Round2 nums2: %v, pntr: %p, len: %d, cap: %d \n", nums2, nums2, len(nums2), cap(nums2))
	nums2 = append(nums2, 5)

	// nums3 := nums2[:3]

	fmt.Printf("Round3 nums2: %v, pntr: %p, len: %d, cap: %d \n", nums2, nums2, len(nums2), cap(nums2))
	cNums2 := make([]int, len(nums2), cap(nums2))
	copy(cNums2, nums2)
	fmt.Printf("cp nums2: %v, pntr: %p, len: %d, cap: %d \n", cNums2, cNums2, len(cNums2), cap(cNums2))

	nums2 = append(nums2, 6)
	fmt.Printf("Round4 nums2: %v, pntr: %p, len: %d, cap: %d \n", nums2, nums2, len(nums2), cap(nums2))
}
