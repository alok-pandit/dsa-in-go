package linkedlist

import "fmt"

type LinkedList interface {
	AppendToEnd(string)
	AppendBeforeStart(string)
	AppendAfter(string, string)
	AppendBefore(string, string)
	Reverse()
	PrintList()
}

func GetChoices() []string {

	return []string{"Insert after tail", "Insert before head", "Insert after specific node", "insert before specific node", "Reverse", "Print"}

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

	newNode := &node{data: data}

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

	newNode := &node{data: data}

	if l.head == nil {
		l.head = newNode
		return
	}

	currentNode := l.head
	newNode.next = currentNode
	l.head = newNode

}

func (l *SinglyLinkedList) AppendAfter(data, nodeData string) {

	newNode := &node{data: data}

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

	newNode := &node{data: data}

	currentNode := l.head

	if currentNode == nil || currentNode.data == nodeData {
		newNode.next = currentNode
		l.head = newNode
		return
	}

	for currentNode.next.data != nodeData {
		currentNode = currentNode.next
	}

	newNode.next = currentNode.next
	currentNode.next = newNode

}

func (l *SinglyLinkedList) PrintList() {

	currentNode := l.head

	if currentNode == nil {
		fmt.Println("List is empty")
		return
	}

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

// * This ensures that the SinglyLinkedList struct implements the Linked List interface
var _ LinkedList = (*SinglyLinkedList)(nil)
