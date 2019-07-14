package minMaxHeap

import (
	"fmt"
	"math"
)

type dummyComparable struct {
	num int
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
	IsEmpty() bool
	RootElement() Comparable
	toArray() []Comparable

	Insert(element Comparable)
	BuildHeap(array []Comparable)
	ExtractRootElement() Comparable
	Heapsort(array []Comparable) []Comparable
	EstatisticaDeOrdem(n int) Comparable
	BFS(level int) ([]Comparable, error)
	VisitTop10(n int) ([]Comparable, error)
	IsMaxHeap() bool
	IsMinHeap() bool

	Size() int
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
	var largest int

	if leftIndex <= this.index && this.biggerElement(leftIndex, rightIndex) == leftIndex {
		largest = leftIndex
	} else {
		largest = index
	}

	if rightIndex <= this.index && this.biggerElement(rightIndex, largest) == rightIndex {
		largest = rightIndex
	}

	if largest != index {
		swap(this.getHeap(), index, largest)
		this.heapfyDown(largest)
	}
}

/**
 * Faz o processo de Heapfy de baixo para cima, levando o elemento
 * adicionado para aposição correta
 */
func (this HeapImpl) heapfyUp(position int) {
	currentIndex := position

	for this.biggerElement(currentIndex, this.parent(currentIndex)) == currentIndex && this.parent(currentIndex) != currentIndex {
		swap(this.getHeap(), currentIndex, this.parent(currentIndex))
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
			this.heap = copyOf(this.heap, len(this.heap)+this._INCREASING_FACTOR)
		}

		this.index++
		this.getHeap()[this.index] = element
		this.heapify(this.index)
	}
}

func (this *HeapImpl) BuildHeap(array []Comparable) {
	if len(array) > 0 {
		this.heap = make([]Comparable, len(array))

		this.index = -1

		for _, v := range array {
			if v != nil {
				this.Insert(v)
			}
		}
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
			swap(this.heap, this._ZERO, i)
			this.index--
			this.heapify(this._ZERO)
		}
		result = this.heap

		// if len(result) > 1 && this.IsMaxHeap() {
		// if len(result) > 1 && this.getHeap()[this._ZERO].CompareTo(this.getHeap()[1]) == 1 {
		// 	fmt.Println("inversing")
		// 	inverse := makeArrayOfComparable(len(array))
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

func (this HeapImpl) getBiggerSmallerIndices(idx int) (smaller, bigger int) {
	leftIndex := this.left(idx)
	rightIndex := this.right(idx)

	bigger = this.biggerElement(leftIndex, rightIndex)

	if bigger == leftIndex {
		smaller = rightIndex
	} else {
		smaller = leftIndex
	}

	return
}

func (this HeapImpl) canIWalk(numbers []Comparable, idx, top int) bool {
	canWalk := len(numbers) < cap(numbers)
	canWalk = canWalk && idx < cap(numbers)
	canWalk = canWalk && this.getHeap()[idx].CompareTo(this.getHeap()[top]) < 0
	return canWalk
}

func (this HeapImpl) VisitTop10(n int) ([]Comparable, error) {
	numbers := make([]Comparable, 0, n)

	bigger, smaller := this.getBiggerSmallerIndices(this._ZERO)

	numbers = append(numbers, this.RootElement())
	numbers = this.visitBranch(smaller, bigger, numbers)
	// visiting right sub-tree
	numbers = append(numbers, this.getHeap()[bigger])
	top, _ := this.getBiggerSmallerIndices(bigger)
	numbers = this.visitBranch(bigger, top, numbers)

	return numbers, nil
}

/**
1. escolher o nó com maior e menor valor
2. Percorra a sub-arvore de menor valor até achar valores maior que a raiz da sub-avores e menor que o maior
3. caso ache algum valor maior que o maior, ou não ache mais valores, ou o número de valores seja = 10 retorne
4. retorne para outra sub-avore ou para a função
*/
func (this HeapImpl) visitBranch(idx int, top int, numbers []Comparable) []Comparable {
	if this.canIWalk(numbers, idx, top) {
		numbers = append(numbers, this.getHeap()[idx])

		bigger, smaller := this.getBiggerSmallerIndices(idx)

		if this.canIWalk(numbers, smaller, top) {
			numbers = this.visitBranch(smaller, bigger, numbers)
			numbers = this.visitBranch(bigger, top, numbers)
		}
	}

	return numbers
}

/***************************
*****  public Access Methods
****************************/

func (this HeapImpl) IsEmpty() bool {
	return this.index == -1
}

func (this HeapImpl) toArray() []Comparable {
	resp := makeArrayOfComparable(this.index + 1)
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
