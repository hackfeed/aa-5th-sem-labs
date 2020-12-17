package ant

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/logrusorgru/aurora"
)

// Compare used to show time difference between Ant algorithm and Brute algorithm.
func Compare(fn string) {
	ant := make([]time.Duration, 0)
	brute := make([]time.Duration, 0)
	for i := 2; i < 11; i++ {
		genData(fn, i)
		w := getWeights(fn)
		col := CreateCol(w)

		start := time.Now()
		col.SearchAnt(100)
		end := time.Now()
		ant = append(ant, end.Sub(start))

		start = time.Now()
		SearchBrute(w)
		end = time.Now()
		brute = append(brute, end.Sub(start))
	}

	logTime("ANT ALGORITHM", ant)
	logTime("BRUTE ALGORITHM", brute)
}

func genData(fn string, n int) {
	os.Remove(fn)
	f, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				str := fmt.Sprintf("%d ", rand.Intn(10)+1)
				f.WriteString(str)
			} else {
				str := fmt.Sprintf("%d ", 0)
				f.WriteString(str)
			}
		}

		str := fmt.Sprintf("\n")
		f.WriteString(str)
	}

	f.Close()
}

func getWeights(fn string) [][]int {
	w := make([][]int, 0)
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		str, err := rd.ReadString('\n')
		if err == io.EOF {
			break
		}
		str = strings.TrimSuffix(str, "\n")
		str = strings.TrimSuffix(str, "\r")
		str = strings.TrimRight(str, " ")
		cur := strings.Split(str, " ")

		path := make([]int, 0)
		for _, i := range cur {
			i, err := strconv.Atoi(i)
			if err != nil {
				fmt.Println(err)
			}
			path = append(path, i)
		}
		w = append(w, path)
	}

	return w
}

func logTime(h string, a []time.Duration) {
	fmt.Printf("%20v\n", aurora.BgRed(h))
	fmt.Printf("%v%18v%v\n", "+", strings.Repeat("-", 18), "+")
	fmt.Printf("|%3v|%14v|\n", aurora.Green("N"), aurora.Green("Time"))
	fmt.Printf("%v%18v%v\n", "+", strings.Repeat("-", 18), "+")
	for i := 0; i < len(a); i++ {
		fmt.Printf("|%3v|%14v|\n", i+2, a[i])
	}
	fmt.Printf("%v%18v%v\n", "+", strings.Repeat("-", 18), "+")
}
