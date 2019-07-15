package minMaxHeap

import "github.com/fopnet/carrerCup_Golang/numberUtil"

type Comparable interface {
	// LessThan(j interface{}) bool
	// EqualTo(j interface{}) bool
	CompareTo(o interface{}) int
}

type Comparator func(o1, o2 Comparable) int

// {
// Compare(i, j Comparable) int
// EqualTo(j interface{}) bool
// }

func CopyOf(a []Comparable, capacity int) []Comparable {
	s := make([]Comparable, len(a), capacity)
	copy(s, a)
	return s
}

func Swap(heap []Comparable, index, largest int) {
	heap[index], heap[largest] = heap[largest], heap[index]
}

func MakeArrayOfComparable(size int) []Comparable {
	return make([]Comparable, size)
}

type inteiro struct {
	num int
}

func (this inteiro) CompareTo(b interface{}) int {
	ib := b.(inteiro)
	switch {
	case this.num > ib.num:
		return 1
	case this.num < ib.num:
		return -1
	default:
		return 0
	}
}

func GenerateAscedingComparables(n int) []Comparable {
	a := []Comparable{}
	for i := 1; i <= n; i++ {
		a = append(a, inteiro{i})
	}
	return a
}
func GenerateOddAscedingComparables(n int) []Comparable {
	a := []Comparable{}
	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			a = append(a, inteiro{i})
		}
	}
	return a
}
func GenerateDescedingComparables(n int) []Comparable {
	var a = []Comparable{}
	for i := n; i > 0; i-- {
		// for i := 1; i <= n; i++ {
		// rnd := numberUtil.RandomRange(0, n)
		// a = append(a, inteiro{rnd})
		a = append(a, inteiro{i})
		// fmt.Print("rnd ", rnd)
	}
	return a

}
func GenerateRandomComparables(n int) []Comparable {
	var a = []Comparable{}
	for i := 1; i <= n; i++ {
		rnd := numberUtil.RandomRange(0, n)
		a = append(a, inteiro{rnd})
	}
	return a
}
