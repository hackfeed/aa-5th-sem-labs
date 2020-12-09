package dict

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"

	"github.com/brianvoe/gofakeit"
)

// CreateArray used to create DictArray with given size.
func CreateArray(n int) DictArray {
	var (
		darr DictArray
		g    Dict
	)

	darr = make(DictArray, n)

	for i := 0; i < n; i++ {
		dup := true
		for dup != false {
			g = Dict{
				"id":       gofakeit.Uint8(),
				"gamertag": gofakeit.Gamertag(),
			}
			dup = g.IsDup(darr[:i])
		}

		darr[i] = g
	}

	return darr
}

// IsDup used to check whether Dict presents in given DictArray.
func (d Dict) IsDup(darr DictArray) bool {
	for _, v := range darr {
		if reflect.DeepEqual(d, v) {
			return true
		}
	}
	return false
}

// Print used to print single Dict.
func (d Dict) Print() {
	fmt.Printf("ID: %v\nGamertag: %v\n", d["id"], d["gamertag"])
}

// Print used to print single DictArray.
func (darr DictArray) Print() {
	for _, d := range darr {
		d.Print()
	}
}

// Pick used to get gamertag with first letter.
func (darr DictArray) Pick(l string) string {
	for _, d := range darr {
		if d["gamertag"].(string)[:1] == l {
			return d["gamertag"].(string)
		}
	}

	i := rand.Int() % len(darr)

	return darr[i]["gamertag"].(string)
}

// Brute used to find value using bruteforce method.
func (darr DictArray) Brute(gt string) Dict {
	var r Dict

	for _, d := range darr {
		if d["gamertag"] == gt {
			return d
		}
	}

	return r
}

// Binary used to find value using binary search method.
func (darr DictArray) Binary(gt string) Dict {
	var (
		l   int = len(darr)
		mid int = l / 2
		r   Dict
	)

	switch {
	case l == 0:
		return r
	case darr[mid]["gamertag"].(string) > gt:
		r = darr[:mid].Binary(gt)
	case darr[mid]["gamertag"].(string) < gt:
		r = darr[mid+1:].Binary(gt)
	default:
		r = darr[mid]
	}

	return r
}

// FAnalysis used to analyse frequency of given DictArray.
func (darr DictArray) FAnalysis() FreqArray {
	var (
		az   string    = "abcdefghijklmnopqrstuvwxyz"
		farr FreqArray = make(FreqArray, len(az))
	)

	for i, v := range az {
		a := Freq{
			l:    string(v),
			cnt:  0,
			darr: make(DictArray, 0),
		}
		farr[i] = a
	}

	for _, v := range darr {
		l := v["gamertag"].(string)[:1]
		for i := range farr {
			if farr[i].l == l {
				farr[i].cnt++
			}
		}
	}

	sort.Slice(farr, func(i, j int) bool {
		return farr[i].cnt > farr[j].cnt
	})

	for i := range farr {
		for j := range darr {
			if darr[j]["gamertag"].(string)[:1] == farr[i].l {
				farr[i].darr = append(farr[i].darr, darr[j])
			}
		}

		sort.Slice(farr[i].darr, func(l, m int) bool {
			return farr[i].darr[l]["gamertag"].(string) < farr[i].darr[m]["gamertag"].(string)
		})
	}

	return farr
}

// Combined used to find value using binary search and frequency analysis method.
func (farr FreqArray) Combined(w string) Dict {
	var (
		l string = w[:1]
		r Dict
	)

	for _, d := range farr {
		if string(d.l) == l {
			r = d.darr.Binary(w)
		}
	}

	return r
}
