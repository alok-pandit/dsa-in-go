package initiator

import (
	"fmt"
	"os"
	"strconv"

	"github.com/alok-pandit/dsa-in-go/src/lists/doublylinkedlist"
	"github.com/alok-pandit/dsa-in-go/src/lists/linkedlist"
	"github.com/alok-pandit/dsa-in-go/src/stacks/arraystack"
	"github.com/alok-pandit/dsa-in-go/src/stacks/linkedliststack"
	"github.com/alok-pandit/dsa-in-go/src/trees/binaryTree"
	"github.com/alok-pandit/dsa-in-go/src/utils"
)

var SLL *linkedlist.SinglyLinkedList
var DLL *doublylinkedlist.DoublyLinkedList
var AS *arraystack.ArrayStack
var LLS *linkedliststack.LinkedListStack
var BT *binaryTree.BinaryTree

func InitializeAllDS() {
	SLL = &linkedlist.SinglyLinkedList{}
	DLL = &doublylinkedlist.DoublyLinkedList{}
	AS = arraystack.NewArrayStack(0)
	LLS = &linkedliststack.LinkedListStack{}
	BT = &binaryTree.BinaryTree{}
}

func GetChoices(ds string) []string {

	dsInt, err := strconv.Atoi(ds)

	if err != nil || dsInt < 1 || dsInt > len(utils.DataStructures)+1 {

		fmt.Println("Invalid choice")

		return nil

	}

	if dsInt == len(utils.DataStructures)+1 {

		os.Exit(0)

		return nil

	}

	switch utils.DataStructures[dsInt-1] {

	case "Singly Linked List":
		return linkedlist.GetChoices()

	case "Doubly Linked List":
		return doublylinkedlist.GetChoices()

	case "Array Stack":
		return arraystack.GetChoices()

	case "Linked List Stack":
		return linkedlist.GetChoices()

	case "Binary Tree":
		return binaryTree.GetChoices()

	default:
		fmt.Println("Invalid choice")
		return nil

	}

}

func handleSinglyLinkedListCases(action string) {
	switch action {
	case "1.1":
		value := utils.GetValue()
		SLL.AppendToEnd(value)
	case "1.2":
		value := utils.GetValue()
		SLL.AppendBeforeStart(value)
	case "1.3":
		value := utils.GetValue()
		node := utils.GetNode()
		SLL.AppendAfter(value, node)
	case "1.4":
		value := utils.GetValue()
		node := utils.GetNode()
		SLL.AppendBefore(value, node)
	case "1.5":
		SLL.DeleteHead()
	case "1.6":
		SLL.DeleteTail()
	case "1.7":
		node := utils.GetNode()
		SLL.DeleteNode(node)
	case "1.8":
		node := utils.GetNode()
		SLL.DeleteOneBefore(node)
	case "1.9":
		node := utils.GetNode()
		SLL.DeleteOneAfter(node)
	case "1.10":
		node := utils.GetNode()
		SLL.DeleteAllBefore(node)
	case "1.11":
		node := utils.GetNode()
		SLL.DeleteAllAfter(node)
	case "1.12":
		SLL.Reverse()
	case "1.13":
		SLL.PrintList()
	}
}

func handleDoublyLinkedListCases(action string) {
	switch action {
	case "2.1":
		value := utils.GetValue()
		DLL.Append(value)
	case "2.2":
		value := utils.GetValue()
		DLL.Prepend(value)
	case "2.3":
		value := utils.GetValue()
		node := utils.GetNode()
		DLL.AppendAfter(node, value)
	case "2.4":
		value := utils.GetValue()
		node := utils.GetNode()
		DLL.AppendBefore(node, value)
	case "2.5":
		DLL.DeleteHead()
	case "2.6":
		DLL.DeleteTail()
	case "2.7":
		node := utils.GetNode()
		DLL.DeleteNode(node)
	case "2.8":
		node := utils.GetNode()
		DLL.DeleteOneBefore(node)
	case "2.9":
		node := utils.GetNode()
		DLL.DeleteOneAfter(node)
	case "2.10":
		node := utils.GetNode()
		DLL.DeleteAllBefore(node)
	case "2.11":
		node := utils.GetNode()
		DLL.DeleteAllAfter(node)
	case "2.12":
		DLL.Reverse()
	case "2.13":
		DLL.PrintLength()
	case "2.14":
		DLL.PrintList()
	}
}

func handleArrayStackCases(action string) {
	switch action {
	case "6.1":
		value := utils.GetValue()
		AS.Push(value)
		AS.PrintStack()
	case "6.2":
		AS.Pop()
		AS.PrintStack()
	case "6.3":
		fmt.Println(AS.Peek())
	case "6.4":
		AS.Length()
	case "6.5":
		AS.IsEmpty()
	case "6.6":
		AS.PrintStack()
	}
}

func handleLinkedListStackCases(action string) {
	switch action {
	case "7.1":
		value := utils.GetValue()
		LLS.Push(value)
	case "7.2":
		fmt.Println("Popped: ", LLS.Pop())
	case "7.3":
		fmt.Println("Peeked: ", LLS.Peek())
	case "7.4":
		fmt.Println("Is empty: ", LLS.IsEmpty())
	case "7.5":
		fmt.Println("Size: ", LLS.Size())
	case "7.6":
		fmt.Println("Printing Stack:")
		LLS.PrintStack()
	}
}

func handleBinaryTreeCases(action string) {
	switch action {
	case "16.1":
		value := utils.GetValue()
		intValue, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Error: Invalid input. Please enter an integer.")
			return
		}
		BT.Insert(intValue)
	case "16.2":
		value := utils.GetValue()
		intValue, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Error: Invalid input. Please enter an integer.")
			return
		}
		BT.Delete(intValue)
	case "16.3":
		value := utils.GetValue()
		intValue, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Error: Invalid input. Please enter an integer.")
			return
		}
		fmt.Println(BT.Search(intValue))
	case "16.4":
		BT.InorderTraversal()
	case "16.5":
		BT.PreorderTraversal()
	case "16.6":
		BT.PostorderTraversal()
	case "16.7":
		fmt.Println(BT.IsEmpty())
	case "16.8":
		fmt.Println(BT.Size())
	}
}

func ExecuteAction(action string) {
	switch action[0:2] {
	case "1.":
		handleSinglyLinkedListCases(action)
	case "2.":
		handleDoublyLinkedListCases(action)
	case "6.":
		handleArrayStackCases(action)
	case "7.":
		handleLinkedListStackCases(action)
	case "16":
		handleBinaryTreeCases(action)
	default:
		fmt.Println("Invalid choice")
	}

}
