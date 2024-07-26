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

// AppendToEnd appends a new node with the given data at the end of the linked list.
//
// Parameters:
// - data: a string representing the data to be stored in the new node.
//
// Return type: None.
func (l *SinglyLinkedList) AppendToEnd(data string) {

	// Create a new node with the given data.
	newNode := &node{data: data, next: nil}

	// If the linked list is empty, make the new node the head.
	if l.head == nil {
		l.head = newNode
		return
	}

	// Traverse the linked list to find the last node.
	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	// Append the new node as the next node of the last node.
	currentNode.next = newNode

}

// AppendBeforeStart appends a new node with the given data before the start of the linked list.
//
// Parameters:
// - data: a string representing the data to be stored in the new node.
//
// Return type: None.
//
// This function does the following:
//  1. Creates a new node with the given data.
//  2. If the linked list is empty, it sets the head of the linked list to the new node.
//  3. Otherwise, it sets the next pointer of the new node to the head of the linked list
//     and updates the head of the linked list to the new node.
func (l *SinglyLinkedList) AppendBeforeStart(data string) {
	// Create a new node with the given data.
	newNode := &node{data: data, next: nil}

	// If the linked list is empty, make the new node the head.
	if l.head == nil {
		l.head = newNode
		return
	}

	// Set the next pointer of the new node to the head of the linked list.
	newNode.next = l.head

	// Update the head of the linked list to the new node.
	l.head = newNode
}

// AppendAfter appends a new node with the given data after the node with the specified nodeData in the SinglyLinkedList.
//
// Parameters:
// - data (string): The data of the new node to be appended.
// - nodeData (string): The data of the node after which the new node will be appended.
//
// Return type: None.
//
// This function does the following:
// 1. Creates a new node with the given data.
// 2. If the linked list is empty, it sets the head of the linked list to the new node.
// 3. If the head node has the given data, it sets the head of the linked list to the new node.
// 4. Otherwise, it traverses the linked list to find the node with the specified nodeData.
// 5. It sets the next pointer of the new node to the node with the specified nodeData.
// 6. It sets the next pointer of the previous node to the new node.
func (l *SinglyLinkedList) AppendAfter(data, nodeData string) {

	// Create a new node with the given data.
	newNode := &node{data: data, next: nil}

	// If the linked list is empty, make the new node the head.
	if l.head == nil {
		l.head = newNode
		return
	}

	currentNode := l.head

	// Traverse the linked list to find the node with the specified nodeData.
	for currentNode.data != nodeData {
		currentNode = currentNode.next
	}

	// Append the new node as the next node of the previous node.
	newNode.next = currentNode.next
	currentNode.next = newNode

}

// AppendBefore appends a new node with the given data before the node with the specified nodeData in the SinglyLinkedList.
//
// Parameters:
// - data (string): The data of the new node to be appended.
// - nodeData (string): The data of the node before which the new node will be appended.
//
// Return type: None.
//
// This function does the following:
// 1. Creates a new node with the given data.
// 2. If the linked list is empty, it sets the head of the linked list to the new node.
// 3. If the head node has the given data, it sets the head of the linked list to the new node.
// 4. Otherwise, it traverses the linked list to find the node with the specified nodeData.
// 5. It sets the next pointer of the previous node to the new node.
// 6. It sets the next pointer of the new node to the node with the specified nodeData.
func (l *SinglyLinkedList) AppendBefore(data, nodeData string) {

	// Create a new node with the given data.
	newNode := &node{data: data, next: nil}

	// If the linked list is empty, make the new node the head.
	if l.head == nil {
		l.head = newNode
		return
	}

	// If the head node has the given data, make the new node the head.
	if l.head.data == nodeData {
		newNode.next = l.head
		l.head = newNode
		return
	}

	// Traverse the linked list to find the node with the specified nodeData.
	currentNode := l.head
	for currentNode.next.data != nodeData {
		currentNode = currentNode.next
	}

	// Append the new node as the next node of the previous node.
	newNode.next = currentNode.next
	currentNode.next = newNode

}

// PrintList prints the elements of the SinglyLinkedList.
//
// This function takes no parameters.
// It does not return anything.
// It prints the elements of the linked list to the console.
func (l *SinglyLinkedList) PrintList() {

	fmt.Println("")

	if l.head == nil {
		// If the list is empty, print a message indicating that the list is empty.
		fmt.Println("List is empty")
		return
	}

	currentNode := l.head

	for currentNode.next != nil {
		// Print the data of the current node followed by an arrow.
		fmt.Print(currentNode.data, " -> ")
		// Move to the next node in the list.
		currentNode = currentNode.next
	}

	// Print the data of the last node in the list.
	fmt.Println(currentNode.data)

}

// Reverse the list
//
// This function reverses the order of the nodes in the SinglyLinkedList.
// It does not take any parameters.
// It does not return any values.
func (l *SinglyLinkedList) Reverse() {
	// Initialize currentNode to the head of the list
	currentNode := l.head

	// Variable to keep track of the previous node
	var prevNode *node

	// Loop through the list until we reach the end
	for currentNode != nil {
		// Store the next node before we update the current node's next pointer
		nextNode := currentNode.next

		// Update the current node's next pointer to point to the previous node
		currentNode.next = prevNode

		// Move the previous node pointer to the current node
		prevNode = currentNode

		// Move the current node pointer to the next node
		currentNode = nextNode
	}

	// Update the head of the list to point to the last node (the previous node)
	l.head = prevNode
}

// DeleteHead deletes the head node from the SinglyLinkedList.
//
// This function does not take any parameters.
// It does not return any values.
//
// This function updates the head of the linked list to point to the next node
// after the current head. This effectively removes the head node from the list.
// If the list is empty, it prints a message indicating that the list is empty.
func (l *SinglyLinkedList) DeleteHead() {
	// If the list is empty, print a message and return.
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}

	// Update the head pointer to point to the next node after the current head.
	l.head = l.head.next
}

// DeleteTail deletes the tail node from the SinglyLinkedList.
//
// This function does not take any parameters.
// It does not return any values.
//
// This function works by iterating through the linked list until it finds the
// second-to-last node. It then updates the next pointer of the second-to-last
// node to nil, effectively deleting the last node in the list.
func (l *SinglyLinkedList) DeleteTail() {

	// If the list is empty, print a message and return.
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}

	// If the list has only one node, delete the node and return.
	if l.head.next == nil {
		l.head = nil
		return
	}

	// Initialize the currentNode to the head of the list.
	currentNode := l.head

	// Loop until we reach the second-to-last node in the list.
	for currentNode.next.next != nil {
		currentNode = currentNode.next
	}

	// Update the next pointer of the second-to-last node to nil.
	currentNode.next = nil

}

// DeleteNode deletes a node from the SinglyLinkedList with the given nodeData.
//
// Parameters:
// - nodeData (string): The data of the node to be deleted.
//
// Return type: None.
//
// This function iterates through the linked list until it finds the node with
// the given data. If the node is found, it updates the next pointer of the
// previous node to skip the current node. If the node is not found, it prints
// a message indicating that the node is not present in the list.
func (l *SinglyLinkedList) DeleteNode(nodeData string) {

	// If the list is empty, print a message and return.
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}

	// If the head node has the given data, update the head and return.
	if l.head.data == nodeData {
		l.head = l.head.next
		return
	}

	// Initialize the currentNode to the head of the list.
	currentNode := l.head

	// Traverse the list until the node with the given data is found.
	for currentNode.next != nil && currentNode.next.data != nodeData {
		currentNode = currentNode.next
	}

	// If the node with the given data is not found, print a message and return.
	if currentNode.next == nil {
		fmt.Println("Node not found in the list")
		return
	}

	// Update the next pointer of the previous node to skip the current node.
	currentNode.next = currentNode.next.next

}

// DeleteOneBefore deletes the node before the node with the given data in the SinglyLinkedList.
//
// Parameters:
// - nodeData (string): The data of the node after which the node will be deleted.
//
// Return type: None.
func (l *SinglyLinkedList) DeleteOneBefore(nodeData string) {
	// If the list is empty or has only one node, return.
	if l.head == nil || l.head.next == nil {
		return
	}

	// If the nodeData is the head or the next node is the nodeData,
	// update the head and return.
	if l.head.data == nodeData || l.head.next.data == nodeData {
		return
	}

	// Initialize the previous node pointer to the head.
	var prev *node = l.head

	// Initialize the current node pointer to the second node.
	currentNode := l.head.next

	// Iterate through the list until the nodeData is found or the end of the list is reached.
	for currentNode != nil && currentNode.data != nodeData {
		prev = currentNode
		currentNode = currentNode.next
	}

	// If the nodeData is not found, return.
	if currentNode == nil {
		return
	}

	// Update the next pointer of the previous node to skip the current node.
	prev.next = currentNode.next
}

// DeleteOneAfter deletes the node after the node with the given data in the SinglyLinkedList.
//
// Parameters:
// - nodeData (string): The data of the node after which the node will be deleted.
//
// Return type: None.
func (l *SinglyLinkedList) DeleteOneAfter(nodeData string) {
	// If the list is empty, print a message and return.
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}

	// If the list has only one element, print a message and return.
	if l.head.next == nil {
		fmt.Println("List has only 1 element")
		return
	}

	// Initialize the currentNode to the head of the list.
	currentNode := l.head

	// Traverse the list until the node with the given data is found.
	for currentNode.next != nil && currentNode.data != nodeData {
		currentNode = currentNode.next
	}

	// If the node with the given data is not found, print a message and return.
	if currentNode.next == nil {
		fmt.Println("Node not found in the list")
		return
	}

	// Point the next pointer of the current node to the node after it.
	currentNode.next = currentNode.next.next
}

// DeleteAllBefore deletes all nodes before the node with the given data in the SinglyLinkedList.
//
// Parameters:
// - nodeData (string): The data of the node after which all nodes will be deleted.
//
// Return type: None.
func (l *SinglyLinkedList) DeleteAllBefore(nodeData string) {
	// If the list is empty, print a message and return.
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}

	// If the list has only one element, print a message and return.
	if l.head.next == nil {
		fmt.Println("List has only 1 element")
		return
	}

	// Initialize the current node to the head of the list.
	currentNode := l.head

	// Traverse the list until we find the node with the given data or reach the end of the list.
	for currentNode.data != nodeData && currentNode.next != nil {
		currentNode = currentNode.next
	}

	// If the given node is not found in the list, print a message and return.
	if currentNode.next == nil && currentNode.data != nodeData {
		fmt.Println("Node not found in the list")
		return
	}

	// Update the head of the list to the node with the given data.
	l.head = currentNode
}

// DeleteAllAfter deletes all nodes after the node with the given data in the SinglyLinkedList.
//
// Parameters:
// - nodeData (string): The data of the node after which all nodes will be deleted.
//
// Return type: None.
func (l *SinglyLinkedList) DeleteAllAfter(nodeData string) {
	// If the list is empty, print a message and return.
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}

	// If the list has only one element, print a message and return.
	if l.head.next == nil {
		fmt.Println("List has only 1 element")
		return
	}

	// Initialize the current node to the head of the list.
	currentNode := l.head

	// Traverse the list until we find the node with the given data or reach the end of the list.
	for currentNode.data != nodeData && currentNode.next != nil {
		currentNode = currentNode.next
	}

	// If the given node is not found in the list, print a message and return.
	if currentNode.next == nil && currentNode.data != nodeData {
		fmt.Println("Node not found in the list")
		return
	}

	// Set the next pointer of the current node to nil, which will delete all nodes after it.
	currentNode.next = nil
}

// * This ensures that the SinglyLinkedList struct
// * implements the Linked List interface
var _ LinkedList = (*SinglyLinkedList)(nil)
