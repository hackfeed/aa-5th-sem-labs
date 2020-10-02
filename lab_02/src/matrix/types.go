package matrix

import "fmt"

// MInt used to represent int matrix
type MInt struct {
	mat  [][]int
	n, m int
}

// PrintMatrix used to pretty print matrix.
func (mat MInt) PrintMatrix() {
	for i := 0; i < mat.n; i++ {
		for j := 0; j < mat.m; j++ {
			fmt.Printf("%3d ", mat.mat[i][j])
		}
		fmt.Printf("\n")
	}
}
