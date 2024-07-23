package linkedlist

import "fmt"

type LinkedList interface {
	AppendHead(string)
	PrintList()
	Reverse()
}

func GetChoices() []string {

	return []string{"Insert at Beginning", "Print List", "Reverse"}

}

type node struct {
	data string
	next *node
}

type SinglyLinkedList struct {
	head *node
}

// Append a new element to the end of the list
func (l *SinglyLinkedList) AppendHead(data string) {
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

// Print the list elements
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
