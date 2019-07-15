package numberUtil

import (
	"math/rand"

	"time"
)

func RandomRange(min int, max int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max+1-min) + min
}

func Random(n int) int {
	// fmt.Println(n)
	// fmt.Println(n)

	rand.Seed(time.Now().UnixNano())

	return rand.Intn(n)
}

func GenerateIntRandom(n int) []int {
	var a = []int{}
	for i := 1; i <= n; i++ {
		rnd := RandomRange(0, n)
		a = append(a, rnd)
	}
	return a
}

func GenerateAscending(n int) []int {
	var a = []int{}
	for i := 1; i <= n; i++ {
		a = append(a, i)
	}
	return a
}

func GenerateDistinctIntRandom(n int) []int {
	var a = []int{}
	sorted := map[int]bool{}
	for i := 1; i <= n; i++ {
		rnd := RandomRange(0, n)
		if _, ok := sorted[rnd]; !ok {
			a = append(a, rnd)
			sorted[rnd] = true
		}
	}
	return a
}
