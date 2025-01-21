package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alok-pandit/dsa-in-go/src/initiator"
	"github.com/alok-pandit/dsa-in-go/src/utils"
)

func main() {

	initiator.InitializeAllDS()

	var dataStructure, subChoice string

	flag := true

	for {

		if flag {

			fmt.Println("\nChoose ur Data Structure:")

			categories := make(map[string][]int)

			for i, ds := range utils.DataStructures {

				words := strings.Fields(ds)

				category := words[len(words)-1]

				if category == "List" {

					categories["Lists"] = append(categories["Lists"], i)

				} else if category == "Set" {

					categories["Sets"] = append(categories["Sets"], i)

				} else if category == "Stack" {

					categories["Stacks"] = append(categories["Stacks"], i)

				} else if category == "Map" {

					categories["Maps"] = append(categories["Maps"], i)

				} else if category == "Trees" || category == "Tree" || category == "BTree" {

					categories["Trees"] = append(categories["Trees"], i)

				} else if category == "Queue" {

					categories["Queues"] = append(categories["Queues"], i)

				}

			}

			categoryOrder := []string{"Lists", "Sets", "Stacks", "Maps", "Trees", "Queues"}

			printCategories(categories, categoryOrder)

			fmt.Print("\nEnter Your Choice: ")

			fmt.Scanln(&dataStructure)

		}

		flag = false

		dsInt, err := strconv.Atoi(dataStructure)

		if err != nil || dsInt < 1 || dsInt > len(utils.DataStructures)+1 {

			fmt.Println("Invalid choice")

			return

		}

		choices := initiator.GetChoices(dsInt)

		if choices == nil {

			return

		}

		fmt.Printf("\nChoose Operation for %s:\n\n", utils.DataStructures[dsInt-1])

		var exitIndex int = utils.PrintChoices(choices)

		fmt.Print("\nEnter Your Choice: ")

		fmt.Scanln(&subChoice)

		if subChoice == strconv.Itoa(exitIndex) {

			flag = true

			continue

		}

		initiator.ExecuteAction(dataStructure + "." + subChoice)

	}

}

func printCategories(categories map[string][]int, categoryOrder []string) {

	for _, category := range categoryOrder {

		indices := categories[category]

		fmt.Printf("\n%s:\n\n", category)

		for _, idx := range indices {

			fmt.Printf("\t%d. %s", idx+1, utils.DataStructures[idx])

			if idx%2 == 0 && idx+1 < len(utils.DataStructures) {

				fmt.Printf("\t")

			} else {

				fmt.Println()

			}

		}

		fmt.Println()

	}

	fmt.Println(len(utils.DataStructures)+1, ". Exit")

}
