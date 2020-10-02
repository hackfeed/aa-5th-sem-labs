package matrix

import "fmt"

// ReadMatrix used to read matrix with n rows and m columns.
func ReadMatrix(n, m int) MInt {
	mat := formResMat(n, m)

	for i := 0; i < mat.n; i++ {
		for j := 0; j < mat.m; j++ {
			fmt.Scanf("%d", &mat.mat[i][j])
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

func formResMat(n, m int) MInt {
	var rmat MInt

	rmat.n = n
	rmat.m = m
	rmat.mat = make([][]int, rmat.n)

	for i := range rmat.mat {
		rmat.mat[i] = make([]int, rmat.m)
	}

	return rmat
}
