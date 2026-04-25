package gheap

import (
	"fmt"
)

// ================PACKAGE SPACE=====================
type IHeap interface {
	AddItem(value int)
	expandHeap(heap *Heap, count int)
	swapItemsHeap(heap *Heap, i, j int)
}

func swapItemsHeap(heap *Heap, i, j int) {
	heap.memory[i], heap.memory[j] =
		heap.memory[j], heap.memory[i]
}

func expandHeap(heap *Heap, count int) {
	heap.memory = append(heap.memory, make([]int, count)...)
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
	fmt.Println(heap.memory)
}

// ================USER SPACE=====================
type Heap struct {
	memory []int
	size   int
}

func (heap *Heap) AddItem(value int) {
	expandHeap(heap, 1)
	heap.memory[heap.size-1] = value
	swapRecursive(heap, heap.size)
}
