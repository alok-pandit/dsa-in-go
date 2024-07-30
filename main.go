package main

import (
	"fmt"
	"strconv"

	initiator "github.com/alok-pandit/dsa-in-go/src"
)

func main() {
	initiator.InitializeAllDS()
	var dataStructure, subChoice string

	flag := true
	for {
		if flag {

			// TODO: This needs to be dynamic instead of hard-coded, so that if another data structure is added, shouldn't mess up the selection logic, especially for Exit!
			fmt.Println("\nChoose ur DS:\n\nLists:\n1. Singly Linked List\t2. Doubly Linked List\n\nSets:\n3. Hashed Set\t4. Tree Set\t5. Linked Hash Set\n\nStacks:\n6. Array Stack\t7. Linked List Stack\n\nMaps:\n8.Hash Map\t9. Tree Map\t10. Linked Hash Map\t11. Hashed Bidi Map\t12. Tree Bidi Map\n\nTrees:\n13. Red Black Trees\t14. AVL Trees\t15. BTree\t16. Binary Heap\n\nQueues:\n17. Linked List Queue\t18. Array Queue\t19. Circular Buffer\t20. Priority Queue\n\n21. Exit")
			fmt.Print("Enter Your Choice: ")
			fmt.Scanln(&dataStructure)

		}

		flag = false

		fmt.Println("\nChoose Operation:")

		choices := initiator.GetChoices(dataStructure)

		var exitIndex int

		for i, o := range choices {
			if i < 9 {
				fmt.Println("", i+1, ". "+o)
			} else {
				fmt.Println(i+1, ". "+o)
			}
			if i == len(choices)-1 {
				exitIndex = i + 2
				fmt.Println(exitIndex, ". Exit")
			}
		}

		fmt.Print("Enter Your Choice: ")
		fmt.Scanln(&subChoice)

		if subChoice == strconv.Itoa(exitIndex) {
			flag = true
			continue
		}

		initiator.ExecuteAction(dataStructure + "." + subChoice)

	}
}
