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
