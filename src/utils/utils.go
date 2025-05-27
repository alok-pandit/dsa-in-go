package utils

import (
	"fmt"
	"strings"
)

var DataStructures = []string{
	"Singly Linked List",
	"Doubly Linked List",
	"Hash Set",
	"Tree Set",
	"Linked Hash Set",
	"Array Stack",
	"Linked List Stack",
	"Hash Map",
	"Tree Map",
	"Linked Hash Map",
	"Hashed Bidi Map",
	"Tree Bidi Map",
	"Red Black Trees",
	"AVL Trees",
	"BTree",
	"Binary Tree",
	"Linked List Queue",
	"Array Queue",
	"Circular Buffer Queue",
	"Priority Queue",
}

func GetValue() string {

	fmt.Print("Enter the value to insert: ")

	var value string

	fmt.Scanln(&value)

	if strings.Trim(value, " ") == "" {

		_ = fmt.Errorf("data cannot be empty")

		return ""

	}

	return value

}

func GetNode() string {

	fmt.Print("Enter the node: ")

	var node string

	fmt.Scanln(&node)

	if strings.Trim(node, " ") == "" {

		_ = fmt.Errorf("node cannot be empty")

		return ""

	}

	return node

}

func PrintChoices(choices []string) int {

	var exitIndex int

	for i, o := range choices {

		if i < 9 {

			fmt.Println("", i+1, ". "+o)

		} else {

			fmt.Println(i+1, ". "+o)

		}
		if i == len(choices)-1 {

			exitIndex = i + 2

			if exitIndex < 9 {

				fmt.Println("", exitIndex, ". Exit")

			} else {

				fmt.Println(exitIndex, ". Exit")

			}

		}

	}

	return exitIndex

}
