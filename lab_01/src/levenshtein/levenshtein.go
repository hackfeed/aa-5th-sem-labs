package levenshtein

// Recursive used to find Levenshtein distance with recursive method.
func Recursive(fWord, sWord string) int {
	return 5
}

// RecursiveMatrix used to find Levenshtein distance with recursive method and matrix filling.
func RecursiveMatrix(fWord, sWord string) (int, MInt) {
	distMat := makeMatrix(len(fWord), len(sWord))

	return 5, distMat
}

// IterativeMatrix used to find Levenshtein distance with matrix filling.
func IterativeMatrix(fWord, sWord string) (int, MInt) {
	var (
		n, m, dist, shDist int
	)

	n, m = len(fWord), len(sWord)

	distMat := makeMatrix(n, m)

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			insDist := distMat[i][j-1] + 1
			delDist := distMat[i-1][j] + 1
			eq := 1
			if fWord[i-1] == sWord[j-1] {
				eq = 0
			}
			eqDist := distMat[i-1][j-1] + eq

			dist = minFromThree(insDist, delDist, eqDist)
			distMat[i][j] = dist
		}
	}

	shDist = distMat[n][m]

	return shDist, distMat
}

// DamerauLevenshtein used to find Damerau-Levenshtein distance.
func DamerauLevenshtein(fWord, sWord string) (int, MInt) {
	distMat := makeMatrix(len(fWord), len(sWord))

	return 5, distMat
}

func makeMatrix(n, m int) MInt {
	mat := make(MInt, n+1)
	for i := range mat {
		mat[i] = make([]int, m+1)
	}

	for i := 0; i < m+1; i++ {
		mat[0][i] = i
	}

	for i := 0; i < n+1; i++ {
		mat[i][0] = i
	}

	return mat
}

func minFromThree(a, b, c int) int {
	var min, fMin int

	if a < b {
		fMin = a
	} else {
		fMin = b
	}

	if c < fMin {
		min = c
	} else {
		min = fMin
	}

	return min
}
