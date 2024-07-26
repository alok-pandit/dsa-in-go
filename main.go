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
			// TODO: This needs to be dynamic instead of hard-coded, so that if another data structure, if added, shouldn't mess up the selection logic, especially for Exit!
			fmt.Println("\nChoose ur DS:\nLists:\n 1. Singly Linked List\n 2. ArrayList\n 3. Doubly Linked List\nSets:\n 4. Hashed Set\n 5. Tree Set\n 6. Linked Hash Set\nStacks:\n 7. Array Stack\n 8. Linked List Stack\nMaps:\n 9.Hash Map\n10. Tree Map\n11. Linked Hash Map\n12. Hashed Bidi Map\n13. Tree Bidi Map\nTrees:\n14. Red Black Trees\n15. AVL Trees\n16. BTree\n17. Binary Heap\nQueues:\n18. Linked List Queue\n19. Array Queue\n20. Circular Buffer\n21. Priority Queue\n\n22. Exit")

			fmt.Print("Enter Your Choice: ")
			fmt.Scanln(&dataStructure)

		}

		flag = false

		fmt.Println("\nChoose Operation:")

		choices := initiator.GetChoices(dataStructure)

		var choiceConter int
		for i, o := range choices {
			fmt.Println(i+1, ". "+o)
			if i == len(choices)-1 {
				choiceConter = i + 2
				fmt.Println(choiceConter, ". Exit")
			}
		}

		fmt.Print("Enter Your Choice: ")
		fmt.Scanln(&subChoice)

		if subChoice == strconv.Itoa(choiceConter) {
			flag = true
			continue
		}

		fmt.Print("\nOutput: ")
		initiator.ExecuteAction(dataStructure + "." + subChoice)
	}
}
