package sort

import "fmt"

// ReadArray used to read array of n elements.
func ReadArray(n int) []int {
	arr := make([]int, n)

	for i := range arr {
		fmt.Scanf("%d", &arr[i])
	}

	return arr
}

// ReadNum used to read a word through EOL symbol.
func ReadNum() int {
	var num int
	fmt.Scanln(&num)

	return num
}
