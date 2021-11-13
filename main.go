package main

import (
	"www.github.com/purnachand-jaddu/goLang/bst"
	"www.github.com/purnachand-jaddu/goLang/heap"
	"www.github.com/purnachand-jaddu/goLang/linkedlist"
	"www.github.com/purnachand-jaddu/goLang/queue"
	"www.github.com/purnachand-jaddu/goLang/set"
	"www.github.com/purnachand-jaddu/goLang/stack"
)

func main() {
	bst.TestBST()
	heap.TestHeap()
	linkedlist.TestLinkedList()
	queue.TestQueue()
	set.TestHashSet()
	stack.TestStack()
}
