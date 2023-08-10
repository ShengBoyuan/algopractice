package sorting

import "math"

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

// 5. MergeSort
// Time Complexity: O(nlogn)
// Space Complexity: O(1)
// Stability: False
// In Place: True
func MergeSort(nums1 []int, nums2 []int, m, n int) []int {
	i, m, n := len(nums1)-1, m-1, n-1

	for m >= 0 && n >= 0 {
		if nums1[m] >= nums2[n] {
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

	return nums1
}

// 6. QuickSort
// Time Complexity: O(nlogn)
// Space Complexity: O(logn)
// Stability: False
// In Place: false
func QuickSort(nums []int, low, high int) {
	if low < high {
		pivot := Partition(nums, low, high)
		QuickSort(nums, low, pivot-1)
		QuickSort(nums, pivot+1, high)
	}
}

func Partition(nums []int, low, high int) int {
	pivotValue := nums[low]

	for low < high {
		for low < high && nums[high] >= pivotValue {
			high--
		}
		nums[low] = nums[high]
		for low < high && nums[low] <= pivotValue {
			low++
		}
		nums[high] = nums[low]
	}

	nums[low] = pivotValue
	return low
}

// 10. RadixSort
// Time Complexity: O(k*n)
// Space Complexity: O(r*n)
// Stability: True
// In Place: False
func RadixSort(nums []int, r int) []int {
	// 0. Preprocessing
	n := len(nums)
	if n < 2 {
		return nums
	}

	// 1. Get max info and init
	max := nums[0]
	k := 0 // length of the max num
	for i := 1; i < n; i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	for max != 0 {
		k++
		max /= 10
	}

	buckets := make([][]int, 10) // r is radix of the nums
	for i := 0; i < 10; i++ {
		buckets[i] = make([]int, 0, n)
	}

	// 2. Process loop
	for i := 0; i < k; i++ {
		for j := 0; j < n; j++ {
			radio := nums[j] / int(math.Pow(10, float64(i))) % 10
			buckets[radio] = append(buckets[radio], nums[j])
		}
		nums = nums[:0] // clear slice
		for t := 0; t < 10; t++ {
			nums = append(nums, buckets[t]...)
			buckets[t] = buckets[t][:0] // clear bucket after remove the nums
		}
	}

	return nums
}
