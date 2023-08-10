package array

// func MatrixReshape(mat [][]int, r int, c int) [][]int {
// 	m, n := len(mat), len(mat[0])
// 	if m*n != r*c {
// 		return mat
// 	}

// 	res := make([][]int, 0)
// 	for i, t := range mat {
// 		for j, num := range t {
// 			index := n*i + j + 1

// 		}
// 	}

// 	return nil
// }

func getIndex(num int, r int, c int) (int, int) {
	return (num / r) - 1, num % c
}
