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
	case "5":
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
		fmt.Print("Enter the value to insert: ")
		var value string
		fmt.Scanln(&value)
		SLL.AppendHead(value)

	case "1.2":
		SLL.PrintList()

	case "1.3":
		SLL.Reverse()
		SLL.PrintList()

	default:
		fmt.Println("Invalid choice")
	}
}
