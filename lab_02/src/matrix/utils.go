package matrix

import "fmt"

// ReadMatrix used to read matrix with n rows and m columns.
func ReadMatrix(n, m int) MInt {
	mat := make(MInt, n)
	for i := range mat {
		mat[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &mat[i][j])
		}
	}

	return mat
}

// ReadNum used to read a word through EOL symbol.
func ReadNum() int {
	var num int
	fmt.Scanln(&num)

	return num
}
