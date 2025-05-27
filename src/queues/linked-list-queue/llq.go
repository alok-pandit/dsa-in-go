package linkedlistqueue

import (
	"errors"
	"fmt"
)

type ILinkedListQueue interface {
	Enqueue(any)
	Dequeue() (any, error)
	Front() (any, error)
	IsEmpty() bool
	Print()
}

func GetChoices() []string {

	return []string{
		"Enqueue",
		"Dequeue",
		"Front",
		"IsEmpty",
		"Print Queue",
	}

}

type node struct {
	data any
	next *node
	prev *node
}

type LinkedListQueue struct {
	head *node
	tail *node
	size int
}

func (llq *LinkedListQueue) Enqueue(data any) {

	newNode := &node{data: data, next: llq.head, prev: nil}

	if llq.head == nil {
		llq.head = newNode
		llq.tail = newNode
	} else {
		llq.head.prev = newNode
		llq.head = newNode
	}

	if llq.tail == nil {
		llq.tail = newNode
	}

	llq.size++

}

func (llq *LinkedListQueue) Dequeue() (any, error) {

	if llq.IsEmpty() {
		return "", errors.New("Queue is empty")
	}

	data := llq.tail.data

	llq.tail = llq.tail.prev

	if llq.tail == nil {
		llq.head = nil
	} else {
		llq.tail.next = nil
	}

	llq.size--

	return data, nil

}

func (llq *LinkedListQueue) Front() (any, error) {

	if llq.IsEmpty() {
		return "", errors.New("Queue is empty")
	}

	return llq.tail.data, nil

}

func (llq *LinkedListQueue) IsEmpty() bool {
	return llq.head == nil
}

func (llq *LinkedListQueue) Print() {

	if llq.head == nil {
		fmt.Println("Queue is empty")
		return
	}

	current := llq.head

	for current != nil {

		fmt.Print(current.data)

		if current.next != nil {
			fmt.Print(" -> ")
		}

		current = current.next

	}

	fmt.Println()

}

// * This ensures that the SinglyLinkedList struct
// * implements the Linked List interface
var _ ILinkedListQueue = (*LinkedListQueue)(nil)
