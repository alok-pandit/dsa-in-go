package trie

type MapTrieNode struct {
	children map[rune]*MapTrieNode
	isEnd    bool
}

type MapTrie struct {
	root *MapTrieNode
}

func NewMapTrie() *MapTrie {
	return &MapTrie{
		root: &MapTrieNode{
			children: make(map[rune]*MapTrieNode),
			isEnd:    false,
		},
	}
}

func (t *MapTrie) Insert(word string) {
	current := t.root
	for _, char := range word {
		if current.children[char] == nil {
			current.children[char] = &MapTrieNode{
				children: make(map[rune]*MapTrieNode),
				isEnd:    false,
			}
		}
		current = current.children[char]
	}
	current.isEnd = true
}

func (t *MapTrie) Search(word string) bool {
	current := t.root
	for _, char := range word {
		if current.children[char] == nil {
			return false
		}
		current = current.children[char]
	}
	return current.isEnd
}

func (t *MapTrie) StartsWith(prefix string) bool {
	current := t.root
	for _, char := range prefix {
		if current.children[char] == nil {
			return false
		}
		current = current.children[char]
	}
	return true
}

func (t *MapTrie) Delete(word string) bool {
	return t.deleteHelper(t.root, word, 0)
}

func (t *MapTrie) deleteHelper(node *MapTrieNode, word string, index int) bool {
	if index == len(word) {
		if !node.isEnd {
			return false
		}
		node.isEnd = false
		return len(node.children) == 0
	}

	char := rune(word[index])
	childNode := node.children[char]
	if childNode == nil {
		return false
	}

	shouldDeleteChild := t.deleteHelper(childNode, word, index+1)

	if shouldDeleteChild {
		delete(node.children, char)
		return !node.isEnd && len(node.children) == 0
	}

	return false
}

func (t *MapTrie) GetAllWords() []string {
	var words []string
	t.getAllWordsHelper(t.root, "", &words)
	return words
}

func (t *MapTrie) getAllWordsHelper(node *MapTrieNode, prefix string, words *[]string) {
	if node.isEnd {
		*words = append(*words, prefix)
	}

	for char, child := range node.children {
		t.getAllWordsHelper(child, prefix+string(char), words)
	}
}
