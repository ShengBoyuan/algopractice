package sorting_test

import (
	"algopractice/sorting"
	"fmt"
	"testing"
)

var Datas map[int][]int

func init() {
	Datas = map[int][]int{
		0: {4, 2, 1, 3, 45, 5, 1, 90, 11},
		1: {3, 4, 1, 5, 6, 3, 1, 4, 3, 100, -2, -1},
		2: {},
		3: {2, 1, 4, 1, 2, 5, 123, 53, 12, 643, -109, -0, -1},
	}
}

func TestSelectSort(t *testing.T) {
	resData := sorting.SelectSort(Datas[0])
	fmt.Println(resData)
}

func TestInsertSort(t *testing.T) {
	resData := sorting.InsertSort(Datas[0])
	fmt.Println(resData)
}

func TestBubbleSort(t *testing.T) {
	resData := sorting.BubbleSort(Datas[0])
	fmt.Println(resData)
}

func TestRadixSort(t *testing.T) {
	resData := sorting.RadixSort(Datas[0], 10)
	fmt.Println(resData)
}
