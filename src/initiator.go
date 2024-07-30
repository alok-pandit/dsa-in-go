package initiator

import (
	"fmt"
	"os"

	"github.com/alok-pandit/dsa-in-go/src/lists/doublylinkedlist"
	"github.com/alok-pandit/dsa-in-go/src/lists/linkedlist"
	"github.com/alok-pandit/dsa-in-go/src/stacks/arraystack"
	"github.com/alok-pandit/dsa-in-go/src/utils"
)

var SLL *linkedlist.SinglyLinkedList
var DLL *doublylinkedlist.DoublyLinkedList
var AS *arraystack.ArrayStack

func InitializeAllDS() {
	SLL = &linkedlist.SinglyLinkedList{}
	DLL = &doublylinkedlist.DoublyLinkedList{}
	AS = arraystack.NewArrayStack(0)
}

func GetChoices(ds string) []string {
	switch ds {
	case "1":
		return linkedlist.GetChoices()
	case "2":
		return doublylinkedlist.GetChoices()
	case "6":
		return arraystack.GetChoices()
	case "21":
		os.Exit(0)
		return nil
	default:
		fmt.Println("Invalid choice")
		return nil
	}
}

func ExecuteAction(action string) {
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

	default:
		fmt.Println("Invalid choice")

	}

}
