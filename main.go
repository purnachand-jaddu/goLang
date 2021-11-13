package main

import (
	"github.com/purnachand-jaddu/goLang/tree/main/bst"
	"github.com/purnachand-jaddu/goLang/tree/main/heap"
	"github.com/purnachand-jaddu/goLang/tree/main/linkedlist"
	"github.com/purnachand-jaddu/goLang/tree/main/queue"
	"github.com/purnachand-jaddu/goLang/tree/main/set"
	"github.com/purnachand-jaddu/goLang/tree/main/stack"
)

func main() {
	bst.TestBST()
	heap.TestHeap()
	linkedlist.TestLinkedList()
	queue.TestQueue()
	set.TestHashSet()
	stack.TestStack()
}
