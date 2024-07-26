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
			fmt.Println("\nChoose ur DS:\n\nLists:\n1. Singly Linked List\t2. ArrayList\t3. Doubly Linked List\n\nSets:\n4. Hashed Set\t5. Tree Set\t6. Linked Hash Set\n\nStacks:\n7. Array Stack\t8. Linked List Stack\n\nMaps:\n9.Hash Map\t10. Tree Map\t11. Linked Hash Map\t12. Hashed Bidi Map\t13. Tree Bidi Map\n\nTrees:\n14. Red Black Trees\t15. AVL Trees\t16. BTree\t17. Binary Heap\n\nQueues:\n18. Linked List Queue\t19. Array Queue\t20. Circular Buffer\t21. Priority Queue\n\n22. Exit")

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
