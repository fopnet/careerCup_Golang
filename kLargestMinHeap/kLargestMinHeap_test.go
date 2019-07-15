package kLargestMinHeap

import (
	"fmt"
	"math"
	"testing"

	"github.com/fopnet/carrerCup_Golang/minMaxHeap"
)

const msg = "O valor esperado deveria ser %v, mas foi retornado %v"

func TestVisitTop10FromHeap(t *testing.T) {

	// a := GenerateOddAscedingComparables(30)
	a := minMaxHeap.GenerateDescedingComparables(31)
	// a := []Comparable{inteiro{11}, inteiro{3}, inteiro{2}, inteiro{1}, inteiro{15}, inteiro{5},
	// 	inteiro{4}, inteiro{45}, inteiro{88}, inteiro{96}, inteiro{50}, inteiro{45}}
	heap := minMaxHeap.NewMinHeap(a)
	// fmt.Println("heap", heap)

	numbers, err := heap.VisitLargestFromHeap(10)
	fmt.Println("Top10 numbers", numbers)
	// fmt.Println("Original numbers", heap)

	// bfs, err := heap.BFS(4)
	// fmt.Println("bfs odds", bfs[:10])

	if err != nil {
		result := minMaxHeap.GenerateDescedingComparables(10)
		t.Errorf(msg, result, err)
	}

}

func TestTop10(t *testing.T) {

	numbers := minMaxHeap.GenerateAscedingComparables GenerateAscedingComparables(50)
	// fmt.Println("rnd numbers", numbers)

	top10 := make([]int, 10)
	for i := 0; i < len(top10); i++ {
		max := math.MinInt64
		iMax := 0

		for j := 0; j < len(numbers); j++ {
			if max < numbers[j] {
				max = numbers[j]
				iMax = j
			}
		}
		top10[i] = max
		numbers[iMax] = math.MinInt64
	}

	fmt.Println("top10", top10)
}
