package main

import (
	"fmt"
)

func main() {
	var nums1 = []int{0,1,2,3,4,5,6,7,8,9,10}
	var nums2 = []int{0,1,2,3,4,5,6,7,8,9,10}


	var i = 0
	var isEqual = true
	if(len(nums1) != len(nums2)) {
		isEqual = false
	} else {
		for i < len(nums1) {
			if nums1[i] != nums2[i] {
				isEqual = false
				break
			}
		}
	}
	
	
	
	fmt.Println("equal", isEqual)



}