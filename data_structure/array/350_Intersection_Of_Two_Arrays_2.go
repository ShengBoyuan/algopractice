package array

// Time: O(n), Space: O(n)
func MyOwnSolution(nums1 []int, nums2 []int) []int {

	tmp := make(map[int]int)

	for _, t := range nums1 {
		tmp[t]++
		// unnecessary guard around map access.
		// if _, ok := tmp[t]; ok {
		// 	tmp[t]++
		// } else {
		// 	tmp[t] = 1
		// }
	}

	index := 0
	for i, t := range nums2 {
		if v, ok := tmp[t]; ok && v > 0 {
			nums2[index], nums2[i] = nums2[i], nums2[index]
			index++
			tmp[t]--
		}
	}

	return nums2[:index]
}

// Time: O(n^2), Space: O(1)
func Another(nums1 []int, nums2 []int) []int {
	index := 0

	for i := 0; i < len(nums1); i++ {
		for j := index; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				nums1[index] = nums1[i]
				nums2[index], nums2[j] = nums2[j], nums2[index] // put used num to the front.
				index++
				break // the num1[i] has been used, so break.
			}
		}
	}

	return nums1[:index]
}
