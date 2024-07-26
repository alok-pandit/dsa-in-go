package initiator

import (
	"fmt"
	"os"
	"strings"

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
		if strings.Trim(value, " ") == "" {
			_ = fmt.Errorf("data cannot be empty")
			return
		}
		SLL.AppendToEnd(value)

	case "1.2":
		fmt.Print("Enter the value to insert before start: ")
		var value string
		fmt.Scanln(&value)
		if strings.Trim(value, " ") == "" {
			_ = fmt.Errorf("data cannot be empty")
			return
		}
		SLL.AppendBeforeStart(value)

	case "1.3":
		fmt.Print("Enter the value to insert: ")
		var value, node string
		fmt.Scanln(&value)
		fmt.Print("Enter the node after which this value is to be inserted:")
		fmt.Scanln(&node)
		if strings.Trim(value, " ") == "" {
			_ = fmt.Errorf("data cannot be empty")
			return
		}
		if strings.Trim(node, " ") == "" {
			_ = fmt.Errorf("nodeData cannot be empty")
			return
		}
		SLL.AppendAfter(value, node)

	case "1.4":
		fmt.Print("Enter the value to insert: ")
		var value, node string
		fmt.Scanln(&value)
		fmt.Print("Enter the node before which this value is to be inserted:")
		fmt.Scanln(&node)
		if strings.Trim(value, " ") == "" {
			_ = fmt.Errorf("data cannot be empty")
			return
		}
		if strings.Trim(node, " ") == "" {
			_ = fmt.Errorf("nodeData cannot be empty")
			return
		}
		SLL.AppendBefore(value, node)

	case "1.5":
		SLL.DeleteHead()

	case "1.6":
		SLL.DeleteTail()

	case "1.7":
		fmt.Print("Enter the node to delete: ")
		var node string
		fmt.Scanln(&node)
		if strings.Trim(node, " ") == "" {
			_ = fmt.Errorf("nodeData cannot be empty")
			return
		}
		SLL.DeleteNode(node)

	case "1.8":
		fmt.Print("Enter the node to delete one before: ")
		var node string
		fmt.Scanln(&node)
		if strings.Trim(node, " ") == "" {
			_ = fmt.Errorf("nodeData cannot be empty")
			return
		}
		SLL.DeleteOneBefore(node)

	case "1.9":
		fmt.Print("Enter the node to delete one after: ")
		var node string
		fmt.Scanln(&node)
		if strings.Trim(node, " ") == "" {
			_ = fmt.Errorf("nodeData cannot be empty")
			return
		}
		SLL.DeleteOneAfter(node)

	case "1.10":
		fmt.Print("Enter the node to delete all before: ")
		var node string
		fmt.Scanln(&node)
		if strings.Trim(node, " ") == "" {
			_ = fmt.Errorf("nodeData cannot be empty")
			return
		}
		SLL.DeleteAllBefore(node)

	case "1.11":
		fmt.Print("Enter the node to delete all after: ")
		var node string
		fmt.Scanln(&node)
		if strings.Trim(node, " ") == "" {
			_ = fmt.Errorf("nodeData cannot be empty")
			return
		}
		SLL.DeleteAllAfter(node)

	case "1.12":
		SLL.Reverse()

	case "1.13":
		SLL.PrintList()

	default:
		fmt.Println("Invalid choice")
	}

	if action != "1.13" {
		SLL.PrintList()
	}

}
