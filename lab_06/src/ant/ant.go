package ant

import (
	"math"
	"math/rand"
	"time"
)

// SearchAnt used to perform Ant algorithm for given colony in d days.
func (c *Colony) SearchAnt(d int) []int {
	s := make([]int, len(c.w))

	for n := 0; n < d; n++ {
		for i := 0; i < len(c.w); i++ {
			a := c.CreateAnt(i)
			a.moveAnt()
			cur := a.getDistance()

			if (s[i] == 0) || (cur < s[i]) {
				s[i] = cur
			}
		}
	}

	return s
}

// SearchBrute used to perform brute algorithm.
func SearchBrute(w [][]int) []int {
	var (
		path = make([]int, 0)
		r    = make([]int, len(w))
	)

	for k := 0; k < len(w); k++ {
		var (
			rts  = make([][]int, 0)
			sum  = math.MaxInt64
			curr = 0
		)
		getRoutes(k, w, path, &rts)

		for i := 0; i < len(rts); i++ {
			curr = 0

			for j := 0; j < len(rts[i])-1; j++ {
				curr += w[rts[i][j]][rts[i][j+1]]
			}

			if curr < sum {
				sum = curr
			}
		}

		r[k] = sum
	}

	return r
}

// CreateAnt used to create single Ant for Ant algorithm.
func (c *Colony) CreateAnt(pos int) *Ant {
	a := Ant{
		col: c,
		vis: make([][]int, len(c.w)),
		isv: make([][]bool, len(c.w)),
		pos: pos,
	}

	for i := 0; i < len(c.w); i++ {
		a.vis[i] = make([]int, len(c.w))
		for j := 0; j < len(c.w[i]); j++ {
			a.vis[i][j] = c.w[i][j]
		}
	}

	for i := range a.isv {
		a.isv[i] = make([]bool, len(c.w))
	}

	return &a
}

// CreateCol used to create Ants' colony.
func CreateCol(w [][]int) *Colony {
	var (
		c = Colony{
			w:  w,
			ph: make([][]float64, len(w), len(w)),
			a:  3.0,
			b:  7.0,
			q:  20.0,
			p:  0.6,
		}
	)

	for i := 0; i < len(c.ph); i++ {
		c.ph[i] = make([]float64, len(c.w[i]))
		for j := range c.ph[i] {
			c.ph[i][j] = 0.5
		}
	}

	return &c
}

func (a *Ant) moveAnt() {
	for {
		prob := a.getProb()
		way := getWay(prob)
		if way == -1 {
			break
		}
		a.follow(way)
		a.updatePh()
	}
}

func (a *Ant) getProb() []float64 {
	var (
		p   = make([]float64, 0)
		sum float64
	)

	for i, l := range a.vis[a.pos] {
		if l != 0 {
			d := math.Pow((float64(1)/float64(l)), a.col.a) * math.Pow(a.col.ph[a.pos][i], a.col.b)
			p = append(p, d)
			sum += d
		} else {
			p = append(p, 0)
		}
	}

	for _, l := range p {
		l /= sum
	}

	return p
}

func (a *Ant) getDistance() int {
	d := 0

	for i, j := range a.isv {
		for k, z := range j {
			if z {
				d += a.col.w[i][k]
			}
		}
	}

	return d
}

func (a *Ant) updatePh() {
	delta := 0.0

	for k := 0; k < len(a.col.ph); k++ {
		for i, j := range a.col.ph[k] {
			if a.col.w[k][i] != 0 {
				if a.isv[k][i] {
					delta = a.col.q / float64(a.col.w[k][i])
				} else {
					delta = 0
				}
				a.col.ph[k][i] = (1 - a.col.p) * (float64(j) + delta)
			}

			if a.col.ph[k][i] <= 0 {
				a.col.ph[k][i] = 0.1
			}
		}
	}
}

func (a *Ant) follow(path int) {
	for i := range a.vis {
		a.vis[i][a.pos] = 0
	}
	a.isv[a.pos][path] = true
	a.pos = path
}

func getWay(p []float64) int {
	var (
		r       *rand.Rand
		sum, rn float64
	)

	for _, j := range p {
		sum += j
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	rn = r.Float64() * sum
	sum = 0

	for i, j := range p {
		if rn > sum && rn < sum+j {
			return i
		}
		sum += j
	}

	return -1
}

func getRoutes(pos int, w [][]int, path []int, rts *[][]int) {
	path = append(path, pos)

	if len(path) < len(w) {
		for i := 0; i < len(w); i++ {
			if !isExist(path, i) {
				getRoutes(i, w, path, rts)
			}
		}
	} else {
		*rts = append(*rts, path)
	}
}

func isExist(a []int, v int) bool {
	for _, val := range a {
		if v == val {
			return true
		}
	}

	return false
}
