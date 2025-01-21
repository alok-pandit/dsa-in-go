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
			fmt.Println("\nChoose ur DS:\n\nLists:\n\n\t1. Singly Linked List\t\t2. Doubly Linked List\n\nSets:\n\n\t3. Hash Set\t\t\t4. Tree Set\n\t5. Linked Hash Set\n\nStacks:\n\n\t6. Array Stack\t\t\t7. Linked List Stack\n\nMaps:\n\n\t8.Hash Map\t\t\t9. Tree Map\n\t10. Linked Hash Map\t\t11. Hashed Bidi Map\n\t12. Tree Bidi Map\n\nTrees:\n\n\t13. Red Black Trees\t\t14. AVL Trees\n\t15. BTree\t\t\t16. Binary Tree\n\nQueues:\n\n\t17. Linked List Queue\t\t18. Array Queue\n\t19. Circular Buffer\t\t20. Priority Queue\n\n21. Exit")
			fmt.Print("\nEnter Your Choice: ")
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
