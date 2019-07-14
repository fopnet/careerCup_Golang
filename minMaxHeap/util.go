package minMaxHeap

func copyOf(a []Comparable, capacity int) []Comparable {
	s := make([]Comparable, len(a), capacity)
	copy(s, a)
	return s
}

func swap(heap []Comparable, index, largest int) {
	heap[index], heap[largest] = heap[largest], heap[index]
}
