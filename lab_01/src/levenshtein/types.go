package levenshtein

import "fmt"

// MInt used to represent int64 matrix
type MInt [][]int

// PrintMatrix used to pretty print matrix.
func (mat MInt) PrintMatrix() {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			fmt.Printf("%3d ", mat[i][j])
		}
		fmt.Printf("\n")
	}
}
