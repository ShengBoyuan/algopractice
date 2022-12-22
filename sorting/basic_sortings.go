package sorting

// 10 basic sortings

// 1. SelectSort
// Time Complexity: O(n^2)
// Space Complexity: O(1)
// Stability: True
// In Place: True
func SelectSort(nums []int) []int {
	n := len(nums)
	tmp := 0

	for i := 0; i < n; i++ {
		min := i
		for j := i; j < n; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}

		tmp = nums[i]
		nums[i] = nums[min]
		nums[min] = tmp
	}

	return nums
}

// 2. InsertSort
// Time Complexity: O(n^2)
// Space Complexity: O(1)
// Stability: True
// In Place: True
func InsertSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}

	for i := 1; i < n; i++ {
		tmp := nums[i]
		k := i - 1
		for k >= 0 && nums[k] > tmp {
			k--
		}
		for j := i; j > k+1; j-- {
			nums[j] = nums[j-1]
		}
		nums[k+1] = tmp
	}

	return nums
}

// 3. BubbleSort
// Time Complexity: O(n^2)
// Space Complexity: O(1)
// Stability: True
// In Place: True
func BubbleSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}

	for i := 0; i < n; i++ {
		flag := false
		tmp := 0

		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				tmp = nums[j+1]
				nums[j+1] = nums[j]
				nums[j] = tmp
				flag = true
			}
		}

		if !flag {
			break
		}
	}

	return nums
}

// 4. ShellSort
// Time Complexity: O(nlogn)
// Space Complexity: O(1)
// Stability: False
// In Place: True
// func ShellSort(nums []int, div int) []int {
// 	n := len(nums)
// 	if n < 2 {
// 		return nums
// 	}

// 	for gap := n / div; gap > 0; gap /= div {
// 		for i := gap; i < n; i += gap {
// 			tmp := 0
// 			if nums[i]
// 		}
// 	}
// 	return nums
// }

// 10. RadixSort
// Time Complexity: O(k*n)
// Space Complexity: O(r*n)
// Stability: True
// In Place: False
// func RadixSort(nums []int, r int) []int {
// 	// 0. Preprocessing
// 	n := len(nums)
// 	if n < 2 {
// 		return nums
// 	}

// 	// 1. Get max info and init
// 	max := nums[0]
// 	k := 0 // length of the max num

// 	for i := 1; i < n; i++ {
// 		if max < nums[i] {
// 			max = nums[i]
// 		}
// 	}
// 	for max != 0 {
// 		k++
// 		max /= 10
// 	}
// 	buckets := make([][]int, 0, r) // r is radix of the nums
// 	res := make([]int, 0, n) // temporarily store results
// 	for i := 0; i < r; i++ {
// 		buckets[i] = make([]int, 0, n)
// 	}

// 	// 2. Process loop
// 	for i := 0; i < r; i++ {
// 		for j := 0; j < n; j++ {

// 		}
// 	}

// 	return nil
// }
