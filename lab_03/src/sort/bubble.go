package sort

// Bubble used to sort array using bubble method.
func Bubble(arr []int) {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
