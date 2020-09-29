package sort

// Selection used to sort array using selection method.
func Selection(arr []int) {
	l := len(arr)
	for i := 0; i < l; i++ {
		minIdx := i
		for j := i; j < l; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}
