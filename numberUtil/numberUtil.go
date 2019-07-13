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
