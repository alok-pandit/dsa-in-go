package trie

type TrieChild struct {
	val   rune
	child *ListTrieNode
	next  *TrieChild
}

type ListTrieNode struct {
	children *TrieChild
	isEnd    bool
}

type ListTrie struct {
	root *ListTrieNode
}

func NewListTrie() *ListTrie {
	return &ListTrie{
		root: &ListTrieNode{
			children: nil,
			isEnd:    false,
		},
	}
}

func (t *ListTrie) Insert(word string) {
	current := t.root
	for _, char := range word {
		child := t.findChild(current, char)
		if child == nil {
			newChild := &TrieChild{
				val: char,
				child: &ListTrieNode{
					children: nil,
					isEnd:    false,
				},
				next: current.children,
			}
			current.children = newChild
			current = newChild.child
		} else {
			current = child.child
		}
	}
	current.isEnd = true
}

func (t *ListTrie) findChild(node *ListTrieNode, char rune) *TrieChild {
	current := node.children
	for current != nil {
		if current.val == char {
			return current
		}
		current = current.next
	}
	return nil
}

func (t *ListTrie) Search(word string) bool {
	current := t.root
	for _, char := range word {
		child := t.findChild(current, char)
		if child == nil {
			return false
		}
		current = child.child
	}
	return current.isEnd
}

func (t *ListTrie) StartsWith(prefix string) bool {
	current := t.root
	for _, char := range prefix {
		child := t.findChild(current, char)
		if child == nil {
			return false
		}
		current = child.child
	}
	return true
}

func (t *ListTrie) Delete(word string) bool {
	return t.deleteHelper(t.root, word, 0)
}

func (t *ListTrie) deleteHelper(node *ListTrieNode, word string, index int) bool {
	if index == len(word) {
		if !node.isEnd {
			return false
		}
		node.isEnd = false
		return node.children == nil
	}

	char := rune(word[index])
	child := t.findChild(node, char)
	if child == nil {
		return false
	}

	shouldDeleteChild := t.deleteHelper(child.child, word, index+1)

	if shouldDeleteChild {
		t.removeChild(node, char)
		return !node.isEnd && node.children == nil
	}

	return false
}

func (t *ListTrie) removeChild(node *ListTrieNode, char rune) {
	if node.children == nil {
		return
	}

	if node.children.val == char {
		node.children = node.children.next
		return
	}

	current := node.children
	for current.next != nil {
		if current.next.val == char {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

func (t *ListTrie) GetAllWords() []string {
	var words []string
	t.getAllWordsHelper(t.root, "", &words)
	return words
}

func (t *ListTrie) getAllWordsHelper(node *ListTrieNode, prefix string, words *[]string) {
	if node.isEnd {
		*words = append(*words, prefix)
	}

	current := node.children
	for current != nil {
		t.getAllWordsHelper(current.child, prefix+string(current.val), words)
		current = current.next
	}
}
