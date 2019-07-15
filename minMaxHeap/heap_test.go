package minMaxHeap

import (
	"github.com/fopnet/carrerCup_Golang/numberUtil"
	"fmt"
	"math"
	"reflect"
	"testing"
)

const msg = "O valor esperado deveria ser %v, mas foi retornado %v"

func TestCopyOf(t *testing.T) {
	a := []Comparable{inteiro{num: 1}, inteiro{num: 3}, inteiro{num: 4}}
	b := CopyOf(a, 5)

	// fmt.Println("b is copy of a", b)

	if cap(b) != 5 {
		t.Fatalf(msg, 5, cap(b))
	}

	if !reflect.DeepEqual(a, b) {
		t.Fatalf(msg, []Comparable{inteiro{num: 1}, inteiro{num: 3}, inteiro{num: 4}}, b)
	}

	cmp := b[0].CompareTo(a[0])
	if cmp != 0 {
		t.Fatalf(msg, 0, cmp)
	}
}
func TestMinHeap(t *testing.T) {
	i1 := inteiro{num: 1}
	a := []Comparable{i1, inteiro{4}, inteiro{3}}

	heap := NewMinHeap(a)

	h := heap.toArray()
	fmt.Println("built heap", h)

	if heap.RootElement() != i1 {
		t.Errorf(msg, i1, h[0])
	}
}

func TestExtractRoot(t *testing.T) {
	i1 := inteiro{num: 1}
	i3 := inteiro{num: 3}
	i4 := inteiro{num: 4}
	a := []Comparable{i1, i4, i3}

	heap := NewMinHeap(a)

	heap.ExtractRootElement()

	if heap.RootElement() != i3 {
		t.Errorf(msg, i3, heap.RootElement())
	}
}

func TestBFS(t *testing.T) {
	i1 := inteiro{num: 1}
	i3 := inteiro{num: 3}
	i4 := inteiro{num: 4}
	a := []Comparable{i1, i4, i3}

	heap := NewMinHeap(a)

	bfs, err := heap.BFS(2)

	fmt.Println("bfs", bfs)
	if err != nil {
		t.Errorf(msg, bfs, err)
	}
}

func TestIsMaxHeap(t *testing.T) {
	heap := NewMaxHeap([]Comparable{})

	if heap.IsMaxHeap() == false {
		t.Errorf(msg, true, false)
	}
}

func TestIsMinHeap(t *testing.T) {
	heap := NewMinHeap([]Comparable{})

	if heap.IsMinHeap() == false {
		t.Errorf(msg, true, false)
	}
}

/**
https://www.careercup.com/question?id=5154165987213312
Implement a function which accepts a number and returns top 10 big numbers the function is called with so far;

If we call the function with 1.. to 100 , for the call function(100)
the function will return 91 to 100 in reverse order since they are top 10 biggest numbers so far
*/
func TestBFS2(t *testing.T) {
	n := 100
	a := numberUtil.

	// fmt.Println("a", a)
	// heap := NewMaxHeap([]Comparable{})
	heap := NewMaxHeap(a)

	top := float64(10) // 91 <-> 100
	ilevel := int(math.Ceil(math.Log2(top - 1)))
	// fmt.Println("level", ilevel)

	// fmt.Println("a", a)
	// heap.Heapsort(a)
	// fmt.Println("heapsort", heap)

	bfs, err := heap.BFS(ilevel)
	fmt.Println("bfs", bfs[:10])
	if err != nil {
		t.Errorf(msg, bfs, err)
	}
}
