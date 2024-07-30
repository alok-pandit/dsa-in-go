package doublylinkedlist

import "fmt"

type IDoublyLinkedList interface {
	Append(string)
	Prepend(string)
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
	PrintLength()
	PrintList()
}

func GetChoices() []string {

	return []string{
		"Append",
		"Prepend",
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
		"Print Length",
		"Print",
	}

}

type node struct {
	data string
	next *node
	prev *node
}

type DoublyLinkedList struct {
	head   *node
	tail   *node
	length int
}

func (l *DoublyLinkedList) Append(data string) {

	if l.head == nil {
		l.head = &node{data: data, next: nil, prev: nil}
		l.tail = l.head
	} else {
		newNode := &node{data: data, next: nil, prev: l.tail}
		l.tail.next = newNode
		l.tail = newNode
	}

	l.length++

}

func (l *DoublyLinkedList) Prepend(data string) {

	if l.head == nil {
		l.head = &node{data: data, next: nil, prev: nil}
		l.tail = l.head
	} else {
		newNode := &node{data: data, next: l.head, prev: nil}
		l.head.prev = newNode
		l.head = newNode
	}

	l.length++

}

func (l *DoublyLinkedList) findNode(nodeData string) *node {

	currentNode := l.head

	for currentNode != nil && currentNode.data != nodeData {
		currentNode = currentNode.next
	}

	return currentNode

}

func (l *DoublyLinkedList) AppendAfter(nodeData, data string) {

	currentNode := l.findNode(nodeData)

	if currentNode == nil {
		return
	}

	if currentNode == l.tail {

		l.Append(data)
		return

	}
	newNode := &node{data: data, next: currentNode.next, prev: currentNode}

	currentNode.next.prev = newNode

	currentNode.next = newNode

	l.length++

}

func (l *DoublyLinkedList) AppendBefore(nodeData, data string) {

	currentNode := l.findNode(nodeData)

	if currentNode == nil {
		return
	}

	if currentNode == l.head {

		l.Prepend(data)
		return

	}

	newNode := &node{data: data, next: currentNode, prev: currentNode.prev}

	currentNode.prev.next = newNode

	currentNode.prev = newNode

	l.length++

}

func (l *DoublyLinkedList) DeleteHead() {

	if l.head == nil {
		return
	}

	l.head = l.head.next

	l.length--

}

func (l *DoublyLinkedList) DeleteTail() {

	if l.head == nil {
		return
	}

	l.tail = l.tail.prev

	l.length--

}

func (l *DoublyLinkedList) DeleteNode(nodeData string) {

	currentNode := l.findNode(nodeData)

	if currentNode == nil {
		return
	}

	currentNode.prev.next = currentNode.next

	currentNode.next.prev = currentNode.prev

	l.length--

}

func (l *DoublyLinkedList) DeleteOneBefore(nodeData string) {

	currentNode := l.findNode(nodeData)

	if currentNode == nil {
		return
	}

	if currentNode == l.head {
		return
	}

	if currentNode.prev == l.head {

		currentNode.prev = nil

		l.head = currentNode

		return

	}

	currentNode.prev.prev.next = currentNode

	currentNode.prev = currentNode.prev.prev

	l.length--

}

func (l *DoublyLinkedList) DeleteOneAfter(nodeData string) {

	currentNode := l.findNode(nodeData)

	if currentNode == nil {
		return
	}

	if currentNode == l.tail {
		return
	}

	if currentNode.next == l.tail {

		currentNode.next = nil

		l.tail = currentNode

		return

	}

	currentNode.next.next.prev = currentNode

	currentNode.next = currentNode.next.next

	l.length--

}

func (l *DoublyLinkedList) DeleteAllBefore(nodeData string) {

	currentNode := l.findNode(nodeData)

	if currentNode == nil {
		return
	}

	l.head = currentNode

	currentNode.prev = nil

	l.length--

}

func (l *DoublyLinkedList) DeleteAllAfter(nodeData string) {

	currentNode := l.findNode(nodeData)

	if currentNode == nil {
		return
	}

	l.tail = currentNode

	currentNode.next = nil

	l.length--

}

func (l *DoublyLinkedList) Reverse() {

	currentNode := l.head

	for currentNode != nil {

		currentNode.next, currentNode.prev = currentNode.prev, currentNode.next
		currentNode = currentNode.prev

	}

	l.head, l.tail = l.tail, l.head

}

func (l *DoublyLinkedList) PrintLength() {

	fmt.Println(l.length)

}

func (l *DoublyLinkedList) PrintList() {

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

// * This ensures that the DoublyLinkedList struct
// * implements the IDoublyLinkedList interface
var _ IDoublyLinkedList = (*DoublyLinkedList)(nil)
