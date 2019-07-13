package minMaxHeap

func copyOf(a []Comparable, capacity int) []Comparable {
	s := make([]Comparable, len(a), capacity)
	copy(s, a)
	return s
}

func swap(heap []Comparable, index, largest int) {
	heap[index], heap[largest] = heap[largest], heap[index]
}

func isAccesibleAndNotFull(numbers []Comparable, idx int) bool {
	return len(numbers) < cap(numbers) && idx < cap(numbers)
}
