package levenshtein

import "math"

// Recursive used to find Levenshtein distance with recursive method.
func Recursive(fWord, sWord string) int {
	fWordRune, sWordRune := []rune(fWord), []rune(sWord)

	n, m := len(fWordRune), len(sWordRune)

	return getDistance(fWordRune, sWordRune, n, m)
}

// RecursiveMatrix used to find Levenshtein distance with recursive method and matrix filling.
func RecursiveMatrix(fWord, sWord string) (int, MInt) {
	var (
		n, m, shDist int
	)

	fWordRune, sWordRune := []rune(fWord), []rune(sWord)

	n, m = len(fWordRune), len(sWordRune)

	distMat := makeMatrixRec(n, m)

	getDistanceRec(fWordRune, sWordRune, n, m, distMat)

	shDist = distMat[n][m]

	return shDist, distMat
}

// IterativeMatrix used to find Levenshtein distance with matrix filling.
func IterativeMatrix(fWord, sWord string) (int, MInt) {
	var (
		n, m, dist, shDist int
	)

	fWordRune, sWordRune := []rune(fWord), []rune(sWord)

	n, m = len(fWordRune), len(sWordRune)

	distMat := makeMatrix(n, m)

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			insDist := distMat[i][j-1] + 1
			delDist := distMat[i-1][j] + 1
			eq := 1
			if fWordRune[i-1] == sWordRune[j-1] {
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
	var (
		n, m, dist, shDist, transDist int
	)

	fWordRune, sWordRune := []rune(fWord), []rune(sWord)

	n, m = len(fWordRune), len(sWordRune)

	distMat := makeMatrix(n, m)

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			insDist := distMat[i][j-1] + 1
			delDist := distMat[i-1][j] + 1
			eq := 1
			if fWordRune[i-1] == sWordRune[j-1] {
				eq = 0
			}
			eqDist := distMat[i-1][j-1] + eq
			transDist = -1
			if i > 1 && j > 1 {
				transDist = distMat[i-2][j-2] + 1
			}

			if transDist != -1 && fWordRune[i-1] == sWordRune[j-2] && fWordRune[i-2] == sWordRune[j-1] {
				dist = minFromFour(insDist, delDist, eqDist, transDist)
			} else {
				dist = minFromThree(insDist, delDist, eqDist)
			}
			distMat[i][j] = dist
		}
	}

	shDist = distMat[n][m]

	return shDist, distMat
}

func getDistanceRec(fWord, sWord []rune, i, j int, mat MInt) int {
	if mat[i][j] != math.MaxInt16 {
		return mat[i][j]
	}

	if i == 0 {
		mat[i][j] = j
		return mat[i][j]
	}
	if j == 0 && i > 0 {
		mat[i][j] = i
		return mat[i][j]
	}
	eq := 1
	if fWord[i-1] == sWord[j-1] {
		eq = 0
	}

	mat[i][j] = minFromThree(
		getDistanceRec(fWord, sWord, i, j-1, mat)+1,
		getDistanceRec(fWord, sWord, i-1, j, mat)+1,
		getDistanceRec(fWord, sWord, i-1, j-1, mat)+eq)
	return mat[i][j]
}

func makeMatrixRec(n, m int) MInt {
	mat := make(MInt, n+1)
	for i := range mat {
		mat[i] = make([]int, m+1)
	}

	for i := 0; i < n+1; i++ {
		for j := 0; j < m+1; j++ {
			mat[i][j] = math.MaxInt16
		}
	}

	return mat
}

func getDistance(fWord, sWord []rune, i, j int) int {
	if i == 0 {
		return j
	}
	if j == 0 && i > 0 {
		return i
	}
	eq := 1
	if fWord[i-1] == sWord[j-1] {
		eq = 0
	}

	return minFromThree(
		getDistance(fWord, sWord, i, j-1)+1,
		getDistance(fWord, sWord, i-1, j)+1,
		getDistance(fWord, sWord, i-1, j-1)+eq)
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

func minFromFour(a, b, c, d int) int {
	min := a

	if b < min {
		min = b
	}

	if c < min {
		min = c
	}

	if d < min {
		min = d
	}

	return min
}

func minFromThree(a, b, c int) int {
	min := a

	if b < min {
		min = b
	}

	if c < min {
		min = c
	}

	return min
}
