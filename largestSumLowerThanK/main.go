/* https://www.careercup.com/question?id=5092397411729408

Given a positive integer array and a positive integer.
Find a pair of integers in the array so that the sum of the pair is the largest one among all pairs
but not larger than the given integer.

e.g.
Input
{90, 85, 75, 60, 120, 150, 125}
d: 250

Output:
{90, 125}
*/
package main

import (
	"fmt"
	"strconv"

	"github.com/fopnet/carrerCup_Golang/minMaxHeap"
)

type Pair struct {
	p1, p2 int
}

func (this Pair) CompareTo(b interface{}) int {
	ib := b.(Pair)
	switch {
	case this.Sum() > ib.Sum():
		return 1
	case this.Sum() < ib.Sum():
		return -1
	default:
		return 0
	}
}

func (this Pair) Sum() int {
	return int(this.p1 + this.p2)
}

func (this Pair) ToString() string {
	return "{" + strconv.Itoa(this.p1) + " , " + strconv.Itoa(this.p2) + "}"
}

func main() {

	arr := []int{90, 85, 75, 60, 120, 150, 125}
	integer := 250
	heap := generateHeapOfPairs(integer, arr)
	fmt.Println("All pairs", heap.ToArray())
	if !heap.IsEmpty() {
		fmt.Println("Largest Pair", heap.RootElement().ToString())
	}

	maxPair := getLargestPair(integer, arr)
	fmt.Println("getLargestPair", maxPair)
}

func generateHeapOfPairs(limit int, arr []int) minMaxHeap.Heap {
	length := len(arr)
	pairs := make([]minMaxHeap.Comparable, 0, length*length)

	for i := 0; i < length; i++ {

		for j := 0; j < length; j++ {
			p := Pair{arr[i], arr[j]}
			if j > i && p.Sum() < limit {
				pairs = append(pairs, p)
			}
		}
	}

	return minMaxHeap.NewMaxHeap(pairs)
}

func getLargestPair(limit int, arr []int) Pair {
	length := len(arr)
	max := Pair{0, 0}
	for i := 0; i < length; i++ {

		for j := 0; j < length; j++ {
			p := Pair{arr[i], arr[j]}
			if j > i && p.Sum() < limit && p.Sum() > max.Sum() {
				max = p
			}
		}
	}

	return max
}
