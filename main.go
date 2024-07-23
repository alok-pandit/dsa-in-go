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
			fmt.Println("\nChoose ur DS:\n1. Linked List\n2. Stack\n3. Queue\n4. Doubly Linked List\n5. Exit")
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

		initiator.ExecuteAction(dataStructure + "." + subChoice)
	}
}
