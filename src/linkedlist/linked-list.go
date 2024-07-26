package linkedlist

import (
	"fmt"
)

type LinkedList interface {
	AppendToEnd(string)
	AppendBeforeStart(string)
	AppendAfter(string, string)
	AppendBefore(string, string)
	DeleteHead()
	DeleteTail()
	DeleteNode(string)
	DeleteOneBefore(string)
	DeleteOneAfter(string)
	DeleteAllBefore(string)
	DeleteAllAfter(string)
	Reverse()
	PrintList()
}

func GetChoices() []string {

	return []string{
		"Insert after tail",
		"Insert before head",
		"Insert after specific node",
		"Insert before specific node",
		"Delete head",
		"Delete tail",
		"Delete specific node",
		"Delete one before",
		"Delete one after",
		"Delete all before",
		"Delete all after",
		"Reverse",
		"Print",
	}

}

type node struct {
	data string
	next *node
}

type SinglyLinkedList struct {
	head *node
}

// Append a new element to the end of the list
func (l *SinglyLinkedList) AppendToEnd(data string) {

	newNode := &node{data: data, next: nil}

	if l.head == nil {
		l.head = newNode
		return
	}

	currentNode := l.head

	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	currentNode.next = newNode

}

func (l *SinglyLinkedList) AppendBeforeStart(data string) {

	newNode := &node{data: data, next: nil}

	if l.head == nil {
		l.head = newNode
		return
	}

	currentNode := l.head
	newNode.next = currentNode
	l.head = newNode

}

func (l *SinglyLinkedList) AppendAfter(data, nodeData string) {

	newNode := &node{data: data, next: nil}

	if l.head == nil {
		l.head = newNode
		return
	}

	currentNode := l.head

	for currentNode.data != nodeData {
		currentNode = currentNode.next
	}

	newNode.next = currentNode.next
	currentNode.next = newNode

}

func (l *SinglyLinkedList) AppendBefore(data, nodeData string) {

	newNode := &node{data: data, next: nil}

	if l.head == nil {
		l.head = newNode
		return
	}

	if l.head.data == nodeData {
		newNode.next = l.head
		l.head = newNode
		return
	}

	currentNode := l.head

	for currentNode.next.data != nodeData {
		currentNode = currentNode.next
	}

	newNode.next = currentNode.next
	currentNode.next = newNode

}

func (l *SinglyLinkedList) PrintList() {

	fmt.Println("")

	if l.head == nil {
		fmt.Println("List is empty")
		return
	}

	currentNode := l.head

	for currentNode.next != nil {
		fmt.Print(currentNode.data, " -> ")
		currentNode = currentNode.next
	}

	fmt.Println(currentNode.data)

}

// Reverse the list
func (l *SinglyLinkedList) Reverse() {

	currentNode := l.head

	var prevNode *node

	for currentNode != nil {
		nextNode := currentNode.next
		currentNode.next = prevNode
		prevNode = currentNode
		currentNode = nextNode
	}

	l.head = prevNode

}

func (l *SinglyLinkedList) DeleteHead() {

	if l.head == nil {
		fmt.Println("List is empty")
		return
	}

	l.head = l.head.next

}

func (l *SinglyLinkedList) DeleteTail() {

	if l.head == nil {
		fmt.Println("List is empty")
		return
	}

	if l.head.next == nil {
		l.head = nil
		return
	}

	currentNode := l.head

	for currentNode.next.next != nil {
		currentNode = currentNode.next
	}

	currentNode.next = nil

}

func (l *SinglyLinkedList) DeleteNode(nodeData string) {

	if l.head == nil {
		fmt.Println("List is empty")
		return
	}

	if l.head.data == nodeData {
		l.head = l.head.next
		return
	}

	currentNode := l.head

	for currentNode.next != nil && currentNode.next.data != nodeData {
		currentNode = currentNode.next
	}

	if currentNode.next == nil {
		fmt.Println("Node not found in the list")
		return
	}

	currentNode.next = currentNode.next.next

}

func (l *SinglyLinkedList) DeleteOneBefore(nodeData string) {

	if l.head == nil {
		fmt.Println("List is empty")
		return
	}

	if l.head.next == nil {
		fmt.Println("List has only 1 element")
		return
	}

	currentNode := l.head

	for currentNode.next != nil && currentNode.next.next != nil && currentNode.next.next.data != nodeData {
		currentNode = currentNode.next
	}

	if currentNode.next.next == nil {
		fmt.Println("Node not found in the list")
		return
	}

	currentNode.next = currentNode.next.next

}

func (l *SinglyLinkedList) DeleteOneAfter(nodeData string) {

	if l.head == nil {
		fmt.Println("List is empty")
		return
	}

	if l.head.next == nil {
		fmt.Println("List has only 1 element")
		return
	}

	currentNode := l.head

	for currentNode.next != nil && currentNode.data != nodeData {
		currentNode = currentNode.next
	}

	if currentNode.next == nil {
		fmt.Println("Node not found in the list")
		return
	}

	currentNode.next = currentNode.next.next

}

func (l *SinglyLinkedList) DeleteAllBefore(nodeData string) {

	if l.head == nil {
		fmt.Println("List is empty")
		return
	}

	if l.head.next == nil {
		fmt.Println("List has only 1 element")
		return
	}

	currentNode := l.head

	for currentNode.data != nodeData {
		currentNode = currentNode.next
	}

	if currentNode.next == nil {
		fmt.Println("Node not found in the list")
		return
	}

	l.head = currentNode

}

func (l *SinglyLinkedList) DeleteAllAfter(nodeData string) {

	if l.head == nil {
		fmt.Println("List is empty")
		return
	}

	if l.head.next == nil {
		fmt.Println("List has only 1 element")
		return
	}

	currentNode := l.head

	for currentNode.data != nodeData {
		currentNode = currentNode.next
	}

	currentNode.next = nil

}

// * This ensures that the SinglyLinkedList struct
// * implements the Linked List interface
var _ LinkedList = (*SinglyLinkedList)(nil)
