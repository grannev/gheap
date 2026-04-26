package gheap

import (
	"fmt"
)

type IHeap interface {
	Insert(value int)
	Pop(value int) int
	Size() int
	Max() int
	Print()

	expandHeap(heap *Heap, count int)
	swapItemsHeap(heap *Heap, i, j int)
}

func swapItemsHeap(heap *Heap, i, j int) {
	heap.memory[i], heap.memory[j] =
		heap.memory[j], heap.memory[i]
}

func expandHeap(heap *Heap, count int) {
	if count == 0 {
		return
	}

	if count < 0 {
		heap.memory = heap.memory[(count * -1):]
	}

	for range count {
		heap.memory = append(heap.memory, 0)
	}
	heap.size += count
}

func swapRecursive(heap *Heap, index int) {
	var (
		parrent_index int
	)
	if index == 1 {
		return
	}
	parrent_index = index / 2
	if heap.memory[index-1] > heap.memory[parrent_index-1] {
		swapItemsHeap(heap, index-1, parrent_index-1)
		swapRecursive(heap, parrent_index)
	}
}

// ================USER SPACE=====================
type Heap struct {
	memory []int
	size   int
}

func (heap *Heap) Insert(value int) {
	expandHeap(heap, 1)
	heap.memory[heap.size-1] = value
	swapRecursive(heap, heap.size)
}

func (heap *Heap) Pop(value int) int {
	mx := heap.Max()
	expandHeap(heap, -1)
	swapRecursive(heap, heap.size)
	return mx
}

func (heap Heap) Print() {
	fmt.Println(heap)
}

func (heap Heap) Size() int {
	return heap.size
}

func (heap Heap) Max() int {
	return heap.memory[0]
}
