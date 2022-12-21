package sorting_test

import (
	"algopractice/sorting"
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	// srcData1 := []int{3, 4, 1, 5, 6, 3, 1, 4, 3, 100, -2, -1}
	srcData2 := []int{}
	resData := sorting.SelectSort(srcData2)
	fmt.Println(resData)
}

func TestInsertSort(t *testing.T) {
	srcData1 := []int{3, 4, 1, 5, 6, 3, 1, 4, 3, 100, -2, -1}
	// srcData2 := []int{}
	resData := sorting.InsertSort(srcData1)
	fmt.Println(resData)
}

func TestBubbleSort(t *testing.T) {
	srcData1 := []int{3, 4, 1, 5, 6, 3, 1, 4, 3, 100, -2, -1}
	// srcData2 := []int{}
	resData := sorting.BubbleSort(srcData1)
	fmt.Println(resData)
}
