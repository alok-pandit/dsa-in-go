package initiator

import (
	"fmt"
	"os"

	"github.com/alok-pandit/dsa-in-go/src/linkedlist"
)

var SLL *linkedlist.SinglyLinkedList

func InitializeAllDS() {
	SLL = &linkedlist.SinglyLinkedList{}
}

func GetChoices(ds string) []string {
	switch ds {
	case "1":
		return linkedlist.GetChoices()
	case "22":
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
		fmt.Print("Enter the value to insert at end: ")
		var value string
		fmt.Scanln(&value)
		SLL.AppendToEnd(value)

	case "1.2":
		fmt.Print("Enter the value to insert before start: ")
		var value string
		fmt.Scanln(&value)
		SLL.AppendBeforeStart(value)

	case "1.3":
		fmt.Print("Enter the value to insert: ")
		var value, node string
		fmt.Scanln(&value)
		fmt.Println("Enter the node after which this value is to be inserted:")
		fmt.Scanln(&node)
		SLL.AppendAfter(value, node)

	case "1.4":
		fmt.Print("Enter the value to insert: ")
		var value, node string
		fmt.Scanln(&value)
		fmt.Println("Enter the node before which this value is to be inserted:")
		fmt.Scanln(&node)
		SLL.AppendBefore(value, node)

	case "1.5":
		SLL.Reverse()

	case "1.6":
		SLL.PrintList()

	default:
		fmt.Println("Invalid choice")
	}

	SLL.PrintList()

}
