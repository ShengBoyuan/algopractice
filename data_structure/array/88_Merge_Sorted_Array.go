package array

// Very basic. But it's very clever to merge from the tail.
func Merge(nums1 []int, m int, nums2 []int, n int) {
	// This way to copy nums1 will cause imperceptible bugs,
	// because nums3 is a slice based on nums1,
	// when change nums1, nums3 will also change implicitly.
	// nums3 := nums[:m]
	i, m, n := len(nums1)-1, m-1, n-1

	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			m--
		} else {
			nums1[i] = nums2[n]
			n--
		}
		i--
	}

	for m >= 0 {
		nums1[i] = nums1[m]
		m--
		i--
	}

	for n >= 0 {
		nums1[i] = nums2[n]
		n--
		i--
	}
}
