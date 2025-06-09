package trie

import (
	"fmt"
	"strings"
)

const arraySize = 26

type ITrie interface {
	Insert(word string)
	Search(word string) bool
	Delete(word string)
	Print()
}

func GetChoices() []string {
	return []string{
		"Insert",
		"Search",
		"Delete",
		"Print",
	}
}

type TrieNode struct {
	children [arraySize]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: [arraySize]*TrieNode{},
			isEnd:    false,
		},
	}
}

func (t *Trie) Insert(word string) {

	curr := t.root

	for _, ascii := range word {

		ascii := ascii - 'a'

		if curr.children[ascii] == nil {

			curr.children[ascii] = &TrieNode{
				children: [arraySize]*TrieNode{},
				isEnd:    false,
			}

		}

		curr = curr.children[ascii]

	}

	curr.isEnd = true

}

func (t *Trie) Search(word string) bool {

	current := t.root

	for _, ascii := range word {

		ascii := ascii - 'a'

		if current.children[ascii] == nil {
			return false
		}

		current = current.children[ascii]

	}

	return current.isEnd

}

func (t *Trie) Delete(word string) {

	t.deleteHelper(t.root, word, 0)

}

func (t *Trie) deleteHelper(node *TrieNode, word string, index int) bool {

	if index == len(word) {

		if !node.isEnd {
			return false
		}

		node.isEnd = false

		return !node.isEnd && !hasChildren(node)

	}

	ascii := word[index] - 'a'

	if node.children[ascii] == nil {
		return false
	}

	shouldDelete := t.deleteHelper(node.children[ascii], word, index+1)

	if shouldDelete {

		node.children[ascii] = nil

		return !node.isEnd && !hasChildren(node)

	}

	return false

}

func hasChildren(node *TrieNode) bool {

	for i := range arraySize {

		if node.children[i] != nil {
			return true
		}

	}

	return false

}

func (t *Trie) Print() {
	t.printNode(t.root, 0, "")
}

func (t *Trie) printNode(node *TrieNode, level int, prefix string) {

	if node == nil {
		return
	}

	indent := strings.Repeat("  ", level)

	branch := ""

	if level > 0 {
		branch = "└─"
	}

	char := ""

	if prefix != "" {
		char = fmt.Sprintf("'%s'", string('a'+rune(prefix[0])))
	}

	endMarker := ""

	if node.isEnd {
		endMarker = " *"
	}

	fmt.Printf("%s%s%s%s\n", indent, branch, char, endMarker)

	for i, child := range node.children {

		if child != nil {
			t.printNode(child, level+1, string(rune(i)))
		}

	}

}

// * This ensures that the Trie struct
// * implements the ITrie interface
var _ ITrie = (*Trie)(nil)
