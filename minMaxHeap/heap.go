package minMaxHeap

import (
	"fmt"
	"math"
)

type dummyComparable struct {
	num int
}

func (this dummyComparable) ToString() string {
	return string(this.num)
}

func (this dummyComparable) CompareTo(b interface{}) int {
	ib := b.(dummyComparable)
	switch {
	case this.num > ib.num:
		return 1
	case this.num < ib.num:
		return -1
	default:
		return 0
	}
}

// https://pt.wikipedia.org/wiki/Heap
type Heap interface {
	RootElement() Comparable
	ToArray() []Comparable

	Insert(element Comparable)
	BuildHeap(array []Comparable)
	ExtractRootElement() Comparable
	Heapsort(array []Comparable) []Comparable
	EstatisticaDeOrdem(n int) Comparable
	BFS(level int) ([]Comparable, error)
	VisitLargestFromHeap(n int) ([]Comparable, error)
	IsMaxHeap() bool
	IsMinHeap() bool

	Size() int
	IsEmpty() bool
	IsFull() bool
	GetComparator() Comparator
	SetComparator(comparator Comparator)
}

type HeapImpl struct {
	heap  []Comparable
	index int

	_ZERO              int
	_INITIAL_SIZE      int
	_INCREASING_FACTOR int
	/**
	 * O comparador é utilizado para fazer as comparações da heap. O ideal é
	 * mudar apenas o comparator e mandar reordenar a heap usando esse
	 * comparator. Assim os metodos da heap não precisam saber se vai funcionar
	 * como max-heap ou min-heap.
	 */
	comparator Comparator
}

// We are forced to call the constructor to get an instance of
func New(comparator Comparator) Heap {
	const INITIAL_SIZE = 20
	initialHeap := make([]Comparable, INITIAL_SIZE)

	// enforce the default value here
	return &HeapImpl{
		heap:               initialHeap,
		index:              -1,
		_ZERO:              0,
		_INITIAL_SIZE:      INITIAL_SIZE,
		_INCREASING_FACTOR: 10,
		comparator:         comparator,
	}
}

func NewMinHeap(array []Comparable) Heap {
	heap := New(getMinHeapComparator())
	heap.BuildHeap(array)
	return heap
}

func NewMaxHeap(array []Comparable) Heap {
	heap := New(getMaxHeapComparator())
	heap.BuildHeap(array)
	return heap
}

func getMaxHeapComparator() Comparator {
	return func(i, j Comparable) int { return i.CompareTo(j) }
}

func getMinHeapComparator() Comparator {
	return func(i, j Comparable) int { return j.CompareTo(i) }
}

func (heap HeapImpl) parent(i int) int {
	x := (i - 1) / 2
	// x := i / 2
	return x
}

/**
 * Deve retornar o indice que representa o filho a esquerda do elemento
 * indexado pela posição i no vetor
 */
func (heap HeapImpl) left(i int) int {
	// return i * 2
	return i*2 + 1
}

/**
 * Deve retornar o indice que representa o filho a direita do elemento
 * indexado pela posição i no vetor
 */
func (heap HeapImpl) right(i int) int {
	return i*2 + 2
	// return (i*2 + 1)
}

/***************************
*****  private  Heap Methods
****************************/

func (this HeapImpl) getHeap() []Comparable {
	return this.heap
}

/**
 * Valida o invariante de uma heap a partir de determinada posição, que pode
 * ser a raiz da heap ou de uma sub-heap. O heapify deve colocar os maiores
 * (comparados usando o comparator) elementos na parte de cima da heap.
 */
func (this HeapImpl) heapify(position int) {

	if position == this._ZERO { // Remove
		this.heapfyDown(this._ZERO)
	} else { // Insert
		this.heapfyUp(position)
	}
}

/**
 * Faz o processo de Heapfy de cima para baixo, levando o elemento
 * adicionado na raiz para aposição correta
 */
func (this HeapImpl) heapfyDown(index int) {
	rightIndex := this.right(index)
	leftIndex := this.left(index)

	// If Leaf Node, Simply return
	if index >= this.Size()/2 {
		return
	}

	// Select minimum from left node and
	// current node i, and store the minimum
	// index in smallest variable
	var smallest int
	if this.getHeap()[leftIndex].CompareTo(this.getHeap()[index]) < this._ZERO {
		smallest = leftIndex
	} else {
		smallest = index
	}

	// If right child exist, compare and
	// update the smallest variable
	if rightIndex < this.Size() {
		if this.getHeap()[rightIndex].CompareTo(this.getHeap()[smallest]) < this._ZERO {
			smallest = rightIndex
		}
	}

	// If Node i violates the min heap
	// property, Swap  current node i with
	// smallest to fix the min-heap property
	// and recursively call heapify for node smallest.
	if smallest != index {
		Swap(this.getHeap(), index, smallest)
		this.heapfyDown(smallest)
	}
	/*

		var largest int
		if rightIndex < this.Size() && leftIndex <= this.index && this.biggerElement(leftIndex, rightIndex) == leftIndex {
			largest = leftIndex
		} else {
			largest = index
		}

		if rightIndex < this.Size() && rightIndex <= this.index && this.biggerElement(rightIndex, largest) == rightIndex {
			largest = rightIndex
		}

		if largest != index {
			Swap(this.getHeap(), index, largest)
			this.heapfyDown(largest)
		}
	*/

}

/**
 * Faz o processo de Heapfy de baixo para cima, levando o elemento
 * adicionado para aposição correta
 */
func (this HeapImpl) heapfyUp(position int) {
	currentIndex := position

	for this.biggerElement(currentIndex, this.parent(currentIndex)) == currentIndex && this.parent(currentIndex) != currentIndex {
		Swap(this.getHeap(), currentIndex, this.parent(currentIndex))
		currentIndex = this.parent(currentIndex)
	}

}

/**
 * Verifica qual o maior elemento com base e seu indice e retorna o indice
 * do mesmo
 */
func (this HeapImpl) biggerElement(IndexOfElem1, IndexOfElem2 int) int {
	if this.comparator(this.getHeap()[IndexOfElem1], this.getHeap()[IndexOfElem2]) > this._ZERO {
		return IndexOfElem1
	} else {
		return IndexOfElem2
	}
}

/***************************
*****  public Heap Methods
****************************/

func (this *HeapImpl) Insert(element Comparable) {
	if element != nil {

		if this.index == len(this.heap)-1 {
			this.heap = CopyOf(this.heap, len(this.heap)+this._INCREASING_FACTOR)
		}

		this.index++
		this.getHeap()[this.index] = element

		this.heapify(this.index)
	}
}

func (this *HeapImpl) BuildHeap(array []Comparable) {
	if len(array) > 0 {
		// this.heap = make([]Comparable, len(array))
		this.heap = array

		this.index = -1
		for _, v := range array {
			if v != nil {
				this.Insert(v)
			}
		}
		// // Calling Heapify for all non leaf nodes
		// for i := (this.Size() / 2) - 1; i >= 0; i-- {
		// 	this.heapfyDown(i)
		// }
	}
}

func (this *HeapImpl) ExtractRootElement() Comparable {
	var root Comparable

	if !this.IsEmpty() {
		root = this.getHeap()[this._ZERO]
		this.getHeap()[this._ZERO] = this.getHeap()[this.index]
		this.index--
		this.heapify(this._ZERO)
	}

	return root
}

/*
 O heapsort é uma técnica de ordenação que se aproveita da característica da heap
 em que sempre o menor ou maior elemento (dependendo do tipo da heap) deve estar na raiz,
 para ordenar os dados.
 Pode ser aplicado in place em um max heap
 ou através de consecutivas remoções em um min heap.

 Utiliza a ordenação da Heap -> Min -> Crescente |  Max -> decrescente
*/
func (this *HeapImpl) Heapsort(array []Comparable) []Comparable {
	result := array
	if len(array) > 1 {
		this.BuildHeap(array)
		for i := len(array) / 2; i >= 0; i-- {
			// for i := len(array) - 1; i >= 0; i-- {
			Swap(this.heap, this._ZERO, i)
			this.index--
			this.heapify(this._ZERO)
		}
		result = this.heap

		// if len(result) > 1 && this.IsMaxHeap() {
		// if len(result) > 1 && this.getHeap()[this._ZERO].CompareTo(this.getHeap()[1]) == 1 {
		// 	fmt.Println("inversing")
		// 	inverse := MakeArrayOfComparable(len(array))
		// 	for i := 0; i < len(array); i++ {
		// 		inverse[i] = this.heap[len(array)-1-i]
		// 	}
		// 	result = inverse
		// }
	}
	return result
}

/*
A n-ésima estatística de ordem de uma heap é o n-ésimo menor elemento da estrutura.
Tendo como exemplo a seguinte min heap: [1, 6, 58, 12, 34, 64, 99, 82];
O elemento que tem estatística de ordem igual a 3 é o número 12,
pois o mesmo é o 3º menor elemento da heap.
*/
func (this HeapImpl) EstatisticaDeOrdem(n int) Comparable {
	var n_Element Comparable

	for n > 0 && n <= this.index {
		n_Element = this.ExtractRootElement()
		n--
	}

	return n_Element
}

/*
Breadth-first search (BFS) ou Encaminhamento em Largura é uma forma de percorrer uma árvore
visitando os nós vizinhos de um determinado nível desta árvore.
*/
func (this HeapImpl) BFS(level int) ([]Comparable, error) {
	if level > this.index || level < 0 {
		return nil, fmt.Errorf("Level %d inexistente", level)
	}

	maxElements := int(math.Pow(2, float64(level)) - 1)
	elementosPorLevel := make([]Comparable, maxElements)

	fmt.Println("maxElements", maxElements)

	for i := 0; i < maxElements && i < len(this.heap); i++ {
		elementosPorLevel[i] = this.heap[i]
	}

	return elementosPorLevel, nil

}

/**
Funcionamento para buscar os n maiores valores
1. Criar uma MinHeapAmostra com todos os valores do array da amostra
2. Criar outra MinHeapResultado de tamanho n, para armanezar os maiores valores, com a primeira metade da MinHeapAmostra
3. Varrer a segunda metade MinHeapAmostra, verificando quais valores são maiores que a raiz da MinHeapResultado
4. Coloque esse valor na MinHeapResultado e chamar o Heapify(0)
*/
func (this HeapImpl) VisitLargestFromHeap(n int) ([]Comparable, error) {
	// minHeap := New(getMinHeapComparator())

	middle := (len(this.heap) / 2) // + 1
	// fmt.Println("heap... ", this.heap)

	fisrtMiddleArray := MakeArrayOfComparable(n)
	copy(fisrtMiddleArray, this.heap[:middle])
	// fmt.Println("arr minheap ", fisrtMiddleArray)

	minHeapResultado := NewMinHeap(fisrtMiddleArray)

	for idx := middle; idx < len(this.heap); idx++ {
		if this.getHeap()[idx].CompareTo(minHeapResultado.RootElement()) > this._ZERO {
			if minHeapResultado.IsFull() {
				minHeapResultado.ExtractRootElement()
			}
			minHeapResultado.Insert(this.getHeap()[idx])
		}
	}

	return minHeapResultado.ToArray(), nil
}

/***************************
*****  public Access Methods
****************************/

func (this HeapImpl) ToArray() []Comparable {
	resp := MakeArrayOfComparable(this.index + 1)
	for i := 0; i <= this.index; i++ {
		resp[i] = this.heap[i]
	}
	return resp
}

func (this HeapImpl) RootElement() Comparable {
	var root Comparable

	if !this.IsEmpty() {
		root = this.getHeap()[this._ZERO]
	}

	return root
}

func (this HeapImpl) IsEmpty() bool {
	return this.index == -1
}

func (this HeapImpl) IsFull() bool {
	return this.index == len(this.heap)-1
}

func (this HeapImpl) Size() int {
	return this.index + 1
}

func (this HeapImpl) GetComparator() Comparator {
	return this.comparator
}

func (this *HeapImpl) SetComparator(comparator Comparator) {
	this.comparator = comparator
}

func (this HeapImpl) IsMaxHeap() bool {

	c1 := dummyComparable{1}
	c2 := dummyComparable{2}

	return this.comparator(c1, c2) == -1
}

func (this HeapImpl) IsMinHeap() bool {

	c1 := dummyComparable{1}
	c2 := dummyComparable{2}

	return this.comparator(c1, c2) == 1
}
