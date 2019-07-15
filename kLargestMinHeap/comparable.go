package minMaxHeap

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

func makeArrayOfComparable(size int) []Comparable {
	return make([]Comparable, size)
}
