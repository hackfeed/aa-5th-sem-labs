package main

func main() {
	n := 4							// (1)
	arr := make([]int, 4)			// (2)
	l := 0							// (3)
	r := n - 1						// (4)
	for i := 0; i < n; i++ {		// (5)
		arr[i] = n - i				// (6)
	}
	for l <= r {					// (7)
		for i := l; i < r; i++ {	// (8)
			if arr[i] > arr[i+1] {	// (9)
				tmp := arr[i]		// (10)
				arr[i] = arr[i+1]	// (11)
				arr[i+1] = tmp		// (12)
			}
		}
		r--							// (13)
		for j := r; j > l; j-- {	// (14)
			if arr[j-1] > arr[j] {	// (15)
				tmp := arr[j]		// (16)
				arr[j] = arr[j-1]	// (17)
				arr[j-1] = tmp		// (18)
			}
		}
		l++							// (19)
	}
}
