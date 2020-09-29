package sort

// Insertion used to sort array using insertion method.
func Insertion(arr []int) {
	l := len(arr)
	for i := 0; i < l; i++ {
		x := arr[i]
		j := i
		for j > 0 && arr[j-1] > x {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = x
	}
}
