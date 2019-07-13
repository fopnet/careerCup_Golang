package main

import (
	"fmt"
	"sort"

	"github.com/fopnet/carrerCup_Golang/numberUtil"
)

/**
https://www.careercup.com/question?id=5710755664494592

Write a function that takes a list L and returns a random sublist of size N of that list.
Assume that the indexes must be in increasing order. That is, you cannot go backwards.

Example:

L = [1, 2, 3, 4, 5]
N = 3

The function should return one of these lists:

[1, 2, 3]
[1, 2, 4]
[1, 2, 5]
[1, 3, 4]
[1, 3, 5]
[1, 4, 5]
[2, 3, 4]
[2, 3, 5]
[2, 4, 5]
[3, 4, 5]
*/
func main() {
	a := []int{1, 2, 3, 4, 5}
	n := 3

	b := solution4(a, n)

	fmt.Println("result b", b)

	// fmt.Println("random", randomRange(1, 3)) // [1-3]
	// fmt.Println("random", random(3)+0) // [0-2]
	// fmt.Println("random", random(4)+1) // [1-3]

}

func solution4(L []int, n int) []int {
	sublist := [][]int{}

	for i := 0; i < len(L); i++ {
		for j := i + 1; j < len(L); j++ {
			for k := j + 1; k < len(L); k++ {
				if L[i] < L[j] && L[j] < L[k] {
					// fmt.Println(L[i], ",", L[j], ",", L[k])
					sublist = append(sublist, []int{L[i], L[j], L[k]})
				}
			}
		}
	}

	rndIdx := numberUtil.RandomRange(0, len(sublist)-1)

	fmt.Println("subList", rndIdx, sublist)

	return sublist[rndIdx]
}

func solution3_Faster(a []int, n int) []int {
	b := []int{}
	rnd := -1

	for i := 0; i < len(a) && len(b) < n; i++ {
		max := i + n - 1
		min := rnd + 1

		rnd = numberUtil.RandomRange(min, max)
		fmt.Println("rnd", min, "<=", rnd, "<=", max)

		b = append(b, a[rnd])
		if len(a)-1-rnd <= n-len(b) && len(b) < n {
			fmt.Println("break ", b, a[rnd:])
			b = append(b, a[rnd+1:]...)
			break
		}
	}
	return b
}
func solution2(a []int, n int) []int {
	b := []int{}
	for i := 0; i < len(a) && len(b) < n; {
		rnd := numberUtil.Random(i+n-i) + i
		fmt.Println("rnd", i+n-i, i, rnd)

		b = append(b, a[rnd])

		i = rnd + 1
	}
	return b
}
func solution1(a []int, n int) []int {
	b := []int{}

	for i := 0; i < n; i++ {
		rnd := numberUtil.Random(len(a))
		b = append(b, a[rnd])
		a = append(a[:rnd], a[rnd+1:]...)
		fmt.Println("a,b", a, b)
	}

	sort.Ints(b)

	return b
}
