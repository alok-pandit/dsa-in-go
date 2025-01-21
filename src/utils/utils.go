package utils

import (
	"fmt"
	"strings"
)

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
