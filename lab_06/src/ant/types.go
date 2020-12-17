package ant

// Ant used to represent sigle ant in colony.
type Ant struct {
	col *Colony
	vis [][]int
	isv [][]bool
	pos int
}

// Colony used to represent ants' environment.
type Colony struct {
	w  [][]int
	ph [][]float64
	a  float64
	b  float64
	q  float64
	p  float64
}
